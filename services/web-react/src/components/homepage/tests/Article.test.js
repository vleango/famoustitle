import React from 'react';
import { shallow } from 'enzyme';

import Article from '../Article';
import { article } from '../../../fixtures/articles';

describe('Components', () => {
    describe('HomePage', () => {
        describe('Article', () => {

            let wrapper = shallow(<Article article={article} />);

            describe('Snapshot', () => {
                it('should correctly render Article', () => {
                    expect(wrapper).toMatchSnapshot();
                });
            });

            describe('article does not contain tags', () => {
                let art = article;
                art.tags = null;
                let wrapper = shallow(<Article article={art} />);
                expect(wrapper).toMatchSnapshot();
            });

        });
    });
});
