import React, { Component } from 'react';
import { connect } from 'react-redux';

import Header from './Header';
import Sidebar from './Sidebar';
import Article from './Article';
import Pagination from '../shared/Pagination';

import { fetchList } from '../../actions/articles';

import './css/homepage.css';

export class HomePage extends Component {

  componentDidMount() {
    this.props.fetchList && this.props.fetchList();
  }

  render() {
    return (
      <div className="canvas">
        <Header />
        <div className="container pt-5">
          <div className="row">
            <div className="col-xl-4">
              <Sidebar />
            </div>
            <div className="col-xl-8 main--content">
              {
                this.props.articles.length === 0 ? (
                  <p>Loading...</p>
                ) : (
                  [
                    this.props.articles.map((article) => {
                      return <Article key={article.id} article={article} />;
                    }),
                    <Pagination key="pagination" {...this.props.pagination} />
                  ]
                )
              }
            </div>
          </div>
        </div>
        <footer className="text-muted">
          created by Tha Leang
        </footer>
      </div>
    );
  }
}

const mapStateToProps = (state) => {
	return {
    pagination: state.articles.index.pagination,
		articles: state.articles.index.resources
	};
};

const mapDispatchToProps = (dispatch) => ({
  fetchList: () => dispatch(fetchList())
});

export default connect(mapStateToProps, mapDispatchToProps)(HomePage);
