import React, { useState } from "react";
import styles from "./Login.module.css";

const Login = ({ login }) => {
  const [name, setName] = useState("");

  const sendLogin = () => {
    if (name.length > 0) {
      login(name);
    }
  };

  const onEnter = (e) => {
    if (e.keyCode === 13) {
      sendLogin();
    }
  };

  return (
    <div className={styles.container}>
      <h2>Enter the username</h2>
      <input
        className={styles.input}
        placeholder="username"
        type="text"
        id="user"
        onChange={(e) => setName(e.target.value)}
        name="user"
        value={name}
        onKeyUp={onEnter}
      />
      <button className={styles.button} onClick={sendLogin}>
        Login
      </button>
    </div>
  );
};

export default Login;
