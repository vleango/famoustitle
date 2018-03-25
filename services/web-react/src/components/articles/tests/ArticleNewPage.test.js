import React from 'react';
import { shallow } from 'enzyme';
import { ArticleNewPage } from '../ArticleNewPage';

import { createMemoryHistory } from 'history'

describe('Components', () => {
  describe('ArticleNewPage', () => {

    let wrapper = shallow(<ArticleNewPage />);

    describe('Snapshot', () => {
      it('should correctly render ArticleNewPage', () => {
        expect(wrapper).toMatchSnapshot();
      });
    })

    describe('Submit', () => {
      describe('Validation', () => {
        const errorMsg = "title or body is blank";
        describe('Title is blank', () => {
          it('should not submit', () => {
            const target = { title: '', body: 'hello' };
            wrapper.setState({ ...target });
            wrapper.instance().onSubmitArticle(null);
            expect(wrapper.state('errorMsg')).toBe(errorMsg);
            expect(wrapper.state('submitting')).toBe(false);
          });
        });
        describe('Body is blank', () => {
          it('should not submit', () => {
            const target = { title: 'hello', body: '' };
            wrapper.setState({ ...target });
            wrapper.instance().onSubmitArticle(null);
            expect(wrapper.state('errorMsg')).toBe(errorMsg);
            expect(wrapper.state('submitting')).toBe(false);
          });
        });
        describe('Title and Body are blank', () => {
          it('should not submit', () => {
            const target = { title: '', body: '' };
            wrapper.setState({ ...target });
            wrapper.instance().onSubmitArticle(null);
            expect(wrapper.state('errorMsg')).toBe(errorMsg);
            expect(wrapper.state('submitting')).toBe(false);
          });
        });
        describe('Title and Body are present', () => {
          it('should submit', () => {
            const action = jest.fn();
            let wrapper = shallow(<ArticleNewPage createItem={action} history={[]} />);
            const target = { title: 'hello', body: 'bye' };
            wrapper.setState({ ...target });
            wrapper.instance().onSubmitArticle(null);
            expect(action).toHaveBeenCalled();
          });
        });
      });
    });

  });
});
