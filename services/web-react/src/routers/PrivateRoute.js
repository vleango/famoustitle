import React, {Fragment} from 'react';
import { connect } from 'react-redux';
import { Route, Redirect } from 'react-router-dom';

import ProfileHeader from '../components/shared/headers/ProfileHeader';

export const PrivateRoute = ({
                                 isAuthenticatedWriter,
                                 component: Component,
                                 ...rest
                             }) => (
    <Route {...rest} component={(props) => (
        isAuthenticatedWriter ? (
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
    isAuthenticatedWriter: !!state.auth.token && state.auth.isWriter
});

export default connect(mapStateToProps)(PrivateRoute);
