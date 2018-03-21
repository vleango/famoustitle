import React from 'react';
import { shallow } from 'enzyme';
import { ArticleListPage } from '../ArticleListPage';
import { articles } from './fixtures/articles';

describe('Components', () => {
  describe('ArticleListPage', () => {
    let wrapper = shallow(<ArticleListPage articles={articles} />);

    it('should correctly render ArticleListPage', () => {
      expect(wrapper).toMatchSnapshot();
    });

  });
});
