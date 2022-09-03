import { callAPI } from "./api";

export const getUsers = async () => {
  const users = await callAPI('/api/users', 'post');

  return users;
}
