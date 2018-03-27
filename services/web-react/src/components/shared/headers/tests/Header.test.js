import React from 'react';
import { shallow } from 'enzyme';

import Header from '../Header';

describe('Shared', () => {
  describe('Headers', () => {
    describe('Header', () => {
      let wrapper = shallow(<Header />);

      it('should correctly render Header', () => {
        expect(wrapper).toMatchSnapshot();
      });
    });

  });
});
