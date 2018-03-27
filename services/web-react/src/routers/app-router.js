import React from 'react';
import { Router, Switch } from 'react-router-dom';
import createHistory from 'history/createBrowserHistory';

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
    <div>
      <Switch>
        <PublicRoute path="/" component={ArticleListPage} exact={true} />
        <PrivateRoute path="/articles/new" component={ArticleNewPage} exact={true} />
        <PublicRoute path="/articles/:id" component={ArticleItemPage} />
        <PublicRoute path="/login" component={LoginPage} />
        <PublicRoute component={NotFoundPage} />
      </Switch>
    </div>
  </Router>
);

export default AppRouter;
