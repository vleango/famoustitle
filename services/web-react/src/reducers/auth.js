const loginReducerDefaultState = {
	token: null
};

export default (state = loginReducerDefaultState, action) => {
  switch(action.type) {
    case 'AUTH_LOGIN':
      return {
        ...state,
        token: action.data.token
      };
    default:
      return state;
  }
};
