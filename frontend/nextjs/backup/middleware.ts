import {NextRequest, NextResponse} from "next/server";

export const middleware = (req: NextRequest) => {
  const session = req.cookies.get('session_id');

  // Check session, redirect to login page if not logged in
  if (!isAssets(req) && !isLoginOrRegisterPage(req)) {
    if (!session) {
      return NextResponse.redirect(new URL('/login', req.url));
    }
  }

  // Don't allow user to access the login page if already logged in
  if (isLoginOrRegisterPage(req)) {
    if (session) {
      return NextResponse.redirect(new URL('/', req.url));
    }
  }
}

const isAssets = (req: NextRequest): boolean => {
  return req.url.includes('.js') || 
    req.url.includes('favicon'); 
}

const isLoginOrRegisterPage = (req: NextRequest): boolean => {
  return req.url.includes('login') ||
    req.url.includes('register'); 
}
