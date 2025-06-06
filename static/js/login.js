function convertLoginToJson() {
    let form = document.getElementById("loginData");
    let formData = {};
    for (let i = 0; i < form.elements.length; i++) {
        let element = form.elements[i];
        if (element.type !== "submit") {
            formData[element.name] = element.value;
        }
    }
    return JSON.stringify(formData);
}

async function sendLoginData(jsonData) {
    let url = "http://localhost:8080/api/login"
    try {
        const response = await fetch(url, {
            method: "POST",
            headers: {
                "Content-Type": "application/json",
            },
            body: jsonData,
        });

        const json = await response.json()
        if (json.token) {
            localStorage.setItem("token", json.token);
            location.href = "/";
        } else {
            alert("Login failed")
        }

    }
    catch (error) {
        console.error('Error:', error)
    }   
}

const loginButton = document.getElementById("loginButton");
loginButton.addEventListener("click", (e) => {
    e.preventDefault();
    const json = convertLoginToJson();
    sendLoginData(json);
})