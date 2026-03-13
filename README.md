# 🚀 URL Shortener - Go & React Monorepo

A high-performance, full-stack URL shortener application built using **Go (Golang)** for the backend and **React (TypeScript)** for the frontend. It features a clean UI, fast redirection logic, and automated API documentation.

## ✨ Features

- **Efficient Shortening:** Converts long URLs into unique, 6-character short codes.
- **Instant Redirection:** High-speed redirection to the original URL via Go Gin.
- **Interactive API Docs:** Fully documented API using **Swagger**.
- **Type-Safe Frontend:** React application built with TypeScript for better developer experience.
- **CORS Configuration:** Seamless communication between the frontend and backend.

## 🛠️ Tech Stack

**Backend:**
- **Language:** Go 1.25.6
- **Framework:** Gin Gonic, GORM
- **Documentation:** Swaggo (Swagger)
- **Database:** PostgreSQL

**Frontend:**
- **Framework:** React 18
- **Language:** TypeScript
- **Tooling:** Vite
- **HTTP Client:** Axios

---

## 📂 Project Structure

```text
url_shortener/
├── backend/           # Go source code
│   ├── cmd/           # Application entry point
│   ├── docs/          # Swagger auto-generated documentation
│   ├── internals/     # API logic, Models, and Helpers
│   └── main.go        # Server setup
├── frontend/          # React source code
│   ├── src/           # Components and Logic
│   ├── package.json   # Dependencies
│   └── vite.config.ts # Vite configuration
└── README.md