# SPA Authentication API with Cookie

This implements how SPA (Single Page Application) use API authentication with a cookie as we know that localStorage and sessionStorage are not secure for storing sensitive information such as token for API authentication. localStorage and sessionStorage could be accessed by a third-party script (XSS). To prevent this, we can use the HttpOnly cookie. With this approach, the third-party script cannot access the cookie from the browser and is only used when the request is sent to the server.

# Todo
* Add CSRF token
* Finish frontend
* HTTPS

# Contributor
Anggit M Ginanjar - <anggit.ginanjar.dev@gmail.com>
