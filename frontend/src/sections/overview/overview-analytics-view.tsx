import {
  Box,
  Typography,
  Paper,
  Button,
  IconButton,

} from "@mui/material";
import { useEffect, useState } from "react";
import { useWebSocket } from "src/hook/WebsocketContext";
import {
  ApplicationLayer,
  MessageType,
  NetworkLayer,
  TrafficOptions,
  TransportLayer,
} from "src/types/ws_sending";
import { PacketItem } from "./layers";
import { CustomSelect } from "src/components/option/CustomSelect";
import { applicationLayerOptions, networkLayerOptions, Service, trafficOptions, transportLayerOptions, wellKnownServicesOptions } from "src/constansts/capture";
import { ServiceSelect } from "./serviceSelector";
import { Iconify } from "src/components/iconify";
import SearchInputWithButton from "src/components/option/SearchInput";
import { validatePort } from "src/utils/valid_port";
import { NetworkInterface } from "src/types/net_interface";
import { getNetworkInterfaces } from "src/services/networkService";
import { NetSelector } from "./netSelector";



export default function OverviewAnalyticsView() {
  const { connected, messages, send, clearMessages } = useWebSocket();
  const [interfaces, setInterfaces] = useState<NetworkInterface[]>([])
  const [selectedInterface, setSelectedInterface] = useState<NetworkInterface | null>(null)
  const [trafficOption, setTrafficOption] = useState<TrafficOptions>(TrafficOptions.Both);
  const [networkLayer, setNetworkLayer] = useState<NetworkLayer>(NetworkLayer.Unknown);
  const [transportLayer, setTransportLayer] = useState<TransportLayer>(TransportLayer.Unknown);
  const [applicationLayer, setApplicationLayer] = useState<ApplicationLayer>(ApplicationLayer.Any)
  const [selectedServices, setSelectedServices] = useState<Service[]>([])
  const [customInput, setCustomInput] = useState<string>("")

  useEffect(()=>{
    fetchNetInterfaces()
  },[])

  const handleApply = () => {
    send({
      type: MessageType.StartCapturing,
      trafficOptions: trafficOption,
      networkLayer,
      transport: transportLayer,
      services: buildServices(),
      interface: selectedInterface!
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

  const fetchNetInterfaces = async()=>{
    const response =  await  getNetworkInterfaces()
    if (response instanceof Error) {
      alert(response.message)
      return 
    }

    setInterfaces(response)
    setSelectedInterface(response[0])
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
          options={trafficOptions}
          value={trafficOption}
          onChange={(val) => {
            setTrafficOption(val as TrafficOptions)
          }}
        />


        <CustomSelect
          label="Network"
          options={networkLayerOptions}
          value={networkLayer}
          onChange={(val) => {
            setNetworkLayer(val as NetworkLayer)
          }}
        />

        <CustomSelect
          label="Transport"
          options={transportLayerOptions}
          value={transportLayer}
          onChange={(val) => {
            setTransportLayer(val as TransportLayer)
          }}
        />

        <CustomSelect
          label="Application"
          options={applicationLayerOptions}
          value={applicationLayer}
          onChange={(val) => {
            setApplicationLayer(val)
          }}
        />

        {applicationLayer === ApplicationLayer.WellKnown ? (
          <>
            <ServiceSelect
              selectedServices={selectedServices}
              addService={addService}
              options={wellKnownServicesOptions} />
          </>
        )
          : applicationLayer == ApplicationLayer.Custom ? (
            <>

              <SearchInputWithButton onChange={(val:string) => { setCustomInput(val) }} onSearch={onSearch} value={customInput} />
            </>
          ) : null}


        <Button variant="contained" onClick={handleApply}>
          Apply
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
                    <Iconify icon={"solar:close-square-bold"} />
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
          width:"40%",
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
