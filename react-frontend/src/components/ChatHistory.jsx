import React, { Component } from "react";
import Message from "./Message";
import MessageTwo from "./MessageTwo";
import uuid from "react-uuid";
export class ChatHistory extends Component {
  render() {
    const messages = this.props.chatHistory.map((msg) => {
      let uniqueKey = uuid.apply();
      return <MessageTwo key={uniqueKey} message={msg.data} uk={uniqueKey} />;
    });

    return (
      <div className="chat-history">
        <h2>Chat History</h2>
        {messages}
      </div>
    );
  }
}

export default ChatHistory;
