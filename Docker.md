# 🐳 Go Microservices + PostgreSQL (Docker Compose Setup)

This project demonstrates a simple **microservices architecture** using **Go (Golang)**, **PostgreSQL**, and **Docker Compose**.

It includes services like:
- `user-service`
- `order-service`
- `payment-service`
- `notification-service`
- `user-db` (PostgreSQL database)

---

## 🧱 Project Structure
├── docker-compose.yml
├── user-service/
│ ├── Dockerfile
│ ├── main.go
│ └── db/
│ └── connection.go
├── order-service/
├── payment-service/
└── notification-service/


---

## 🚀 Commands Used (with Explanation)

### 🔹 1. Build and Run All Services

```bash
docker compose up -d --build
```
Explanation:

docker compose → runs the services defined in docker-compose.yml.

up → starts all containers.

--build → rebuilds images before starting (useful when code changes).

-d → detached mode (runs in the background).

🧠 Use this whenever you update your Go code or Dockerfile.

---
### 🔹 2. Check Running Containers
```bash
docker ps
```

Explanation:
Lists all running Docker containers along with:
- container name

- port mappings

- uptime

- status (healthy/unhealthy)
---
### 🔹 3. View Container Logs

```bash
docker logs user-service
```

### 🔹 4. Restart a Specific Service
``` bash
docker compose up -d --build user-service
```
Explanation:
Rebuilds and restarts only the user-service without affecting other containers.

### 🔹 5. Connect to PostgreSQL Inside Container
```bash
docker exec -it user-db psql -U user -d userdb
```

Explanation:

- docker exec → run a command inside a running container.

- -it → interactive terminal mode.

- user-db → the name of the Postgres container.

- psql → PostgreSQL CLI tool.

- -U user → connects using username user.

- -d userdb → connects to database userdb.