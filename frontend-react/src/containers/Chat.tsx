import React, {useState} from "react";
import ChatBox from "../components/ChatBox";
import ChatHistory from "../components/ChatHistory";
import {enterRoom,sendMsgToRoom,receiveMsgFromRoom} from "../libs/chatRoom";

enterRoom()

function Chat() {
    let [msg,setMsg]=useState("nono");
    let [msgList,setMsgList]=useState([{text:"ann"}]);
    console.log("test first")
    const changeMsg=(e:React.ChangeEvent<HTMLInputElement>)=>{
        let typedValue=e.target.value;
        setMsg(typedValue);
        console.log(`change msg${msg}`)
    }
    console.log("test second")
    const sendMsg = (e:React.FormEvent<HTMLInputElement>) => {
        e.preventDefault();
        sendMsgToRoom(msg);
        setMsg("");
    };

    receiveMsgFromRoom((msg)=>{
        console.log(msg)
        setMsgList([...msgList,{text: msg.message}])
    })

    return (
        <>
            <ChatHistory msgList={msgList}/>
            <ChatBox msg={msg} onChangeMsg={changeMsg} onSendMsg={sendMsg}/>
        </>

    );
}

export default Chat;