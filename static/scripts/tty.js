const terminal = new Terminal({
    convertEol: true,
    cursorBlink: true,
});

const host = window.location.host; 

terminal.open(document.getElementById('terminal'));
const socket = new WebSocket(`wss://${host}/dashboard/ws`);
socket.onmessage = function(event) {
    const out = event.data
    terminal.write(out.replace(`/\r/g`, `\r\n`));
    printPrompt();
};

let currentCommand = ``;
const prompt = `desky>>`;

terminal.onData(function(data) {
    if (data == `\r`) {
        socket.send(currentCommand);
        terminal.write(`\n`)
        currentCommand = ``;
    } else if (data == `\u007f`) {
        if (currentCommand.length > 0) {
            currentCommand = currentCommand.slice(0, -1);
            terminal.write(`\b \b`)
        }
    } else {
        currentCommand += data;
        terminal.write(data)
    }
});

terminal.attachCustomKeyEventHandler((event) => {
    if (event.ctrlKey && event.key == `l`) {
        event.preventDefault();
        terminal.clear();
        return false;
    };
    return true;
});

printPrompt();

function printPrompt () {
    terminal.write(`\r\n${prompt} `);
}