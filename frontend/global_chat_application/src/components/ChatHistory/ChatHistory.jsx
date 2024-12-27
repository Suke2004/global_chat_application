import React,{Component} from "react";
import './ChatHistory.css'
import Message from '../Message'

class ChatHistory extends Component{
    render(){
        const messages = this.props.ChatHistory.map(msg =>< Message key = {msg.timeStamp} message = {msg.data} />)
        console.log(messages)
        return(
            <div className="chatHistory">
                <h2>Chat History</h2>
                {messages}
            </div>
        );
    };
}

export default ChatHistory;