# SPA Authentication API with Cookie

This implements how SPA (Single Page Application) use API authentication with a cookie as we know that `localStorage` and `sessionStorage` are not secure for storing sensitive information such as token for API authentication. `localStorage` and `sessionStorage` could be accessed by a third-party script (XSS). To prevent this, we can use the `HttpOnly` cookie. With this approach, the third-party script cannot access the cookie from the browser and is only used when the request is sent to the server.

# Usage
## Prerequisite
* Go version: `go1.17.2 linux/amd64`
* Docker Engine: `20.10+`
* Docker Compose: `v2.4.1`

### Backend
* First thing first, start with run the backend server
* Go to `backend` directory project, and run command:
```
$ cd backend
$ make run
```

### Frontend
* Go to `frontend/nextjs` directory, and run command:
```
$ cd frontend/nextjs
$ npm run dev
```

###

# Todo
* Add CSRF token
* Dockerize the backend

# Contributor
Anggit M Ginanjar - <anggit.ginanjar.dev@gmail.com>
