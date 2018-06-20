import authReducer from '../auth';
import articleReducer from "../articles";

describe('Reducers', () => {
    describe('Auth', () => {

        describe('Default', () => {
            it('should return the default state', () => {
                const action = {
                    type: 'something else'
                };
                const state = authReducer(undefined, action);
                expect(state).toEqual({
                    token: null,
                    firstName: "",
                    lastName: "",
                    email: "",
                    isWriter: false
                });
            });
        });

        describe('AUTH_TOKEN', () => {
            it('returns user info', () => {
                const action = {
                    type: "AUTH_TOKEN",
                    data: {
                        token: "123",
                        firstName: "tha",
                        lastName: "leang",
                        email: "tha@test.com",
                        isWriter: true
                    }
                };

                const state = authReducer(undefined, action);
                expect(state).toEqual({
                    token: "123",
                    firstName: "tha",
                    lastName: "leang",
                    email: "tha@test.com",
                    isWriter: true
                });
            });
        });

        describe('AUTH_LOGOUT', () => {
            it('should return the default state', () => {
                const action = {
                    type: 'AUTH_LOGOUT'
                };
                const state = authReducer(undefined, action);
                expect(state).toEqual({
                    token: null,
                    firstName: "",
                    lastName: "",
                    email: "",
                    isWriter: false
                });
            });
        })
    });
});
