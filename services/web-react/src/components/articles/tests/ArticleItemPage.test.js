import React from 'react';
import { shallow } from 'enzyme';
import { ArticleItemPage } from '../ArticleItemPage';

describe('Components', () => {
    describe('ArticleItemPage', () => {

        let wrapper = shallow(<ArticleItemPage />);

        describe('Snapshot', () => {
            it('should correctly render ArticleItemPage', () => {
                expect(wrapper).toMatchSnapshot();
            });
        });
    });
});