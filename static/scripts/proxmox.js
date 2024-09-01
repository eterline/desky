const reboot_buttons = document.querySelectorAll('.reboot-btn');
const shutdown_buttons = document.querySelectorAll('.shutdown-btn');
const start_buttons = document.querySelectorAll('.start-btn');


let d = function reloadPage(){
    window.location.reload();
}
// Добавляем обработчик события на каждую кнопку
reboot_buttons.forEach(button => {
    button.addEventListener('click', () => {
        const id = button.getAttribute('data-id');
        const type = button.getAttribute('data-type');
        let url = " ";
        if (type == "lxc") {
            url = `/api/pct/${id}/reboot`;
        } else {
            url = `/api/qm/${id}/reboot`;
        };
        fetch(url, {
            method: 'GET',
            headers: {
                'Content-Type': 'application/json'
            }
        });
        setInterval(d, 5000);
    });
});

shutdown_buttons.forEach(button => {
    button.addEventListener('click', () => {
        const id = button.getAttribute('data-id');
        const type = button.getAttribute('data-type');
        let url = " ";
        if (type == "lxc") {
            url = `/api/pct/${id}/shutdown`;
        } else {
            url = `/api/qm/${id}/shutdown`;
        };
        fetch(url, {
            method: 'GET',
            headers: {
                'Content-Type': 'application/json'
            }
        });
        setInterval(d, 5000);
    });
});

start_buttons.forEach(button => {
    button.addEventListener('click', () => {
        const id = button.getAttribute('data-id');
        const type = button.getAttribute('data-type');
        let url = " ";
        if (type == "lxc") {
            url = `/api/pct/${id}/start`;
        } else {
            url = `/api/qm/${id}/start`;
        };
        fetch(url, {
            method: 'GET',
            headers: {
                'Content-Type': 'application/json'
            }
        });
        setInterval(d, 5000);
    });
});