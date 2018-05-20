import React, {Fragment} from 'react';
import { connect } from 'react-redux';
import { Route, Redirect } from 'react-router-dom';

import ProfileHeader from '../components/shared/headers/ProfileHeader';

export const PrivateRoute = ({
                                 isAuthenticated,
                                 component: Component,
                                 ...rest
                             }) => (
    <Route {...rest} component={(props) => (
        isAuthenticated ? (
            <Fragment>
                <ProfileHeader />
                <Component {...props} />
            </Fragment>
        ) : (
            <Redirect to="/login" />
        )
    )} />
);

const mapStateToProps = (state) => ({
    isAuthenticated: !!state.auth.token
});

export default connect(mapStateToProps)(PrivateRoute);
