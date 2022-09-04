// UNUSED MIDDLEWARE
// IT SEEMS LIKE WE CANNOT USE FETCH OR AXIOS FOR CALLING API REQUEST IN MIDDLEWARE

// import {NextRequest, NextResponse} from "next/server";
// import {isAuthenticated} from "./lib/auth";

// export const middleware = (req: NextRequest) => {
//   // const session = req.cookies.get('session_id');

//   // // Check session, redirect to login page if not logged in
//   // if (!isAssets(req) && !isLoginOrRegisterPage(req)) {
//   //   if (!session) {
//   //     return NextResponse.redirect(new URL('/login', req.url));
//   //   }
//   // }

//   // // Don't allow user to access the login page if already logged in
//   // if (isLoginOrRegisterPage(req)) {
//   //   if (session) {
//   //     return NextResponse.redirect(new URL('/', req.url));
//   //   }
//   // } 

//   isAuthenticated()
//   .then((res) => {
//     console.log('res:', res)
//   })
//   .catch((err) => {
//     console.error('err:', err)
//   })
// }

// // const isAssets = (req: NextRequest): boolean => {
// //   return req.url.includes('.js') || 
// //     req.url.includes('favicon'); 
// // }

// // const isLoginOrRegisterPage = (req: NextRequest): boolean => {
// //   return req.url.includes('login') ||
// //     req.url.includes('register'); 
// // }
