import React from 'react';
import { shallow } from 'enzyme';

import { ProfileHeader } from '../ProfileHeader';

describe('Shared', () => {
  describe('Headers', () => {
    describe('ProfileHeader', () => {

      let wrapper = shallow(<ProfileHeader />);

      describe('Snapshot', () => {
        it('should correctly render ProfileHeader', () => {
          expect(wrapper).toMatchSnapshot();
        });
      });

      describe('toggle', () => {
        it('should toggle isOpen', () => {
          wrapper.instance().toggle();
          expect(wrapper.state('isOpen')).toEqual(true);
          wrapper.instance().toggle();
          expect(wrapper.state('isOpen')).toEqual(false);
        });
      });

      describe('onLogout', () => {
        const logout = jest.fn();
        wrapper = shallow(<ProfileHeader startLogout={logout} />);
        wrapper.instance().onLogout();
        expect(logout).toHaveBeenCalled();
      });
    });

  });
});
