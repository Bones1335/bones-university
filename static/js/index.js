const enroll = document.getElementById("enrollment");
enroll.addEventListener("click", () => {
    location.href = "/enrollment";
});

const loginButton = document.getElementById("login");
loginButton.addEventListener("click", () => {
    location.href = "/login";
})

const logoutButton = document.getElementById("logout");
logoutButton.addEventListener("click", () => {
    localStorage.removeItem("token");
    location.href = "/"
})

document.addEventListener("DOMContentLoaded", async () => {
    const token = localStorage.getItem('token');

    if (token) {
        enroll.style.display = "none";
        loginButton.style.display = "none";
        logoutButton.style.display = "block";
    } else {
        enroll.style.display = "block";
        loginButton.style.display = "block";
        logoutButton.style.display = "none";        
    }
})