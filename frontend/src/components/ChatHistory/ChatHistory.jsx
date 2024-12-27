import React from "react";
import './ChatHistory.css';
import Message from '../Message/Message';

const ChatHistory = ({ chatHistory }) => {
    const messages = chatHistory.map((msg) => (
        <Message key={msg.timeStamp} message={msg.data} />
    ));

    return (
        <div className="chat-history">
            <h2 className="chat-history-title">Chat History</h2>
            <div className="chat-history-messages">{messages}</div>
        </div>
    );
};

export default ChatHistory;
