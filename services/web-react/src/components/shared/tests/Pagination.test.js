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
        expect(wrapper.childAt(1).prop('active')).toBe(false);
      });

      it('should set the previousArrow hidden', () => {
        expect(wrapper.childAt(0).prop('hidden')).toBe(true);
      });
    });

    describe('currentPage < totalPages', () => {
      beforeEach(() => {
        wrapper = shallow(<Pagination currentPage={7} totalPages={10} />);
      });
      it('should set the currentPage active', () => {
        expect(wrapper.childAt(2).prop('active')).toBe(true);
      });

      it('should not set the previousArrow hidden', () => {
        expect(wrapper.childAt(0).prop('hidden')).toBe(false);
      });
    });

    describe('currentPage > totalPages', () => {
      beforeEach(() => {
        wrapper = shallow(<Pagination currentPage={20} totalPages={10} />);
      });
      it('should not render pagination', () => {
        expect(wrapper.html()).toBe("");
      });
    });

    describe('currentPage == totalPages', () => {
      beforeEach(() => {
        wrapper = shallow(<Pagination currentPage={10} totalPages={10} />);
      });
      it('should set the currentPage active', () => {
        expect(wrapper.childAt(wrapper.children().length - 2).prop('active')).toBe(true);
      });

      it('should not set the previousArrow hidden', () => {
        expect(wrapper.childAt(0).prop('hidden')).toBe(false);
      });

      it('should set the nextArrow hidden', () => {
        expect(wrapper.childAt(wrapper.children().length - 1).prop('hidden')).toBe(true);
      });
    });

  });

});
