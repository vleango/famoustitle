import React, { Component, Fragment } from 'react';
import { connect } from 'react-redux';
import { Button } from 'reactstrap';
import { Form, FormGroup, Input } from 'reactstrap';

import { createItem } from '../../actions/articles';
import Header from '../shared/headers/Header';

export class ArticleNewPage extends Component {

  constructor(props) {
    super(props);
    this.state = {
      title: '',
      body: '',
      submitting: false
    };
  }

  onSubmitArticle = async (e) => {
    e.preventDefault();
    try {
      this.setState({ submitting: true });
      await this.props.createItem({ article: { title: this.state.title, body: this.state.body }});
      this.props.history.push('/');
    }
    catch(error) {
      this.setState({ submitting: false });
      console.log(error);
    }
  }

  onInputChange = (e) => {
    const field = e.target.name;
    const value = e.target.value;
    this.setState(() => ({ [field]: value }));
  }

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
              <Input type="textarea" rows="20"
                name="body"
                value={this.state.body}
                placeholder="Add your article"
                onChange={this.onInputChange} />
            </FormGroup>
            <div className="clearfix">
              <Button color="primary float-right" disabled={this.state.submitting} size="lg">Save</Button>
            </div>
          </Form>
        </div>
      </Fragment>
    );
  }
}

const mapDispatchToProps = (dispatch) => ({
  createItem: async (data) => await dispatch(createItem(data))
});

export default connect(null, mapDispatchToProps)(ArticleNewPage);
