import { callAPI } from "./api";
import { Users } from "./interfaces"

export const login = async (data: Users) => {
  const auth = await callAPI('/auth/login', 'POST', data);

  return auth;
}
