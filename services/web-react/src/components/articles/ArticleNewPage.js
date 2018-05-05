import React, { Component, Fragment } from 'react';
import { connect } from 'react-redux';
import { Button } from 'reactstrap';
import { Form, FormGroup, Input } from 'reactstrap';
import { split, map, trim, uniq } from 'lodash';

import { createItem } from '../../actions/articles';
import Header from '../shared/headers/Header';

export class ArticleNewPage extends Component {

  constructor(props) {
    super(props);
    this.state = {
      title: "",
      body: "",
      tags: "",
      submitting: false,
      errorMsg: "",
      token: props.token
    };
  }

  onSubmitArticle = async (e) => {
    e && e.preventDefault();
    const { title, body } = this.state;
    if(title === "" || body === "") {
      this.setState({ errorMsg: "title or body is blank" });
      return;
    }

    try {
      this.setState({ submitting: true, errorMsg: "" });
      let tags = split(this.state.tags, ',');
      let trimmedTags = map(tags, (tag) => { return trim(tag).toLowerCase() });
      await this.props.createItem({ token: this.state.token, article: { title: this.state.title, body: this.state.body, tags: uniq(trimmedTags) }});
      this.props.history.push('/');
    }
    catch(error) {
      this.setState({ submitting: false });
      console.log(error);
    }
  };

  onInputChange = (e) => {
    const field = e.target.name;
    const value = e.target.value;
    this.setState(() => ({ [field]: value }));
  };

  render() {
    return (
      <Fragment>
        <Header resourceTitle="Create New Article" />
        <div className="container">
          <Form onSubmit={this.onSubmitArticle} autoComplete="off">
            <FormGroup>
              <Input type="text"
                name="title"
                value={this.state.title}
                placeholder="Title"
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

            <div className="clearfix">
              <Button color="primary float-right" disabled={this.state.submitting} size="lg">Save</Button>
            </div>
          </Form>
        </div>
      </Fragment>
    );
  }
}

const mapStateToProps = (state) => {
	return {
    token: state.auth.token
	};
};

const mapDispatchToProps = (dispatch) => ({
  createItem: async (data) => await dispatch(createItem(data))
});

export default connect(mapStateToProps, mapDispatchToProps)(ArticleNewPage);
