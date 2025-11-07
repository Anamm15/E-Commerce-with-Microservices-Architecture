# 🛍️ Simple E-Commerce — Monolith & Microservices

This project is a simple e-commerce web application that I built to understand and compare **Monolithic** and **Microservices** architectures.  
It started as a small experiment, but it evolved into a working system where users can browse products, manage carts, and simulate checkout — all while exploring two different backend designs.

---

## 🚀 Overview

The main goal of this project is to learn how an e-commerce application behaves when built using two different architectural approaches:

- **Monolithic Architecture:**  
  Everything (API, business logic, and database) is managed in one codebase using Node.js and Express.

- **Microservices Architecture:**  
  The application is divided into smaller independent services written in Go (Gin framework).  
  Each service has its own database and communicates through REST and **Kafka** for event-driven messaging.

This setup helped me explore scalability, fault tolerance, and service communication patterns.

---

## ⚙️ Tech Stack

### 🖥️ Frontend
- **React + TypeScript**  
- Axios for API communication  
- Basic state management (React Context or simple hooks)  
- Styled Components / TailwindCSS for UI styling

### 🧱 Backend (Monolith)
- **Node.js + Express**
- **Sequelize ORM**
- MySQL / PostgreSQL as the database
- RESTful API endpoints for all main features

### 🔗 Backend (Microservices)
- **Go + Gin Framework**
- **GORM ORM**
- **Kafka** for asynchronous messaging and event handling
- Separate services for auth, product, and order management
- Each service runs independently and communicates through REST APIs or Kafka topics

### 🗄️ Database
- MySQL (Monolith)
- Separate PostgreSQL databases (Microservices)

---

## 🧩 Microservices Architecture

- **Auth Service** – Manages user registration, login, and authentication  
- **Product Service** – Handles product catalog, pricing, and inventory  
- **Order Service** – Processes orders and checkout  
- **Kafka Broker** – Used for event-driven communication between services (e.g., order events, product updates)  
- **API Gateway (Optional)** – Can be added for routing and request aggregation

---

## 🧠 What I Learned Soon

- How to design and build both **monolithic** and **microservices** architectures  
- Managing **data consistency** and **communication** between services  
- Setting up **Kafka** for event-driven microservices  
- How Docker simplifies running multiple services  
- The impact of architecture choice on deployment and scalability

---

## 💡 Future Improvements

- Add real payment gateway integration  
- Implement distributed tracing and logging (Jaeger / Prometheus)  
- Improve frontend UI and UX  
- Add CI/CD pipeline for automated deployment  
- Deploy both architectures to cloud for comparison (e.g., AWS, GCP, or Render)

---

## 📦 How to Run

### 🖥️ Frontend
```bash
cd frontend
npm install
npm run dev
```


### 🧱 Monolithic Version
```bash
cd backend-monolith
npm install
npm run dev
```


### ⚙️ Microservices Version
```bash
cd backend-microservices
docker-compose up --build
```

---

## 🧑‍💻 Author
Choirul Anam
Computer Science student exploring full-stack development, backend architecture, and distributed systems.

---

## 📄 License
This project is open source and free to use for learning and educational purposes.


