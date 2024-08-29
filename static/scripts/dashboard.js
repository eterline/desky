function fetchStats() {
    fetch('/api/system')
        .then(response => response.json())
        .then(data => {
            document.getElementById('cpu-progress').value = data.cpu.load;
            document.getElementById('mem-progress').value = data.mem.load;
            document.getElementById('cpu-usage').innerText = data.cpu.load+'%';
            document.getElementById('mem-usage').innerText = data.mem.load+'%';
        })
        .catch(error => console.error('Ошибка:', error));
}

setInterval(fetchStats, 5000);
fetchStats();