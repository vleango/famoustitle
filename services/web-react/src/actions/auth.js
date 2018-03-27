import axios from 'axios';

import { ROOT_API_URL } from './Base';

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

export const login = (data) => ({
  type: 'AUTH_LOGIN',
  data: {
    token: data.token,
    firstName: data.first_name,
    lastName: data.last_name
   }
});
