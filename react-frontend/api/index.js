let socket = new WebSocket("ws://localhost:8080/ws");

let connect = (cb) => {
  console.log("attempting connection");
  socket.onopen = () => {
    console.log("Houston, we have landing.");
  };

  socket.onmessage = (msg) => {
    console.log("transmission: ", msg.data);
    cb(msg);
  };

  socket.onclose = (event) => {
    console.log(`socket closed connection: ${event}`);
  };
};

socket.onerror = (err) => {
  console.log(`socket error: ${err}`);
};

let sendMSG = (msg) => {
  socket.send(msg);
};

export { connect, sendMSG };
