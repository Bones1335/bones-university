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

    this.setupNavigation();
  }

  setupNavigation() {
    document.addEventListener('click', async (e) => {
      if (e.target.matches('[data-screen]')) {
        const screenId = e.target.dataset.screen;
        showScreen(screenId);

        if (screenId === 'admin_dashboard_screen') {
          this.loadAdminDashboard();
        } else if (screenId === 'enrollment_screen') {
          this.loadEnrollment();
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

  async loadEnrollment(e) {console.log("you are logged in to the enrollment screen.")}

  async loadAdminDashboad(e) {console.log("you are logged in to the admin dashboard.")}

  async loadStudentDashboad(e) {console.log("you are logged in to the student dashboard.")}
}

document.addEventListener('DOMContentLoaded', () => {
  const app = new UniversityApp();

  window.universityApp = app;
})
