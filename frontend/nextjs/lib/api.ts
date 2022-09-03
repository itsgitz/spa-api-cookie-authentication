import axios from 'axios';

// NOTES: use `fetch` API instead
// with axios, I faced this problem:
// `[TypeError: adapter is not a function]`
export const callAPI = async (path: string, method: string, body: any = {}) => {
  const url: string = process.env.NEXT_PUBLIC_API_ENDPOINT;

  let options: any = {
    method: method,
    url: `${url}${path}`,
    withCredentials: true,
    headers: {
      'Content-Type': 'application/json'
    }
  }

  if (body) {
    options.data = body;
  }

  const response = await axios(options);

  return response.data;
}
