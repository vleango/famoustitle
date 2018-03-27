import React from 'react';
import { shallow } from 'enzyme';

import { LoginPage } from '../LoginPage';

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

        describe('email and password are given', () => {
          describe('Success', () => {
            const startLogin = jest.fn(() => Promise.resolve());
            wrapper = shallow(<LoginPage startLogin={startLogin} history={[]} />);
            wrapper.setState({ email: 'test@test.com', password: 'hogehoge' });
            wrapper.instance().onSubmitLogin();
            expect(wrapper.state('submitting')).toBe(true);
          });
        });
      });

    });

  });
});
