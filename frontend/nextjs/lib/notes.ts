import {useEffect, useState} from "react";
import {callAPI} from "./api";
import {ErrorState, Notes} from "./interfaces";

export const getNotes = async () => {
  const notes = await callAPI('/api/notes', 'GET');

  return notes;
}

export const useNotes = () => {
  const [notes, setNotes] = useState<Notes[]>([]);
  const [error, setError] = useState<ErrorState>();

  useEffect(() => {
    if (notes.length == 0) {
      getNotes().then((res) => {
        if (res.error) { 
          setError(res.error);
        } else {
          setNotes(res.data);
        }
      })
    } 
  })

  return [
    notes,
    error
  ]
}
