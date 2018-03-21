import axios from 'axios';

import { ROOT_API_URL } from './Base';

export const fetchList = () => {
  return async (dispatch, getState) => {
    try {
      const response = await axios.get(`${ROOT_API_URL}/articles`);
      dispatch(list({articles: response.data.articles}));
    }
    catch (error) {
      debugger;
      console.log(error);
    }
  }
};

export const list = (data) => ({
  type: 'ARTICLE_LIST',
  articles: data.articles
});
