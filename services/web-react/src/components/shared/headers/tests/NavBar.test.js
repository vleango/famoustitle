import React from 'react';
import { shallow } from 'enzyme';
import { NavBar } from '../NavBar';

let wrapper;

describe('Shared', () => {
    describe('NavBar', () => {
        wrapper = shallow(<NavBar />);

        it('should correctly render NavBar', () => {
            expect(wrapper).toMatchSnapshot();
        });
    });
});
