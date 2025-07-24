import './style.css';
import './app.css';

const form = document.getElementById("yvacForm") as HTMLFormElement

form.addEventListener("submit", function(event) {
    event.preventDefault();

    const formData = new FormData(form);

    const data = {
        url: formData.get("url") as string,
        start: formData.get("start") as string,
        end: formData.get("end") as string,
        filename: formData.get("filename") as string
    };

    handleFormData(data);
})

function handleFormData(data: {
    url: string,
    start: string,
    end: string,
    filename: string
}) {
    console.log(data)
}
