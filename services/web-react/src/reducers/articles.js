import { forEach } from 'lodash';

const articlesReducerDefaultState = {
    index: {
        pagination: {
            currentPage: 0,
            totalPages: 0
        },
        resources: [],
        archives: {},
        tags: [],
        selected: {}
    },
    show:  {
        resource: null,
        editable_id: null
    },
    edit:  { resource: null }
};

export default (state = articlesReducerDefaultState, action) => {
    switch(action.type) {
        case 'ARTICLE_LIST':
            let tags = action.data.tags.buckets.map((bucket) => {
                return bucket["key"];
            });
            return {
                ...state,
                index: {
                    ...state.index,
                    resources: action.data.articles,
                    tags: tags,
                    selected: action.data.selected
                }
            };
        case 'ARCHIVE_ARTICLES_LIST':
            let archives = {};
            forEach(action.data.archives.buckets, (bucket) => {
                archives[bucket["key_as_string"]] = bucket["doc_count"];
            });

            return {
                ...state,
                index: {
                    ...state.index,
                    archives: archives
                }
            };
        case 'ARTICLE_ITEM':
            return {
                ...state,
                show: {
                    ...state.show,
                    resource: action.data.article
                }
            };
        case 'ARTICLE_EDITABLE':
            return {
                ...state,
                show: {
                    ...state.show,
                    editable_id: action.data.editable_id
                }
            };
        case 'ARTICLE_UPDATE':
            return {
                ...state
            };
        default:
            return state;
    }
};
