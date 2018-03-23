import React, { Component, Fragment } from 'react';
import { connect } from 'react-redux';
import { Link } from 'react-router-dom';
import { Button } from 'reactstrap';
import { Form, FormGroup, Input } from 'reactstrap';
import moment from 'moment';
import { includes, pull, uniq } from 'lodash';

import { fetchItem, updateItem } from '../../actions/articles';
import Header from '../shared/headers/Header';

export class ArticleItemPage extends Component {

  constructor(props) {
    super(props);
    this.state = {
      article: null,
      title: '',
      body: '',
      editMode: [],
      editModeClass: ''
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
        body: nextProps.article.body
      });
    }
  }

  onMouseOver = () => {
    this.setState({ editModeClass: 'outline' })
  }

  onMouseLeave = () => {
    this.setState({ editModeClass: '' })
  }

  onTextClicked = (e) => {
    let editMode = this.state.editMode;
    editMode.push(e.target.dataset.name);
    this.setState({ editMode: uniq(editMode) });
  }

  onCancelClicked = (e) => {
    const field = e.target.dataset.name;
    let inputs = this.state.editMode;
    pull(inputs, field); // remove cancel mode for input
    const value = this.state.article[field]; // get value to revert text
    this.setState({
      editMode: inputs,
      [field]: value
    });
  }

  onSavedClicked = (e) => {
    const field = e.target.dataset.name;
    let inputs = this.state.editMode;
    pull(inputs, field); // remove cancel mode for input

    // saved the field to article
    const article = this.state.article;
    article[field] = this.state[field];

    // save to backend
    this.props.updateItem(this.props.match.params.id, { article: { [field]: article[field] }})

    this.setState({
      editMode: inputs,
      article: article
    });

  }

  onInputChange = (e) => {
    const field = e.target.name;
    const value = e.target.value;
    this.setState(() => ({ [field]: value }));
  }

  render() {
    return (
      <div>
        <Header resourceTitle={ this.state.title } />
        { this.state.article && (
          <Fragment>
            <div className="container">
              <div className="clearfix">
                <Button tag={Link} to={`/`} className="float-right ml-3" color="danger">Delete</Button>{' '}
              </div>


              { includes(this.state.editMode, 'title') ? (
                <Form>
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

              { includes(this.state.editMode, 'body') ? (
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
                  <Button color="primary" size="sm" data-name="body" className="ml-1" onClick={this.onSavedClicked}>Save</Button>
                </Form>
              ) : (
                <p
                  data-name="body"
                  className={this.state.editModeClass}
                  onMouseOver={this.onMouseOver}
                  onClick={this.onTextClicked}
                  onMouseLeave={this.onMouseLeave}>
                  { this.state.article.body }
                </p>
              )}

            </div>
          </Fragment>
        ) }
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
  fetchItem: (data) => dispatch(fetchItem(data)),
  updateItem: (id, data) => dispatch(updateItem(id, data))
});

export default connect(mapStateToProps, mapDispatchToProps)(ArticleItemPage);
