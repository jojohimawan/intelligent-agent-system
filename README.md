# Intelligent Agent System: Kafka Producer

This module is the **Kafka producer** for the _Intelligent Agent System_ layer of the **CReATE Fleet Management System.** It is designed to run on a **System on Module (SoM)** embedded in a vehicle, where it collects sensor and telemetry data and publishes it to Kafka. Downstream services can then consume this data for further processing and analytics.

### Current Features
- Reads data from serial port and virtual CAN network interface.
- Parses NMEA sentence and OBD2 CAN frames.
- Publishes data to Kafka topics.
- Supports Protocol Buffers.
- Integrates with Schema Registry.
- Pipeline concurrency.

### Setup

**1. Clone the repository**
```bash
git clone <the repository url>
cd intelligent-agent-system
```

**2. Compile the Protocol Buffers**
```bash
protoc \
  --go_out=. --go_opt=module=github.com/jojohimawan/intelligent-agent-system \
  --go-grpc_out=. --go-grpc_opt=module=github.com/jojohimawan/intelligent-agent-system \
  api/location.proto
```

**3. Create a virtual CAN interface**
```bash
sudo modprobe vcan
sudo ip link add dev vcan0 type vcan
sudo ip link set up vcan0
```

**4. Verify Kafka and Schema Registry connectivity**
<br> Ensure that:
- Kafka broker is running and accessible.
- Schema Registry is up and reachable.


**5. Run the server**
```bash
go run cmd/server/main.go
```

**6. Test with sending CAN frame**
```bash
cansend vcan0 98DAF115#04410C1770000000
```

### Dev Notes
- Uses pipeline concurrency model with independent goroutines for reading, parsing, and publishing.
- Each stage logs recoverable errors without blocking the pipeline.
- Additional data sources or stages can be added.
- Graceful shutdown by context cancellation and Kafka/serial connection cleanup.

### Requirements
- Go >= 1.24
- Kafka Broker
- Confluent Schema Registry
- Serial Device (e.g. Arduino Uno, GPS module emitting NMEA sentences)
- Virtual CAN Network Interface (e.g. Linux's can-utils)

### Debts
- ~~Multi-topic publishing.~~
- Worker pool for Kafka publishing.
- Graceful shutdown.
- VCAN connection cleanup.
- VCAN connection error handling.
- Environment variables adjustment.

<br>

> _Authored by Jordan Himawan._
>
> _Cyber Security Research Group, C304 - D4 Building._ <br>
> _Politeknik Elektronika Negeri Surabaya._ <br>
>
> Last change: October 6th, 2025.