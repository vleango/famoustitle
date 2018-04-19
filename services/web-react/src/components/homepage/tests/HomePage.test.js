import React from 'react';
import { shallow } from 'enzyme';

import { HomePage } from '../HomePage';

describe('Components', () => {
  describe('HomePage', () => {
    describe('HomePage', () => {

      let wrapper = shallow(<HomePage />);

      describe('Snapshot', () => {
        it('should correctly render HomePage', () => {
          expect(wrapper).toMatchSnapshot();
        });
      });
    });
  });
});
