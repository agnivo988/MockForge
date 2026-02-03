#  MockForge

>  MockForge is under active development. APIs and features may evolve.

**MockForge** is an **OpenAPI-driven API mocking and testing platform** built for modern frontend and backend teams.

It allows developers to instantly mock APIs from OpenAPI specs, test endpoints using a **Postman-like UI**, simulate real-world scenarios, and switch environments — all without restarting servers.

Think **Postman + Mock Server + Scenario Engine**, fully open-source.

---

##  Project Bio

MockForge helps teams build, test, and iterate faster by removing dependency on real APIs during development and testing.

Designed for speed, CI/CD workflows, and extensibility, it fits seamlessly into modern engineering stacks.

---

##  Features

-  Mock APIs directly from **OpenAPI / Swagger specs**
-  **Postman-like API testing UI**
-  Scenario-based responses (success, error, edge cases)
-  Hot-reload specs (no restart needed)
-  Multiple environments (dev / staging / prod)
-  CLI-first workflow (WIP)
-  Docker & CI/CD ready
-  Built for frontend + backend teams

---

## 🏗 Tech Stack

### Backend
- **Go**
- Gin
- kin-openapi (OpenAPI parsing)

### Frontend
- **Next.js (App Router)**
- React
- Tailwind CSS

---

## 📂 Project Structure

```
MockForge/
├── backend/
│   ├── cmd/
│   ├── internal/
│   │   ├── mock/
│   │   ├── scenario/
│   │   └── specstore/
│   └── specs/
├── frontend/
│   ├── app/
│   ├── components/
│   └── styles/
├── cli/
├── docker/
└── README.md
```

---

## ⚡ Getting Started

### 1️⃣ Start Backend

```bash
cd backend
go run cmd/server/main.go
```

Backend runs at:
```
http://localhost:8080
```

### 2️⃣ Start Frontend

```bash
cd frontend
npm install
npm run dev
```

Frontend runs at:
```
http://localhost:3000
```

---

## 📄 Upload OpenAPI Spec

Upload your OpenAPI spec using the UI or via API.

**API Endpoint:**
```
POST /api/spec/upload
```

**Supported formats:**
- `.yaml`
- `.json`

**Specs are:**
- Validated before loading
- Hot-reloaded automatically
- Environment-aware

---

## 🧪 API Testing UI

MockForge provides a modern, Postman-like interface:

- Select HTTP method (GET / POST / PUT / DELETE)
- Enter request URL
- Edit headers
- Add request body
- Send requests to mocked APIs
- View formatted responses & status codes

Designed to feel like a real developer tool.

---

## 🎭 Scenario Engine (Core Feature)

MockForge supports scenario-based mocking, allowing responses to change dynamically based on request conditions.

**Scenarios can match on:**
- Query parameters
- Headers
- Request body
- Active environment

**Examples:**
- Return `500` when `?fail=true`
- Simulate auth failures
- Mock edge cases
- Add artificial delays

Scenario rules are designed to be loaded from YAML and hot-reloaded.

---

## 🧑‍💻 CLI (Work in Progress)

```bash
mockforge serve spec.yaml --env dev
```

**Planned CLI features:**
- Start mock server
- Validate OpenAPI specs
- Switch environments
- Export and replay traffic

---

## 🐳 Docker Support

```bash
docker-compose up
```

**Perfect for:**
- CI pipelines
- Preview environments
- Team-wide mock servers

---

## 🤝 Contributing

MockForge is open-source and contributor-friendly.

**You can contribute by:**
- Improving UI/UX
- Extending scenario rule operators
- Enhancing CLI capabilities
- Writing docs or examples

**How to contribute:**
1. Fork the repository
2. Create a feature branch
3. Submit a pull request 🚀

---

## 💡 Why MockForge?

- ✅ Saves frontend & backend development time
- ✅ Works well in CI/CD pipelines
- ✅ No vendor lock-in
- ✅ Built with modern tools
- ✅ Designed for extensibility and scale

---

## 📜 License

MIT License

---

## ❤️ Author

Built with passion by **Agniva Mukherjee**  
Open-source contributor & full-stack developer.

---

**⭐ If you find MockForge useful, consider giving it a star on GitHub!**
