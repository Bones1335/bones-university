import { authService } from "./services/auth.js";
import { userService } from "./services/users.js";
import { showScreen, showError, clearError } from "./ui/screens.js";

class UniversityApp {
  constructor() {
    this.init();
  }

  init() {
    this.setupEventListeners();
    this.checkAuthStatus();
  }

  setupEventListeners() {
    document.getElementById('login_form').addEventListener('submit', async (e) => {
      e.preventDefault();
      await this.handleLogin(e);
    });

    document.getElementById('enrollment_form').addEventListener('submit', async (e) => {
      e.preventDefault();
      await this.handleEnrollment(e);
    });

    this.setupNavigation();
  }

  setupNavigation() {
    document.addEventListener('click', async (e) => {
      if (e.target.matches('[data-screen]')) {
        const screenId = e.target.dataset.screen;
        showScreen(screenId);

        if (screenId === 'admin_dashboard_screen') {
          this.loadAdminDashboard();
        } else if (screenId === 'student_dashboard_screen') {
          this.loadStudentDashboad();
        } else if (screenId === 'login_screen') {
          authService.logout();
          this.checkAuthStatus();
        }
      }
    });
  }

  checkAuthStatus() {
    if (authService.isAuthenticated()) {
      showScreen('home_screen');
    } else {
      showScreen('login_screen');
    }
  }

  async handleLogin(e) {
    clearError('login_error');

    const formData = new FormData(e.target);
    const username = formData.get('login_username');
    const password = formData.get('login_password');

    try {
      await authService.login(username, password);
      showScreen('home_screen');
      e.target.reset();
    } catch (error) {
      showError('login_error', error.message);
    }
  }

  async handleEnrollment(e) {
    clearError('enrollment_error');

    const formData = new FormData(e.target);
    const userData = {
      enrollment_last_name: formData.get('enrollment_last_name'),
      enrollment_first_name: formData.get('enrollment_first_name'),
      enrollment_personal_email: formData.get('enrollment_personal_email'),
      enrollment_password: formData.get('enrollment_password'),
    }

    try {
      const response = await userService.createUser(userData);
      alert(`Here's your new username to login in to the university website: ${response.username}`);
      showScreen('login_screen');
      e.target.reset();
    } catch (error) {
      showError('enrollment_error', error.message);
    }
  }

  async loadAdminDashboad(e) {console.log("you are logged in to the admin dashboard.")}

  async loadStudentDashboad(e) {console.log("you are logged in to the student dashboard.")}
}

document.addEventListener('DOMContentLoaded', () => {
  const app = new UniversityApp();

  window.universityApp = app;
});
