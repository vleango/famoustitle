import React, { Component } from 'react';
import { connect } from 'react-redux';
import { Link } from 'react-router-dom';
import { Button, ListGroup, ListGroupItem, Media } from 'reactstrap';
import { truncate } from 'lodash';
import moment from 'moment';

import { fetchList } from '../../actions/articles';
import Header from '../shared/headers/Header';

export class ArticleListPage extends Component {

  componentDidMount() {
    this.props.fetchList && this.props.fetchList();
  }

  render() {
    return (
      <div>
        <Header />
        <div className="clearfix p-3">
          <Button tag={Link} to={`/articles/new`} className="float-right" color="primary">Add</Button>{' '}
        </div>
        {
          this.props.articles.map((article) => {
            return (
              <ListGroup key={article.id}>
                <ListGroupItem tag={Link} to={`/articles/${article.id}`} action>
                  <Media className="mb-5">
                    <Media>
                      <Media object src="http://via.placeholder.com/128x128" alt="Generic placeholder image">
                      </Media>
                    </Media>
                    <Media body className="ml-4">
                      <Media heading>
                        {article.title}
                      </Media>
                      <p>{ article.author }</p>
                      <p>{ moment(article.created_at).format('MM-DD-YYYY HH:mm') }</p>
                      {truncate(article.body, {
                        'length': 150
                      })}
                    </Media>
                  </Media>
                </ListGroupItem>
              </ListGroup>
            )
          })
        }
      </div>
    );
  }
}

const mapStateToProps = (state) => {
	return {
		articles: state.articles.index.articles
	};
};

const mapDispatchToProps = (dispatch) => ({
  fetchList: () => dispatch(fetchList())
});

export default connect(mapStateToProps, mapDispatchToProps)(ArticleListPage);
