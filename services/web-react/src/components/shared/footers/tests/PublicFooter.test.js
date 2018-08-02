import React from 'react';
import { shallow } from 'enzyme';

import { PublicFooter } from '../PublicFooter';

describe('Shared', () => {
    describe('Footers', () => {
        describe('PublicFooter', () => {
            let wrapper = shallow(<PublicFooter />);

            it('should correctly render PublicFooter', () => {
                expect(wrapper).toMatchSnapshot();
            });
        });

    });
});
