import { useEffect, useState } from "react";
import {
  getInterfaces,
  startScan,
  stopScan,
   type NetworkInterface,
} from "../api/scannerApi";
import { useWebSocket } from "../hooks/useWebSocket";

export default function Dashboard() {
  const [interfaces, setInterfaces] = useState<NetworkInterface[]>([]);
  const [selectedIf, setSelectedIf] = useState("");
  const { messages } = useWebSocket("ws://localhost:8081/ws");

  useEffect(() => {
    loadInterfaces();
  }, []);

  const loadInterfaces = async () => {
    const data = await getInterfaces();
    setInterfaces(data)
    console.log(data)
  };

  const handleStart = async () => {
    await startScan({
      interface: selectedIf,
      snaplen: 65535,
      promisc: true,
      filter: "",
    });
    alert("Scan started");
  };

  const handleStop = async () => {
    await stopScan();
    alert("Scan stopped");
  };

  return (
    <div style={{ padding: 20 }}>
      <h2>Network Scanner</h2>

      <select
        value={selectedIf}
        onChange={(e) => setSelectedIf(e.target.value)}
      >
        <option value="">Select Interface</option>
        {interfaces.map((i) => (
          <option key={i.name} value={i.name}>
            {i.name} ({i.addresses?.[0]})
          </option>
        ))}
      </select>

      <div style={{ marginTop: 10 }}>
        <button onClick={handleStart}>Start</button>
        <button onClick={handleStop} style={{ marginLeft: 10 }}>
          Stop
        </button>
      </div>

      <h3>Live Packets</h3>
      <div
        style={{
          height: 400,
          overflow: "auto",
          background: "#111",
          color: "#0f0",
          padding: 10,
        }}
      >
        {messages.map((msg, i) => (
          <pre key={i}>{JSON.stringify(msg, null, 2)}</pre>
        ))}
      </div>
    </div>
  );
}
