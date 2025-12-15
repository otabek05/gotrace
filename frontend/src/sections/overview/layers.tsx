import {
  Box,
  Typography,
  Collapse,
  Divider,
} from "@mui/material";
import { useState } from "react";
import { ParsedPacket } from "src/types/ws_receiving";
import { Iconify } from "src/components/iconify";

interface PacketItemProps {
  packet: ParsedPacket;
  index: number;
}

const LayerItem = ({ label, data }: { label: string; data: any }) => {
  const [open, setOpen] = useState(false);
  const isObject = typeof data === "object" && data !== null;
  const directionIcon = open ? "material-symbols:arrow-downward" : "material-symbols:arrow-upward";

  return (
    <Box sx={{ width: "100%", pl: 2 }}>
      <Box
        onClick={() => setOpen(!open)}
        sx={{
          display: "flex",
          justifyContent: "space-between",
          alignItems: "center",
          cursor: "pointer",
          color: "inherit",
          py: 0.5,
        }}
      >

        {isObject && (
          <>
            <Typography variant="body2" sx={{ fontWeight: "bold" }}>
              {label.toUpperCase()}
            </Typography>
            <Iconify icon={directionIcon} />
          </>
        )}
      </Box>

      {isObject && (
        <Collapse in={open} timeout="auto" unmountOnExit>
          {Object.entries(data).map(([key, value]) => (
            <LayerItem key={key} label={key} data={value} />
          ))}
        </Collapse>
      )}

      {/* When the data is not an object, we render the value */}
      {!isObject && (
        <Typography sx={{ fontFamily: "monospace", fontSize: 12 }}>
          {`${label.toUpperCase()}: ${String(data).toUpperCase()}`}
        </Typography>
      )}
    </Box>
  );
};




export const PacketItem = ({ packet, index }: PacketItemProps) => {
  const [open, setOpen] = useState(false);

  const isIncomingReq = packet.direction === "incoming"

  const directionIcon = isIncomingReq ? 'material-symbols:arrow-downward' : 'material-symbols:arrow-upward';
  const backgroundColor = isIncomingReq ? '#3498db' : '#2ecc71';

    const getEthernetDetails = ()=>{
      if (packet.ethernet)
       return (
        <Box sx={{ display: 'flex', flexDirection: 'column', pr: 2 }}>
          <Typography sx={{ fontFamily: 'monospace', fontSize: 12 }}>
            <strong>SRC MAC:</strong> {packet.ethernet.src_mac.toUpperCase()}
          </Typography>
          <Typography sx={{ fontFamily: 'monospace', fontSize: 12 }}>
            <strong>DST MAC:</strong> {packet.ethernet.dst_mac.toUpperCase()}
          </Typography>
        </Box>
      );
    }

  const getIpDetails = () => {
    if (packet.ipv4) {
      return (
        <Box sx={{ display: 'flex', flexDirection: 'column', pr: 2 }}>
          <Typography sx={{ fontFamily: 'monospace', fontSize: 12 }}>
            <strong>SRC IP:</strong> {packet.ipv4.src_ip.toUpperCase()}
          </Typography>
          <Typography sx={{ fontFamily: 'monospace', fontSize: 12 }}>
            <strong>DST IP:</strong> {packet.ipv4.dst_ip.toUpperCase()}
          </Typography>
        </Box>
      );
    }
    if (packet.ipv6) {
      return (
        <Box sx={{ display: 'flex', flexDirection: 'column', pr: 2 }}>
          <Typography sx={{ fontFamily: 'monospace', fontSize: 12 }}>
            <strong>SRC IP:</strong> {packet.ipv6.src_ip.toUpperCase()}
          </Typography>
          <Typography sx={{ fontFamily: 'monospace', fontSize: 12 }}>
            <strong>DST IP:</strong> {packet.ipv6.dst_ip.toUpperCase()}
          </Typography>
        </Box>
      );
    }
    return null;
  };

  const getTcpUdpDetails = () => {
    if (packet.tcp) {
      return (
        <Box sx={{ display: 'flex', flexDirection: 'column', pr: 2 }}>
          <Typography sx={{ fontFamily: 'monospace', fontSize: 12 }}>
            <strong>SRC PORT:</strong> {packet.tcp.src_port}
          </Typography>
          <Typography sx={{ fontFamily: 'monospace', fontSize: 12 }}>
            <strong>DST PORT:</strong> {packet.tcp.dst_port}
          </Typography>
        </Box>
      );
    }
    if (packet.udp) {
      return (
        <Box sx={{ display: 'flex', flexDirection: 'column', pr: 2 }}>
          <Typography sx={{ fontFamily: 'monospace', fontSize: 12 }}>
            <strong>SRC Port:</strong> {packet.udp.src_port}
          </Typography>
          <Typography sx={{ fontFamily: 'monospace', fontSize: 12 }}>
            <strong>DST Port:</strong> {packet.udp.dst_port}
          </Typography>
        </Box>
      );
    }

    return null;
  };

  return (
    <Box sx={{ width: '100%', mb: 1 }}>
      <Box
        sx={{
          display: 'flex',
          justifyContent: 'start',
          alignItems: 'center',
          padding: '8px 16px',
          backgroundColor: '#1f2327',
          borderRadius: '4px',
          cursor: 'pointer',
          border: '1px solid #30363d',
          gap: 1
        }}
        onClick={() => setOpen(!open)}
      >
        <Typography
          sx={{
            fontFamily: 'monospace',
            fontSize: 12,
            width: '5%',
            color: 'gray',
            borderRightWidth: 1,
            borderRightColor: 'white',
            textAlign: 'center'
          }}
        >
          #{index + 1}
        </Typography>
        <Iconify
          icon={directionIcon}
          width={20}
          height={20}
          color={backgroundColor}
        />
        <Typography sx={{ fontFamily: 'monospace', fontSize: 12, fontWeight: 'bold', width: '12%' }}>
          {packet.timestamp}
        </Typography>
        {getEthernetDetails()}
        {getIpDetails()}
        {getTcpUdpDetails()}
      </Box>

      <Collapse in={open} timeout="auto" unmountOnExit>
        <Box sx={{ pl: 2, pt: 1 }}>
          {packet.ethernet && <LayerItem label="ETHERNET" data={packet.ethernet} />}
          {packet.ipv4 && <LayerItem label="IPv4" data={packet.ipv4} />}
          {packet.ipv6 && <LayerItem label="IPv6" data={packet.ipv6} />}
          {packet.tcp && <LayerItem label="TCP" data={packet.tcp} />}
          {packet.udp && <LayerItem label="UDP" data={packet.udp} />}
          {packet.icmp && <LayerItem label="ICMP" data={packet.icmp} />}
          {packet.arp && <LayerItem label="ARP" data={packet.arp} />}
          {packet.dns && <LayerItem label="DNS" data={packet.dns} />}
          {packet.http && <LayerItem label="HTTP" data={packet.http} />}
          {packet.https && <LayerItem label="HTTPS" data={packet.https} />}
          {packet.dhcp && <LayerItem label="DHCP" data={packet.dhcp} />}
          {packet.application && <LayerItem label={packet.application.protocol} data={packet.application.data} />}
        </Box>
      </Collapse>

      <Divider sx={{ my: 0.5, bgcolor: '#30363d' }} />
    </Box>
  );
};