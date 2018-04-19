import React from 'react';
import { shallow } from 'enzyme';

import Pagination from '../Pagination';

let wrapper;

describe('Shared', () => {
  describe('Pagination', () => {
    wrapper = shallow(<Pagination />);

    it('should correctly render Pagination', () => {
      expect(wrapper).toMatchSnapshot();
    });

    describe('currentPage == 1', () => {
      beforeEach(() => {
        wrapper = shallow(<Pagination currentPage={1} totalPages={10} />);
      });
      it('should set the currentPage active', () => {
        expect(wrapper.childAt(1).html().includes('active')).toBe(true);
      });

      it('should set the previousArrow disabled', () => {
        expect(wrapper.childAt(0).html().includes('disabled')).toBe(true);
      });
    });

    describe('currentPage < totalPages', () => {
      beforeEach(() => {
        wrapper = shallow(<Pagination currentPage={7} totalPages={10} />);
      });
      it('should set the currentPage active', () => {
        expect(wrapper.childAt(2).html().includes('active')).toBe(true);
      });

      it('should not set the previousArrow disabled', () => {
        expect(wrapper.childAt(0).html().includes('disabled')).not.toBe(true);
      });

      it('should not set the nextArrow disabled', () => {
        expect(wrapper.childAt(wrapper.children().length - 1).html().includes('disabled')).not.toBe(true);
      });

    });

    describe('currentPage > totalPages', () => {
      beforeEach(() => {
        wrapper = shallow(<Pagination currentPage={20} totalPages={10} />);
      });
      it('should not render pagination', () => {
        expect(wrapper.html()).toBe(null);
      });
    });

    describe('currentPage == totalPages', () => {
      beforeEach(() => {
        wrapper = shallow(<Pagination currentPage={10} totalPages={10} />);
      });
      it('should set the currentPage active', () => {
        expect(wrapper.childAt(wrapper.children().length - 2).html().includes('active')).toBe(true);
      });

      it('should not set the previousArrow disabled', () => {
        expect(wrapper.childAt(0).html().includes('disabled')).not.toBe(true);
      });

      it('should set the nextArrow disabled', () => {
        expect(wrapper.childAt(wrapper.children().length - 1).html().includes('disabled')).toBe(true);
      });
    });

  });

});
