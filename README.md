# Exinity | Symbol Data Provider

A gRPC-based real-time symbol data service that streams 1-minute candlestick (OHLCV) data by aggregating live trade data from Binance.

---

## 📦 Project Layout

### ⏱ Milestones (estimated hours)

| Task                                            | Duration |
|-------------------------------------------------|----------|
| Requirement analysis & architecture design      | 1h       |
| Core architecture implementation & edge cases   | 2h       |
| Service development                             | 5h       |
| Testing & documentation                         | 0h       |
| Terraform & Kubernetes integration              | 0.5h     |
| Deployment & hand-off                           | 0.5h     |

---

### ✅ Completed Tasks

- [x] Requirements gathering
- [x] Project layout & Git repository initialization
- [x] System design and architecture skeleton
- [x] Docker integration & database setup
- [x] PostgreSQL schema design
- [x] Internal services implementation
- [x] Worker (processor) implementation
- [x] gRPC server implementation
- [ ] Unit test coverage
- [x] Terraform & Kubernetes manifests
- [x] Project documentation
- [x] Project delivery to Exinity team

---

### ⚠️ Known Limitations / Missing parts

Due to time constraints:

- Unit tests could not be implemented due to timing.
- Integration, benchmark, contract, and e2e tests were omitted
- Simpler design choices were made for clarity and speed
- Used the *sugared* version of `zap` logger for ease of use — performance tuning was not prioritized
- I could not complete the terraform part, so it is not working; files are created without check & run.
- K8S configurations are done according to the `Minikube` on the local development environment.

---

## 🚀 Getting Started

### 📂 Access the Code

The source code is available at:  
👉 [GitHub Repository](https://github.com/AkyurekDogan/exinity-task)

---

### 🧰 Prerequisites

- Docker installed on your system
- (For developer mode) Go v1.23.2+
- (Optional) VS Code for IDE support
- If you want to run k8s mode, you need to install `minikube` on local k8s provider.

---

## 🧪 Running the Project

### Step 1 – Setup

Ensure Docker is installed. Choose one of the following methods to run the project.

---

### Step 2 – Run Options

#### ✅ Option 1: Developer Mode (Recommended)

1. Install Go v1.23.2+
2. (Optional) Open in VS Code. The `.vscode/launch.json` config supports debugging.
3. Run:
   ```
    bash
    make local-setup
   ```
   This sets up:
   - A `.env` file (copied from `.env.dist`)
   - A local PostgreSQL database
   - The database schema via initialization scripts

4. DB will be accessible at:
   - Host: `localhost`
   - Port: `5432`
   - Username: `postgress`
   - Password: `mypassword123!`

5. Start the app:
   ```
    bash
    make run
   ```
   Or use VS Code’s debug tools.

> The gRPC server will be running at `localhost:50051`

---

#### 🐳 Option 2: Docker Compose

1. Run with Make:
   ```
    bash
    make compose-up
   ```

2. Or run manually:
   - Update the `database.host` field in `config.yml` to:
     ```
      yaml
      host: go-exinity-task-postgress
     ```
   - Run:
     ```
      bash
      docker-compose -f docker-compose.yml up
     ```

> The gRPC server will be running at `localhost:50051`

---

#### 🐳 Option 3: K8S 

Before running the following commands you need to be sure that following items are done properly 
- API docker image is on the local repository otherwise please use `.dockerfile` to create the build using make file or docker commands.
- Minikube uses the custom acessibility logs for local access please consider minikube settings for `IP address` or `localhost` accesibility.

1. Run with Make:
   ```
    bash
    make kubernetes-setup
   ```

> The gRPC server will be running at `localhost:30051`
> If you want to access to database inside the k8s you can use the following command
   ```
   bash
   kubectl port-forward svc/postgres 5432:5432 -n exinity-task
   ```
---

### Step 3 – Test the gRPC API

You can test the running gRPC server using one of the following tools:

#### 🧪 Option 1: grpcurl (CLI)
```
 bash
 grpcurl -plaintext   -proto internal/app/proto/candle.proto   -d '{"symbols": ["BTCUSDT", "PEPEUSDT", "ETHUSDT"]}'   localhost:50051 candle.CandleService/SubscribeCandles
```

#### 🧪 Option 2: Postman

You can also use [Postman](https://www.postman.com/) with its built-in gRPC client for a more user-friendly experience.

---

## 🧱 Architecture Diagram

![diagram-export-4-9-2025-2_55_10-PM](https://github.com/user-attachments/assets/a28740bb-1a6e-42c1-8aab-1c2abeb7568b)

---

## 📸 Screenrecordings

https://github.com/user-attachments/assets/2d545a79-01c7-4634-b20d-6407de453403

https://github.com/user-attachments/assets/64b87be5-1dad-4d02-81a1-79a6046289ac

https://github.com/user-attachments/assets/d6f3c6d8-c1fe-4311-9f80-b4a198d9025e


## 🤝 Contributing

This is a sample project submitted to Exinity.

---

## 📬 Contact

For questions or feedback, feel free to open an issue or reach out via *akyurek.dogan.dgn@gmail.com*

---
