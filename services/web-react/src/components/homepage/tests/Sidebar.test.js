import React from 'react';
import { shallow } from 'enzyme';

import Sidebar from '../Sidebar';

describe('Components', () => {
  describe('HomePage', () => {
    describe('Sidebar', () => {

      let wrapper = shallow(<Sidebar />);

      describe('Snapshot', () => {
        it('should correctly render Sidebar', () => {
          expect(wrapper).toMatchSnapshot();
        });
      });
    });
  });
});
