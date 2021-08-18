import React, { Component } from "react";

import { connect, sendMSG } from "../api";

// individual components
import Header from "../src/components/Header";
import ChatHistory from "./components/ChatHistory";
import ChatInput from "./components/ChatInput.jsx";
import Message from "./components/Message";

class App extends Component {
  constructor(props) {
    super(props);
    this.state = {
      chatHistory: [],
    };
  }

  componentDidMount() {
    connect((msg) => {
      console.log("new message");
      this.setState((prevState) => ({
        chatHistory: [...this.state.chatHistory, msg],
      }));
      console.log(this.state);
    });
  }

  send(event) {
    if (event.keyCode == 13) {
      sendMSG(event.target.value);
      event.target.value = "";
    }
  }

  render() {
    return (
      <div className="app">
        <Header />
        <ChatHistory chatHistory={this.state.chatHistory} />
        <ChatInput send={this.send} />
      </div>
    );
  }
}
export default App;
