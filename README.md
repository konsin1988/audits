# AuditSO Go Server & Client

A Go-based web application with a **chi-powered REST API** backend and a **SolidJS (Vite)** Single Page Application (SPA) frontend. It connects to multiple databases to process and export data in various file formats (`pdf`, `xlsx`, `doc`).

## 🏗️ Architecture Overview

```text
├── config/         # Database and server configurations
├── handlers/       # HTTP route handlers (API logic)
├── services/       # Core business logic (Export service management)
├── repositories/   # Data providers interacting with multiple DBs
├── exporters/      # PDF, XLSX, and DOC generation logic
└── ui/             # SolidJS + TypeScript + Vite frontend SPA
```

---

## 🚀 Backend API

The backend runs on port `:8014` and implements `go-chi/chi` for routing.

### Key Endpoints

| Endpoint | Method | Description |
| :--- | :--- | :--- |
| `/auditso/api/health/` | `GET` | Returns `{"status":"ok"}` health check |
| `/auditso/api/{db}/{entity}/{format}/` | `GET` | Fetches data from `{db}` and `{entity}`, exporting it as `{format}` |

*Supported export formats: `pdf`, `xlsx`, `doc`.*

### Static Assets & SPA Routing
*   **Static Assets:** Served securely under `/auditso/assets` mapped directly to `./ui/assets`.
*   **SPA Catch-All:** Any route matching `/auditso/*` that isn't a direct file or API endpoint falls back to `./ui/index.html` to allow client-side routing.

---

## 🎨 Frontend Client

The frontend is an SPA bundle managed in the `ui` directory utilizing **SolidJS**, **Vite**, **TypeScript**, and **Tailwind CSS**.

### Client-Side Routes (`routes.ts`)

The SolidJS router mirrors the subfolder context structure:
*   `/` ➜ `Home` Component
*   `/:db/:docname/` ➜ Dynamic Registry Loader (`lazy` loaded)
*   `**` ➜ 404 Error Component fallback

### Available Scripts

Manage your frontend using `pnpm` inside the client directory:

```bash
pnpm install   # Install project dependencies
pnpm dev       # Run local Vite development server
pnpm build     # Build production assets into the /ui directory
```

---

## 🛠️ Development & Deployment

1. Make sure your database credentials match the mapping in `config.LoadDBConfigs()`.
2. Compile and run the main entry point:
   ```bash
   go run main.go
   ```
