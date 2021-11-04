import {call, fork, put, takeLatest} from 'redux-saga/effects';
import {CommonAction, User} from '../../../@types/global';
import {getUser, login} from '../../api/user';
import {actionTypes} from '../actions/user';

function* fetchUser(action: CommonAction<string, string>) {
  try {
    const user: User = yield call(getUser, action.payload);
    yield put({type: actionTypes.FETCH_USER_SUCCEEDED, payload: {isLoggedIn: true, user: user}});
  } catch (e) {
    yield put({
      type: actionTypes.FETCH_USER_FAILED,
      payload: {isLoggedIn: true, message: e, user: null},
    });
  }
}

function* loginUser(action: CommonAction<string, string>) {
  try {
    const payload = action.payload as any;
    const user: User = yield call(login, payload.userId, payload.password);
    yield put({type: actionTypes.LOGIN_SUCCESS, payload: {isLoggedIn: true, user: user}});
  } catch (e) {
    yield put({
      type: actionTypes.LOGIN_FAIL,
      payload: {isLoggedIn: true, message: e, user: null},
    });
  }
}

function* watchFetchUser() {
  yield takeLatest(actionTypes.FETCH_USER, fetchUser);
}

function* watchLoginUser() {
  yield takeLatest(actionTypes.LOGIN, loginUser);
}

export default [fork(watchFetchUser), fork(watchLoginUser)];
