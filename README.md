# 🚀 Go Log Metrics Pipeline


## 🧠 Learn Go Concurrency: The Pipeline Pattern

This project demonstrates the **pipeline concurrency pattern**, one of the most powerful and idiomatic patterns in Go for streaming data processing.

### 🧩 What is the Pipeline Pattern?

In Go, a pipeline is a series of **stages** connected by **channels**, where each stage runs in its own **goroutine** and performs a transformation or action on the data.

```
┌────────┐   ┌────────┐   ┌────────┐   ┌────────────┐
│ Tailer │ → │ Parser │ → │ Filter │ → │ Prometheus │
└────────┘   └────────┘   └────────┘   └────────────┘
    log        JSON         error         metrics
  strings      logs          only          exposed
```

Each stage is fully concurrent and communicates via channels.

---

### 💡 Why Use It?

✅ **Decoupling:** Each stage does one thing — easy to test, debug, and extend  
✅ **Scalability:** Each stage runs independently and can be scaled if needed  
✅ **Streaming:** Perfect for handling unbounded streams of data (like logs)  
✅ **Backpressure Friendly:** Naturally supports blocking when downstream is slow  
✅ **CPU Efficiency:** Leverages goroutines, not OS threads — lightweight & fast

---

### 👨‍💻 Where It's Used in This Project

| Stage     | Description                          |
|-----------|--------------------------------------|
| `tailer`  | Emits new log lines periodically     |
| `parser`  | Parses JSON strings into objects     |
| `filter`  | Filters logs by level (`error`, etc.)|
| `metrics` | Counts and exports metrics           |

---

This pattern is widely used in real-world systems like:
- Log processors
- Stream transformers
- ETL pipelines
- Event dispatchers

## 🔧 Features
- 🧵 Go channels & goroutines
- 🪵 Simulated log streaming
- 📦 JSON parsing + filtering
- 📊 Prometheus metrics export

## ▶️ Quickstart

### 1. Install dependencies
```bash
go mod tidy
```

## 🧭 Architecture Diagram

This project uses a concurrency pipeline in Golang to process logs and expose metrics via Prometheus.
![Pipeline Architecture](./flow.png)

