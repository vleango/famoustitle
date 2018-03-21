import React, { Component } from 'react';
import { connect } from 'react-redux';
import ReactMarkdown from 'react-markdown';
import { Button, Form, FormGroup, Input } from 'reactstrap';

import Header from '../shared/headers/Header';

export class ArticleEditPage extends Component {

  constructor(props) {
    super(props);
    this.state = {
      title: props.article.title,
      body: props.article.body
    };
  }

  onSubmitArticle = (e) => {
  }

  onInputChange = (e) => {
    const field = e.target.name;
    const value = e.target.value;
    this.setState(() => ({ [field]: value }));
  }

  render() {
    return (
      <div>
        <Header resourceTitle={this.props.article.title} />
        <div className="container">
            <Form onSubmit={this.onSubmitArticle}>
              <div className="row">
                <div className="col">
                  <FormGroup>
                    <Input type="text"
                      name="title"
                      value={this.state.title}
                      placeholder="Title"
                      onChange={this.onInputChange} />
                  </FormGroup>
                  <FormGroup>
                    <Input type="textarea"
                      rows="20"
                      name="body"
                      value={this.state.body}
                      placeholder="Add your article"
                      onChange={this.onInputChange} />
                  </FormGroup>
                </div>
                <div className="col">
                  <Button className="float-right" color="primary" size="lg">Save</Button>
                  <h2>{this.state.title}</h2>
                  <hr />
                  <ReactMarkdown source={this.state.body} />
                </div>
              </div>
            </Form>
        </div>
      </div>
    );
  }
}

const mapStateToProps = (state, props) => {
	return {
		article: state.articles.show
	};
};

export default connect(mapStateToProps, null)(ArticleEditPage);
