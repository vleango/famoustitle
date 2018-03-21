import React from 'react';
import { shallow } from 'enzyme';
import { ArticleEditPage } from '../ArticleEditPage';
import { article } from './fixtures/articles';

describe('Components', () => {
  describe('ArticleEditPage', () => {
    let wrapper = shallow(<ArticleEditPage article={article} />);

    it('should correctly render ArticleEditPage', () => {
      expect(wrapper).toMatchSnapshot();
    });

  });
});
