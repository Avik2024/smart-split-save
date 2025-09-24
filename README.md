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

```plaintext
smart-split-save/
├── backend/
│   ├── auth-service/
│   ├── goal-service/
│   ├── transaction-service/
│   ├── rule-engine/
│   └── notification-service/
├── frontend/
│   ├── public/
│   └── src/
│       ├── components/
│       ├── pages/
│       ├── services/
│       └── styles/
├── docs/
│   ├── architecture.md
│   ├── scope.md
│   └── day1-log.md
├── infra/
│   ├── docker-compose.yml
│   ├── k8s/
│   └── terraform/
├── scripts/
├── .gitignore
├── README.md
└── LICENSE

```

---

## 📈 Architecture Diagram

![Architecture Diagram](./docs/architecture.md)  
*(Detailed architecture diagram and technical explanation will be included in `docs/architecture.md`)*

---

## 🗓 Project Roadmap

This project will be developed over **5 months** following a Level III Engineer plan:
1. Planning & architecture (Month 1)
2. Backend microservices & core APIs (Month 2)
3. Adaptive rules engine & ML integration (Month 3)
4. Frontend dashboard (Month 4)
5. Deployment, monitoring & documentation (Month 5)

---

## 🎓 Skills Showcased

- System design & microservices
- Event-driven architecture
- Real-time data processing
- Adaptive rule engines & ML integration
- Cloud deployment & Kubernetes orchestration
- Observability with Prometheus & Grafana
- Modern React.js UI development

---

## 🔗 Links

- [Architecture Document](./docs/architecture.md)
- [Scope Document](./docs/scope.md)
- [Day 1 Planning Log](./docs/day1-log.md)

---

## 📜 License
This project is licensed under the MIT License — see the [LICENSE](./LICENSE) file for details.








