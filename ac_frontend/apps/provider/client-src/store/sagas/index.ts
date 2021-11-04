import {all} from '@redux-saga/core/effects';
import user from './user';

const sagas = [...user];
export default function* rootSatga(): Generator {
  yield all(sagas);
}
