export const actionTypes = {
  FETCH_USER: 'FETCH_USER',
  FETCH_USER_SUCCEEDED: 'FETCH_USER_SUCCEEDED',
  FETCH_USER_FAILED: 'FETCH_USER_FAILED',
  LOGIN: 'LOGIN',
  LOGIN_SUCCESS: 'LOGIN_SUCCESS',
  LOGIN_FAIL: 'LOGIN_FAIL',
  REGISTER: 'REGISTER',
  REGISTER_SUCCESS: 'REGISTER_SUCCESS',
  REGISTER_FAIL: 'REGISTER_FAIL',
};

export const fetchUser = (userId: string) => ({
  type: actionTypes.FETCH_USER,
  payload: userId,
});

export const login = (userId: string, password: string) => ({
  type: actionTypes.LOGIN,
  payload: {userId, password},
});

export const register = (
  userName: string,
  password: string,
  firstName: string,
  lastName: string
) => ({
  type: actionTypes.REGISTER,
  payload: {userName, password, firstName, lastName},
});
