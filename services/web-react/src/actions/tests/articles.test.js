import axios from 'axios';
import { ROOT_API_URL } from '../Base';
import { fetchList, fetchArticlesArchiveList, createItem, fetchItem, updateItem, removeItem } from '../articles';

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
          data: {articles: 'hi', selected: {}},
        });
      });

      describe('contains tag', () => {
        it('should call GET /articles?tag={tag}', async () => {
            await fetchList({tag: "rails"})(dispatch, getState);
            expect(axios.get).toHaveBeenLastCalledWith(`${ROOT_API_URL}/articles?tag=rails`);
            expect(dispatch.mock.calls[0][0]).toEqual({
                type: 'ARTICLE_LIST',
                data: {articles: 'hi', selected: {"tag": "rails"}},
            });
        });
      });

      describe('contains date', () => {
          it('should call GET /articles?tag={date}', async () => {
              await fetchList({date: "2018-01-01"})(dispatch, getState);
              expect(axios.get).toHaveBeenLastCalledWith(`${ROOT_API_URL}/articles?date=2018-01-01`);
              expect(dispatch.mock.calls[0][0]).toEqual({
                  type: 'ARTICLE_LIST',
                  data: {articles: 'hi', selected: {"date": "2018-01-01"}},
              });
          });
      });

      describe('contains match', () => {
          it('should call GET /articles?tag={match}', async () => {
              await fetchList({match: "web"})(dispatch, getState);
              expect(axios.get).toHaveBeenLastCalledWith(`${ROOT_API_URL}/articles?match=web`);
              expect(dispatch.mock.calls[0][0]).toEqual({
                  type: 'ARTICLE_LIST',
                  data: {articles: 'hi', selected: {"match": "web"}},
              });
          });
      });

    });

      describe('fetchArticlesArchiveList', () => {
          beforeEach(() => {
              axios.get = jest.fn((url) => Promise.resolve({data: {articles: 'hi'}}));
          });

          it('should call GET /articles/archives', async () => {
              await fetchArticlesArchiveList()(dispatch, getState);
              expect(axios.get).toHaveBeenLastCalledWith(`${ROOT_API_URL}/articles/archives`);
              expect(dispatch.mock.calls[0][0]).toEqual({
                  type: 'ARTICLES_ARCHIVE_LIST',
                  data: {articles: 'hi'},
              });
          });
      });

    describe('createItem', () => {
      beforeEach(() => {
        axios.post = jest.fn((url) => Promise.resolve({ data: { articles: 'hi'} }));
      });

      it('should call POST /articles', async () => {
        await createItem({someData: true})(dispatch, getState);
        expect(axios.post).toHaveBeenLastCalledWith(`${ROOT_API_URL}/articles`, {someData: true});
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
