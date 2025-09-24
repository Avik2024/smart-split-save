# 🚀 Smart Split & Save

**Smart Split & Save** is a real‑time fintech platform that allows users to create shared savings goals with friends, family, or colleagues — automatically splitting contributions and applying adaptive savings rules.  

This project demonstrates **full end‑to‑end engineering skills**: microservices architecture, event-driven backend, caching strategies, adaptive rule engines, cloud deployment, monitoring, and a modern frontend dashboard.

---

## 📌 Project Purpose

Managing shared savings is often tedious and manual. Smart Split & Save solves this by:
- Enabling **real-time shared goals** (e.g., holidays, rent, gifts).
- **Automatically splitting payments** based on contributions.
- Applying **adaptive saving rules** (round-ups, % salary savings, custom rules).
- Providing a **beautiful dashboard** with real-time tracking.
- Offering **notifications and insights** to keep users motivated.

---

## 🎯 Core Features

- **User Management** — registration, authentication, profile management.
- **Shared Goals** — create, join, track progress.
- **Automatic Splitting** — real-time contribution calculation.
- **Adaptive Rules Engine** — ML-powered savings suggestions.
- **Notifications** — real-time updates via WebSocket.
- **Dashboard** — interactive charts and goal summaries.

---

## 🛠 Tech Stack

| Component              | Technology |
|------------------------|------------|
| Backend                | Go |
| Event Streaming       | Kafka |
| Database              | PostgreSQL / Cassandra |
| Caching               | Redis |
| Machine Learning      | Python + TensorFlow |
| Frontend              | React.js + Tailwind CSS + D3.js |
| Containerization      | Docker |
| Orchestration        | Kubernetes |
| Deployment            | AWS / GCP |
| Monitoring            | Prometheus + Grafana |

---

## 📂 Project Structure


smart-split-save/
│
├── backend/ # All backend microservices
│ ├── auth-service/
│ ├── goal-service/
│ ├── transaction-service/
│ ├── rule-engine/
│ └── notification-service/
│
├── frontend/ # React frontend app
│ ├── public/
│ └── src/
│ ├── components/
│ ├── pages/
│ ├── services/
│ └── styles/
│
├── docs/ # Project documentation
│ ├── architecture.md
│ ├── scope.md
│ └── day1-log.md
│
├── infra/ # Infrastructure & deployment configs
│ ├── docker-compose.yml
│ ├── k8s/
│ └── terraform/
│
├── scripts/ # Utility scripts
│
├── .gitignore
├── README.md
└── LICENSE
