# Intelligent Agent System: Kafka Producer

This module is the **Kafka producer** for the _Intelligent Agent System_ layer of the **CReATE Fleet Management System.** It is designed to run on a **System on Module (SoM)** embedded in a vehicle, where it collects sensor and telemetry data and publishes it to Kafka. Downstream services can then consume this data for further processing and analytics.

### Current Features
- Publishes data to Kafka topics.
- Supports Protocol Buffers.
- Integrates with Schema Registry.

### Setup

**1. Clone the repository**
```
git clone <the repository url>
cd intelligent-agent-system
```

**2. Compile the Protocol Buffers**
```
protoc \
  --go_out=. --go_opt=module=github.com/jojohimawan/intelligent-agent-system \
  --go-grpc_out=. --go-grpc_opt=module=github.com/jojohimawan/intelligent-agent-system \
  api/location.proto
```

**3. Verify Kafka and Schema Registry connectivity**
Ensure that:
- Kafka broker is running and accessible.
- Schema Registry is up and reachable.

**4. Run the server**
```
go run cmd/server/main.go
```

### Requirements
- Go >= 1.24
- Kafka Broker
- Confluent Schema Registry

<br>

> _Authored by Jordan Himawan._
>
> _Cyber Security Research Group, C304 - D4 Building._
> _Politeknik Elektronika Negeri Surabaya._