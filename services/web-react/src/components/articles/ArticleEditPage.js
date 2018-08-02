import React, {Component, Fragment} from 'react';
import { connect } from 'react-redux';
import { Link } from 'react-router-dom';
import { Button } from 'reactstrap';
import { Form, FormGroup, Input } from 'reactstrap';
import FontAwesomeIcon from '@fortawesome/react-fontawesome';
import { split } from 'lodash';
import Spinner from '../shared/Spinner';
import {Helmet} from "react-helmet";

import './css/ArticleItemPage.css';
import moment from "moment/moment";
import ReactMarkdown from 'react-markdown';

import faUser from "@fortawesome/fontawesome-free-solid/faUser";
import faCalendarAlt from "@fortawesome/fontawesome-free-solid/faCalendarAlt";
import faTag from "@fortawesome/fontawesome-free-solid/faTag";
import fontawesome from "@fortawesome/fontawesome/index";
import {itemEditable, updateItem} from "../../actions/articles";
import { toastInProgress, toastSuccess, toastFail } from '../shared/Toast';

fontawesome.library.add(faUser, faCalendarAlt, faTag);

export class ArticleEditPage extends Component {

    constructor(props) {
        super(props);
        this.state = {
            author: "",
            title: "",
            subtitle: "",
            body: "",
            tags: "",
            date: "",
            submitting: false,
            errorMsg: ""
        };
    }

    async componentDidMount() {
        window.scrollTo(0, 0);

        if(this.props.itemEditable) {
            try {
                const response = await this.props.itemEditable(this.props.match.params.id);
                this.setState({
                    author: response["author"],
                    title: response["title"],
                    subtitle: response["subtitle"] || "",
                    body: response["body"],
                    tags: response["tags"] || "",
                    date: response["updated_at"]
                });
            } catch (error) {
                let err = error.toString();
                if(error.response && error.response.data["message"].length > 0) {
                    err = error.response.data["message"];
                }
                this.setState({
                    errorMsg: err
                })
            }
        }
    }

    displayInfo = (enableTagLinks) => {
        return this.state.author && (
            <div className="pt-3 pb-5">
                <FontAwesomeIcon className="mr-2" icon="user"/>
                <span className="mr-5">{this.state.author}</span>

                <FontAwesomeIcon className="mr-2" icon="calendar-alt"/>
                <span className="mr-5">{ moment(this.state.date).format('MM-DD-YYYY HH:mm') }</span>

                { this.state.tags && this.state.tags.length > 0 && (
                    <Fragment>
                        <FontAwesomeIcon className="mr-2" icon="tag"/>
                        {
                            split(this.state.tags, ",").map((tag) => {
                                return enableTagLinks ? (
                                    <Link key={tag} className="mr-2" to={`/?tag=${tag}`}>{tag}</Link>
                                ) : (
                                    <span key={tag} className="mr-2">{tag}</span>
                                )
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
                <ReactMarkdown source={ this.state.body } />
            </div>
        )
    }

    onInputChange = (e) => {
        const field = e.target.name;
        const value = e.target.value;
        this.setState(() => ({ [field]: value }));
    };

    onSubmitEditArticle = async (e) => {
        e && e.preventDefault();

        const toastID = toastInProgress("Saving in progress...");
        try {
            this.setState({ submitting: true, errorMsg: "" });
            const {author, title, subtitle, body, tags} = this.state;
            await this.props.updateItem({id: this.props.match.params.id, article: { author, title, subtitle, body, tags: split(tags, ",") }});
            toastSuccess("Success!", toastID);
            this.setState({ submitting: false, errorMsg: "" });
        } catch (error) {
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
            <Fragment>
                <Helmet>
                    <title>Edit an article - FamousTitle.com</title>
                </Helmet>

                <div className="container">

                    <Form onSubmit={this.onSubmitEditArticle} autoComplete="off">
                        <div className="row">
                            <div className="col-md">
                                <div className="mt-5">
                                    <FormGroup>
                                        <Input type="text"
                                               name="title"
                                               value={this.state.title}
                                               placeholder="Title"
                                               onChange={this.onInputChange} />
                                    </FormGroup>
                                    <FormGroup>
                                        <Input type="text"
                                               name="subtitle"
                                               value={this.state.subtitle}
                                               placeholder="Subtitle"
                                               onChange={this.onInputChange} />
                                    </FormGroup>
                                    <FormGroup>
                                        <Input type="text"
                                               name="tags"
                                               value={this.state.tags}
                                               placeholder="Tag"
                                               onChange={this.onInputChange} />
                                    </FormGroup>
                                    <FormGroup>
                                        <Input type="textarea" rows="20"
                                               name="body"
                                               value={this.state.body}
                                               placeholder="Add your article"
                                               onChange={this.onInputChange} />
                                    </FormGroup>

                                    { this.state.errorMsg && <p>{this.state.errorMsg}</p> }

                                </div>
                            </div>
                            <div className="col-md">
                                <div className="mt-5">
                                    { this.state.errorMsg === "" && this.state.author === "" && <Spinner /> }
                                    { this.state.author && (
                                        <Fragment>
                                            <h3>{this.state.title}</h3>
                                            <p>{this.state.subtitle}</p>
                                            { this.displayInfo(false) }
                                            { this.displayBody() }
                                            <div className="clearfix">
                                                <Button color="primary float-right" disabled={this.state.submitting} size="lg">Save</Button>
                                            </div>
                                        </Fragment>
                                    ) }

                                    { this.state.errorMsg && <p>{this.state.errorMsg}</p> }
                                </div>
                            </div>
                        </div>
                    </Form>
                </div>
            </Fragment>
        );
    }
}

const mapDispatchToProps = (dispatch) => ({
    itemEditable: async (data) => await dispatch(itemEditable(data)),
    updateItem: async (data) => await dispatch(updateItem(data))
});

export default connect(null, mapDispatchToProps)(ArticleEditPage);
