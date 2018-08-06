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
import {random} from 'lodash';
import Highlight from 'react-highlight'

import { fetchItem, removeItem, itemEditable } from '../../actions/articles';
import './css/ArticleItemPage.css';
import './css/darcula.css';

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
            <div className="pt-3 article-item-subtitle">
                <FontAwesomeIcon className="mr-2" icon="user"/>
                <span className="mr-5">{this.state.article.author}</span>

                <FontAwesomeIcon className="mr-2" icon="calendar-alt"/>
                <span className="mr-5">{ moment(this.state.article.created_at).format('MM-DD-YYYY') }</span>

                {/*{ this.state.article.tags && this.state.article.tags.length > 0 && (*/}
                    {/*<Fragment>*/}
                        {/*<FontAwesomeIcon className="mr-2" icon="tag"/>*/}
                        {/*{*/}
                            {/*this.state.article.tags.map((tag) => {*/}
                                {/*return [*/}
                                    {/*<Link key={tag} className="mr-2" to={`/?tag=${tag}`}>{tag}</Link>*/}
                                {/*]*/}
                            {/*})*/}
                        {/*}*/}
                    {/*</Fragment>*/}
                {/*)}*/}
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

    CodeRenderer = (props) => {
        return (
            <Highlight>
                {props.value}
            </Highlight>
        )
    };

    displayBody() {
        return (
            <div className={`article-item-body body-markdown spacing`}>
                <ReactMarkdown source={ this.state.article.body } renderers={{link: this.LinkRenderer, image: this.ImgRenderer, code: this.CodeRenderer}}/>
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

    displayArtwork = () => {
        const art = this.artwork();
        return (
            <div className="art-image mt-5 mb-5" style={{
                background: `linear-gradient(to bottom, rgba(0, 0, 0, 0),rgba(0, 0, 0, 0),rgba(0, 0, 0, 0.8)), url('${art.url}')`
            }}>
                <div style={{lineHeight: '3rem', fontFamily: "'Alex Brush', cursive", position: 'absolute', left: '35px', bottom: '8px'}}>
                    <p style={{fontSize: '3rem'}}>{art.title}</p>
                    <p style={{fontSize: '3rem'}}>{art.artist}</p>
                </div>
            </div>
        );
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
                            { this.displayArtwork() }
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

                <footer className="split-footer pb-5">
                    <Link to="/">FamousTitle.com</Link>
                </footer>

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
