# Findtree - A Distributed Model for Solving the Missing Phone Problem
A localized tracking backend designed to help locate lost or stolen phones on a university campus without relying on constant internet access.
Now I wanted to use Go and the encoding/binary but let's try gRPC.
I got inspired by the Bitcoin project on the Networking chapter of Mastering Bitcoin Chapter 2.
It's also my first gRPC project (nyehehehe) :)

## The Problem
Commercial tracking networks (like Apple's Find My) assume everyone has unmetered cellular data, constant power, and omnipresent OS coverage. In a Nigerian university context, students turn off mobile data to save data bundles, school Wi-Fi is unreliable, and constant background GPS tracking kills phone batteries that can't easily be recharged during power outages.

## A Solution?
This system moves the workload to an offline, opportunistic peer-to-peer network.

1. **Phones (Kotlin)** use BLE and local Wi-Fi to sniff nearby devices completely offline, storing sightings in a local SQLite DB.
2. When a student temporarily gets internet access (mobile data or building Wi-Fi), the app flushes those logged sightings to this **Go Backend**.
3. **The Go Backend** acts as a lightweight aggregator, indexing where phones were last seen using Geohashes and coordinating search alerts across campus building routers.

---

## Core Features (What it actually does)

* **gRPC Ingestion:** Handles incoming batch uploads from phones. We use gRPC streams instead of standard REST/JSON because raw bytes save massive amounts of mobile data for the students in a resource constrained environment.
* **Redis Spatial Indexing:** Incoming logs are converted into 6-character Geohashes and thrown into Redis Sorted Sets. Data automatically expires after 24 hours so the server's RAM doesn't bloat.
* **Gateway Clustering (`memberlist`):** If we deploy this on local machines across multiple campus buildings (e.g., Science Faculty, Main Gate), the Go instances form an internal gossip cluster to sync stolen device alerts instantly without a single heavy database bottleneck.
* **Anti-Stalking Cryptography:** To prevent people from using this system to stalk students, we don't broadcast plain text device IDs. The server generates a cryptographically blinded challenge. Only a phone that *actually* passed the missing device offline can solve it and upload an encrypted location update.
* **Network graph traversal:** This implies that the system should be able to propagate the missing devices that were reported stolen across all nodes in the network.


---

## Project Structure

```text
go-backend/
├── cmd/
│   ├── gateway/       # Runs on local campus building routers/machines
│   └── server/        # Runs on the central cloud/admin control plane
└── internal/
    ├── crypto/        # Ed25519 signature verification & target blinding
    ├── gossip/        # memberlist clustering config
    ├── rpc/           # Protobuf generated code & gRPC endpoints
    ├── spatial/       # Geohash and coordinate helpers
    └── storage/       # Redis connection and timeseries tracking

```

---

## Quick Start (Local Mock Cluster)

If you want to test how the building gateways talk to each other on your machine, clear out your terminal and run these:

**Node 1 (Science Faculty):**

```bash
PORT=8081 GOSSIP_PORT=7001 NODE_NAME="science" go run cmd/gateway/main.go

```

**Node 2 (Main Gate):**

```bash
PORT=8082 GOSSIP_PORT=7002 NODE_NAME="maingate" JOIN_PEER="127.0.0.1:7001" go run cmd/gateway/main.go

```

---

# Contributions
Please suggest a better name than Findtree or Odysseus and any code contributions would be really nice.