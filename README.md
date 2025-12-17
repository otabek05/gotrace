# GoTracer

GoTracer is a network analysis and inspection tool inspired by Wireshark, designed to run with a **web-based UI** while performing low-level network operations in **Golang**.  
The backend communicates with the frontend via **WebSocket**, enabling real-time packet streaming, statistics, and analysis directly in the browser.

---

## Overview

GoTracer focuses on **deep network visibility** and **real-time inspection**, making it suitable for developers, network engineers, and security learners who want to understand how data flows through a system.

Key goals:
- Real-time packet tracing
- Clear visualization of incoming and outgoing traffic
- Layer-by-layer protocol inspection
- Network scanning and diagnostics
- Browser-based usability without native GUI dependencies

---

## Core Features

### Packet Tracer (Wireshark-inspired)

- Real-time packet capture using raw sockets
- Live streaming of packets to the browser via WebSocket
- Direction detection (Incoming / Outgoing)
- Timestamped packet logs
- Interface-based capture

#### Supported TCP/IP Layers
- Link Layer (Ethernet)
- Network Layer (IPv4)
- Transport Layer (TCP / UDP)
- Application Layer (DNS, HTTP, Raw payload)
- Basic encryption awareness (TLS traffic detection)

This allows inspection of **almost all 5 layers of the TCP/IP model**.

---

## Traffic Analysis

- Monitor background traffic without interrupting applications
- Identify encrypted vs unencrypted payloads
- Track packet sizes and directions
- Live bandwidth usage (bytes in / bytes out)
- Filter traffic by:
  - Interface
  - Direction
  - IP address
  - Port / service

---

## LAN Device Discovery (In Progress)

Features under development:
- Discover devices on the same local network
- Detect IP and MAC addresses
- Identify device vendors (via MAC OUI)
- Scan open TCP ports
- Test reachability of discovered devices
- Display results in real-time in the web UI

Planned techniques:
- ARP scanning
- ICMP probing
- TCP connect scanning

---

## WiFi Analyzer (In Progress)

Planned functionality:
- Analyze nearby WiFi networks
- Display SSID, signal strength, channel usage
- Detect currently connected WiFi network
- Identify channel congestion
- Basic WiFi security checks (encryption type visibility)

This module aims to provide visibility similar to desktop WiFi analyzer tools, but inside a web interface.

---

## Network Utilities (Planned)

GoTracer will integrate commonly used network tools into the web UI:

- Port scanning (similar to nmap)
- Host discovery
- Traceroute visualization
- Service detection
- Latency and reachability testing

All tools will:
- Run on the Go backend
- Stream results live via WebSocket
- Be controllable from the browser

---

## Architecture

- **Backend**: Golang
  - Packet capture (pcap / raw sockets)
  - Network scanning
  - Protocol parsing
  - WebSocket server

- **Frontend**: Web (React)
  - Live packet visualization
  - Filters and controls
  - Graphs and statistics
  - Interactive network tools

---

## Use Cases

- Learning how network protocols work
- Debugging encrypted vs plaintext traffic
- Monitoring background network activity
- Inspecting LAN devices and services
- Lightweight alternative to desktop packet analyzers
- Educational network security experimentation

---

## Status

- Packet capture: implemented
- WebSocket streaming: implemented
- UI visualization: active development
- LAN discovery: in progress
- WiFi analyzer: in progress
- Network utilities: planned

---

## Vision

GoTracer aims to bring powerful network inspection tools traditionally limited to desktop applications into a **browser-first experience**, while keeping performance and low-level access through Go.
