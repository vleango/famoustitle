import React from 'react';
import { shallow } from 'enzyme';
import { ArticleEditPage } from '../ArticleEditPage';

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
    describe('ArticleEditPage', () => {

        let wrapper = shallow(<ArticleEditPage />);

        describe('Snapshot', () => {
            it('should correctly render ArticleEditPage', () => {
                expect(wrapper).toMatchSnapshot();
            });
        });

        describe('onSubmitEditArticle', () => {
            it('Success', async () => {
                const updateItem = jest.fn(() => Promise.resolve());
                let wrapper = shallow(<ArticleEditPage updateItem={updateItem} history={[]} match={{params: {id: 1}}} />);
                wrapper.setState({ author: 'Tha', title: 'my title', subtitle: 'my sub', body: 'my body', tags: 'tag1,tag2' });
                await wrapper.instance().onSubmitEditArticle();
                expect(wrapper.state('submitting')).toBe(false);
                expect(toastInProgress).toHaveBeenCalledWith("Saving in progress...");
                expect(toastSuccess).toHaveBeenCalledWith("Success!", 1);
                expect(toastFail).not.toHaveBeenCalled();
            });

            it('Fail', async () => {
                const updateItem = jest.fn(() => Promise.reject());
                let wrapper = shallow(<ArticleEditPage updateItem={updateItem} history={[]} match={{params: {id: 1}}} />);
                wrapper.setState({ author: 'Tha', title: 'my title', subtitle: 'my sub', body: 'my body', tags: 'tag1,tag2' });
                await wrapper.instance().onSubmitEditArticle();
                expect(wrapper.state('submitting')).toBe(false);
                expect(wrapper.state('errorMsg')).toBe("server error");
                expect(toastInProgress).toHaveBeenCalledWith("Saving in progress...");
                expect(toastSuccess).not.toHaveBeenCalled();
                expect(toastFail).toHaveBeenCalledWith("server error", 1);
            });
        });
    });
});
