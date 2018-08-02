import React, {Component, Fragment} from 'react';
import { connect } from 'react-redux';
import { Link } from 'react-router-dom';
import moment from 'moment';
import ReactMarkdown from 'react-markdown';
import FontAwesomeIcon from '@fortawesome/react-fontawesome';
import { Button } from 'reactstrap';
import { Modal } from 'reactstrap';
import Spinner from '../shared/Spinner';
import {Helmet} from "react-helmet";

import { fetchItem, removeItem, itemEditable } from '../../actions/articles';
import './css/ArticleItemPage.css';

import faUser from "@fortawesome/fontawesome-free-solid/faUser";
import faCalendarAlt from "@fortawesome/fontawesome-free-solid/faCalendarAlt";
import faTag from "@fortawesome/fontawesome-free-solid/faTag";
import fontawesome from "@fortawesome/fontawesome/index";
import { toastInProgress, toastSuccess, toastFail } from '../shared/Toast';
fontawesome.library.add(faUser, faCalendarAlt, faTag);

export class ArticleItemPage extends Component {

    constructor(props) {
        super(props);
        this.state = {
            imgModal: false,
            article: null,
            errorMsg: "",
            editable_id: null,
            deleting: false
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
                    errorMsg: error.toString()
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

    imgModalToggle = (props) => {
        this.setState({
            modalData: {alt: props.target.alt, src: props.target.src},
            imgModal: !this.state.imgModal,
        });
    };

    LinkRenderer = (props) => {
        return <a href={props.href} rel="noopener noreferrer" target="_blank">{props.children}</a>
    };

    ImgRenderer = (props) => {
        return <img className={"pointer"} onClick={this.imgModalToggle} alt={props.alt} src={props.src}>{props.children}</img>
    };

    displayBody() {
        return (
            <div className={`body-markdown spacing`}>
                <ReactMarkdown source={ this.state.article.body } renderers={{link: this.LinkRenderer, image: this.ImgRenderer}}/>
            </div>
        )
    }

    onDeleteArticle = async () => {
        const toastID = toastInProgress("Deleting in progress...");
        try{
            this.setState({ deleting: true, errorMsg: "" });
            await this.props.removeItem({id: this.props.match.params.id });
            toastSuccess("Success!", toastID);
            this.setState({ submitting: false, errorMsg: "" });
            this.props.history.push('/');
        } catch(error) {
            let msg = "server error";
            if(error && error.response) {
                msg = error.response.statusText;
            }
            this.setState({ submitting: false, errorMsg: msg });
            toastFail(msg, toastID);
        }
    };

    render() {
        return (
            <div className="canvas">
                <Helmet>
                    {this.state.article && <title>{this.state.article.title} - FamousTitle.com</title>}
                </Helmet>

                <div className="container article-show-container pt-5 pb-5">
                    { this.state.editable_id && (
                        <div className="clearfix">
                            <Button disabled={this.state.deleting} onClick={this.onDeleteArticle} className="float-right" color="info">Delete</Button>
                            <Button tag={Link} to={`/articles/${this.state.editable_id}/edit`} className="float-right mr-4" color="info">Edit</Button>
                        </div>
                    ) }

                    {/* message for when loading article */}
                    { !this.state.article && <Spinner /> }
                    { this.state.errorMsg }

                    {/* after article loads */}
                    { this.state.article && (
                        <Fragment>
                            <h1>{this.state.article.title}</h1>
                            { this.displayInfo() }
                            { this.displayBody() }
                        </Fragment>
                    ) }

                    {/* Img Modal */}
                    <Modal isOpen={this.state.imgModal} toggle={this.imgModalToggle} centered={true} size={'article-img-dialog'}>
                        {this.state.modalData && (
                            <a rel="noopener noreferrer" onClick={this.imgModalToggle} target="_blank" href={this.state.modalData.src}>
                                <img className='modal-img' alt={this.state.modalData.alt} src={this.state.modalData.src} />
                            </a>
                        )}
                    </Modal>
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
    itemEditable: async (data) => await dispatch(itemEditable(data)),
    removeItem: async (data) => await dispatch(removeItem(data))
});

export default connect(mapStateToProps, mapDispatchToProps)(ArticleItemPage);
