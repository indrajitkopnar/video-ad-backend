<<<<<<< HEAD
# video-ad-backend
=======
# 📺 Video Ad Tracking Backend (GoLang)

A production-ready backend service for tracking and analyzing video advertisement clicks in real-time. Built with **Golang**, **PostgreSQL**, **Redis**, **Kafka (optional)**, and deployed using **Docker** + **Kubernetes**.

---
## 🚀 Features

- ✅ RESTful API for listing ads, recording clicks, and viewing analytics
- ✅ Rate-limiting (IP-based) to avoid abuse
- ✅ Deduplication logic (AdID + IP + timestamp)
- ✅ Retry logic and backup queue for DB failures
- ✅ Click-through rate (CTR) analytics
- ✅ Structured logging using `logrus`
- ✅ Real-time analytics powered by Redis
- ✅ Metrics exposed for Prometheus
- ✅ Docker & Kubernetes support

---

## 📦 Tech Stack

| Layer            | Technology                  |
|------------------|------------------------------|
| Language         | Go (Golang 1.22+)            |
| Web Framework    | Gin                          |
| Database         | PostgreSQL                   |
| Cache            | Redis                        |
| Message Broker   | Kafka *(future enhancement)* |
| Monitoring       | Prometheus + Grafana         |
| Containerization | Docker                       |
| Orchestration    | Kubernetes + Minikube        |

---

## 📘 API Documentation

### `GET /ads`
Returns list of active ads.

```json
[
  {
    "id": 1,
    "image_url": "https://example.com/ad1.jpg",
    "target_url": "https://example.com/click1"
  },
  ...
]
```

### `POST /ads/click`
Tracks a click event.

#### Request (JSON):
```json
{
  "ad_id": 1,
  "timestamp": "2025-04-15T16:55:00Z",
  "ip": "192.168.1.100",
  "playback_time": 12
}
```

#### Response:
```json
{
  "message": "Click received"
}
```

### `GET /ads/analytics`
Returns click count, impressions, and CTR per ad.

```json
[
  {
    "ad_id": 1,
    "click_count": 50,
    "impressions": 250,
    "ctr": 20
  }
]
```

---

## ⚙️ Project Structure

```
video-ad-backend/
├── cmd/                  # Entry point
├── config/               # Environment & config
├── controllers/          # API handlers
├── database/             # Postgres & Redis clients
├── models/               # Data models
├── middleware/           # Rate limiter
├── services/             # Click processor with retry logic
├── kafka/                # Kafka producer (optional)
├── utils/                # Logger, metrics, context
├── k8s/                  # Kubernetes manifests
│   ├── backend/
│   ├── postgres/
│   ├── redis/
│   ├── prometheus/
│   └── grafana/
├── Dockerfile
├── docker-compose.yml *(optional)*
└── README.md
```

---

## 🧪 How to Run

### ✅ Locally (Development)
```bash
# Start Postgres & Redis manually or via docker-compose
cp .env.example .env

# Run Go app
go run cmd/main.go
```

### ✅ With Docker
```bash
docker build -t video-ad-backend:latest .
docker run -p 8081:8081 --env-file .env video-ad-backend
```

### ✅ With Kubernetes (Minikube)
```bash
minikube start --driver=docker

# Enable Docker env to build image locally
eval $(minikube docker-env)
docker build -t video-ad-backend:latest .

# Apply manifests
kubectl apply -f k8s/

# Get backend service URL
minikube service video-ad-backend --url
```

---

## 📊 Monitoring

- Prometheus: `http://localhost:9090`
- Grafana: `http://localhost:3000`
  - Default login: `admin/admin`

> Prometheus config: `k8s/prometheus/configmap.yaml`
> Grafana deployment: `k8s/grafana/`

---

---


---

## 🧑‍💻 Author

**Indrajit Kopnar**  
[GitHub](https://github.com/indrajitkopnar)  |  Pune, India  
MEAN & Golang Developer | Cloud, Kafka, IoT Integrations


---

>>>>>>> a36caee (Initial complete commit: video ad tracking backend with Docker, Redis, PostgreSQL, Prometheus, Grafana, K8s)
