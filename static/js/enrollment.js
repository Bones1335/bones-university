function convertUserToJSON() {
    let form = document.getElementById("enrollmentForm");
    let formData = {};
    for (let i = 0; i < form.elements.length; i++) {
        let element = form.elements[i];
        if (element.type !== "submit") {
            formData[element.name] = element.value;
        }
    }

    return JSON.stringify(formData);
}

async function sendUserData(jsonData) {
    let url = "http://localhost:8080/api/users"
    try {
        const response = await fetch(url, {
            method: "POST",
            headers: {
                "Content-Type": "application/json"
            },
            body: jsonData,
        });

        await response.json();
    } catch (error) {
        console.error('Error: ', error)
    }
}

const submitEnrollment = document.getElementById("submitEnrollment");
submitEnrollment.addEventListener("click", (e) => { 
    e.preventDefault();
    const json = convertUserToJSON(); 
    sendUserData(json);
})

