import './style.css';
import './app.css';
import { GetData } from "../wailsjs/go/main/App.js"
import { ytDataModel } from './ytData.model';

const form = document.getElementById("yvacForm") as HTMLFormElement

form.addEventListener("submit", function(event) {
    event.preventDefault();

    const formData = new FormData(form);

    const data: ytDataModel = {
        url: formData.get("url") as string,
        startHH: formData.get("startHH") as string,
        startMM: formData.get("startMM") as string,
        startSS: formData.get("startSS") as string,
        endHH: formData.get("endHH") as string,
        endMM: formData.get("endMM") as string,
        endSS: formData.get("endSS") as string,
        name: formData.get("filename") as string
    };

    handleFormData(data);
})

function handleFormData(data: ytDataModel) {
    GetData(data.url, data.startHH, data.startMM, data.startSS, data.endHH, data.startMM, data.endSS, data.name)
}
