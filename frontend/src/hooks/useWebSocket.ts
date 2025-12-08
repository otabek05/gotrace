import { useEffect, useRef, useState } from "react";


export function useWebSocket(url: string) {
  const wsRef = useRef<WebSocket | null>(null);
  const [messages, setMessages] = useState<any[]>([]);
  const reconnectTimer = useRef<number | null>(null);

  useEffect(() => {
    function connect() {
      wsRef.current = new WebSocket(url);

      wsRef.current.onmessage = (event) => {
        const data = JSON.parse(event.data);
        setMessages((prev) => [data, ...prev]);
      };

      wsRef.current.onclose = () => {
        console.log("WebSocket closed, reconnecting...");
        reconnectTimer.current = window.setTimeout(connect, 2000);
      };

      wsRef.current.onerror = () => {
        wsRef.current?.close();
      };
    }

    connect();

    return () => {
      if (reconnectTimer.current) {
        clearTimeout(reconnectTimer.current);
      }
      wsRef.current?.close();
    };
  }, [url]);

  return { messages };
}
