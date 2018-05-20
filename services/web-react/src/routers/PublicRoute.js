import React, {Fragment} from 'react';
import { connect } from 'react-redux';
import { Route } from 'react-router-dom';

import ProfileHeader from '../components/shared/headers/ProfileHeader';

export const PrivateRoute = ({
                                 isAuthenticated,
                                 component: Component,
                                 ...rest
                             }) => (
    <Route {...rest} component={(props) => (
        <Fragment>
            <ProfileHeader />
            <Component {...props} />
        </Fragment>
    )} />
);

const mapStateToProps = (state) => ({
    isAuthenticated: !!state.auth.token
});

export default connect(mapStateToProps)(PrivateRoute);
