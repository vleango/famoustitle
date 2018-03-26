import { createStore, combineReducers, applyMiddleware, compose } from 'redux';
import thunk from 'redux-thunk';
import { persistStore, persistReducer } from 'redux-persist';
import storage from 'redux-persist/lib/storage';

import ArticleReducer from '../reducers/articles';
import AuthReducer from '../reducers/auth';

const composeEnhancers = window.__REDUX_DEVTOOLS_EXTENSION__COMPOSE__ || compose;

const persistConfig = {
  key: 'auth',
  storage,
}

const persistedReducer = persistReducer(persistConfig, AuthReducer);

export default () => {
  const store = createStore(
    combineReducers({
      articles: ArticleReducer,
      auth: persistedReducer
    }),
    composeEnhancers(
      applyMiddleware(thunk)
    )
  );

  let persistor = persistStore(store);

  return { store, persistor }
};
