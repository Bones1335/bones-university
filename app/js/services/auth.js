import { apiClient } from '../api/client.js';
import { showScreen } from "../ui/screens.js";

export class AuthService {
  constructor() {
    this.currentUser = null;
  }

  setAuthToken(token) {
    localStorage.setItem('authToken', token);
  }

  clearAuthToken() {
    localStorage.removeItem('authToken');
  }

  getAuthToken() {
    return localStorage.getItem('authToken');
  }

  async login(username, password) {
    try {
      const response = await apiClient.post('/login', {"login_username": username, "login_password": password});
      this.setAuthToken(response.token);
      this.currentUser = response.username;
    } catch (error) {
      throw new Error('Login failed. Please check your credentials.');
    }
  }

  logout() {
    this.clearAuthToken();
    this.currentUser = null;
    showScreen('home_screen');
  }

  isAuthenticated() {
    return !!this.getAuthToken();
  }
}

export const authService = new AuthService();
