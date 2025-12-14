import { ApplicationLayer, NetworkLayer, TrafficOptions, TransportLayer } from "src/types/ws_sending";

export const trafficOptions = [
    { value: TrafficOptions.Incoming, label: 'Incoming' },
    { value: TrafficOptions.Outgoing, label: 'Outgoing' },
    { value: TrafficOptions.Both, label: 'Both' },
];

export const networkLayerOptions = [
    { value: NetworkLayer.IPv4, label: 'IPv4' },
    { value: NetworkLayer.IPv6, label: 'IPv6' },
    { value: NetworkLayer.ICMP, label: 'ICMP' },
    { value: NetworkLayer.Unknown, label: 'Any' },
];


export const transportLayerOptions = [
    { value: TransportLayer.TCP, label: 'TCP' },
    { value: TransportLayer.UDP, label: 'UDP' },
    { value: TransportLayer.Unknown, label: 'Any' },
];


export const applicationLayerOptions = [
    { value: ApplicationLayer.Any, label: "Any" },
    { value: ApplicationLayer.WellKnown, label: "WellKnown" },
    { value: ApplicationLayer.Custom, label: "Custom" }
]


export const wellKnownServicesOptions = [
  // Web
  { value: '80', label: 'HTTP' },
  { value: '443', label: 'HTTPS' },

  // Mail
  { value: '25', label: 'SMTP' },
  { value: '110', label: 'POP3' },
  { value: '143', label: 'IMAP' },
  { value: '587', label: 'SMTP (Submission)' },
  { value: '465', label: 'SMTP (SMTPS)' },

  // Databases
  { value: '3306', label: 'MySQL' },
  { value: '5432', label: 'PostgreSQL' },
  { value: '1433', label: 'Microsoft SQL Server' },
  { value: '1521', label: 'Oracle DB' },
  { value: '27017', label: 'MongoDB' },
  { value: '27018', label: 'MongoDB (Secondary)' },
  { value: '6379', label: 'Redis' },
  { value: '5984', label: 'CouchDB' },

  // Messaging
  { value: '1883', label: 'MQTT' },
  { value: '8883', label: 'MQTT (TLS)' },
  { value: '9092', label: 'Kafka' },
  { value: '2181', label: 'Zookeeper' },

  // Remote Access & Administration
  { value: '22', label: 'SSH' },
  { value: '23', label: 'Telnet' },
  { value: '3389', label: 'RDP' },
  { value: '5900', label: 'VNC' },

  // DNS & Directory
  { value: '53', label: 'DNS' },
  { value: '389', label: 'LDAP' },
  { value: '636', label: 'LDAPS' },

  // File & SMB
  { value: '21', label: 'FTP' },
  { value: '445', label: 'SMB / CIFS' },
];


 export interface Service {
  value: string;
  label: string;
}