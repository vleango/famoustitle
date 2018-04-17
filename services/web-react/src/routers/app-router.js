import React from 'react';
import { Router, Route, Switch } from 'react-router-dom';
import createHistory from 'history/createBrowserHistory';

import HomePage from '../components/homepage/HomePage';
import ArticleListPage from '../components/articles/ArticleListPage';
import ArticleItemPage from '../components/articles/ArticleItemPage';
import ArticleNewPage from '../components/articles/ArticleNewPage';

import LoginPage from '../components/auth/LoginPage';

import PublicRoute from './PublicRoute';
import PrivateRoute from './PrivateRoute';
import NotFoundPage from '../components/shared/errors/NotFoundPage';

export const history = createHistory();

export const AppRouter = () => (
  <Router history={history}>
    <Switch>
      <Route path="/" component={HomePage} exact={true} />
      <PublicRoute path="/articles" component={ArticleListPage} exact={true} />
      <PrivateRoute path="/articles/new" component={ArticleNewPage} exact={true} />
      <PublicRoute path="/articles/:id" component={ArticleItemPage} />
      <PublicRoute path="/login" component={LoginPage} />
      <PublicRoute component={NotFoundPage} />
    </Switch>
  </Router>
);

export default AppRouter;
