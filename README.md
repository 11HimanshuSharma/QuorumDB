# QuorumDB

QuorumDB is a production-grade distributed read replica database system built in Go, designed to explore the internal architecture of modern distributed databases such as PostgreSQL, CockroachDB, Aurora, and Redis replication systems.

The project focuses on implementing core distributed systems and database engineering concepts from scratch, including:

* Write Ahead Logging (WAL)
* Crash Recovery
* Read Replication
* Replica Synchronization
* Snapshotting
* Replication Lag Tracking
* Distributed Consensus (Raft)
* Automatic Failover
* Leader Election
* Durable Storage
* Observability & Metrics
* Chaos Testing



---
# why i am doing

Modern distributed databases are extremely complex systems involving storage engines, consensus protocols, replication pipelines, recovery mechanisms, and distributed coordination.

QuorumDB is an fun and engineering-focused attempt to deeply understand and implement these systems from first principles using Go.

The project prioritizes:

* production-oriented architecture
* systems-level engineering
* distributed systems correctness
* performance and durability
* real-world fault tolerance patterns

---



---

# Project Goal

The goal of QuorumDB is not to build a toy key-value store, but to engineer a production-oriented distributed storage system capable of handling replication, durability, failover, and consistency challenges encountered in real-world infrastructure systems.

The project is being developed incrementally, following the same architectural evolution used by real database systems:

1. Durable Storage Engine
2. WAL-based Persistence
3. Replication Protocol
4. Read Replica Architecture
5. Snapshot & Recovery Systems
6. Consensus & Leader Election
7. Fault Tolerance & Failover
8. Observability & Performance Engineering

---

# Architecture Overview

```text
                ┌─────────────┐
                │   Client    │
                └──────┬──────┘
                       │
              ┌────────▼────────┐
              │ Query Router    │
              │ (Go Service)    │
              └──────┬──────────┘
         writes ─────┘ └───── reads
                   primary      replicas
                ┌─────────┐   ┌─────────┐
                │ Node A  │   │ Node B  │
                │ Leader  │   │ Replica │
                └────┬────┘   └────┬────┘
                     │             │
                     └────WAL──────┘
```

---

# Core Features

## Storage Engine

* In-memory memtable
* Durable WAL persistence
* Crash recovery via WAL replay
* Thread-safe concurrent operations

## Replication

* Primary-replica architecture
* WAL shipping
* Replica synchronization
* Replication lag tracking

## Distributed Systems

* Raft consensus
* Leader election
* Automatic failover
* Split-brain prevention

## Reliability

* Snapshotting
* WAL segmentation
* Corruption detection
* Checksums
* Recovery pipelines

## Observability

* Prometheus metrics
* p95/p99 latency tracking
* Structured logging
* Distributed tracing

---

# Tech Stack

* Go (Golang)
* gRPC / TCP Networking
* Hashicorp Raft
* Prometheus
* Grafana
* WAL-based storage engine



---
This project explores concepts used in:

* PostgreSQL WAL replication
* CockroachDB distributed consensus
* Redis replication
* Kafka append-only logs
* Aurora-style read replicas
* etcd leader election

---

# Disclaimer

QuorumDB is a learning and engineering project intended for understanding distributed systems and database internals. It is not production-ready software.
