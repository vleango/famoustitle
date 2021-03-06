import React from 'react';
import { shallow } from 'enzyme';
import { ArticleItemPage } from '../ArticleItemPage';

import { toastInProgress, toastSuccess, toastFail } from '../../shared/Toast';

jest.mock('../../shared/Toast', () => ({
    toastInProgress: jest.fn((message, id) => { return 1 }),
    toastSuccess: jest.fn((message, id) => { return 2 }),
    toastFail: jest.fn((message, id) => { return 3 })
}));

beforeEach(() => {
    jest.clearAllMocks();
});

describe('Components', () => {
    describe('ArticleItemPage', () => {

        let wrapper = shallow(<ArticleItemPage />);

        describe('Snapshot', () => {
            it('should correctly render ArticleItemPage', () => {
                expect(wrapper).toMatchSnapshot();
            });
        });

        describe('onDeleteArticle', () => {
            it('Success', async () => {
                const removeItem = jest.fn(() => Promise.resolve());
                let wrapper = shallow(<ArticleItemPage removeItem={removeItem} history={[]} match={{params: {id: 1}}} />);
                wrapper.setState({ author: 'Tha', title: 'my title', subtitle: 'my sub', body: 'my body', tags: 'tag1,tag2' });
                await wrapper.instance().onDeleteArticle();
                expect(wrapper.state('submitting')).toBe(false);
                expect(toastInProgress).toHaveBeenCalledWith("Deleting in progress...");
                expect(toastSuccess).toHaveBeenCalledWith("Success!", 1);
                expect(toastFail).not.toHaveBeenCalled();
            });

            it('Fail', async () => {
                const removeItem = jest.fn(() => Promise.reject());
                let wrapper = shallow(<ArticleItemPage removeItem={removeItem} history={[]} match={{params: {id: 1}}} />);
                wrapper.setState({ author: 'Tha', title: 'my title', subtitle: 'my sub', body: 'my body', tags: 'tag1,tag2' });
                await wrapper.instance().onDeleteArticle();
                expect(wrapper.state('submitting')).toBe(false);
                expect(wrapper.state('errorMsg')).toBe("server error");
                expect(toastInProgress).toHaveBeenCalledWith("Deleting in progress...");
                expect(toastSuccess).not.toHaveBeenCalled();
                expect(toastFail).toHaveBeenCalledWith("server error", 1);
            });
        });
    });
});