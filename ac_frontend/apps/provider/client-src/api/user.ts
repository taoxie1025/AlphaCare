import {User} from '../../@types/global';
import axios from 'axios';

export function getUser(userId: string): Promise<User> {
  // TODO: implement actual fetch user
  return Promise.resolve({
    userId,
  });
}

export function login(userId: string, password: string): Promise<User> {
  // TODO: implement actual change password

  return axios.post('/provider/login', {userId, password}).then((data) => {
    return ({} as any) as User;
  });
}

export function register(
  userName: string,
  password: string,
  firstName: string,
  lastName: string
): Promise<User> {
  // TODO: implement actual change password
  return axios.post('/provider/register', {userName, password}).then((data) => {
    return ({} as any) as User;
  });
}
