import React from 'react';
import { Router, Switch } from 'react-router-dom';
import createHistory from 'history/createBrowserHistory';

import HomePage from '../components/homepage/HomePage';
import ArticleItemPage from '../components/articles/ArticleItemPage';
import ArticleNewPage from '../components/articles/ArticleNewPage';
import ArticleEditPage from '../components/articles/ArticleEditPage';

import LoginPage from '../components/auth/LoginPage';
import RegisterPage from '../components/auth/RegisterPage';

import PublicRoute from './PublicRoute';
import PrivateRoute from './PrivateRoute';
import NotFoundPage from '../components/shared/errors/NotFoundPage';

export const history = createHistory();

export const AppRouter = () => (
    <Router history={history}>
        <Switch>
            <PublicRoute path="/" component={HomePage} exact={true} />
            <PublicRoute path="/articles" component={HomePage} exact={true} />
            <PrivateRoute path="/articles/new" component={ArticleNewPage} exact={true} />
            <PrivateRoute path="/articles/:id/edit" component={ArticleEditPage} />
            <PublicRoute path="/articles/:id" component={ArticleItemPage} />
            <PublicRoute path="/login" component={LoginPage} />
            <PublicRoute path="/register" component={RegisterPage} />
            <PublicRoute component={NotFoundPage} />
        </Switch>
    </Router>
);

export default AppRouter;
