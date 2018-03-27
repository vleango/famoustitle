import axios from 'axios';
import { ROOT_API_URL } from '../Base';
import { fetchList, createItem, fetchItem, updateItem, removeItem } from '../articles';

// thunk methods
let dispatch, getState;

describe('Actions', () => {
  describe('Articles', () => {
    beforeEach(() => {
      dispatch = jest.fn();
      getState = jest.fn();
    });

    describe('fetchList', () => {
      beforeEach(() => {
        axios.get = jest.fn((url) => Promise.resolve({ data: { articles: 'hi'} }));
      });

      it('should call GET /articles', async () => {
        await fetchList()(dispatch, getState);
        expect(axios.get).toHaveBeenLastCalledWith(`${ROOT_API_URL}/articles`);
        expect(dispatch.mock.calls[0][0]).toEqual({
          type: 'ARTICLE_LIST',
          data: {articles: 'hi'}
        });
      });
    });

    describe('createItem', () => {
      beforeEach(() => {
        axios.post = jest.fn((url) => Promise.resolve({ data: { articles: 'hi'} }));
      });

      it('should call POST /articles', async () => {
        await createItem()(dispatch, getState);
        expect(axios.get).toHaveBeenLastCalledWith(`${ROOT_API_URL}/articles`);
      });
    });

    describe('fetchItem', () => {
      beforeEach(() => {
        axios.get = jest.fn((url) => Promise.resolve({ data: { article: 'art1'} }));
      });

      it('should call GET /articles/:id', async () => {
        const id = 6;
        await fetchItem(id)(dispatch, getState);
        expect(axios.get).toHaveBeenLastCalledWith(`${ROOT_API_URL}/articles/${id}`);
        expect(dispatch.mock.calls[0][0]).toEqual({
          type: 'ARTICLE_ITEM',
          data: {article: {article: 'art1'}}
        });
      });
    });

    describe('updateItem', () => {
      beforeEach(() => {
        axios.post = jest.fn((url) => Promise.resolve({ data: { article: 'art1'} }));
      });

      it('should call POST /articles/:id', async () => {
        const id = 6;
        await updateItem(id)(dispatch, getState);
        expect(axios.get).toHaveBeenLastCalledWith(`${ROOT_API_URL}/articles/${id}`);
        expect(dispatch.mock.calls[0][0]).toEqual({
          type: 'ARTICLE_ITEM',
          data: {}
        });
      });
    });

    describe('removeItem', () => {
      beforeEach(() => {
        axios.delete = jest.fn((url) => Promise.resolve({ data: { article: 'art1'} }));
      });

      it('should call DELETE /articles/:id', async () => {
        const id = 6;
        await removeItem(id)(dispatch, getState);
        expect(axios.get).toHaveBeenLastCalledWith(`${ROOT_API_URL}/articles/${id}`);
      });
    });

  });
});
