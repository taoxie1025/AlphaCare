export interface CommonAction<T, P> {
  readonly type: T;
  readonly payload: P;
}

export interface User {
  userId: string;
  // TODO
}

export interface UserState {
  isLoggedIn: boolean;
  error?: Error;
  user?: User;
}

export interface AppState {
  user?: UserState;
}
