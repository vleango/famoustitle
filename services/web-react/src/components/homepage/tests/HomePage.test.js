import React from 'react';
import { shallow } from 'enzyme';

import { HomePage } from '../HomePage';
import { articles } from '../../../fixtures/articles';

let wrapper;

describe('Components', () => {
  describe('HomePage', () => {
    describe('HomePage', () => {

      describe('Snapshot', () => {
        it('should correctly render HomePage', () => {
          wrapper = shallow(<HomePage articles={articles} />);
          expect(wrapper).toMatchSnapshot();
        });
      });
    });
  });
});
