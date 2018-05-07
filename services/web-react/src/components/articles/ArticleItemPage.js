import React, { Component, Fragment } from 'react';
import { connect } from 'react-redux';
import { Button } from 'reactstrap';
import { Form, FormGroup, Input } from 'reactstrap';
import moment from 'moment';
import {includes, map, pull, split, trim, uniq, compact} from 'lodash';
import ReactMarkdown from 'react-markdown';
import FontAwesomeIcon from '@fortawesome/react-fontawesome';

import { fetchItem, updateItem, removeItem } from '../../actions/articles';
import Header from '../shared/headers/Header';

import './css/ArticleItemPage.css';

export class ArticleItemPage extends Component {

  constructor(props) {
    super(props);
    this.state = {
      article: null,
      title: '',
      body: '',
      tags: [],
      editMode: [],
      previewMode: [],
      editModeClass: '',
      submitting: false
    };
  }

  componentDidMount() {
    this.props.fetchItem && this.props.fetchItem(this.props.match.params.id);
  }

  componentWillReceiveProps = (nextProps) => {
    if(nextProps.article) {
      this.setState({
        article: nextProps.article,
        title: nextProps.article.title,
        body: nextProps.article.body,
        tags: nextProps.article.tags
      });
    }
  };

  onMouseOver = () => {
    if(!this.props.isAuthenticated) {
      return
    }

    this.setState({ editModeClass: 'outline' })
  };

  onMouseLeave = () => {
    if(!this.props.isAuthenticated) {
      return
    }

    this.setState({ editModeClass: '' })
  };

  onTextClicked = (e) => {
    if(!this.props.isAuthenticated) {
      return
    }

    let editMode = this.state.editMode;
    editMode.push(e.currentTarget.dataset.name);
    this.setState({ editMode: uniq(editMode) });
  };

  onCancelClicked = (e) => {
    const field = e.target.dataset.name;
    let inputs = this.state.editMode;
    pull(inputs, field); // remove cancel mode for input
    const value = this.state.article[field]; // get value to revert text
    this.setState({
      editMode: inputs,
      [field]: value
    });
  };

  onPreviewClicked = (e) => {
    let previewMode = this.state.previewMode;
    previewMode.push(e.target.dataset.name);
    this.setState({ previewMode: uniq(previewMode) });
  };

  onPreviewTextClicked = (e) => {
    let previewMode = this.state.previewMode;
    pull(previewMode, e.currentTarget.dataset.name);
    this.setState({ previewMode: previewMode });
  };

  onPreviewExitClicked = (e) => {
    this.onPreviewTextClicked(e);
  };

  onSavedClicked = (e) => {
    const field = e.target.dataset.name;
    let inputs = this.state.editMode;

    // check if empty (can't be empty)
    const value = this.state[field];
    if(value === "" && field !== 'tags') {
      return;
    }

    pull(inputs, field); // remove cancel mode for input

    // save to backend
    const rawTags = split(this.state.tags, ',');
    const trimmedTags = map(rawTags, (tag) => { return trim(tag).toLowerCase() });
    let tags = uniq(compact(trimmedTags));

    const { title, body } = this.state;
    const article = { title, body, tags };
    this.props.updateItem(this.props.match.params.id, { article: article } );

    this.setState({
      editMode: inputs,
      article: article
    });

  };

  onRemoveClicked = async (e) => {
    try {
      this.setState({ submitting: true });
      await this.props.removeItem(this.props.match.params.id);
      this.props.history.push('/');
    }
    catch (e) {
      this.setState({ submitting: false });
    }
  };

  onInputChange = (e) => {
    const field = e.target.name;
    const value = e.target.value;
    this.setState(() => ({ [field]: value }));
  };

  onSubmitChanges = (e) => {
    e.preventDefault();
    // TODO need to be able to save after clicking enter
  };

