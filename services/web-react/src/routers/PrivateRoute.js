import React from 'react';
import { connect } from 'react-redux';
import { Route, Redirect } from 'react-router-dom';

export const PrivateRoute = ({
  isAuthenticated,
  component: Component,
  ...rest
 }) => (
  <Route {...rest} component={(props) => (
    isAuthenticated ? (
      <div>
        <Component {...props} />
      </div>
    ) : (
      <Redirect to="/login" />
    )
  )} />
);

const mapStateToProps = (state) => ({
  isAuthenticated: !!state.auth.token
});

export default connect(mapStateToProps)(PrivateRoute);
