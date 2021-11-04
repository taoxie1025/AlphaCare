import {configureStore} from '@reduxjs/toolkit';
import userReducer from './reducers/user';
import createSagaMiddleware from 'redux-saga';
import rootSatga from './sagas';

export default function initStore() {
  const sagaMiddleware = createSagaMiddleware();
  const store = configureStore({
    reducer: {
      user: userReducer,
    },
    middleware: [sagaMiddleware],
  });

  sagaMiddleware.run(rootSatga);
  return store;
}
