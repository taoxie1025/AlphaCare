import {CommonAction, UserState} from '../../../@types/global';
import {isEqual} from 'lodash';
import {actionTypes} from '../actions/user';

const initialState: UserState = {
  isLoggedIn: false,
};

export default (state: UserState = initialState, action: CommonAction<string, any>): any => {
  if (isEqual(state, action.payload)) {
    return state;
  }

  switch (action.type) {
    case actionTypes.FETCH_USER: {
      return action.payload;
    }
    case actionTypes.LOGIN:
    case actionTypes.LOGIN_SUCCESS:
    case actionTypes.LOGIN_FAIL: {
      return action.payload;
    }
    default: {
      return state;
    }
  }
};
