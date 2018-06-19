const loginReducerDefaultState = {
    token: null,
    firstName: "",
    lastName: ""
};

export default (state = loginReducerDefaultState, action) => {
    switch(action.type) {
        case 'AUTH_TOKEN':
            return {
                ...state,
                token: action.data.token,
                firstName: action.data.firstName,
                lastName: action.data.lastName,
                email: action.data.email
            };
        case 'AUTH_LOGOUT':
            return loginReducerDefaultState;
        default:
            return state;
    }
};
