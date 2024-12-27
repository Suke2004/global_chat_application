import React from "react";
import './ChatInput.css';

const ChatInput = ({ send }) => (
    <div className="chat-input">
        <input
            type="text"
            className="chat-input-field"
            onKeyDown={send}
            placeholder="Type something..."
            aria-label="Chat input"
        />
    </div>
);

export default ChatInput;
