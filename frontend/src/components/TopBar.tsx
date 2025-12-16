import { Box, Typography } from "@mui/material";
import { styled } from "@mui/material/styles";
import { Icon } from "@iconify/react";
import { useNavigate, useLocation } from "react-router-dom";


const menuItems = [
  {
    label: "Packet Capture",
    path: "/capture",
    icon: "mdi:access-point-network",
    color: "#4caf50",
  },
  {
    label: "Port Scan",
    path: "/port-scan",
    icon: "mdi:lan-connect",
    color: "#2196f3",
  },

  {
    label: "LAN",
    path: "/port-scan",
    icon: "mdi:lan-connect",
    color: "#2196f3",
  },

  {
    label: "OS",
    path: "/port-scan",
    icon: "mdi:lan-connect",
    color: "#2196f3",
  },

  {
    label: "Wifi",
    path: "/nmap",
    icon: "mdi:radar",
    color: "#ff9800",
  },
  {
    label: "Other Tools",
    path: "/tools",
    icon: "mdi:tools",
    color: "#9c27b0",
  },
];


export default function TopBar() {
  const navigate = useNavigate();
  const location = useLocation();

  return (
    <TopBarRoot>
      <Title>GoTrace</Title>

      <Menu>
        {menuItems.map((item) => {
          const active = location.pathname === item.path;

          return (
            <ToolItem
              key={item.path}
              active={active}
              color={item.color}
              onClick={() => navigate(item.path)}
            >
              <ToolIcon icon={item.icon} />
              <ToolLabel>{item.label}</ToolLabel>
            </ToolItem>
          );
        })}
      </Menu>
    </TopBarRoot>
  );
}


const TopBarRoot = styled(Box)(() => ({
  position: "fixed",
  top: 0,
  left: 0,
  right: 0,
  height: 72,
  padding: "0 24px",
  display: "flex",
  alignItems: "center",
  gap: 32,
  backgroundColor: "#0f172a",
  borderBottom: "1px solid #1e293b",
  zIndex: 1200,
}));

const Title = styled(Typography)(() => ({
  color: "#e5e7eb",
  fontSize: 18,
  fontWeight: 600,
  letterSpacing: "0.5px",
}));

const Menu = styled(Box)(() => ({
  display: "flex",
  gap: 20,
}));

const ToolItem = styled(Box, {
  shouldForwardProp: (prop) =>
    prop !== "active" && prop !== "color",
})<{
  active: boolean;
  color: string;
}>(({ active, color }) => ({
  display: "flex",
  alignItems: "center",
  gap: 10,
  padding: "10px 16px",
  borderRadius: 12,
  cursor: "pointer",
  color: active ? color : "#cbd5f5",
  backgroundColor: active ? "rgba(255,255,255,0.12)" : "transparent",
  transform: active ? "scale(1.05)" : "scale(1)",
  transition:
    "transform 0.25s ease, background-color 0.25s ease, box-shadow 0.25s ease",
    /*
  boxShadow: active
    ? `0 0 12px ${color}55`
    : "none",
    */
  "&:hover": {
    backgroundColor: "rgba(255,255,255,0.1)",
    transform: "scale(1.08)",
  },
}));

const ToolIcon = styled(Icon)(() => ({
  fontSize: 28,
}));

const ToolLabel = styled(Typography)(() => ({
  fontSize: 14,
  fontWeight: 500,
}));
