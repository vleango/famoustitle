import React, {Fragment} from 'react';
import { connect } from 'react-redux';
import { Route } from 'react-router-dom';
import { NavBar } from "../components/shared/headers/NavBar";

export const PrivateRoute = ({
                                 isAuthenticated,
                                 component: Component,
                                 ...rest
                             }) => (
    <Route {...rest} component={(props) => (
        <Fragment>
            <NavBar />
            <Component {...props} />
        </Fragment>
    )} />
);

const mapStateToProps = (state) => ({
    isAuthenticated: !!state.auth.token
});

export default connect(mapStateToProps)(PrivateRoute);
