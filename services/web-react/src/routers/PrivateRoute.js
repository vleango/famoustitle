import React, {Fragment} from 'react';
import { connect } from 'react-redux';
import { Route, Redirect } from 'react-router-dom';
import { NavBar } from "../components/shared/headers/NavBar";

export const PrivateRoute = ({
                                 isAuthenticatedWriter,
                                 component: Component,
                                 ...rest
                             }) => (
    <Route {...rest} component={(props) => (
        isAuthenticatedWriter ? (
            <Fragment>
                <NavBar />
                <Component {...props} />
            </Fragment>
        ) : (
            <Redirect to="/login" />
        )
    )} />
);

const mapStateToProps = (state) => ({
    isAuthenticatedWriter: !!state.auth.token && state.auth.isWriter
});

export default connect(mapStateToProps)(PrivateRoute);
