import React from "react";

const ChatInput = (props) => {
  return (
    <div className="chatInput">
      <input type="text" onKeyDown={props.send} />
    </div>
  );
};

export default ChatInput;
