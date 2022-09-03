import { callAPI } from "./api";
import { ErrorState, LogoutState, Users } from "./interfaces"

export const login = async (data: Users) => {
  const auth = await callAPI('/auth/login', 'POST', data);

  return auth;
}

export const logout = async () => {
  const auth = await callAPI('/auth/logout', 'GET');

  const result: LogoutState | ErrorState = auth;

  return result;
}

export const isAuthenticated = async () => {
  const auth = await callAPI('/auth/is_login', 'GET');

  return auth;
}
