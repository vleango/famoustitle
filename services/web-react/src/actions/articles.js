import axios from 'axios';

import { ROOT_API_URL } from './Base';

export const fetchList = (filters = {}) => {
    return async (dispatch, getState) => {
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
        }
        catch (error) {
            console.log(error);
        }
    }
};

export const fetchArticlesArchiveList = () => {
    return async (dispatch, getState) => {
        try {
            const response = await axios.get(`${ROOT_API_URL}/articles/archives`);
            let data = { ...response.data };
            dispatch(archives(data));
        } catch (error) {
            console.log(error);
        }
    }
};

export const createItem = (data) => {
    return async (dispatch, getState) => {
        return new Promise(async(resolve, reject) => {
            try {
                let token = "";
                if(getState().auth && getState().auth.token) {
                    token = getState().auth.token;
                }

                const response = await axios.post(`${ROOT_API_URL}/articles`, data, {
                    headers: {"Authorization": `Bearer ${token}`}
                });
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
            let token = "";
            if(getState().auth && getState().auth.token) {
                token = getState().auth.token;
            }

            await axios.post(`${ROOT_API_URL}/articles/${id}`, data, {
                headers: {"Authorization": `Bearer ${token}`}
            });
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
                let token = "";
                if(getState().auth && getState().auth.token) {
                    token = getState().auth.token;
                }

                const response = await axios.delete(`${ROOT_API_URL}/articles/${id}`, {
                    headers: {"Authorization": `Bearer ${token}`}
                });
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
