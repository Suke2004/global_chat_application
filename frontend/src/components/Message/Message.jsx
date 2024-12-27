import React from "react";
import './Message.css';

const Message = ({ message }) => {
    const parsedMessage = JSON.parse(message);

    return (
        <div className="message">
            {parsedMessage.body}
        </div>
    );
};

export default Message;
