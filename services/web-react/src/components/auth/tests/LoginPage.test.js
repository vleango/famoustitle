import React from 'react';
import { shallow } from 'enzyme';

import { LoginPage } from '../LoginPage';

import { toastInProgress, toastSuccess, toastFail } from '../../shared/Toast';

jest.mock('../../shared/Toast', () => ({
    toastInProgress: jest.fn((message, id) => { return 1 }),
    toastSuccess: jest.fn((message, id) => { return 2 }),
    toastFail: jest.fn((message, id) => { return 3 })
}));

beforeEach(() => {
    jest.clearAllMocks();
});

describe('Components', () => {
    describe('Auth', () => {
        describe('LoginPage', () => {

            let wrapper = shallow(<LoginPage />);

            describe('Snapshot', () => {
                it('should correctly render LoginPage', () => {
                    expect(wrapper).toMatchSnapshot();
                });
            });

            describe('onSubmitLogin', () => {
                describe('email is blank', () => {
                    it('sets errorMsg', () => {
                        wrapper.setState({ email: '', password: 'hogehoge' });
                        wrapper.instance().onSubmitLogin();
                        expect(wrapper.state('errorMsg')).not.toBe("");
                    });
                });

                describe('password is blank', () => {
                    wrapper.setState({ email: 'test@test.com', password: '' });
                    wrapper.instance().onSubmitLogin();
                    expect(wrapper.state('errorMsg')).not.toBe("");
                });

                describe('email and password are given', async () => {
                    it('Success', async () => {
                        const startLogin = jest.fn(() => Promise.resolve());
                        let wrapper = shallow(<LoginPage startLogin={startLogin} history={[]} />);
                        wrapper.setState({ email: 'test@test.com', password: 'hogehoge' });
                        await wrapper.instance().onSubmitLogin();
                        expect(wrapper.state('submitting')).toBe(true);
                        expect(toastInProgress).toHaveBeenCalledWith("Logging in...");
                        expect(toastSuccess).toHaveBeenCalledWith("Success!", 1);
                        expect(toastFail).not.toHaveBeenCalled();
                    });

                    it('Fail', async () => {
                        const startLogin = jest.fn(() => Promise.reject());
                        let wrapper = shallow(<LoginPage startLogin={startLogin} history={[]} />);
                        wrapper.setState({ email: 'test@test.com', password: 'hogehoge' });
                        await wrapper.instance().onSubmitLogin();
                        expect(wrapper.state('submitting')).toBe(false);
                        expect(toastInProgress).toHaveBeenCalledWith("Logging in...");
                        expect(toastSuccess).not.toHaveBeenCalled();
                        expect(toastFail).toHaveBeenCalledWith("email and/or password was incorrect", 1);
                    });
                });
            });

        });

    });
});
