import React from 'react';
import { Router, Route } from 'react-router-dom';
import { shallow } from 'enzyme';
import { Input } from 'reactstrap';
import ReactMarkdown from 'react-markdown';

import { ArticleItemPage } from '../ArticleItemPage';
import { article } from './fixtures/articles';

let wrapper;

describe('Components', () => {
  describe('ArticleItemPage', () => {

    beforeEach(() => {
      wrapper = shallow(<ArticleItemPage article={article} />);
    });

    describe('Snapshot', () => {
      it('should correctly render ArticleItemPage', () => {
        expect(wrapper).toMatchSnapshot();
      });
    });

    describe('isAuthenticated', () => {
      beforeEach(() => {
        wrapper = shallow(<ArticleItemPage article={article} isAuthenticated={true} />);
      });

      describe('EditMode', () => {
        describe('contains title', () => {
          it('shows title form', () => {
            const target = { article: article, editMode: ['title'] };
            wrapper.setState({ ...target });
            expect(wrapper.containsMatchingElement(<Input name="title"></Input>)).toBeTruthy();
          });
        });

        describe('contains body', () => {
          it('shows body form', () => {
            const target = { article: article, editMode: ['body'] };
            wrapper.setState({ ...target });
            expect(wrapper.containsMatchingElement(<Input type="textarea" name="body"></Input>)).toBeTruthy();
          });
        });

        describe('does not contain title', () => {
          it('does not show title form', () => {
            const target = { article: article, editMode: [''] };
            wrapper.setState({ ...target });
            expect(wrapper.containsMatchingElement(<Input name="title"></Input>)).toBeFalsy();
          });
        });

        describe('does not contain body', () => {
          it('does not show body form', () => {
            const target = { article: article, editMode: [''] };
            wrapper.setState({ ...target });
            expect(wrapper.containsMatchingElement(<Input type="textarea" name="body"></Input>)).toBeFalsy();
          });
        })
      });

      describe('onMouseOver', () => {
        it('sets the editModeClass', () => {
          wrapper.instance().onMouseOver();
          expect(wrapper.state('editModeClass')).toBe('outline');
        });
      });

      describe('onMouseLeave', () => {
        it('sets the editModeClass', () => {
          wrapper.instance().onMouseLeave();
          expect(wrapper.state('editModeClass')).toBe('');
        });
      });

      describe('onTextClicked', () => {
        it('sets editMode for the name in e', () => {
          const e = { currentTarget: { dataset: { name: 'hello' } }};
          wrapper.instance().onTextClicked(e);
          expect(wrapper.state('editMode')).toEqual(['hello']);
        })

        describe('uniqueness', () => {
          it('saves the editMode with uniquely', () => {
            wrapper.setState({editMode: ['hello']});
            const e = { currentTarget: { dataset: { name: 'hello' } }};
            wrapper.instance().onTextClicked(e);
            expect(wrapper.state('editMode')).toEqual(['hello']);
          });
        });
      });

      describe('onCancelClicked', () => {
        describe('editMode', () => {
          it('removes the field name from editMode', () => {
            wrapper.setState({editMode: ['title'], article: article});
            const e = { target: { dataset: { name: 'title' } }};
            wrapper.instance().onCancelClicked(e);
            expect(wrapper.state('editMode')).toEqual([]);
            expect(wrapper.state('title')).toEqual(article.title);
          });
        });
      });

      describe('onSavedClicked', () => {
        describe('empty field', () => {
          it('does not call updateItem', () => {
            const updateItem = jest.fn();
            wrapper = shallow(<ArticleItemPage article={article} updateItem={updateItem} />);
            const e = { target: { dataset: { name: 'title' } }};
            wrapper.instance().onSavedClicked(e);
            expect(updateItem).not.toHaveBeenCalled();
          });
        });

        describe('editMode', () => {
          it('removes the field name from editMode, call updateItem, and setState of editMode and article', () => {
            const updateItem = jest.fn();
            const match = {params: {id: 1}};
            wrapper = shallow(<ArticleItemPage article={article} updateItem={updateItem} match={match} />);
            wrapper.setState({editMode: ['title'], article: article, title: 'hello'});
            const e = { target: { dataset: { name: 'title' } }};
            wrapper.instance().onSavedClicked(e);
            expect(wrapper.state('editMode')).toEqual([]);
            expect(updateItem).toHaveBeenCalled();
            expect(wrapper.state('article').title).toEqual('hello');
          });
        });

        describe('onRemoveClicked', () => {
          describe('success', () => {
            it('calls removeItem', async () => {
              const removeItem = jest.fn(() => Promise.resolve());
              const match = {params: {id: 1}};
              wrapper = shallow(<ArticleItemPage article={article} removeItem={removeItem} match={match} history={[]} />);
              await wrapper.instance().onRemoveClicked();
              expect(removeItem).toHaveBeenCalled();
            });
          });

          describe('error', () => {
            it('resets submitting to false', async () => {
              const removeItem = jest.fn(() => Promise.reject());
              const match = {params: {id: 1}};
              wrapper = shallow(<ArticleItemPage article={article} removeItem={removeItem} match={match} history={[]} />);
              await wrapper.instance().onRemoveClicked();
              expect(removeItem).toHaveBeenCalled();
              expect(wrapper.state('submitting')).toBe(false);
            });
          });
        });

        describe('onSubmitChanges', () => {
          it('calls preventDefault on e', () => {
            const e = { preventDefault: jest.fn() }
            wrapper.instance().onSubmitChanges(e);
            expect(e.preventDefault).toHaveBeenCalled();
          });
        });

      });

      describe('onPreviewClicked', () => {
        it('sets previewMode for the element uniquely', () => {
          wrapper.setState({ previewMode: ['body'] });
          const e = { target: { dataset: { name: 'body' } } };
          wrapper.instance().onPreviewClicked(e);
          expect(wrapper.state('previewMode')).toEqual(['body']);
        });
      });

      describe('onPreviewTextClicked', () => {
        it('removes the element from previewMode', () => {
          wrapper.setState({ previewMode: ['body'] });
          const e = { currentTarget: { dataset: { name: 'body' } } };
          wrapper.instance().onPreviewTextClicked(e);
          expect(wrapper.state('previewMode')).toEqual([]);
        });
      });

      describe('onPreviewExitClicked', () => {
        it('removes the element from preview Mode', () => {
          wrapper.setState({ previewMode: ['body'] });
          const e = { currentTarget: { dataset: { name: 'body' } } };
          wrapper.instance().onPreviewTextClicked(e);
          expect(wrapper.state('previewMode')).toEqual([]);
        });
      });

      describe('displayBody', () => {
        describe('previewMode has body', () => {
          it('displays body in previewMode', () => {
            const displaySpy = jest.spyOn(wrapper.instance(), 'displayPreview');
            const displayForm = jest.spyOn(wrapper.instance(), 'displayForm');
            const displayMarkdown = jest.spyOn(wrapper.instance(), 'displayMarkdown');
            wrapper.setState({ previewMode: ['body'], editMode: [] });
            wrapper.instance().displayBody();
            expect(displaySpy).toHaveBeenCalled();
            expect(displayForm).not.toHaveBeenCalled();
            expect(displayMarkdown).not.toHaveBeenCalled();
          });
        });

        describe('previewMode does not have body but editMode does', () => {
          it('displays body in the Form', () => {
            const displaySpy = jest.spyOn(wrapper.instance(), 'displayPreview');
            const displayForm = jest.spyOn(wrapper.instance(), 'displayForm');
            const displayMarkdown = jest.spyOn(wrapper.instance(), 'displayMarkdown');
            wrapper.setState({ previewMode: [], editMode: ['body'] });
            wrapper.instance().displayBody();
            expect(displaySpy).not.toHaveBeenCalled();
            expect(displayForm).toHaveBeenCalled();
            expect(displayMarkdown).not.toHaveBeenCalled();
          });
        });

        describe('previewMode does not have body and editMode does not have body', () => {
          it('displays body in markDown', () => {
            const displaySpy = jest.spyOn(wrapper.instance(), 'displayPreview');
            const displayForm = jest.spyOn(wrapper.instance(), 'displayForm');
            const displayMarkdown = jest.spyOn(wrapper.instance(), 'displayMarkdown');
            wrapper.setState({ article: {body: 'test'}, previewMode: [], editMode: [] });
            wrapper.instance().displayBody();
            expect(displaySpy).not.toHaveBeenCalled();
            expect(displayForm).not.toHaveBeenCalled();
            expect(displayMarkdown).toHaveBeenCalled();
          });
        });
      });

    });

  });
});
