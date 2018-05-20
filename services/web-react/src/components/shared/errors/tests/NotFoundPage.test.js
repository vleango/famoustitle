import React from 'react';
import { shallow } from 'enzyme';

import NotFoundPage from '../NotFoundPage';

describe('Shared', () => {
    describe('Errors', () => {
        describe('NotFoundPage', () => {
            let wrapper = shallow(<NotFoundPage />);

            it('should correctly render NotFoundPage', () => {
                expect(wrapper).toMatchSnapshot();
            });
        });

    });
});
