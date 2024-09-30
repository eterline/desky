const reboot_buttons = document.querySelectorAll('.reboot-btn');
const shutdown_buttons = document.querySelectorAll('.shutdown-btn');
const start_buttons = document.querySelectorAll('.start-btn');

const virt = document.querySelector('.virt');
const firstItem = virt.querySelector('.dev-block');

let reloadPage = function reloadPage(){
    window.location.reload();
}

window.addEventListener('load', () => {
    const itemStyles = window.getComputedStyle(firstItem);
    const itemWidth = firstItem.offsetWidth + 
                      parseFloat(itemStyles.marginLeft) + 
                      parseFloat(itemStyles.marginRight);
    virt.style.width = `${itemWidth*2+15}px`;
});

async function reqApi(button, exec)  {
    const id = button.getAttribute('data-id');
    const type = button.getAttribute('data-type');
    const host = button.getAttribute('data-host');
    let url = " ";
    if (type == "lxc") {
        url = `/api/proxmox/${host}/pct/${id}/${exec}`;
    } else {
        url = `/api/proxmox/${host}/qm/${id}/${exec}`;
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