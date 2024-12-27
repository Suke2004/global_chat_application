import React, { Component } from "react";
import './Message.css';

class Message extends Component{
    constructor(props){
        super(props); //When we are importing props from other parent functions we have to use super
        let temp = JSON.parse(this.props.Message);
        this.state = {
            message:temp
        }
    }

    render(){
        return(
            <div className="Message">
                {this.state.message.body}
            </div>
        );
    }
}

export default Message;