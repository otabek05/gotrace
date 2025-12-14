
export const WELL_KNOWN_SERVICES = {
  HTTP: 80,
  HTTPS: 443,
  DNS: 53,
  SSH: 22,
  FTP: 21,
  SMTP: 25,
  POP3: 110,
  IMAP: 143,
  DHCP: 67,
} as const;

export type ServiceName = keyof typeof WELL_KNOWN_SERVICES;


export enum ServiceFilterMode {
  All = "ALL",
  WellKnown = "WELL_KNOWN",
  Custom = "CUSTOM",
}


