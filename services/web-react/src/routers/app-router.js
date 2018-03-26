import React from 'react';
import { Router, Route, Switch } from 'react-router-dom';
import createHistory from 'history/createBrowserHistory';

import ArticleListPage from '../components/articles/ArticleListPage';
import ArticleItemPage from '../components/articles/ArticleItemPage';
import ArticleNewPage from '../components/articles/ArticleNewPage';

import LoginPage from '../components/auth/LoginPage';

import PrivateRoute from './PrivateRoute';
import NotFoundPage from '../components/shared/errors/NotFoundPage';

export const history = createHistory();

export const AppRouter = () => (
  <Router history={history}>
    <div>
      <Switch>
        <Route path="/" component={ArticleListPage} exact={true} />
        <PrivateRoute path="/articles/new" component={ArticleNewPage} exact={true} />
        <Route path="/articles/:id" component={ArticleItemPage} />
        <Route path="/login" component={LoginPage} />
        <Route component={NotFoundPage} />
      </Switch>
    </div>
  </Router>
);

export default AppRouter;
