import axios from 'axios';

import { ROOT_API_URL } from './Base';

export const fetchList = (filters = {}) => {
    return async (dispatch, getState) => {
        return new Promise(async(resolve, reject) => {
            try {

                let resource = "articles";
                if(filters['tag']) {
                    resource = `articles?tag=${filters['tag']}`
                } else if(filters['date']) {
                    resource = `articles?date=${filters['date']}`
                } else if(filters['match']) {
                    resource = `articles?match=${filters['match']}`
                }

                const response = await axios.get(`${ROOT_API_URL}/${resource}`);
                let data = { ...response.data, selected: filters };
                dispatch(list(data));
                resolve(data);
            }
            catch (error) {
                reject(error);
            }
        });
    }
};

export const fetchArticlesArchiveList = () => {
    return async (dispatch, getState) => {
        return new Promise(async(resolve, reject) => {
            try {
                const response = await axios.get(`${ROOT_API_URL}/articles/archives`);
                let data = { ...response.data };
                dispatch(archives(data));
                resolve(data);
            } catch (error) {
                reject(error);
            }
        });
    }
};

export const createItem = (data) => {
    return async (dispatch, getState) => {
        return new Promise(async(resolve, reject) => {
            try {
                const response = await axios.post(`${ROOT_API_URL}/articles`, data, authHeader(getState));
                resolve(response.body);
            }
            catch (error) {
                reject(error);
            }
        });
    }
};

export const fetchItem = (id) => {
    return async (dispatch, getState) => {
        return new Promise(async(resolve, reject) => {
            try {
                const response = await axios.get(`${ROOT_API_URL}/articles/${id}`);
                resolve(response.body);
                dispatch(item({article: response.data}));
            }
            catch (error) {
                reject(error);
            }
        });
    };
};

export const updateItem = (id, data) => {
    return async (dispatch, getState) => {
        return new Promise(async(resolve, reject) => {
            try {
                const response = await axios.post(`${ROOT_API_URL}/articles/${id}`, data, authHeader(getState));
                resolve(response.body);
            }
            catch (error) {
                console.log(error);
                reject(error);
            }
        });
    }
};

export const removeItem = (id) => {
    return async (dispatch, getState) => {
        return new Promise(async(resolve, reject) => {
            try {
                const response = await axios.delete(`${ROOT_API_URL}/articles/${id}`, authHeader(getState));
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

export const archives = (data) => ({
    type: 'ARTICLES_ARCHIVE_LIST',
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

const authHeader = (state) => {
    let token = "";
    if(state() && state().auth && state().auth.token) {
        token = state().auth.token;
    }
    return {
        headers: {"Authorization": `Bearer ${token}`}
    };
};
