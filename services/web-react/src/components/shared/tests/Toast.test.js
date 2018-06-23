import React from 'react';
import { shallow } from 'enzyme';
import Toast from '../Toast';

let wrapper;

describe('Shared', () => {
    describe('Toast', () => {
        wrapper = shallow(<Toast />);
        it('should correctly render Toast', () => {
            expect(wrapper).toMatchSnapshot();
        });
    });

});
