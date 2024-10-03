const restart_buttons = document.querySelectorAll('.restart-btn');

const start_buttons = document.querySelectorAll('.start-btn');
const stop_buttons = document.querySelectorAll('.stop-btn');

const enable_buttons = document.querySelectorAll('.enable-btn');
const disable_buttons = document.querySelectorAll('.disable-btn');


let reloadPage = function reloadPage(){
    window.location.reload();
}

async function reqApi(button, exec)  {
    const service = button.getAttribute('data-service');
    const url = `/api/systemd/${service}/${exec}`;
    console.log(url)
    const response = await fetch(url, {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json'
        }
    });
    if (response.ok) {
        setInterval(alert("Success!"), 500)
        setInterval(reloadPage, 2000);
    } else {
        alert("Error request")
    }
}



restart_buttons.forEach(button => {
    button.addEventListener('click', () => {
        const result = confirm("Restart service?");
        if (result) {reqApi(button, "restart")}
    });
});

start_buttons.forEach(button => {
    button.addEventListener('click', () => {
        const result = confirm("Start service?");
        if (result) {reqApi(button, "start")}
    });
});
stop_buttons.forEach(button => {
    button.addEventListener('click', () => {
        const result = confirm("Stop service?");
        if (result) {reqApi(button, "stop")}
    });
});

enable_buttons.forEach(button => {
    button.addEventListener('click', () => {
        const result = confirm("Enable service?");
        if (result) {reqApi(button, "enable")}
    });
});
disable_buttons.forEach(button => {
    button.addEventListener('click', () => {
        const result = confirm("Disable service?");
        if (result) {reqApi(button, "disable")}
    });
});