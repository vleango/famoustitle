import React from 'react';
import { shallow } from 'enzyme';

import Header from '../Header';

describe('Components', () => {
  describe('HomePage', () => {
    describe('Header', () => {

      let wrapper = shallow(<Header />);

      describe('Snapshot', () => {
        it('should correctly render Header', () => {
          expect(wrapper).toMatchSnapshot();
        });
      });
    });
  });
});
