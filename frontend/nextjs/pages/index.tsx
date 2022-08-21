import Head from 'next/head'
import {useRouter} from 'next/router';
import { useEffect, useState } from 'react';
import { Notes } from '../lib/interfaces';
import { getNotes } from '../lib/notes';

export default function Home({}) {
  const router = useRouter();
  const [notes, setNotes] = useState<Notes[]>([]);

  useEffect(() => {
    if (notes.length === 0) {
      getNotes().then((res) => {
        if (res.error) {
          return router.push('/login');
        }

        setNotes(res.data);
      })
    } 
  }, [notes]);

  return (
    <div className="container mx-auto">
      <Head>
        <title>Take a notes application - by Anggit M Ginanjar (itsgitz)</title>
        <link rel="icon" href="/favicon.ico" />
      </Head>

      <main className=''>
        <h1 className='text-3xl font-bold underline'>Hello!</h1>

        <Notes notes={notes} />
      </main> 
    </div>
  )
}

const Notes = (props: any) => {
  const notes: Notes[] = props.notes;

  if (notes.length > 0) {
    return (
      <>
        <div className='py-3' />
        <div>
          {notes.map((note: Notes, key: number) => {
            return (
              <div key={key}>{note.content}</div>
            )
          })}
        </div>
      </> 
    );
  } 

  return null;
}
