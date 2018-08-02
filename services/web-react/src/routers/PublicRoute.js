import React, {Fragment} from 'react';
import { connect } from 'react-redux';
import { Route } from 'react-router-dom';
import { NavBar } from "../components/shared/headers/NavBar";
import { PublicFooter } from '../components/shared/footers/PublicFooter';

export const PrivateRoute = ({
                                 isAuthenticated,
                                 component: Component,
                                 ...rest
                             }) => (
    <Route {...rest} component={(props) => (
        <Fragment>
            <NavBar />
            <Component {...props} />
            <PublicFooter />
        </Fragment>
    )} />
);

const mapStateToProps = (state) => ({
    isAuthenticated: !!state.auth.token
});

export default connect(mapStateToProps)(PrivateRoute);
