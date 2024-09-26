const reboot_buttons = document.querySelectorAll('.reboot-btn');
const shutdown_buttons = document.querySelectorAll('.shutdown-btn');
const start_buttons = document.querySelectorAll('.start-btn');


let reloadPage = function reloadPage(){
    window.location.reload();
}

async function reqApi(button, exec)  {
    const id = button.getAttribute('data-id');
    const type = button.getAttribute('data-type');
    let url = " ";
    if (type == "lxc") {
        url = `/api/proxmox/pct/${id}/${exec}`;
    } else {
        url = `/api/proxmox/qm/${id}/${exec}`;
    };
    const response = await fetch(url, {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json'
        }
    });
    if (response.ok) {
        setInterval(alert("Success!"), 500)
        setInterval(reloadPage, 5000);
    } else {
        alert("Error request")
    }
}

reboot_buttons.forEach(button => {
    button.addEventListener('click', () => {
        const result = confirm("Reboot Virtual device?");
        if (result) {reqApi(button, "reboot")}
    });
});

shutdown_buttons.forEach(button => {
    button.addEventListener('click', () => {
        const result = confirm("Shutdown Virtual device?");
        if (result) {reqApi(button, "shutdown")}
    });
});

start_buttons.forEach(button => {
    button.addEventListener('click', () => {
        const result = confirm("Start Virtual device?");
        if (result) {reqApi(button, "start")}
    });
});


const container = document.querySelector('.virt');

container.addEventListener('wheel', (event) => {
    event.preventDefault();
    container.scrollLeft += event.deltaY;
});