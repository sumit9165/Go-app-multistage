# Docker Image Optimization Multi-Stage Builds & Docker Hub  
-----
> Built with Go · Served on Alpine · Pushed to Docker Hub
- We build a image with single stage and mltistage Dockerfile
- It optimize and reduces the size of application.
- Also run container with both size of images.
---

## 🚀 Live Demo

Click **"Run Simulation"** to fetch real metrics from the Go backend and see the difference between a naive Docker image and an optimized one.

| Metric | 🐳 Large Image | ⚡ Optimized Image |
|---|---|---|
| Image Size | 845 GB | 16 MB |

<img width="1655" height="492" alt="day-35-docker-push-multistage" src="https://github.com/user-attachments/assets/9857a57c-a0b9-4253-8ad9-bd37b2fdf557" />

---

## 🗂️ Project Structure

```
docker-image-demo/
├── main.go              # Go HTTP server + JSON API
├── go.mod               # Go module (no external deps)
└── public/
    ├── index.html       # Frontend UI
    ├── style.css        # Glassmorphism dark theme
    └── script.js        # Fetches images and updates DOM
```

---

## 🛠️ Tech Stack

- **Backend** — Go 1.21 (standard library only)
- **Frontend** — Vanilla HTML, CSS, JavaScript
- **Containerization** — Docker with multi-stage build
- **Base Image** — Alpine Linux (minimal, ~5 MB)

---

## 🏗️ Multi-Stage Dockerfile

The core concept of this project. Two stages — one to **build**, one to **run**:

```dockerfile
# Stage 1: Build
FROM golang:1.21-alpine AS builder
WORKDIR /app
COPY . .
RUN go mod download
RUN go build -o server .

# Stage 2: Run
FROM alpine:latest
WORKDIR /app
COPY --from=builder /app .
EXPOSE 3000
CMD ["./server"]
```

**Why this matters:**  
The `golang` build image is ~850 MB. The final `alpine` image is only ~16–17 MB.  
That's a **98% reduction** — same app, just smarter packaging.

---

## ⚙️ Getting Started

### Run Locally (without Docker)

```bash
git clone https://github.com/sumit9165/Go-app-multistage
cd docker-image-demo

go run main.go
```

Open [http://localhost:3000](http://localhost:3000)

------------
### Run with Docker

```bash
# Build the optimized image
docker build -t docker-image-demo .

# Run the container
docker run -p 3000:3000 docker-image-demo
```
<img width="1905" height="537" alt="day-35-docker-run-multistage" src="https://github.com/user-attachments/assets/4ac25047-1199-47a9-835b-622682146b3c" />

-----------------
Open [http://localhost:3000](http://localhost:3000)

- Final Result
  
  -----------
<img width="1918" height="1017" alt="day-35-multistage-app-run" src="https://github.com/user-attachments/assets/16d9d981-6257-41b9-8ad5-be4843abbd87" />

-----------
### Pull from Docker Hub

```bash
docker image tag new-app-gp:latest sumit9165/multistage-go-app:latest
docker push sumit9165/multistage-go-app:latest
```

<img width="1655" height="492" alt="day-35-docker-push-multistage" src="https://github.com/user-attachments/assets/9bfa988a-40e8-4ebf-8935-bfbdb1dcd193" />

-----------------------------
### Pull from Docker Hub

```bash
docker pull sumit9165/multistage-go-app:latest
docker run -p 3000:3000 sumit9165/multistage-go-app:latest
```

-----

## 💡 Key Learnings

- **Multi-stage builds** let you use a heavy build environment without shipping it in production
- **Alpine / Distroless** base images remove everything non-essential — no shell, no package manager, just your binary
- **Go compiles to a single static binary** — no runtime dependencies needed, making it perfect for minimal containers
- **Smaller images** = faster CI/CD pipelines, lower storage costs, and quicker deployments
-------------------------
## 👤 Author

**SUMIT** 

---
