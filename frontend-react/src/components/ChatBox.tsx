import React, {ChangeEventHandler, FormEventHandler} from 'react';

interface ChatBoxProps{
    onChangeMsg:ChangeEventHandler
    onSendMsg:FormEventHandler
    msg:string
}

function ChatBox({msg,onSendMsg,onChangeMsg}:ChatBoxProps){
    return (
        <div className="chat-composer">
            <form onSubmit={onSendMsg}>
                <input
                    className="form-control"
                    placeholder="Type & hit enter"
                    onChange={onChangeMsg}
                    value={msg}
                />
            </form>
        </div>
    );
}

export default ChatBox;