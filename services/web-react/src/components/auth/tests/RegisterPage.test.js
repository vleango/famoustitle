import React from 'react';
import { shallow } from 'enzyme';

import { RegisterPage } from '../RegisterPage';

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
        describe('RegisterPage', () => {

            let wrapper = shallow(<RegisterPage />);

            describe('Snapshot', () => {
                it('should correctly render RegisterPage', () => {
                    expect(wrapper).toMatchSnapshot();
                });
            });

            describe('onSubmitRegister', () => {
                describe('first_name is blank', () => {
                    it('sets errorMsg', () => {
                        wrapper.setState({ first_name: "", last_name: "Leang", email: "tha@test.com", password: "hogehoge", password_confirmation: "hogehoge" });
                        wrapper.instance().onSubmitRegister();
                        expect(wrapper.state('errorMsg')).toBe("required field is missing");
                    });
                });

                describe('last_name is blank', () => {
                    it('sets errorMsg', () => {
                        wrapper.setState({ first_name: "Tha", last_name: "", email: "tha@test.com", password: "hogehoge", password_confirmation: "hogehoge" });
                        wrapper.instance().onSubmitRegister();
                        expect(wrapper.state('errorMsg')).toBe("required field is missing");
                    });
                });

                describe('email is blank', () => {
                    it('sets errorMsg', () => {
                        wrapper.setState({ first_name: "Tha", last_name: "Leang", email: "", password: "hogehoge", password_confirmation: "hogehoge" });
                        wrapper.instance().onSubmitRegister();
                        expect(wrapper.state('errorMsg')).toBe("required field is missing");
                    });
                });

                describe('password is blank', () => {
                    it('sets errorMsg', () => {
                        wrapper.setState({ first_name: "Tha", last_name: "Leang", email: "tha@test.com", password: "", password_confirmation: "hogehoge" });
                        wrapper.instance().onSubmitRegister();
                        expect(wrapper.state('errorMsg')).toBe("required field is missing");
                    });
                });

                describe('password_confirmation is blank', () => {
                    it('sets errorMsg', () => {
                        wrapper.setState({ first_name: "Tha", last_name: "Leang", email: "tha@test.com", password: "hogehoge", password_confirmation: "" });
                        wrapper.instance().onSubmitRegister();
                        expect(wrapper.state('errorMsg')).toBe("required field is missing");
                    });
                });

                describe('passwords do not match', () => {
                    it('sets errorMsg', () => {
                        wrapper.setState({ first_name: "Tha", last_name: "Leang", email: "tha@test.com", password: "hogehoge", password_confirmation: "piyopiyo" });
                        wrapper.instance().onSubmitRegister();
                        expect(wrapper.state('errorMsg')).toBe("passwords does not match");
                    });
                });

                describe('password less than 6 characters', () => {
                    it('sets errorMsg', () => {
                        wrapper.setState({ first_name: "Tha", last_name: "Leang", email: "tha@test.com", password: "12345", password_confirmation: "12345" });
                        wrapper.instance().onSubmitRegister();
                        expect(wrapper.state('errorMsg')).toBe("password is less than 6 characters");
                    });
                });

                describe('first_name, last_name, email, password, and password_confirmation are given', () => {
                   it('Success', async () => {
                       const startRegister = jest.fn(() => Promise.resolve());
                       let wrapper = shallow(<RegisterPage startRegister={startRegister} history={[]} />);

                       wrapper.setState({ first_name: "Tha", last_name: "Leang", email: "tha@test.com", password: "123456", password_confirmation: "123456" });
                       await wrapper.instance().onSubmitRegister();
                       expect(wrapper.state('submitting')).toBe(true);
                       expect(wrapper.state('errorMsg')).toBe("");
                       expect(toastInProgress).toHaveBeenCalledWith("Creating your account...");
                       expect(toastSuccess).toHaveBeenCalledWith("Success!", 1);
                       expect(toastFail).not.toHaveBeenCalled();
                   });

                   it('Fail', async () => {
                       const startRegister = jest.fn(() => Promise.reject());
                       let wrapper = shallow(<RegisterPage startRegister={startRegister} history={[]} />);

                       wrapper.setState({ first_name: "Tha", last_name: "Leang", email: "tha@test.com", password: "123456", password_confirmation: "123456" });
                       await wrapper.instance().onSubmitRegister();
                       expect(wrapper.state('submitting')).toBe(false);
                       expect(wrapper.state('errorMsg')).toBe("something went wrong...");
                       expect(toastInProgress).toHaveBeenCalledWith("Creating your account...");
                       expect(toastSuccess).not.toHaveBeenCalled();
                       expect(toastFail).toHaveBeenCalledWith("something went wrong...", 1);
                   });
                });
            });

        });

    });
});
