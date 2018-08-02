import articleReducer from '../articles';

describe('Reducers', () => {
    describe('Articles', () => {

        describe('Default', () => {
            it('should return the default state', () => {
                const action = {
                    type: 'something else'
                };
                const state = articleReducer(undefined, action);
                expect(state).toEqual({
                    edit: {
                        resource: null
                    },
                    index: {
                        pagination: {
                            currentPage: 0,
                            totalPages: 0
                        },
                        resources: [],
                        selected: {},
                        archives: {},
                        tags: []
                    },
                    show: {
                        editable_id: null,
                        resource: null
                    }
                });

            });
        });

        describe('ARTICLE_LIST', () => {
           it('returns index', () => {
              const action = {
                  type: "ARTICLE_LIST",
                  data: {
                      tags: {
                          buckets: [
                              {key: "123"},
                              {key: "456"}
                          ]
                      },
                      articles: "art",
                      selected: "select"
                  }
              };

               const state = articleReducer(undefined, action);
               expect(state).toEqual({
                   edit: {
                       resource: null
                   },
                   index: {
                       pagination: {
                           currentPage: 0,
                           totalPages: 0
                       },
                       resources: "art",
                       selected: "select",
                       archives: {},
                       tags: ["123", "456"]
                   },
                   show: {
                       editable_id: null,
                       resource: null
                   }
               });
           });
        });

        describe('ARCHIVE_ARTICLES_LIST', () => {
            it('returns index', () => {
                const action = {
                    type: "ARCHIVE_ARTICLES_LIST",
                    data: {
                        archives: {
                            buckets: [
                                {
                                    "doc_count": 1,
                                    "key_as_string": "hello"
                                }
                            ]
                        }
                    }
                };

                const state = articleReducer(undefined, action);
                expect(state).toEqual({
                    edit: {
                        resource: null
                    },
                    index: {
                        pagination: {
                            currentPage: 0,
                            totalPages: 0
                        },
                        archives: {"hello": 1},
                        resources: [],
                        selected: {},
                        tags: []
                    },
                    show: {
                        editable_id: null,
                        resource: null
                    }
                });
            });
        });

        describe('ARTICLE_ITEM', () => {
            it('returns show', () => {
                const action = {
                    type: "ARTICLE_ITEM",
                    data: {
                        article: "art"
                    }
                };

                const state = articleReducer(undefined, action);
                expect(state).toEqual({
                    edit: {
                        resource: null,
                    },
                    index: {
                        pagination: {
                            currentPage: 0,
                            totalPages: 0
                        },
                        resources: [],
                        selected: {},
                        archives: {},
                        tags: []
                    },
                    show: {
                        editable_id: null,
                        resource: "art"
                    }
                });
            });
        });

        describe('ARTICLE_UPDATE', () => {
            it('returns state', () => {
                const action = {
                    type: "ARTICLE_UPDATE"
                };

                const state = articleReducer({"hi": "bye"}, action);
                expect(state).toEqual({
                    "hi": "bye"
                });
            });
        });

        describe('ARTICLE_EDITABLE', () => {
            it('returns state', () => {
                const action = {
                    type: "ARTICLE_EDITABLE",
                    data: {
                        editable_id: "art"
                    }
                };

                const state = articleReducer({"hi": "bye"}, action);
                expect(state).toEqual({
                    "hi": "bye",
                    "show": {
                        "editable_id": "art"
                    }
                });
            });
        });
    });
});
