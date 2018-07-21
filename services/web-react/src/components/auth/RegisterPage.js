import React, { Component, Fragment } from 'react';
import { connect } from 'react-redux';
import { Button } from 'reactstrap';
import { Form, FormGroup, Input } from 'reactstrap';
import { toastInProgress, toastSuccess, toastFail } from '../shared/Toast';
import {Helmet} from "react-helmet";

import { startRegister } from '../../actions/auth';

export class RegisterPage extends Component {

    constructor(props) {
        super(props);
        this.state = {
            first_name: "",
            last_name: "",
            email: "",
            password: "",
            password_confirmation: "",
            submitting: false,
            errorMsg: ""
        };
    }

    onInputChange = (e) => {
        const field = e.target.name;
        const value = e.target.value;
        this.setState(() => ({ [field]: value }));
    };

    onSubmitRegister = async (e) => {
        e && e.preventDefault();
        const { first_name, last_name, email, password, password_confirmation } = this.state;
        if(first_name === "" || last_name === "" || email === "" || password === "" || password_confirmation === "") {
            this.setState({ errorMsg: "required field is missing" });
            return;
        }

        if(password.length < 6) {
            this.setState({ errorMsg: "password is less than 6 characters" });
            return;
        }

        if(password !== password_confirmation) {
            this.setState({ errorMsg: "passwords does not match" });
            return;
        }

        const toastID = toastInProgress("Creating your account...");

        try {
            this.setState({ submitting: true, errorMsg: "" });
            await this.props.startRegister({
                user: {
                    first_name: this.state.first_name,
                    last_name: this.state.last_name,
                    email: this.state.email
                },
                password: this.state.password,
                password_confirmation: this.state.password_confirmation
            });
            toastSuccess("Success!", toastID);
            this.props.history.push('/');
        }
        catch (error) {
            let msg = "something went wrong...";
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
                    <title>Create an account - FamousTitle.com</title>
                </Helmet>

                <div className="container">
                    <Form onSubmit={this.onSubmitRegister} autoComplete="off">
                        <FormGroup>
                            <Input type="text"
                                   name="first_name"
                                   value={this.state.first_name}
                                   placeholder="First name"
                                   onChange={this.onInputChange} />
                            <Input type="text"
                                   name="last_name"
                                   value={this.state.last_name}
                                   placeholder="Last name"
                                   onChange={this.onInputChange} />
                            <Input type="email"
                                   name="email"
                                   value={this.state.email}
                                   placeholder="Email"
                                   onChange={this.onInputChange} />
                        </FormGroup>
                        <FormGroup>
                            <Input type="password"
                                   name="password"
                                   value={this.state.password}
                                   placeholder="password"
                                   onChange={this.onInputChange} />
                            <Input type="password"
                                   name="password_confirmation"
                                   value={this.state.password_confirmation}
                                   placeholder="password confirmation"
                                   onChange={this.onInputChange} />
                        </FormGroup>

                        { this.state.errorMsg && <p>{this.state.errorMsg}</p> }

                        <div className="clearfix">
                            <Button color="primary float-right" disabled={this.state.submitting} size="lg">Register</Button>
                        </div>
                    </Form>
                </div>
            </Fragment>
        );
    }

}

const mapDispatchToProps = (dispatch) => ({
    startRegister: async (data) => await dispatch(startRegister(data))
});

export default connect(null, mapDispatchToProps)(RegisterPage);
