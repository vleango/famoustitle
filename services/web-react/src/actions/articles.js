import axios from 'axios';

import { ROOT_API_URL } from './Base';

export const fetchList = () => {
  return async (dispatch, getState) => {
    try {
      const response = await axios.get(`${ROOT_API_URL}/articles`);
      dispatch(list(response.data));
    }
    catch (error) {
      console.log(error);
    }
  }
};

export const createItem = (data) => {
  return async (dispatch, getState) => {
    return new Promise(async(resolve, reject) => {
      try {
        const response = await axios.post(`${ROOT_API_URL}/articles`, data);
        resolve(response.body);
      }
      catch (error) {
        console.log(error);
        reject(error);
      }
    });
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
};

export const updateItem = (id, data) => {
  return async (dispatch, getState) => {
    try {
      await axios.post(`${ROOT_API_URL}/articles/${id}`, data);
      dispatch(item({})); // if success don't need to return anything since the page has already been updated
    }
    catch (error) {
      console.log(error);
    }
  }
};

export const removeItem = (id) => {
  return async (dispatch, getState) => {
    return new Promise(async(resolve, reject) => {
      try {
        const response = await axios.delete(`${ROOT_API_URL}/articles/${id}`);
        resolve(response.body);
      }
      catch (error) {
        console.log(error);
        reject(error);
      }
    });
  }
};

export const list = (data) => ({
  type: 'ARTICLE_LIST',
  data
});

export const item = (data) => ({
  type: 'ARTICLE_ITEM',
  data
});

export const update = (data) => ({
  type: 'ARTICLE_UPDATE',
  data
});
