var socket = new WebSocket("ws://localhost:8080/websocket");
var log = document.getElementById("log");

socket.onmessage = function (event) {
  var frame = JSON.parse(event.data);
  switch(frame.type) {
    case "status":
      log.insertAdjacentHTML('beforeend', frame.status + "<br />");
      break;
    case "countdown":
      log.insertAdjacentHTML('beforeend', "rolling in " + frame.secondsLeft + "<br />");
      break;
    case "roll":
      log.insertAdjacentHTML('beforeend', "roll for " + frame.candidates.join(', ') + ": <b>" + frame.winner + "</b><br />");
      break;
  }
}

socket.onopen = function (event) {
  log.insertAdjacentHTML('beforeend', "opened<br />");
};

socket.onclose = function (event) {
  log.insertAdjacentHTML('beforeend', "closed<br />");
};

socket.onerror = function (event) {
  console.log(event);
};


var rollRequest = {
  "type": "rollRequest",
  "candidates": ["ben", "jeff", "kevin", "vincent"]
};

var requestRoll = function () {
  socket.send(JSON.stringify(rollRequest));
};
