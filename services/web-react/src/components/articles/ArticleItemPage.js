import React, {Component, Fragment} from 'react';
import { connect } from 'react-redux';
import { Link } from 'react-router-dom';
import moment from 'moment';
import ReactMarkdown from 'react-markdown';
import FontAwesomeIcon from '@fortawesome/react-fontawesome';
import { Button } from 'reactstrap';

import { fetchItem, itemEditable } from '../../actions/articles';
import './css/ArticleItemPage.css';

import faUser from "@fortawesome/fontawesome-free-solid/faUser";
import faCalendarAlt from "@fortawesome/fontawesome-free-solid/faCalendarAlt";
import faTag from "@fortawesome/fontawesome-free-solid/faTag";
import fontawesome from "@fortawesome/fontawesome/index";
fontawesome.library.add(faUser, faCalendarAlt, faTag);

export class ArticleItemPage extends Component {

    constructor(props) {
        super(props);
        this.state = {
            article: null,
            loadingText: "Loading...",
            editable_id: null
        };
    }

    async componentDidMount() {
        window.scrollTo(0, 0);

        if(this.props.itemEditable) {
            try {
                const response = await this.props.itemEditable(this.props.match.params.id);
                this.setState({
                    editable_id: response["id"]
                });
            } catch (error) {}
        }

        if(this.props.fetchItem) {
            try {
                await this.props.fetchItem(this.props.match.params.id);
            } catch(error) {
                this.setState({
                    loadingText: error.toString()
                });
            }
        }
    }

    componentWillReceiveProps = (nextProps) => {
        if(nextProps.article) {
            this.setState({
                article: nextProps.article
            });
        }
    };

    displayInfo = () => {
        return (
            <div className="pt-3 pb-5">
                <FontAwesomeIcon className="mr-2" icon="user"/>
                <span className="mr-5">{this.state.article.author}</span>

                <FontAwesomeIcon className="mr-2" icon="calendar-alt"/>
                <span className="mr-5">{ moment(this.state.article.created_at).format('MM-DD-YYYY HH:mm') }</span>

                { this.state.article.tags && this.state.article.tags.length > 0 && (
                    <Fragment>
                        <FontAwesomeIcon className="mr-2" icon="tag"/>
                        {
                            this.state.article.tags.map((tag) => {
                                return [
                                    <Link key={tag} className="mr-2" to={`/?tag=${tag}`}>{tag}</Link>
                                ]
                            })
                        }
                    </Fragment>
                )}
            </div>
        );
    };

    displayBody() {
        return (
            <div className={`body-markdown spacing`}>
                <ReactMarkdown source={ this.state.article.body } />
            </div>
        )
    }

    render() {
        return (
            <div className="canvas">
                <div className="container pt-5 pb-5">

                    { this.state.editable_id && (
                        <div className="clearfix">
                            <Button tag={Link} to={`/articles/${this.state.editable_id}/edit`} className="float-right" color="info">Edit</Button>{' '}
                        </div>
                    ) }

                    {/* message for when loading article */}
                    { !this.state.article && <p>{this.state.loadingText}</p> }

                    {/* after article loads */}
                    { this.state.article && (
                        <Fragment>
                            <h1>{this.state.article.title}</h1>
                            { this.displayInfo() }
                            { this.displayBody() }
                        </Fragment>
                    ) }
                </div>
            </div>
        );
    }
}

const mapStateToProps = (state, props) => {
    return {
        article: state.articles.show.resource
    };
};

const mapDispatchToProps = (dispatch) => ({
    fetchItem: async (data) => await dispatch(fetchItem(data)),
    itemEditable: async (data) => await dispatch(itemEditable(data))
});

export default connect(mapStateToProps, mapDispatchToProps)(ArticleItemPage);
