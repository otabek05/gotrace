import React, { useEffect, useState, useRef } from "react";

function App() {
  const [messages, setMessages] = useState<string[]>([]);
  const [input, setInput] = useState<string>(""); // 입력 상태
  const ws = useRef<WebSocket | null>(null);

  useEffect(() => {
    // WebSocket 연결
    ws.current = new WebSocket("ws://localhost:8080/ws");

    ws.current.onopen = () => {
      console.log("WebSocket connected");
    };

    ws.current.onmessage = (event) => {
      console.log("Received:", event.data);
      setMessages((prev) => [...prev, event.data]);
    };

    ws.current.onclose = () => {
      console.log("WebSocket disconnected");
    };

    return () => {
      ws.current?.close();
    };
  }, []);

  const sendMessage = () => {
    if (ws.current && ws.current.readyState === WebSocket.OPEN && input.trim() !== "") {
      ws.current.send(input);
      setInput(""); // 전송 후 input 초기화
    }
  };

  const handleInputChange = (e: React.ChangeEvent<HTMLInputElement>) => {
    setInput(e.target.value);
  };

  const handleKeyPress = (e: React.KeyboardEvent<HTMLInputElement>) => {
    if (e.key === "Enter") {
      sendMessage();
    }
  };

  return (
    <div style={{ padding: "20px" }}>
      <h1>WebSocket React Example</h1>

      <input
        type="text"
        value={input}
        onChange={handleInputChange}
        onKeyPress={handleKeyPress}
        placeholder="Type a message..."
        style={{ width: "300px", marginRight: "10px" }}
      />
      <button onClick={sendMessage}>Send Message</button>

      <div style={{ marginTop: "20px" }}>
        <h2>Incoming Messages:</h2>
        {messages.map((msg, idx) => (
          <div key={idx}>{msg}</div>
        ))}
      </div>
    </div>
  );
}

export default App;
