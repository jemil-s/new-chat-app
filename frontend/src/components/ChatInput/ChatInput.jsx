import React, { useState } from "react";
import styles from "./ChatInput.module.css";

const ChatInput = ({ send }) => {
  const [message, setMessage] = useState("");

  const sendMessage = () => {
    if (message.length > 0) {
      send(message);
      setMessage("");
    }
  };

  const onEnter = (e) => {
    if (e.keyCode === 13) {
      sendMessage();
    }
  };

  return (
    <div className={styles.container}>
      <div className={styles.wrapper}>
        <input
          onChange={(e) => setMessage(e.target.value)}
          value={message}
          className={styles.input}
          type="text"
          onKeyUp={onEnter}
        />
        <button className={styles.button} onClick={sendMessage}>
          send
        </button>
      </div>
    </div>
  );
};

export default ChatInput;
