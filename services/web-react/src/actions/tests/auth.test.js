import axios from 'axios';
import {ROOT_API_URL} from '../Base';
import {startLogin, startRegister, startLogout} from '../auth';

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
                axios.post = jest.fn((url) => Promise.resolve({
                    data: {
                        email: 'bob@hope.com',
                        first_name: 'Bob',
                        last_name: 'Hope',
                        token: '123',
                        is_writer: true
                    }
                }));
            });

            it('should call POST /login', async () => {
                await startLogin({email: "email"})(dispatch, getState);
                expect(axios.post).toHaveBeenLastCalledWith(`${ROOT_API_URL}/tokens`, {email: "email"});
                expect(dispatch.mock.calls[0][0]).toEqual({
                    type: 'AUTH_TOKEN',
                    data: {
                        email: 'bob@hope.com',
                        token: '123',
                        firstName: 'Bob',
                        lastName: 'Hope',
                        isWriter: true
                    }
                });
            });
        });

        describe('startRegister', () => {
            beforeEach(() => {
                axios.post = jest.fn((url) => Promise.resolve({
                    data: {
                        email: 'bob@hope.com',
                        first_name: 'Bob',
                        last_name: 'Hope',
                        token: '123',
                        is_writer: true
                    }
                }));
            });

            it("should call POST/users", async () => {
               await startRegister({email: "email"})(dispatch, getState);
               expect(axios.post).toHaveBeenLastCalledWith(`${ROOT_API_URL}/users`, {email: "email"});
               expect(dispatch.mock.calls[0][0]).toEqual({
                   type: 'AUTH_TOKEN',
                   data: {
                       email: 'bob@hope.com',
                       token: '123',
                       firstName: 'Bob',
                       lastName: 'Hope',
                       isWriter: true
                   }
               });
            });
        });

        describe('startLogout', () => {
            describe('Success', () => {
                it('should call DELETE /logout', async () => {
                    await startLogout()(dispatch, getState);
                    expect(dispatch.mock.calls[0][0]).toEqual({
                        type: 'AUTH_LOGOUT'
                    });
                });
            });
        });
    });
});
