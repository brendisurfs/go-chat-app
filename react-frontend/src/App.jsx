import React, { Component } from "react";

import { connect, sendMSG } from "../api";

// individual components
import Header from "../src/components/Header";
import ChatHistory from "./components/ChatHistory";
import ChatInput from "./components/ChatInput.jsx";

class App extends Component {
  constructor(props) {
    super(props);
    this.state = {
      chatHistory: [],
    };
  }

  componentDidMount() {
    connect((msg) => {
      this.setState((prevState) => ({
        chatHistory: [...this.state.chatHistory, msg],
      }));
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
