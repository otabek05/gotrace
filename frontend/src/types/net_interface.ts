

export interface Address {
  IP: string;
  Netmask: string;
  Broadaddr: string;
  P2P: string;
}

export interface NetworkInterface {
  Name: string;
  Description: string;
  Flags: number;
  Addresses: Address[];
}

