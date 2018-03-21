import React, { Component } from 'react';
import { connect } from 'react-redux';
import { Link } from 'react-router-dom';
import { Button } from 'reactstrap';
import moment from 'moment';

import Header from '../shared/headers/Header';

export class ArticleItemPage extends Component {
  render() {
    return (
      <div>
        <Header resourceTitle={this.props.article.title} />
        <div className="container">
          <div className="clearfix">
            <Button tag={Link} to={`/`} className="float-right ml-3" color="danger">Delete</Button>{' '}
            <Button tag={Link} to={`/articles/edit/${this.props.article.id}`} className="float-right" color="info">Edit</Button>{' '}
          </div>
          <h1>{this.props.article.title}</h1>
          <p>{this.props.article.author}</p>
          <p>{ moment(this.props.article.created_at).format('MM-DD-YYYY HH:mm') }</p>
          <p>{ this.props.article.body }</p>
        </div>
      </div>
    );
  }
}

const mapStateToProps = (state, props) => {
	return {
		article: state.articles.show
	};
};

export default connect(mapStateToProps, null)(ArticleItemPage);
