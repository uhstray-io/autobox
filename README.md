# autobox

GO, Python, NixOS, and Docker driven Orchestration Box for Setting up Business-as-Code, Managing local Automation and Setting up Custom Network Configurations.

## Target Local Systems

- pFSense Gateway/Firewall
- NixOS Linux Servers
- Docker Containers

## Requirements

- Use GO to develop Webserver Resources and REST API
- Use Python to develop Scripting and Automation
- Use GO to develop for Dioxus to develop GUI Resources
- Use NixOS to develop OS & Networking Configurations and Package Mangement (using Nix Flakes for reproducibility)
- Use Docker to develop Container Configurations
- Use PXE to access devices on LAN
- Use Git, NixOS (flakes), and Docker for Version Control
- Use any Git repository for Code Hosting
- Use GitHub Actions for CI/CD

## Features

- Test Local Firewall Rules & Configurations
- Test Local NixOS Configurations
- Test Local Docker Container Configurations
- Test Local Network Routes & Configurations
- Deploy Business-as-Code Applications
  - Application List
    - Caddy for Reverse Proxy and SSL Certificate Management
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
- Setup Automated Backups of deployed systems with a selected cloud provider (Digital Ocean, Vercel)
- Use Cloudflare for DNS, SSL, and CDN

## Autobox Admin Web Application Architecture

```mermaid
graph TB

    subgraph "Frontend - Rust/WASM"
        A[Rust Web UI<br/>Dioxus Framework]
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
        K[Python Script Runner]
    end

    subgraph "Operations Stack"
        OO[Operations]
        R[Wisbot<br/>Search/LLM/Events]
        W[Observability & Security]
        Q[Caddy<br/>Reverse Proxy/SSL]
        AA[GitHub Actions CI/CD]
        subgraph "Observability & Security Stack"
            W1[Authentik]
            W2[WireGuard VPN]
            U4[Grafana]
            X[pfSense Firewall]
        end
        V[Cloud Services]
        subgraph "External Services"
            AB[Cloudflare<br/>DNS/SSL/CDN]
            AC[Cloud Backup<br/>Digital Ocean/Vercel]
        end
    end

    
    subgraph "Infrastructure Ecosystem"
        II[Infrastructure]
        subgraph "Infrastructure Services"
            L[NixOS Config Manager<br/>with Flakes]
            M[Docker Orchestrator]
            N[Network Controller]
            O[PXE Boot Server]
            P[Git Version Control]
        end
        subgraph "Infrastructure Systems"
            Y[NixOS Machines]
            Z[Docker Containers]
        end
    end

    subgraph "Data Ecosystem"
        T[Databases]
        subgraph "Database Stack"
            T1[PostgreSQL]
            T2[Postgres Vector DB]
            T3[Prometheus]
            T4[Model Context Protocol]
        end
    end

    subgraph "Business Resources"
        U[Web Resources]
        subgraph "Web/Cloud Stack"
            U1[Wiki.js - Documentation]
            U2[Nextcloud - Cloud Storage & Documents]
            U3[NocoDB - Project Management]
        end
    end

    subgraph "AI Ecosystem"
        S[WisLLM<br/>Agentic LLM]
        subgraph "LLM Stack"
            S1[Ollama]
            S2[LangChain]
            S3[vLLM]
        end
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
    F --> K
    G --> H
    G <--> I

    K --> I

    H <--> II
    I <--> II

    II --> L
    II --> M
    II --> N
    II --> O
    II --> P

    L --> Y
    M --> Z
    N --> X
    O --> X
    O --> Y

    J --> R
    J --> S
    J --> T
    J --> U
    J --> OO

    OO --> W
    OO --> Q
    OO --> V
    OO --> AA

    S --> S1
    S --> S2
    S --> S3

    T --> T1
    T --> T2
    T --> T3
    T --> T4

    U --> U1
    U --> U2
    U --> U3

    W --> W1
    W --> W2
    W --> U4

    Q --> AB
    V --> AC

    style A fill:#dea584
    style F fill:#00ADD8
    style L fill:#5277C3
    style M fill:#2496ED
    style K fill:#3776AB
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