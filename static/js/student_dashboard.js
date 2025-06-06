console.log('student_dashboard')
const logoutButton = document.getElementById("logout");
logoutButton.addEventListener("click", () => {
    localStorage.removeItem("token");
    location.href = "/"
})