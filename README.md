# SPA Authentication API with Cookie

This implements how SPA (Single Page Application) use API authentication with cookie.
As we know that `localStorage` and `sessionStorage` is not secure for storing
sensitive information such as `token` for API authentication. `localStorage` and `sessionStorage`
could be accessed by third party script (XSS). For prevent this, we can use `HttpOnly` cookie.
With this approach, the third party script unable to access the cookie from browser and the cookie only used
when request is send to server.

# Todo
* Add CSRF token
* Finish frontend
* HTTPS

# Contributor
Anggit M Ginanjar - <anggit.ginanjar.dev@gmail.com>
