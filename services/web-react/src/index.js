import React from 'react';
import ReactDOM from 'react-dom';
import { Provider } from 'react-redux'

import 'bootstrap/dist/css/bootstrap.css';
import './index.css';

import AppRouter from './routers/app-router';
import configureStore from './config/store';

const store = configureStore();

const jsx = (
  <Provider store={store}>
    <AppRouter />
  </Provider>
);

ReactDOM.render(jsx, document.getElementById('root'));
