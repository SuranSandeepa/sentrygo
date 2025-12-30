# SentryGo 🛰️

A lightweight, real-time infrastructure monitoring dashboard built with the **GOTH stack** (Go, Templ/HTML, HTMX). 

SentryGo allows you to monitor the status of various web services and APIs. It uses background workers to ping targets and updates the dashboard automatically without a full page refresh.



## 🚀 Features
- **Real-time Monitoring:** Background Goroutines check service health every 30 seconds.
- **Dynamic UI:** Powered by HTMX for seamless, "no-refresh" status updates.
- **Full CRUD:** Add and remove services instantly via the dashboard.
- **Containerized:** Fully orchestrated using Docker and Docker Compose.
- **Database Persistence:** Uses PostgreSQL to store service history and status.

## 🛠️ Tech Stack
- **Backend:** Go (Golang)
- **Database:** PostgreSQL
- **Frontend:** HTML + Tailwind CSS + HTMX
- **DevOps:** Docker, Docker Compose

## 📦 Getting Started

### Prerequisites
- Docker & Docker Compose installed.

### Installation & Running
1. Clone the repository:
   ```bash
   git clone [https://github.com/SuranSandeepa/sentrygo.git](https://github.com/SuranSandeepa/sentrygo.git)
   cd sentrygo
   ```

1. Start the entire stack:
   ```bash
   docker-compose up --build
   ```

<img width="1228" height="692" alt="image" src="https://github.com/user-attachments/assets/11991fe9-742c-49ea-8580-f250003419c3" />

   
   
