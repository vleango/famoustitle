import axios from 'axios';

import { ROOT_API_URL } from './Base';

export const startLogin = (data) => {
  return async (dispatch, getState) => {
    try {
      dispatch(login({token: '456'}));
    }
    catch (error) {
      console.log(error);
    }
  }
};

export const login = (data) => ({
  type: 'AUTH_LOGIN',
  data: { token: data.token }
});
