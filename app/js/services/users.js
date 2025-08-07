import { apiClient } from "../api/client.js";

export class UserService {
  async createUser(userData) {
    return await apiClient.post('/users', userData);
  }
}

export const userService = new UserService();
