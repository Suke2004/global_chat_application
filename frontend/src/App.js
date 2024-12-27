import React, { useState, useEffect } from 'react';
import Header from './components/Header/Header';
import ChatInput from './components/ChatInput/ChatInput';
import ChatHistory from './components/ChatHistory/ChatHistory';
import { connect, sendMsg } from './api';
import './App.css';

const App = () => {
    const [chatHistory, setChatHistory] = useState([]);

    useEffect(() => {
        connect((msg) => {
            console.log("New message from user");
            setChatHistory((prevHistory) => [...prevHistory, msg]);
        });
    }, []); // Empty dependency array ensures this runs only once

    const handleSend = (event) => {
        if (event.keyCode === 13) { // Enter key
            sendMsg(event.target.value);
            event.target.value = '';
        }
    };

    return (
        <div className="App">
            <Header />
            <ChatHistory chatHistory={chatHistory} />
            <ChatInput send={handleSend} />
        </div>
    );
};

export default App;
