import {
  Box,
  Typography,
  Paper,
  Button,
  IconButton,

} from "@mui/material";
import { useEffect, useState } from "react";
import { useWebSocket } from "../../hook/WebSocketContext";
import { NetworkInterface } from "../../types/netInterface";
import { ApplicationLayer, ApplicationLayerList, NetworkLayer, NetworkLayerList, TrafficOptions, TrafficOptionsList, TransportLayer, TransportLayerList } from "../../types/wsTX";
import { Service, wellKnownServicesOptions } from "../../contstants/capture";
import { validatePort } from "../../utils/validatePort";
import { getIPFromDomain, getNetworkInterfaces } from "../../services/networkService";
import { NetSelector } from "./NetworkSelector";
import { CustomSelect } from "../../components/options/CustomSelect";
import { ServiceSelect } from "./ServiceSelector";
import { PacketItem } from "./Layers";
import SearchInputWithButton from "../../components/options/SearchInput";
import { Icon } from "@iconify/react";
import SpeedIndicator from "./SpeedIndicator";
import { isIPV4 } from "../../utils/validateIP";


export default function CaptureView() {
  const { connected, messages, send, clearMessages, stopCapturing, speed } = useWebSocket();
  const [interfaces, setInterfaces] = useState<NetworkInterface[]>([])
  const [selectedInterface, setSelectedInterface] = useState<NetworkInterface | null>(null)
  const [trafficOption, setTrafficOption] = useState<TrafficOptions>("any");
  const [networkLayer, setNetworkLayer] = useState<NetworkLayer>("any");
  const [transportLayer, setTransportLayer] = useState<TransportLayer>("any");
  const [applicationLayer, setApplicationLayer] = useState<ApplicationLayer>("any")
  const [selectedServices, setSelectedServices] = useState<Service[]>([])
  const [customInput, setCustomInput] = useState<string>("")
  const [domain, setDomain] = useState<string>("")
  const [selectedIP, setSelectedIP] = useState<string>("")


  useEffect(() => {
    fetchNetInterfaces()
  }, [])

  const handleApply = async () => {

    let selectedIPAddress = domain
    if (domain !== "" && !isIPV4(domain)) {
      const response = await getIPFromDomain(domain)
      if (!response.isSuccess()) {
        console.log(response)
        return 
      }
      selectedIPAddress = response.data ? response.data![0] : ""
    }

    send({
      type: "start_capturing",
      trafficOptions: trafficOption,
      networkLayer,
      transport: transportLayer,
      services: applicationLayer != "any" ? buildServices() : null,
      interface: selectedInterface!,
      ipv4: selectedIPAddress != "" ? selectedIPAddress : null
    });
  };


  const addService = (service: Service) => {
    if (selectedServices.length >= 10) {
      alert("Maximum of 10 application filters allowed.");
      return;
    }

    if (!validatePort(service.value))
      return


    setSelectedServices((prev) => {
      const exist = prev.some((s) => s.value === service.value)
      if (!exist)
        return [...prev, service]

      return prev
    })
  }


  const removeService = (service: Service) => {
    setSelectedServices((prev) => prev.filter((s) => s.value !== service.value))
  }

  const onSearch = () => {
    const newService: Service = {
      label: customInput,
      value: customInput
    }

    addService(newService)
    setCustomInput("")
  }

  const buildServices = (): string[] => {
    return selectedServices.map(service => service.value)
  }

  const fetchNetInterfaces = async () => {
    const response = await getNetworkInterfaces()
    if (response instanceof Error) {
      alert(response.message)
      return
    }

    setInterfaces(response?.data ?? [])
    setSelectedInterface(response!.data![0])
  }

  return (
    <Box sx={{ p: 2, display: "flex", flexDirection: "column" }}>
      <Box
        sx={{
          display: "flex",
          gap: 2,
          flexWrap: "wrap",
          position: "sticky",
          top: 0,
          zIndex: 10,
          pb: 2,
        }}
      >


        <NetSelector
          label="Network Interface"
          interfaces={interfaces}
          value={selectedInterface}
          onChange={setSelectedInterface}
        />

        <CustomSelect
          label="Traffic"
          options={TrafficOptionsList}
          value={trafficOption}
          onChange={(val) => {
            setTrafficOption(val as TrafficOptions)
          }}
        />


        <CustomSelect
          label="Network"
          options={NetworkLayerList}
          value={networkLayer}
          onChange={(val) => {
            setNetworkLayer(val as NetworkLayer)
          }}
        />

        <CustomSelect
          label="Transport"
          options={TransportLayerList}
          value={transportLayer}
          onChange={(val) => {
            setTransportLayer(val as TransportLayer)
          }}
        />

        <SearchInputWithButton
          placeholder="Filter By Domain or IP "
          onChange={(val: string) => { setDomain(val) }}
          value={domain} />


        <CustomSelect
          label="Application"
          options={ApplicationLayerList}
          value={applicationLayer}
          onChange={(val) => {
            setApplicationLayer(val)
          }}
        />


        {applicationLayer == "well-known" ? (
          <>
            <ServiceSelect
              selectedServices={selectedServices}
              addService={addService}
              options={wellKnownServicesOptions} />
          </>
        )
          : applicationLayer == "custom" ? (
            <>
              <SearchInputWithButton onChange={(val: string) => { setCustomInput(val) }} onSearch={onSearch} value={customInput} />
            </>
          ) : null}


        <Button variant="contained" onClick={handleApply}>
          Start
        </Button>

        <Button variant="contained" color="error" onClick={stopCapturing}>
          Stop
        </Button>

        <Button variant="outlined" onClick={clearMessages}>
          Clear
        </Button>
      </Box>

      <Box>
        {selectedServices.length > 0 && (
          <>
            <Box mt={2} display="flex" flexWrap="wrap" gap={1}>
              {selectedServices.map((service) => (
                <Box
                  key={service.value}
                  display="flex"
                  alignItems="center"
                  bgcolor="grey.200"
                  px={1}
                  py={0.5}
                  borderRadius={1}
                >
                  <Typography variant="body2">{service.label}</Typography>
                  <IconButton
                    size="small"
                    onClick={() => removeService(service)}
                    sx={{ ml: 0.5 }}
                  >
                    <Icon icon={"solar:close-square-bold"} />
                  </IconButton>
                </Box>
              ))}
            </Box>
          </>
        )}

      </Box>

      <Typography variant="body2" sx={{ mt: 1, opacity: 0.6 }}>
        WebSocket Status: {connected ? "🟢 Connected" : "🔴 Disconnected"}
      </Typography>

      {/* PACKETS */}
      <Paper
        sx={{
          p: 2,
          mt: 1,
          width: "40%",
          flexGrow: 1,
          maxHeight: "75vh",
          overflowY: "auto",
          bgcolor: "#0d1117",
          color: "#c9d1d9",
          borderRadius: 2,
          border: "1px solid #30363d",
          fontFamily: "monospace",
          "&::-webkit-scrollbar": { display: "none" },
        }}
      >
        {messages.length === 0 ? (
          <Typography sx={{ color: "#8b949e" }}>
            No packets yet...
          </Typography>
        ) : (
          messages.map((msg, idx) => (
            <PacketItem key={idx} packet={msg} index={idx} />
          ))
        )}
      </Paper>
    </Box>
  );
}
