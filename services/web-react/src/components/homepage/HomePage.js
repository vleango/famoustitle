import React, { Component } from 'react';
import { connect } from 'react-redux';

import Header from './Header';
import Sidebar from './Sidebar';
import Article from './Articles';
import Pagination from '../shared/Pagination';

import './css/homepage.css';

export class HomePage extends Component {

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
              <Article />
              <Article />
              <Article />
              <Article />
              <Article />
              <Pagination currentPage={7} totalPages={10} />
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

// const mapStateToProps = (state) => {
// 	return {
// 		articles: state.articles.index.resources,
//     isAuthenticated: !!state.auth.token
// 	};
// };
//
// const mapDispatchToProps = (dispatch) => ({
//   fetchList: () => dispatch(fetchList())
// });

export default connect(null, null)(HomePage);
