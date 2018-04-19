import React from 'react';
import { shallow } from 'enzyme';

import Article from '../Article';

describe('Components', () => {
  describe('HomePage', () => {
    describe('Article', () => {

      let wrapper = shallow(<Article />);

      describe('Snapshot', () => {
        it('should correctly render Article', () => {
          expect(wrapper).toMatchSnapshot();
        });
      });
    });
  });
});
