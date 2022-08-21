import {callAPI} from "./api";

export const getNotes = async () => {
  const notes = await callAPI('/api/notes', 'GET');

  return notes;
}
