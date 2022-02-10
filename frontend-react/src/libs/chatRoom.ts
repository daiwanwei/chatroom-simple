const socket = new WebSocket("ws://localhost:8080/room");

export function enterRoom(){
    console.log("attempt to connect");
    socket.onopen = () => {
        console.log("connect successfully")
    }
    socket.onclose = (event) => {
        console.log(`socket close:${event}`)
    }
    socket.onerror = (err) => {
        console.log(`socket err: ${err}`)
    }
}

export function receiveMsgFromRoom(cb:(msg:Message)=>void){
    socket.onmessage = (e) => {
        console.log(`get msg event:${e}`)
        if (e.data){
            let m=JSON.parse(e.data) as Message;
            console.log(m.message)
            cb(m)
        }
    }
}

export function sendMsgToRoom(msg:string){
    socket.send(JSON.stringify({
        message:msg,
    }));
}

interface Message{
    name:string
    message:string
    when:string
}