package config

import "os"

var ElasticSearchHost string

func init() {
	switch os.Getenv("APP_ENV") {
	case "development":
		ElasticSearchHost = "http://datastore-es:9200"
	case "test":
		ElasticSearchHost = "http://datastore-es-test:9200"
	case "ci":
		ElasticSearchHost = "http://localhost:9200"
	case "production":
		ElasticSearchHost = "https://search-tech-writer-production-whcygtwifvuk2aw2zfxfs64ysq.us-west-2.es.amazonaws.com"
	}
}
