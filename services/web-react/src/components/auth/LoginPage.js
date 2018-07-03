import React, { Component, Fragment } from 'react';
import { connect } from 'react-redux';
import { Button } from 'reactstrap';
import { Form, FormGroup, Input } from 'reactstrap';
import { toastInProgress, toastSuccess, toastFail } from '../shared/Toast';

import { startLogin } from '../../actions/auth';

export class LoginPage extends Component {

    constructor(props) {
        super(props);
        this.state = {
            email: "",
            password: "",
            submitting: false,
            errorMsg: ""
        };
    }

    onInputChange = (e) => {
        const field = e.target.name;
        const value = e.target.value;
        this.setState(() => ({ [field]: value }));
    };

    onSubmitLogin = async (e) => {
        e && e.preventDefault();
        const { email, password } = this.state;
        if(email === "" || password === "") {
            this.setState({ errorMsg: "email or password is blank" });
            return;
        }

        const toastID = toastInProgress("Logging in...");

        try {
            this.setState({ submitting: true, errorMsg: "" });
            await this.props.startLogin({ email: this.state.email, password: this.state.password });
            toastSuccess("Success!", toastID);
            this.props.history.push('/');
        }
        catch (error) {
            let msg = "email and/or password was incorrect";
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
                <div className="container">
                    <Form onSubmit={this.onSubmitLogin} autoComplete="off">
                        <FormGroup>
                            <Input type="email"
                                   name="email"
                                   value={this.state.email}
                                   placeholder="Enter your email"
                                   onChange={this.onInputChange} />
                        </FormGroup>
                        <FormGroup>
                            <Input type="password"
                                   name="password"
                                   value={this.state.password}
                                   placeholder="Enter your password"
                                   onChange={this.onInputChange} />
                        </FormGroup>

                        { this.state.errorMsg && <p>{this.state.errorMsg}</p> }

                        <div className="clearfix">
                            <Button color="primary float-right" disabled={this.state.submitting} size="lg">Login</Button>
                        </div>
                    </Form>
                </div>
            </Fragment>
        );
    }

}

const mapDispatchToProps = (dispatch) => ({
    startLogin: async (data) => await dispatch(startLogin(data))
});

export default connect(null, mapDispatchToProps)(LoginPage);
