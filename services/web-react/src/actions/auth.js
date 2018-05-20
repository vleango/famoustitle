
export const startLogin = (data) => {
    return async (dispatch, getState) => {
        try {
            dispatch(login({first_name: 'Tha', last_name: 'Leang', token: '456'}));
        }
        catch (error) {
            console.log(error);
        }
    }
};

export const startLogout = (data) => {
    return async (dispatch, getState) => {
        try {
            dispatch(logout({}));
        }
        catch (error) {
            console.log(error);
        }
    }
};

export const login = (data) => ({
    type: 'AUTH_LOGIN',
    data: {
        token: data.token,
        firstName: data.first_name,
        lastName: data.last_name
    }
});

export const logout = (data) => ({
    type: 'AUTH_LOGOUT'
});
