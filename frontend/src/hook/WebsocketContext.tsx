import React, { createContext, useContext, useRef, useState, useCallback, useEffect } from "react";
import { InternetSpeed, ParsedPacket, WSReceiving } from "src/types/ws_receiving";
import { MessageType, WSIncomingMessage, WSOutgoingMessage } from "src/types/ws_sending";

interface WebSocketContextType {
  connected: boolean;
  messages: ParsedPacket[];
  connect: (url: string) => void;
  disconnect: () => void;
  send: (msg: WSOutgoingMessage) => void;
  startCapturing: () => void;
  stopCapturing: () => void;
  clearMessages: ()=> void;
}

const WebSocketContext = createContext<WebSocketContextType | undefined>(undefined);

export const WebSocketProvider = ({ children }: { children: React.ReactNode }) => {
  const ws = useRef<WebSocket | null>(null);
  const reconnectTimer = useRef<NodeJS.Timeout | null> (null)
  const reconnectAttempts = useRef(0)
  const shouldReconnect = useRef(true)

  const [connected, setConnected] = useState<boolean>(false);
  const [messages, setMessages] = useState<ParsedPacket[]>([]);
  const [speed, setSpeed] = useState<InternetSpeed | null>(null)
  const urlRef = useRef<string>("")

  const connect = useCallback((url: string) => {
    urlRef.current = url;
    shouldReconnect.current = true;

    if (ws.current) ws.current.close();
    ws.current = new WebSocket(url);

    ws.current.onopen = () => {
      setConnected(true);
      reconnectAttempts.current = 0
      console.log("WS connected")
    } 

    ws.current.onclose = () => {
      setConnected(false);
      ws.current = null
      if (shouldReconnect.current) {
        scheduleReconnect();
      }
    } 

    ws.current.onmessage = (event) => {
      try {
        const data: WSReceiving = JSON.parse(event.data);
        if (data.type == "packets" && data.packets) {
            setMessages((prev) => [...prev, data.packets!]);
            console.log(data)
        }else if (data.type == "speed") {

        }

      } catch {
        console.log("raw:", event.data);
      }
    };
  }, []);


    const scheduleReconnect = () => {
    if (reconnectTimer.current) return;

    const delay = Math.min(1000 * 2 ** reconnectAttempts.current, 10000);
    reconnectAttempts.current++;

    console.log(`WS reconnect in ${delay}ms`);

    reconnectTimer.current = setTimeout(() => {
      reconnectTimer.current = null;
      if (urlRef.current) {
        connect(urlRef.current);
      }
    }, delay);
  };

  const disconnect = () => {
    shouldReconnect.current = false;
    reconnectTimer.current && clearTimeout(reconnectTimer.current);
    reconnectTimer.current = null;
    ws.current?.close();
  };


  const send = (msg: WSOutgoingMessage) => ws.current?.send(JSON.stringify(msg));

  const startCapturing = () => send({ type: MessageType.StartCapturing });
  const stopCapturing = () => send({ type: MessageType.StopCapturing});

  const clearMessages = ()=> setMessages([])

  useEffect(() => {
    connect("ws://localhost:8080/ws");

    return () => {
      disconnect();
    };
  }, [connect]);

  return (
    <WebSocketContext.Provider
      value={{
        connected,
        messages,
        connect,
        disconnect,
        send,
        startCapturing,
        stopCapturing,
        clearMessages,
      }}
    >
      {children}
    </WebSocketContext.Provider>
  );
}; 

export const useWebSocket = () => {
  const ctx = useContext(WebSocketContext);
  if (!ctx) throw new Error("useWebSocket must be used inside WebSocketProvider");
  return ctx;
};
