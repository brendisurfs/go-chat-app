import react, { component } from "react";

import { connect, sendmsg } from "./api";

export default class app extends component {
  constructor(props) {
    super(props);
    connect();
  }

  send() {
    console.log("hello");
    sendmsg("hello");
  }

  render() {
    return (
      <div classname="app">
        <button onclick={this.send}>click ya boy</button>
      </div>
    );
  }
}
