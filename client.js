if ("WebSocket" in window) {
  var ws;
  if (!ws) { ws = new WebSocket("ws://localhost:8080/"); }
  ws.onmessage = function(evt) { };
  ws.onclose = function() { 
    ws.send('closed')
    ws = null;
  };
  ws.onopen = function() { 
    ws.send('opening');
  };
}
