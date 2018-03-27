import axios from 'axios';
import { ROOT_API_URL } from '../Base';
import { startLogin, startLogout } from '../auth';

// thunk methods
let dispatch, getState;

describe('Actions', () => {
  describe('Auth', () => {
    beforeEach(() => {
      dispatch = jest.fn();
      getState = jest.fn();
    });

    describe('startLogin', () => {
      beforeEach(() => {
        axios.post = jest.fn((url) => Promise.resolve({ data: { first_name: 'Bob', last_name: 'Hope', token: '123'} }));
      });

      it('should call POST /login', async () => {
        await startLogin()(dispatch, getState);
        expect(axios.post).toHaveBeenLastCalledWith(`${ROOT_API_URL}/login`);
        expect(dispatch.mock.calls[0][0]).toEqual({
          type: 'AUTH_LOGIN',
          data: {
            token: '123',
            firstName: 'Bob',
            lastName: 'Hope'
          }
        });
      });
    });

    describe('startLogout', () => {
      describe('Success', () => {
        beforeEach(() => {
          axios.delete = jest.fn((url) => Promise.resolve({}));
        });

        it('should call DELETE /logout', async () => {
          await startLogout()(dispatch, getState);
          expect(axios.delete).toHaveBeenLastCalledWith(`${ROOT_API_URL}/logout`);
          expect(dispatch.mock.calls[0][0]).toEqual({
            type: 'AUTH_LOGOUT'
          });
        });
      });

      describe('Error', () => {
        beforeEach(() => {
          axios.delete = jest.fn((url) => Promise.reject({}));
        });

        it('should call DELETE /logout', async () => {
          await startLogout()(dispatch, getState);
          expect(axios.delete).toHaveBeenLastCalledWith(`${ROOT_API_URL}/logout`);
          expect(dispatch.mock.calls[0][0]).toEqual({
            type: 'AUTH_LOGOUT'
          });
        });
      });
    });

  });
});
