import './style.css';
import './app.css';
import { GetData } from "../wailsjs/go/main/App.js"
import { dataModel } from './assets/data.model';

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

function handleFormData(data: dataModel}) {
    console.log(data)
    GetData(data.url, data.startHH)
}
