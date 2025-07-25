import './style.css';
import './app.css';
import { DownloadAndTrim } from "../wailsjs/go/main/App.js"
import { ytDataModel } from './ytData.model';

const form = document.getElementById("yvacForm") as HTMLFormElement;
document.getElementById("resetBtn")?.addEventListener("click", onReset);

form.addEventListener("submit", function(event) {
    event.preventDefault();

    const formData = new FormData(form);

    const data: ytDataModel = {
        Url: formData.get("url") as string,
        StartHH: formData.get("startHH") as string,
        StartMM: formData.get("startMM") as string,
        StartSS: formData.get("startSS") as string,
        EndHH: formData.get("endHH") as string,
        EndMM: formData.get("endMM") as string,
        EndSS: formData.get("endSS") as string,
        Name: formData.get("filename") as string
    };

    handleFormData(data);
})

function handleFormData(data: ytDataModel) {
    DownloadAndTrim(data)
}

function onReset() {
    form.reset();
}
