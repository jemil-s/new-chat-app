import React from "react";
import styles from "./Message.module.css";

const Message = ({ message }) => {
  console.log(message, "message");

  const text = message;

  console.log(text);

  return (
    <div>
      {text.user ? (
        <div className={styles.message}>
          <div className={styles.user}>{text.user}: </div>
          <div className={styles.text}>{text.body}</div>
        </div>
      ) : (
        <div>{text.body}</div>
      )}
    </div>
  );
};

export default Message;
