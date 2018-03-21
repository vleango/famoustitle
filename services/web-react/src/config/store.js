import { createStore, combineReducers, applyMiddleware, compose } from 'redux';
import thunk from 'redux-thunk';

import ArticleReducer from '../reducers/articles';

const composeEnhancers = window.__REDUX_DEVTOOLS_EXTENSION__COMPOSE__ || compose;

export default () => {
  const store = createStore(
    combineReducers({
      articles: ArticleReducer
    }),
    composeEnhancers(
      applyMiddleware(thunk)
    )
  );

  return store;
};
