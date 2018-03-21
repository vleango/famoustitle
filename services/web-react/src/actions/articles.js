import axios from 'axios';

import { ROOT_API_URL } from './Base';

export const fetchList = () => {
  return async (dispatch, getState) => {
    try {
      const response = await axios.get(`${ROOT_API_URL}/articles`);
      dispatch(list({articles: response.data.articles}));
    }
    catch (error) {
      console.log(error);
    }
  }
};

export const fetchItem = (id) => {
  return async (dispatch, getState) => {
    try {
      const response = await axios.get(`${ROOT_API_URL}/articles/${id}`);
      dispatch(item({article: response.data}));
    }
    catch (error) {
      console.log(error);
    }
  }
}

export const list = (data) => ({
  type: 'ARTICLE_LIST',
  data: { articles: data.articles }
});

export const item = (data) => ({
  type: 'ARTICLE_ITEM',
  data: data
});
