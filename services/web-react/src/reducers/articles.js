import { forEach } from 'lodash';

const articlesReducerDefaultState = {
	index: {
    pagination: {
      currentPage: 0,
      totalPages: 0
    },
    resources: [],
    archives: {},
    tags: []
  },
  show:  { resource: null }
};

export default (state = articlesReducerDefaultState, action) => {
  switch(action.type) {
    case 'ARTICLE_LIST':

      let tags = action.data.tags.buckets.map((bucket) => {
          return bucket["key"];
      });

      let archives = {};
      forEach(action.data.archives.buckets, (bucket) => {
          archives[bucket["key_as_string"]] = bucket["doc_count"];
      });

      return {
        ...state,
        index: {
          ...state.index,
          resources: action.data.articles,
          archives: archives,
          tags: tags
        }
      };
    case 'ARTICLE_ITEM':
      return {
        ...state,
        show: { resource: action.data.article }
      };
    case 'ARTICLE_UPDATE':
      return {
        ...state
      };
    default:
      return state;
  }
};
