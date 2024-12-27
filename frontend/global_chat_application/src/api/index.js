var socket = new WebSocket('ws://localhost:8080');

let connect = (cb)=>{
    console.log("connecting...")

    socket.onopen = () => {
        console.log("websocket connected successfully");
    }

    socket.onclose = () =>{
        console.log("websocket disconnected",event);
    }

    socket.onmessage = (msg) =>{
        console.log("message received ",msg);
        cb(msg);
    }
    socket.onerror = (error) => {
        console.log("websocket error ",error);
    }
}

let sendMsg = (msg)=>{
    console.log("message sent:= ",msg);
    socket.send(msg)
}

export {connect,sendMsg}  //exporting the functions to be used in other files