import React from 'react';
import { shallow } from 'enzyme';
import { ArticleItemPage } from '../ArticleItemPage';
import { article } from './fixtures/articles';

describe('Components', () => {
  describe('ArticleItemPage', () => {
    let wrapper = shallow(<ArticleItemPage article={article} />);

    it('should correctly render ArticleItemPage', () => {
      expect(wrapper).toMatchSnapshot();
    });

  });
});
