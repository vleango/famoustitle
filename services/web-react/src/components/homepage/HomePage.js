import React, { Component } from 'react';
import { connect } from 'react-redux';
import queryString from 'query-string';

import Header from './Header';
import Sidebar from './Sidebar';
import Article from './Article';
import Pagination from '../shared/Pagination';

import { fetchList, fetchArticlesArchiveList } from '../../actions/articles';

import './css/homepage.css';

export class HomePage extends Component {

    constructor(props) {
        super(props);
        this.state = {
            loading: true
        };
    }

    async componentDidMount() {
        const parsed = queryString.parse(this.props.location.search);
        this.props.fetchArticlesArchiveList && this.props.fetchArticlesArchiveList();
        if(this.props.fetchList) {
            await this.props.fetchList(parsed);
            this.setState({loading: false});
        }
    }

    updateFilter = async (key, value) => {
        let route = "";
        let selected = {};

        this.setState({loading: true});
        if(this.props.selected && this.props.selected[key] !== value) {
            route = `/?${key}=${value}`;
            selected = {[key]: value};
        }

        this.props.history.push(route);
        if (this.props.fetchList) {
            await this.props.fetchList(selected);
            this.setState({loading: false});
        }
    };

    mainContent() {
        if(this.state.loading) {
            return <p>Loading...</p>
        }

        return (
            this.props.articles.length === 0 ? (
                <p>Results not found</p>
            ) : (
                [
                    this.props.articles.map((article) => {
                        return <Article key={article.id} article={article} updateFilter={this.updateFilter} />;
                    }),
                    <Pagination key="pagination" {...this.props.pagination} />
                ]
            )
        );
    }

    render() {
        return (
            <div className="canvas">
                <Header />
                <div className="container pt-5">
                    <div className="row">
                        <div className="col-xl-4">
                            <Sidebar updateFilter={this.updateFilter} />
                        </div>
                        <div className="col-xl-8 main--content">
                            { this.mainContent() }
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
        articles: state.articles.index.resources,
        selected: state.articles.index.selected
    };
};

const mapDispatchToProps = (dispatch) => ({
    fetchList: (filters) => dispatch(fetchList(filters)),
    fetchArticlesArchiveList: () => dispatch(fetchArticlesArchiveList())
});

export default connect(mapStateToProps, mapDispatchToProps)(HomePage);
