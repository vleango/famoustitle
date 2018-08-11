package elasticsearch

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/vleango/config"
	"github.com/vleango/lib/models"
	"io/ioutil"
	"net/http"
	"os"
	"time"
)

var (
	indexName     = ""
	ErrSaveFailed = errors.New("saved failed")
)

func init() {
	indexName = fmt.Sprintf("famoustitle_articles_%v", os.Getenv("APP_ENV"))
}

type ESResponse struct {
	Took         int          `json:"took"`
	TimedOut     bool         `json:"timed_out"`
	Shards       Shards       `json:"_shards"`
	Hits         Hits         `json:"hits"`
	Aggregations Aggregations `json:"aggregations"`
}

type Shards struct {
	Total      int `json:"total"`
	Successful int `json:"successful"`
	Skipped    int `json:"skipped"`
	Failed     int `json:"failed"`
}

type Hits struct {
	Total    int       `json:"total"`
	MaxScore float64   `json:"max_score"`
	Data     []HitData `json:"hits"`
}

type HitData struct {
	Index  string        `json:"_index"`
	Type   string        `json:"_type"`
	ID     string        `json:"_id"`
	Found  bool          `json:"found"`
	Source Source        `json:"_source"`
	Sort   []interface{} `json:"sort"`
}

type Source struct {
	ID        string    `json:"id"`
	Author    string    `json:"author"`
	Title     string    `json:"title"`
	Subtitle  *string   `json:"subtitle"`
	Body      string    `json:"body"`
	Tags      []string  `json:"tags"`
	ImgUrl    *string   `json:"img_url"`
	Published bool      `json:"published"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type Aggregations struct {
	Archives Archives `json:"archives"`
	Tags     Tags     `json:"tags"`
}

type Archives struct {
	Buckets []Bucket `json:"buckets"`
}

type Bucket struct {
	KeyAsString string      `json:"key_as_string,omitempty"`
	Key         interface{} `json:"key"`
	DocCount    int         `json:"doc_count"`
}

type Tags struct {
	DocCountErrorUpperBound int      `json:"doc_count_error_upper_bound"`
	SumOtherDocCount        int      `json:"sum_other_doc_count"`
	Buckets                 []Bucket `json:"buckets"`
}

func ArticleCreate(item models.Article) (*models.Article, error) {
	url := fmt.Sprintf("%v/%v/default/%v", config.ElasticSearchHost, indexName, item.ID)

	b, err := json.Marshal(item)
	if err != nil {
		return nil, err
	}
	var jsonStr = []byte(string(b))

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))
	req.Header.Set("Content-Type", "application/json")

	client := config.ESClient
	resp, err := client.Do(req)
	defer resp.Body.Close()
	if err != nil {
		return nil, err
	}

	return &item, nil
}

func ArticleDestroy(item models.Article) (*models.Article, error) {
	url := fmt.Sprintf("%v/%v/default/%v", config.ElasticSearchHost, indexName, item.ID)
	req, err := http.NewRequest("DELETE", url, nil)
	client := config.ESClient
	resp, err := client.Do(req)
	defer resp.Body.Close()
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != 200 {
		return nil, ErrSaveFailed
	}

	return &item, nil
}

func ArticleFindAll(params ...map[string]string) ([]models.Article, Aggregations, error) {
	str := ``
	if len(params) > 0 {
		str = matchStr(params[0])
	}

	var jsonStr = []byte(`
{
  "query": {
    "bool": {
      "must": [
        {
          "match": {
            "published": true
          }
        }` +
		str + `
      ]
    }
  }, 
  "sort": {
    "created_at": {
      "order": "desc"
    }
  },
  "aggs": {
    "tags": {
      "terms": {
        "min_doc_count": 0,
        "field": "tags.keyword",
        "size": 10
      }
    }
  }
}
`)
	return find(jsonStr)
}

func ArticleArchives() (Aggregations, error) {
	var jsonStr = []byte(`
{
  "aggs": {
    "archives": {
      "date_histogram": {
		"min_doc_count": 0,
        "field": "created_at",
        "interval": "month"
      }
    }
  }
}
`)

	_, aggs, err := find(jsonStr)
	return aggs, err
}

func ArticleFind(id string) (*models.Article, error) {
	url := fmt.Sprintf("%v/%v/default/%v", config.ElasticSearchHost, indexName, id)
	req, err := http.NewRequest("GET", url, bytes.NewBuffer([]byte{}))
	req.Header.Set("Content-Type", "application/json")

	client := config.ESClient
	resp, err := client.Do(req)
	defer resp.Body.Close()
	if err != nil {
		return nil, err
	}

	body, _ := ioutil.ReadAll(resp.Body)

	esResponse := HitData{}
	err = json.Unmarshal(body, &esResponse)
	if err != nil {
		return nil, err
	}

	if !esResponse.Found {
		return nil, fmt.Errorf("record not found")
	}

	article := models.Article(esResponse.Source)
	return &article, nil
}

func ArticleUpdate(item models.Article) (*models.Article, error) {
	url := fmt.Sprintf("%v/%v/default/%v", config.ElasticSearchHost, indexName, item.ID)

	b, err := json.Marshal(item)
	if err != nil {
		return nil, err
	}
	var jsonStr = []byte(string(b))

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))
	req.Header.Set("Content-Type", "application/json")

	client := config.ESClient
	resp, err := client.Do(req)
	defer resp.Body.Close()
	if err != nil {
		return nil, err
	}

	return &item, nil
}

func matchStr(params map[string]string) string {
	str := ``
	if val, ok := params["tag"]; ok {
		str = "," + matchStrTag(val)
	} else if val, ok := params["date"]; ok {
		str = "," + matchStrDate(val)
	} else if val, ok := params["match"]; ok {
		str = "," + matchStrMatch(val)
	}

	return str
}

func matchStrTag(val string) string {
	return `
	{
		"term": {
			"tags.keyword": "` + val + `"
  		}
	}`
}

func matchStrDate(val string) string {
	return `
	{
		"range": {
			"created_at": {
				"gte" : "` + val + `||/M",
				"lt" :  "` + val + `||+1M/M"
			}
		}
	}`
}

func matchStrMatch(val string) string {
	return `
	{
		"multi_match": {
			"query": "` + val + `",
			"fields": ["title", "body", "tags.keyword"]
		}
	}`
}

func find(jsonStr []byte) ([]models.Article, Aggregations, error) {
	url := fmt.Sprintf("%v/%v/_search", config.ElasticSearchHost, indexName)
	req, err := http.NewRequest("GET", url, bytes.NewBuffer(jsonStr))
	req.Header.Set("Content-Type", "application/json")

	client := config.ESClient
	resp, err := client.Do(req)
	defer resp.Body.Close()
	if err != nil {
		return []models.Article{}, Aggregations{}, err
	}

	body, _ := ioutil.ReadAll(resp.Body)
	esResponse := ESResponse{}
	esResponse.Aggregations.Archives.Buckets = make([]Bucket, 0)
	esResponse.Aggregations.Tags.Buckets = make([]Bucket, 0)

	err = json.Unmarshal(body, &esResponse)
	if err != nil {
		return []models.Article{}, Aggregations{}, err
	}

	articles := make([]models.Article, 0)
	for _, hitData := range esResponse.Hits.Data {
		articles = append(articles, models.Article(hitData.Source))
	}

	return articles, esResponse.Aggregations, nil
}
