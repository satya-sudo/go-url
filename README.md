# 🔗 URL Shortener (Go + Docker)

A production-style **URL shortener** built with **Go microservices**, **Postgres**, and **Redis**, managed via **Docker Compose**.

---

## 📦 Architecture

The system is split into independent services:

- **Gateway**
    - Single entrypoint for clients.
    - Validates JWT tokens.
    - Proxies requests to Auth, CRUD, and Redirect services.

- **Auth Service**
    - Handles `signup`, `login`, and `me`.
    - Stores users in Postgres.
    - Issues JWTs for authentication.

- **CRUD Service**
    - Handles short URL creation, deletion, and stats.
    - Persists short URLs in Postgres.
    - Trusts `X-User-*` headers injected by Gateway.

- **Redirect Service**
    - Handles `GET /:code` (public).
    - Uses Redis for fast lookups and hit counting.
    - Falls back to Postgres if cache miss.

- **Job Runner**
    - Background worker that syncs hit counters from Redis → Postgres.

- **Postgres**
    - Persistent storage for users and URLs.

- **Redis**
    - Fast cache for redirects and hit counts.

---

## ⚙️ Tech Stack

- **Go** (1.21)
- **Gin / Fiber** (HTTP framework)
- **Postgres 15**
- **Redis 7**
- **Docker Compose**

---

## 🗄 Database Schema

### Users
```sql
CREATE TABLE IF NOT EXISTS users (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    email TEXT UNIQUE NOT NULL,
    password_hash TEXT NOT NULL,
    role TEXT DEFAULT 'user',
    created_at TIMESTAMPTZ DEFAULT now()
);
```

### URLs
```sql
CREATE TABLE IF NOT EXISTS urls (
    short_code TEXT PRIMARY KEY,
    long_url TEXT NOT NULL,
    user_id UUID REFERENCES users(id) ON DELETE CASCADE,
    created_at TIMESTAMPTZ DEFAULT now(),
    expiration_at TIMESTAMPTZ,
    hits BIGINT DEFAULT 0
);
```

---

## 🚀 Running Locally

### 1. Clone the repo
```bash
git clone https://github.com/yourname/url-shortener-go.git
cd url-shortener-go
```

### 2. Start services
```bash
docker-compose up --build
```

### 3. Verify services
- Gateway → http://localhost:8080
- Auth → http://localhost:8081
- CRUD → http://localhost:8082
- Redirect → http://localhost:8083

---

## 🔑 API Overview

### Auth
- `POST /auth/signup` → register user
- `POST /auth/login` → login, returns JWT
- `GET /auth/me` → get current user (requires JWT)

### CRUD
- `POST /shorten` → create short URL (requires JWT)
- `DELETE /shorten/:id` → delete short URL (requires JWT)
- `GET /shorten/:id/stats` → get stats (requires JWT)

### Redirect
- `GET /:code` → redirect to original URL, increments hit count

---

## 🔄 Background Job

- `jobRunner` runs every 1 minute.
- Reads `hits:*` keys from Redis.
- Updates Postgres counters.
- Resets Redis keys.

---

## 🛠 Development Notes

- Services share `DATABASE_URL` and `REDIS_ADDR`.
- Gateway enforces JWT validation for all protected routes.
- Auth is the only service issuing JWTs.
- All image/service names are **lowercase** (required by Docker).

---

## 📌 TODO / Future Work

- Add request rate limiting.
- Support custom short codes.
- Add metrics & tracing.
- Add admin panel for managing URLs.

---

## 📝 License

MIT
