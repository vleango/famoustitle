// const articlesReducerDefaultState = {
// 	index: {
//     pagination: {
//       currentPage: 1,
//       totalPages: 1
//     },
//     resources: [],
//     archives: {
//       '2018-04-24T14:52:28.254839633Z': 11,
//       '2018-03-24T14:52:28.254839633Z': 40
//     },
//     tags: [
//       'logo',
//       'business',
//       'corporate',
//       'e-commerce',
//       'agency'
//     ]
//   },
//   show:  { resource: null }
// };

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
          resources: action.data.articles
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
