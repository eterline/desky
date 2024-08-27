    // Total seconds to wait
    var seconds = 6;
    
    function countdown() {
        seconds = seconds - 1;
        if (seconds < 0) {
            window.location = "/dashboard/panel";
        } else {
            document.getElementById("timer").innerHTML = seconds;
            window.setTimeout("countdown()", 1000);
    }
}

countdown()