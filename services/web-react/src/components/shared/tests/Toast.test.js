import React from 'react';
import { shallow } from 'enzyme';
import Toast, {toastInProgress, toastSuccess, toastFail} from '../Toast';
import { toast } from 'react-toastify';

jest.mock('react-toastify');


let wrapper;

describe('Shared', () => {
    describe('Toast', () => {
        wrapper = shallow(<Toast />);
        it('should correctly render Toast', () => {
            expect(wrapper).toMatchSnapshot();
        });

        describe('toastInProgress', () => {
            it('calls toast', () => {
                toastInProgress("hello");
                expect(toast).toHaveBeenCalledWith("hello", {autoClose: false});
            });
        });
    });

    describe('toastSuccess', () => {

        beforeEach(() => {
            jest.clearAllMocks();
        });

        describe('calls with toastID', () => {
            it('call update with toastID', () => {
                const spy = jest.spyOn(toast, 'update');
                toastSuccess("hi", 1);
                expect(spy).toHaveBeenCalledWith(1, {autoClose: 3000, render: 'hi', type: 'success'});
            });
        });

        describe('calls without ToastID', () => {
            it('calls success', () => {
                const spy = jest.spyOn(toast, 'success');
                toastSuccess("hi");
                expect(spy).toHaveBeenCalledWith("hi");
            })
        });
    });

    // describe('toastFail', () => {
    //     describe('calls with toastID', () => {
    //         it('call update with toastID', () => {
    //             const spy = jest.spyOn(toast, 'update');
    //             toastFail("hi", 1);
    //             expect(spy).toHaveBeenCalledWith(1, {autoClose: 3000, render: 'hi', type: 'error'});
    //         });
    //     });
    //
    //     describe('calls without ToastID', () => {
    //         it('calls fail', () => {
    //             const spy = jest.spyOn(toast, 'error');
    //             toastFail("hi");
    //             expect(spy).toHaveBeenCalledWith("hi");
    //         })
    //     });
    // });

});
