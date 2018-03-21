const articlesReducerDefaultState = {
	index: [],
  show: null
};

export default (state = articlesReducerDefaultState, action) => {
  switch(action.type) {
    case 'ARTICLE_LIST':
      return {
        ...state,
        index: action.articles
      };
    default:
      return state;
  }
};
