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
    go.mod
    go.sum
  frontend/        # React frontend
    public/
    src/
    package.json
    README.md
README.md          # (this file)
```

---

## Prerequisites
- Go 1.18+
- Node.js 16+
- npm (comes with Node.js)

---

## Backend Setup (Go)
1. **Install dependencies:**
   ```sh
   cd backend
   go mod tidy
   ```
2. **Create a `.env` file** in `backend/` (optional, for custom port):
   ```env
   SERVER_ADDRESS=localhost
   SERVER_PORT=8080
   ```
   If not set, defaults may be used in code.
3. **Run the backend server:**
   ```sh
   go run ./cmd/api/main.go
   ```
   The API will be available at `http://localhost:8080/analyze?url=...`

---

## Frontend Setup (React)
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

## Usage
1. Open the React app in your browser.
2. Enter a full URL (e.g., `https://www.example.com`) and click "Analyze".
3. View the analysis results in a modern, two-column summary card.

---

## Example API Call
```
GET http://localhost:8080/analyze?url=https://www.example.com
```

---