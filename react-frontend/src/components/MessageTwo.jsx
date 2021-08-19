import React from "react";
import uuid from "react-uuid";

const MessageTwo = (props) => {
  let messageData = JSON.parse(props.message);
  let keyid = props.uk;
  return (
    <div className="message">
      {keyid}: {messageData.body}
    </div>
  );
};

export default MessageTwo;
