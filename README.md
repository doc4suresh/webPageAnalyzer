# Web Page Analyzer With GO Lang

A web application to analyze web pages for their HTML version, title, heading structure, accessible/inaccessible links, and login form detection. Built with a Go (Gin) backend and a React frontend.

---

## Features
- Analyze any public web page by URL
- Detect HTML version and page title
- Count all heading tags (h1-h6)
- List accessible and inaccessible links
- Detect presence of login forms
- Modern, responsive UI with Bootstrap

---

## Tech Stack
- **Backend:** Go (Gin, Colly)
- **Frontend:** React, Bootstrap 5, Bootstrap Icons

---

## Folder Structure
```
webPageAnalyzer/
  backend/         # Go backend (API server)
    cmd/
    internal/
      server/
      service/
      util/
    .env
    go.mod
    go.sum
  frontend/        # React frontend
    public/
    src/
    .env
    package.json
README.md          # (this file)
```

## Run without Docker
---

### Prerequisites
- Go 1.18+
- Node.js 16+
- npm (comes with Node.js)

---

### Backend Setup (Go)
1. **Install dependencies:**
   ```sh
   cd backend
   go mod tidy
   ```
2. **Create a `.env` file** in `backend/` (optional, for custom port):
   ```env
   SERVER_ADDRESS=localhost
   SERVER_PORT=8080
   URL_RETRY_LIMIT=3
   URL_RETRY_DELAY=2
   ```
   If not set, defaults may be used in code.
3. **Run the backend server:**
   ```sh
   go run ./cmd/api/main.go
   ```
   The API will be available at `http://localhost:8080/analyze?url=...`

---

### Frontend Setup (React)
1. **Install dependencies:**
   ```sh
   cd frontend
   npm install
   ```
2. **Create a `.env` file** in `frontend/`:
   ```env
   REACT_APP_BACKEND_API_URL=http://localhost:8080
   ```
3. **Start the React app:**
   ```sh
   npm start
   ```
   The app will open at [http://localhost:3000](http://localhost:3000)

---

### Usage
1. Open the React app in your browser.
2. Enter a full URL (e.g., `https://www.example.com`) and click "Analyze".
3. View the analysis results in a modern, two-column summary card.

---

### Example API Call
```
GET http://localhost:8080/analyze?url=https://www.example.com
```

## Run with Docker
---

### Prerequisites
- Install Docker Desktop for Windows or Mac.
- For Linux, install Docker Engine for your distribution.

---

### 1. **Backend (Go) Dockerfile**

**`backend/dockerfile`**
```dockerfile
# Build stage
FROM golang:1.21-alpine AS builder

WORKDIR /app
COPY . .
RUN go mod tidy
RUN CGO_ENABLED=0 GOOS=linux go build -o webanalyzer ./cmd/api/main.go

# Final stage
FROM alpine:latest
WORKDIR /app
COPY --from=builder /app/webanalyzer .
COPY .env . # if you use .env for config (optional)
EXPOSE 8080
CMD ["./webanalyzer"]
```

---

### 2. **Frontend (React) Dockerfile**

**`frontend/dockerfile`**
```dockerfile
# Build stage
FROM node:20-alpine AS builder

WORKDIR /app
COPY . .
RUN npm install
RUN npm run build

# Production stage
FROM nginx:alpine
COPY --from=builder /app/build /usr/share/nginx/html
COPY nginx.conf /etc/nginx/conf.d/default.conf # Optional: custom nginx config
EXPOSE 80
CMD ["nginx", "-g", "daemon off;"]
```

---

### 3. **Build Docker Images**

**`backend/Dockerfile`**
```
docker build -t backend Dockerfile
```

**`frontend/Dockerfile`**
```
docker build -t frontend Dockerfile
```

Now you will have `backend` and `frontend` Docker images in your local Docker repository. You can run these images as containers using:

### 3. **Run Docker Images**

```sh
docker run -p 8080:8080 backend
docker run -p 3000:80 frontend
```

- The backend will be available at `http://localhost:8080`
- The frontend will be available at `http://localhost:3000`

---


