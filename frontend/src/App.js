import { useEffect, useState } from "react";
import "./App.css";
//import { sendMsg, connect } from "./api";
import Header from "./components/Header/Header";
import ChatHistory from "./components/ChatHistory/ChatHistory";
import ChatInput from "./components/ChatInput/ChatInput";
import { useCallback } from "react";
import Login from "./components/Login/Login";

function App() {
  const [chatHistoryMsg, setChatHistoryMsg] = useState([]);
  const [logged, setLogged] = useState("");
  const [socket, setSocket] = useState(null);
  const [query, setQuery] = useState("http://192.168.0.135:8080/allMessages");

  useEffect(() => {
    fetch(query, { method: "GET" })
      .then((res) => res.json())
      .then((res) => {
        if (res && res.length > 0) {
          setChatHistoryMsg((state) => [...state, ...res]);
        }
      })
      .catch((err) => console.log("err", err));
  }, [query]);

  const send = (message) => {
    console.log("send message");
    socket.send(message);
  };

  const enterChat = (name) => {
    setLogged(name);
  };

  useEffect(() => {
    if (logged && !socket) {
      const websocket = new WebSocket(
        `ws://192.168.0.135:8080/ws?user=${logged}`
      );

      setSocket(websocket);

      //setSendMsg(socket.send)

      websocket.addEventListener("open", () => {
        console.log(`${logged} Successfully Connected`);
      });

      websocket.onmessage = (msg) => {
        console.log(msg, "message from socket");
        const message = JSON.parse(msg.data);
        setChatHistoryMsg((state) => [...state, message]);
      };

      websocket.onclose = (event) => {
        console.log("Socket Closed Connection: ", event);
      };

      websocket.onerror = (error) => {
        console.log("Socket Error: ", error);
      };
    }
  }, [logged, socket]);

  return (
    <div className="App">
      <Header />
      <div>
        {logged ? (
          <div>
            <ChatHistory messages={chatHistoryMsg} />
            <ChatInput send={send} />
          </div>
        ) : (
          <Login login={enterChat} />
        )}
      </div>
    </div>
  );
}

export default App;
