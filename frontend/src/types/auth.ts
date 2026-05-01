export interface AuthResponse {
  access_token: string;
  expires_in: number;
  refresh_token: string;
}

export interface User {
  id: string;
  email: string;
  role: string;
}

export interface LoginError {
  message: string;
  code?: string;
}
