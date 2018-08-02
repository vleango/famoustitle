import React from 'react';
import { shallow } from 'enzyme';
import { ArticleNewPage } from '../ArticleNewPage';

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
    describe('ArticleNewPage', () => {

        let wrapper = shallow(<ArticleNewPage />);

        describe('Snapshot', () => {
            it('should correctly render ArticleNewPage', () => {
                expect(wrapper).toMatchSnapshot();
            });
        });

        describe('Submit', () => {
            describe('Validation', () => {
                const errorMsg = "title or body is blank";
                describe('Title is blank', () => {
                    it('should not submit', () => {
                        const target = { title: '', body: 'hello' };
                        wrapper.setState({ ...target });
                        wrapper.instance().onSubmitArticle(null);
                        expect(wrapper.state('errorMsg')).toBe(errorMsg);
                        expect(wrapper.state('submitting')).toBe(false);
                    });
                });
                describe('Body is blank', () => {
                    it('should not submit', () => {
                        const target = { title: 'hello', body: '' };
                        wrapper.setState({ ...target });
                        wrapper.instance().onSubmitArticle(null);
                        expect(wrapper.state('errorMsg')).toBe(errorMsg);
                        expect(wrapper.state('submitting')).toBe(false);
                    });
                });
                describe('Title and Body are blank', () => {
                    it('should not submit', () => {
                        const target = { title: '', body: '' };
                        wrapper.setState({ ...target });
                        wrapper.instance().onSubmitArticle(null);
                        expect(wrapper.state('errorMsg')).toBe(errorMsg);
                        expect(wrapper.state('submitting')).toBe(false);
                    });
                });
                describe('Title and Body are present', () => {
                    it('should submit', () => {
                        const action = jest.fn();
                        let wrapper = shallow(<ArticleNewPage createItem={action} history={[]} />);
                        const target = { title: 'hello', body: 'bye' };
                        wrapper.setState({ ...target });
                        wrapper.instance().onSubmitArticle(null);
                        expect(action).toHaveBeenCalled();
                    });
                });
                describe('Tags', () => {
                    describe("submits tags as array", () => {
                        const action = jest.fn();
                        const calledArgs = { "article": {"body": "bye", "subtitle": "", "tags": ["tag1", "tag2"], "title": "hello"}};
                        let wrapper = shallow(<ArticleNewPage createItem={action} history={[]} />);
                        const target = { title: 'hello', body: 'bye', tags: "tag1,tag2" };
                        wrapper.setState({ ...target });
                        wrapper.instance().onSubmitArticle(null);
                        expect(action).toHaveBeenCalledWith(calledArgs);
                    });
                    describe("trims tags", () => {
                        const action = jest.fn();
                        const calledArgs = { "article": {"body": "bye", "subtitle": "", "tags": ["tag1", "tag2"], "title": "hello"}};
                        let wrapper = shallow(<ArticleNewPage createItem={action} history={[]} />);
                        const target = { title: 'hello', body: 'bye', tags: " tag1 , tag2" };
                        wrapper.setState({ ...target });
                        wrapper.instance().onSubmitArticle(null);
                        expect(action).toHaveBeenCalledWith(calledArgs);
                    });
                    describe("lowercase tags", () => {
                        const action = jest.fn();
                        const calledArgs = { "article": {"body": "bye", "subtitle": "", "tags": ["tag1", "tag2"], "title": "hello"}};
                        let wrapper = shallow(<ArticleNewPage createItem={action} history={[]} />);
                        const target = { title: 'hello', body: 'bye', tags: " TAG1,tag2" };
                        wrapper.setState({ ...target });
                        wrapper.instance().onSubmitArticle(null);
                        expect(action).toHaveBeenCalledWith(calledArgs);
                    });
                    describe("submits only unique tags", () => {
                        const action = jest.fn();
                        const calledArgs = { "article": {"body": "bye", "subtitle": "", "tags": ["tag1", "tag2"], "title": "hello"}};
                        let wrapper = shallow(<ArticleNewPage createItem={action} history={[]} />);
                        const target = { title: 'hello', body: 'bye', tags: "tag1,tag2,tag1" };
                        wrapper.setState({ ...target });
                        wrapper.instance().onSubmitArticle(null);
                        expect(action).toHaveBeenCalledWith(calledArgs);
                    });
                });
            });

            describe('title and body are present', () => {
                it('Success', async () => {
                    const createItem = jest.fn(() => Promise.resolve());
                    let wrapper = shallow(<ArticleNewPage createItem={createItem} history={[]} />);
                    wrapper.setState({ title: 'my title', body: 'my body' });
                    await wrapper.instance().onSubmitArticle();
                    expect(wrapper.state('submitting')).toBe(true);
                    expect(toastInProgress).toHaveBeenCalledWith("Saving in progress...");
                    expect(toastSuccess).toHaveBeenCalledWith("Success!", 1);
                    expect(toastFail).not.toHaveBeenCalled();
                });

                it('Fail', async () => {
                    const createItem = jest.fn(() => Promise.reject());
                    let wrapper = shallow(<ArticleNewPage createItem={createItem} history={[]} />);
                    wrapper.setState({ title: 'my title', body: 'my body' });
                    await wrapper.instance().onSubmitArticle();
                    expect(wrapper.state('submitting')).toBe(false);
                    expect(wrapper.state('errorMsg')).toBe("server error");
                    expect(toastInProgress).toHaveBeenCalledWith("Saving in progress...");
                    expect(toastSuccess).not.toHaveBeenCalled();
                    expect(toastFail).toHaveBeenCalledWith("server error", 1);
                });
            });
        });

    });
});
