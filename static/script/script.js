const enroll = document.getElementById("enrollment");
enroll.addEventListener("click", () => {
    location.href = "/enrollment";
});

const submitEnrollment = document.getElementById("submitEnrollment");
submitEnrollment.addEventListener("click", () => {
    let form = document.getElementById("enrollmentForm");
    let formData = {};
    for (let i = 0; i < form.elements.length; i++) {
        let element = form.elements[i];
        if (element.type !== "submit") {
            formData[element.name] = element.value;
        }
    }
})