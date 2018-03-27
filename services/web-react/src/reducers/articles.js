const articlesReducerDefaultState = {
	index: { resources: [] },
  show:  { resource: null }
};

export default (state = articlesReducerDefaultState, action) => {
  switch(action.type) {
    case 'ARTICLE_LIST':
      return {
        ...state,
        index: { resources: action.data.articles },
        show: { resource: null }
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
