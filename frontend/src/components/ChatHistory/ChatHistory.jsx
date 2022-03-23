import React, { useRef } from "react";
import styles from "./ChatHistory.module.css";
import Message from "../Message/Message";
import { useEffect } from "react";
const ChatHistory = ({ messages }) => {
  console.log(messages, "messages");

  const ref = useRef();

  useEffect(() => {
    if (ref.current) {
      ref.current.scrollIntoView({ behavior: "smooth" });
    }
  }, [messages]);

  return (
    <div className={styles.container}>
      <h2 className={styles.title}>Chat history</h2>
      {messages.map((message, index) => (
        <div key={message.id + index}>
          <Message message={message} />
        </div>
      ))}
      <div ref={ref} />
    </div>
  );
};

export default ChatHistory;
