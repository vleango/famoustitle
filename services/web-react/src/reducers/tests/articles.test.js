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

        describe('ARTICLES_ARCHIVE_LIST', () => {
            it('returns index', () => {
                const action = {
                    type: "ARTICLES_ARCHIVE_LIST",
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
    });
});
