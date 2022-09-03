import Head from 'next/head'
import {useRouter} from 'next/router';
import {useEffect} from 'react';
import { isAuthenticated, logout } from '../lib/auth';
import { Notes } from '../lib/interfaces';
import { useNotes } from '../lib/notes';

export default function Home({}) {
  const router = useRouter();
  const [notes, notesError] = useNotes(); 
  
  // TODO: use middleware
  useEffect(() => {
    isAuthenticated()
    .then((res) => {
      if (!res.login) {
        console.log('User is not logged in. Redirecting ...')
        router.push('/login');
      }
    })
    .catch((err) => {
      console.error(err)
    })
  })

  const submitLogout = async () => {
    const auth = await logout();

    if (auth) {
      router.push('/login');
    }
  }

  
  return (
    <div className="container mx-auto">
      <Head>
        <title>Take a notes application - by Anggit M Ginanjar (itsgitz)</title>
        <link rel="icon" href="/favicon.ico" />
      </Head>

      <main className=''>
        <h1 className='text-3xl font-bold underline'>Take a note today</h1>

        <div className='py-3'>
          <a onClick={submitLogout}>Logout</a>
        </div>

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