  displayTags = () => {
      if(includes(this.state.editMode, 'tags')) {
          return (
              <Form onSubmit={this.onSubmitChanges} autoComplete="off">
                  <FormGroup>
                      <Input type="text"
                             name="tags"
                             value={this.state.tags || ""}
                             placeholder="Add Tags"
                             onChange={this.onInputChange} />
                  </FormGroup>
                  <Button color="info" size="sm" data-name="tags" onClick={this.onCancelClicked}>Cancel</Button>
                  <Button color="primary" size="sm" data-name="tags" className="ml-1" onClick={this.onSavedClicked}>Save</Button>
              </Form>
          )
      } else {
          return (
              <div
                   data-name="tags"
                   className={this.state.editModeClass}
                   onMouseOver={this.onMouseOver}
                   onClick={this.onTextClicked}
                   onMouseLeave={this.onMouseLeave}>
                  <FontAwesomeIcon className="mr-2" icon="tag"/>
                  {this.state.article.tags && this.state.article.tags.map((tag) => {
                      return [
                          <span key={tag} className="mr-2">{tag}</span>
                      ]
                  })}
              </div>
          )
      }
  };

  displayBody = () => {
    if(includes(this.state.previewMode, 'body')) {
      return this.displayPreview();
    } else if (includes(this.state.editMode, 'body')) {
      return this.displayForm();
    } else {
      return this.displayMarkdown();
    }
  };

  displayPreview() {
    return (
      <Fragment>
        <Button color="info" size="sm" data-name="body" onClick={this.onPreviewExitClicked}>Exit Preview</Button>
        <div data-name="body" onClick={this.onPreviewTextClicked}><ReactMarkdown source={ this.state.body } /></div>
        <Button color="info" size="sm" data-name="body" onClick={this.onPreviewExitClicked}>Exit Preview</Button>
      </Fragment>
    );
  }

  displayForm() {
    return (
      <Form>
        <FormGroup>
          <Input type="textarea"
            rows="20"
            name="body"
            value={this.state.body}
            placeholder="Add your article"
            onChange={this.onInputChange} />
        </FormGroup>
        <Button color="info" size="sm" data-name="body" onClick={this.onCancelClicked}>Cancel</Button>
        <Button color="info" size="sm" data-name="body" className="ml-1" onClick={this.onPreviewClicked}>Preview</Button>
        <Button color="primary" size="sm" data-name="body" className="ml-1" onClick={this.onSavedClicked}>Save</Button>
      </Form>
    );
  }

  displayMarkdown() {
    return (
      <div
        data-name="body"
        className={`body-markdown ${this.state.editModeClass}`}
        onMouseOver={this.onMouseOver}
        onClick={this.onTextClicked}
        onMouseLeave={this.onMouseLeave}>
        <ReactMarkdown source={ this.state.article.body } />
      </div>
    )
  }

  render() {
    return (
      <div>
        <Header resourceTitle={ this.state.title } />
        { this.state.article && (
          <Fragment>
            <div className="canvas">
              <div className="container">
                { this.props.isAuthenticated && (
                  <div className="clearfix">
                    <Button onClick={this.onRemoveClicked} disabled={this.state.submitting} className="float-right ml-3" color="danger">Delete</Button>{' '}
                  </div>
                ) }

                { includes(this.state.editMode, 'title') ? (
                  <Form onSubmit={this.onSubmitChanges} autoComplete="off">
                    <FormGroup>
                      <Input type="text"
                        name="title"
                        value={this.state.title}
                        placeholder="Title"
                        onChange={this.onInputChange} />
                    </FormGroup>
                    <Button color="info" size="sm" data-name="title" onClick={this.onCancelClicked}>Cancel</Button>
                    <Button color="primary" size="sm" data-name="title" className="ml-1" onClick={this.onSavedClicked}>Save</Button>
                  </Form>
                ) : (
                  <h1
                    data-name="title"
                    className={this.state.editModeClass}
                    onMouseOver={this.onMouseOver}
                    onClick={this.onTextClicked}
                    onMouseLeave={this.onMouseLeave}>
                    {this.state.article.title}
                  </h1>
                )}

                <p>{this.state.article.author}</p>
                <p>{ moment(this.state.article.created_at).format('MM-DD-YYYY HH:mm') }</p>

                { this.displayTags() }

                { this.displayBody() }

              </div>
            </div>
          </Fragment>
        ) }
      </div>
    );
  }
}

const mapStateToProps = (state, props) => {
	return {
	    isAuthenticated: !!state.auth.token,
        article: state.articles.show.resource
	};
};

const mapDispatchToProps = (dispatch) => ({
  fetchItem: (data) => dispatch(fetchItem(data)),
  updateItem: (id, data) => dispatch(updateItem(id, data)),
  removeItem: async (id) => await dispatch(removeItem(id))
});

export default connect(mapStateToProps, mapDispatchToProps)(ArticleItemPage);
