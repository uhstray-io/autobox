# autobox

GO, Python, NixOS, and Docker driven Orchestration Box for Setting up Business-as-Code, Managing local Automation and Setting up Custom Network Configurations.

## Target Local Systems

- pFSense Gateway/Firewall
- NixOS Linux Machines
- Docker Containers

## Requirements

- Use GO to develop Webserver Resources
- Use Python to develop Scripting and Automation
- Use GO to develop for Networking Resources
- Use Rust to develop GUI Resources
- Use NixOS to develop OS & Networking Configurations and Package Mangement (using Nix Flakes for reproducibility)
- Use Docker to develop Container Configurations
- Use PXE to access devices on LAN
- Setup Automated Backups of deployed systems with a selected cloud provider (Digital Ocean, Vercel)
- Use Cloudflare for DNS, SSL, and CDN
- Use Git, NixOS (flakes), and Docker for Version Control
- Use GitHub for Code Hosting
- Use GitHub Actions for CI/CD

## Features

- Test Local Firewall Rules & Configurations
- Test Local NixOS Configurations
- Test Local Docker Container Configurations
- Test Local Network Routes & Configurations
- Deploy Business-as-Code Applications
  - Application List
    - Wisbot for Search, LLM Orchestration, Event Management and Automation
    - WisLLM for Agentic LLM Management
      - Ollama
      - LangChain
      - vLLM
    - Databases
      - PostgreSQL
      - Postgres Vector DB
      - Prometheus
      - Model Context Protocol
    - Website Resources
      - Wiki.js for Documentation
    - Nextcloud for Local Cloud and File Management
    - NocoDB for Project Management
    - Grafana for Monitoring and Visualization
    - Authentik for Authentication and Authorization
    - WireGuard for VPN
    - Pi-hole for Ad Blocking
    - Home Assistant for IoT Device Automation

## Autobox Admin Web Application Architecture

```mermaid
graph TB
    subgraph "Frontend - Rust/WASM"
        A[Rust Web UI<br/>Yew Framework]
        B[Admin Dashboard]
        C[Component Testing UI]
        D[Deployment Manager]
        E[Real-time Monitoring]
    end

    subgraph "Backend - GO API"
        F[GO REST API Server]
        G[WebSocket Handler]
        H[Component Manager]
        I[Test Orchestrator]
        J[Deployment Engine]
    end

    subgraph "Core Services"
        K[NixOS Config Manager]
        L[Docker Orchestrator]
        M[Network Controller]
        N[PXE Boot Server]
    end

    subgraph "Target Systems"
        O[pFSense Firewall]
        P[NixOS Machines]
        Q[Docker Containers]
    end

    subgraph "External Services"
        R[GitHub Actions CI/CD]
        S[Cloudflare DNS/SSL]
        T[Cloud Backup<br/>Digital Ocean/Vercel]
    end

    A --> B
    A --> C
    A --> D
    A --> E

    B --> F
    C --> F
    D --> F
    E --> G

    F --> H
    F --> I
    F --> J
    G --> H

    H --> K
    H --> L
    H --> M
    I --> K
    I --> L
    J --> K
    J --> L
    J --> N

    K --> P
    L --> Q
    M --> O
    N --> O
    N --> P

    J --> R
    J --> S
    J --> T

    style A fill:#e34c26
    style F fill:#00ADD8
    style K fill:#5277C3
    style L fill:#2496ED
```

### Admin Application Flow

```mermaid
sequenceDiagram
    participant Admin as Administrator
    participant UI as Rust Web UI
    participant API as GO API Server
    participant Engine as Deployment Engine
    participant Target as Target Systems

    Admin->>UI: Access Dashboard
    UI->>API: GET /api/components
    API-->>UI: Component Status List
    
    Admin->>UI: Select Components to Test
    UI->>API: POST /api/test
    API->>Engine: Initialize Test Suite
    Engine->>Target: Execute Tests
    Target-->>Engine: Test Results
    Engine-->>API: Test Status
    API-->>UI: Real-time Updates (WebSocket)
    
    Admin->>UI: Deploy Configuration
    UI->>API: POST /api/deploy
    API->>Engine: Start Deployment
    Engine->>Target: Deploy Components
    Target-->>Engine: Deployment Status
    Engine-->>API: Progress Updates
    API-->>UI: Live Progress (WebSocket)
```