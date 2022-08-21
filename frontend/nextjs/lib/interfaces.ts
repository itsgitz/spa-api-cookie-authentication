export interface ErrorState {
  error: string,
  message: string,
}

export interface LoginState {
  login: boolean,
  message: string,
}

export interface Users {
  id?: number,
  username?: string,
  password?: string,
}

export interface Notes {
  id: number,
  content: string,
  created_at: string
}
