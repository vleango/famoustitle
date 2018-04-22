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
      return {
        ...state,
        index: {
          ...state.index,
          resources: action.data.articles,
          archives: action.data.archives,
          tags: action.data.tags
        }
      };
    case 'ARTICLE_ITEM':
      return {
        ...state,
        show: { resource: action.data.article }
      }
    case 'ARTICLE_UPDATE':
      return {
        ...state
      }
    default:
      return state;
  }
};
