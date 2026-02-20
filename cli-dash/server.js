const WebSocket = require('ws');

function connectToServer(onOpen, onMessage, onClose) {
  const ws = new WebSocket('ws://localhost:8080/ws');
  ws.on('open', onOpen);
  ws.on('message', onMessage);
  ws.on('close', onClose);
  return ws;
}

module.exports = { connectToServer };

