const articlesReducerDefaultState = {
	index: { articles: [] },
  show:  { resource: null }
};

export default (state = articlesReducerDefaultState, action) => {
  switch(action.type) {
    case 'ARTICLE_LIST':
      return {
        ...state,
        index: { articles: action.data.articles },
        show: { resource: null }
      };
    case 'ARTICLE_ITEM':
      return {
        ...state,
        show: { resource: action.data.article }
      }
    default:
      return state;
  }
};
