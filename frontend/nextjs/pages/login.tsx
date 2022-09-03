import Head from "next/head"
import {useRouter} from "next/router";
import { useEffect, useState } from "react";
import { Users } from "../lib/interfaces";
import { login, isAuthenticated } from "../lib/auth";
import {Message} from "../components/Message";

export default function Login() {
  const [username, setUsername] = useState("");
  const [password, setPassword] = useState("");
  const [messageType, setMessageType] = useState("");
  const [message, setMessage] = useState("");

  const router = useRouter();

  // TODO: use middleware
  useEffect(() => {
    isAuthenticated()
    .then((res) => {
      if (res.login) {
        console.log('User is already logged in. Redirecting ...')
        router.push('/');
      }
    })
    .catch((err) => {
      console.error(err)
    })
  })

  const submitLogin = async (e: any) => {
    e.preventDefault();

    const data: Users = {
      username: username,
      password: password
    };

    const auth = await login(data);

    if (auth.login) {
      setMessageType('success');
      setMessage(auth.message);

      setTimeout(() => {
        router.push('/');
      }, 1000);

    } else {
      setMessageType('error');
      setMessage(auth.message);
    }

    e.target.reset();
  }

  return (
    <div className="container mx-auto">
      <Head>
        <title>Login to notes application - by Anggit M Ginanjar (itsgitz)</title>
      </Head>
      <h1 className="text-3xl font-bold underline py-3">Login Page!</h1>
      <div className="py-3 max-w-md">
        <form onSubmit={(e: any) => { submitLogin(e); }}>
          <div className="py-3">
            <input
              className="border-b-2 border-gray-300 py-2 w-full focus:outline-none"
              type="text"
              placeholder="Username"
              onChange={(e: any) => setUsername(e.target.value)}
            />
          </div> 
          <div className="py-3">
            <input 
              className="border-b-2 border-gray-300 py-2 w-full focus:outline-none"
              type="password"
              placeholder="Password"
              onChange={(e: any) => setPassword(e.target.value)}
            />
          </div> 
          <div className="py-3">
            <input className="py-2 px-4 bg-blue-400 rounded font-bold text-white w-full" type="submit" value="Login" />
          </div>
        </form> 

        <div className="py-2"></div>

        <Message type={messageType} message={message} />
      </div> 
    </div>
  )
}
