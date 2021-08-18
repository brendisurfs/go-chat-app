import React, { Component } from "react";

import { connect, sendMSG } from "../api";

// individual components
import Header from "../src/components/Header";
import ChatHistory from "./components/ChatHistory";

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

  send() {
    sendMSG("hello");
  }

  render() {
    return (
      <div className="app">
        <Header />
        <ChatHistory chatHistory={this.state.chatHistory} />
        <button onClick={this.send}>click ya boy</button>
      </div>
    );
  }
}
export default App;
