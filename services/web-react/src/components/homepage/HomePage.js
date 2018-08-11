import React, {Component, Fragment} from 'react';
import { connect } from 'react-redux';
import { Link } from 'react-router-dom';
import queryString from 'qs';

import Article from './Article';
import Pagination from '../shared/Pagination';
import Spinner from '../shared/Spinner';
import {Helmet} from "react-helmet";

import { fetchList, fetchArchiveArticlesList } from '../../actions/articles';

import './css/homepage.css';
import './css/theme.css';
import {random} from "lodash";

export class HomePage extends Component {

    constructor(props) {
        super(props);
        this.state = {
            search: "",
            loading: true
        };
    }

    async componentDidMount() {
        window.scrollTo(0, 0);

        const parsed = queryString.parse(this.props.location.search, { ignoreQueryPrefix: true });
        // this.props.fetchArchiveArticlesList && await this.props.fetchArchiveArticlesList();
        if(this.props.fetchList) {
            await this.props.fetchList(parsed);
            this.setState({loading: false});
        }
    }

    updateFilter = async (key, value) => {
        let route = "";
        this.setState({loading: true});
        if(this.props.selected && this.props.selected[key] !== value) {
            route = `/?${key}=${value}`;
        }

        this.props.history.push(route);
    };

    artwork = () => {
        const art = [
            { title: 'Mona Lisa', artist: 'Leonardo da Vinci', url: 'https://upload.wikimedia.org/wikipedia/commons/thumb/e/ec/Mona_Lisa%2C_by_Leonardo_da_Vinci%2C_from_C2RMF_retouched.jpg/687px-Mona_Lisa%2C_by_Leonardo_da_Vinci%2C_from_C2RMF_retouched.jpg'},
            { title: 'The Starry Night', artist: 'Vincent van Gogh', url: 'https://upload.wikimedia.org/wikipedia/commons/thumb/e/ea/Van_Gogh_-_Starry_Night_-_Google_Art_Project.jpg/1280px-Van_Gogh_-_Starry_Night_-_Google_Art_Project.jpg' },
            { title: 'Starry Night Over the Rhône', artist: 'Vincent van Gogh', url: 'https://upload.wikimedia.org/wikipedia/commons/thumb/9/94/Starry_Night_Over_the_Rhone.jpg/991px-Starry_Night_Over_the_Rhone.jpg' },
            { title: 'A Sunday Afternoon on the Island of La Grande Jatte', artist: 'Georges Seurat', url: 'https://upload.wikimedia.org/wikipedia/commons/thumb/6/67/A_Sunday_on_La_Grande_Jatte%2C_Georges_Seurat%2C_1884.png/1024px-A_Sunday_on_La_Grande_Jatte%2C_Georges_Seurat%2C_1884.png' }
        ];

        const index = random(0, art.length-1);
        return art[index];
    };

    displaySidebar = () => {
        const art = this.artwork();
        return (
            <section className="sidebar col-lg-5 col-12 sidebar-container" style={{
                background: `linear-gradient(to bottom, rgba(0, 0, 0, 0),rgba(0, 0, 0, 0),rgba(0, 0, 0, 0.8)), url('${art.url}')`
            }}>
                <div className="site-info">
                    <div className="primary-info">
                        <h1>FamousTitle</h1>
                    </div>
                    <div className="secondary-info">
                        <p className="secondary-text">Art meets a software developer's blog</p>
                    </div>
                </div>
            </section>
        );
    };

    displayMainContent = () => {
        if(this.state.loading) {
            return <Spinner />
        }

        let position = 0;

        return (
            this.props.articles.length === 0 ? (
                <p>Results not found</p>
            ) : (
                [
                    this.props.articles.map((article) => {
                        position += 1;
                        return (
                            <Fragment key={position}>
                                <Article key={article.id} article={article} updateFilter={this.updateFilter} />
                            </Fragment>
                        );
                    }),
                    <Pagination key="pagination" {...this.props.pagination} />
                ]
            )
        );
    };

    render() {
        return (
            <div className="">
                <Helmet>
                    <title>FamousTitle.com</title>
                </Helmet>

                <div className="jPanelMenu-panel menu-container-1">
                    <main className="container left-container">
                        <div className="row">
                            { this.displaySidebar() }
                            <section className="col-lg-7 col-12 ml-auto main-content">
                                <div className="sub-nav">
                                    <Link to="/" className="select-posts active">Posts</Link>
                                </div>

                                <div className="home-page-posts animated fadeIn ">
                                    { this.displayMainContent() }
                                </div>

                                <footer className="split-footer">
                                    <Link to="/">FamousTitle.com</Link>
                                </footer>

                            </section>

                        </div>

                    </main>

                </div>

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
    fetchArchiveArticlesList: () => dispatch(fetchArchiveArticlesList())
});

export default connect(mapStateToProps, mapDispatchToProps)(HomePage);
