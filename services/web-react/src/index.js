import React from 'react';
import ReactDOM from 'react-dom';
import { Provider } from 'react-redux'
import { PersistGate } from 'redux-persist/integration/react';

import 'bootstrap/dist/css/bootstrap.css';
import './index.css';

import AppRouter from './routers/app-router';
import configureStore from './config/store';

const storeConfig = configureStore();

const jsx = (
  <Provider store={storeConfig.store}>
    <PersistGate loading={<p>Loading...</p>} persistor={storeConfig.persistor}>
      <AppRouter />
    </PersistGate>
  </Provider>
);

ReactDOM.render(jsx, document.getElementById('root'));
