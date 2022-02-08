import React from "react";

interface ChatHistoryProps{
    msgList: Message[]
}

interface Message{
    text:string
}

function ChatHistory({msgList}:ChatHistoryProps){
    return (
        <div className="chat-window">
            <div className="box">
                <div className="inner">
                    {Array.isArray(msgList) &&
                    msgList.map((msg, index) => (
                        <p key={index} className="message">
                            {msg.text}
                        </p>
                    ))}
                    {/* define ref and call it if component is updated */}
                    {/*<div*/}
                    {/*    className="reference"*/}
                    {/*    ref={node => (this.messageListEnd = node)}*/}
                    {/*/>*/}
                </div>
            </div>
        </div>
    );
}

export default ChatHistory;