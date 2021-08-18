let socket = new Websocket("ws://localhost:8080/ws");

let connect = () => {
  console.log("attempting connection");
  socket.onopen = () => {
    console.log("Houston, we have landing.");
  };

  socket.onmessage = (msg) => {
    console.log(msg);
  };

  socket.onclose = (event) => {
    console.log(`socket closed connection: ${event}`);
  };
};

let sendMSG = (msg) => {
  console.log(`sending msg: ${msg}`);
  socket.send(msg);
};

export { connect, sendMSG };
