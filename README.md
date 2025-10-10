# 👤 Users Service — GameGear E-commerce

บริการสำหรับจัดการผู้ใช้และการยืนยันตัวตนทั้งหมดภายในระบบ **GameGear E-commerce**
เป็นหนึ่งใน microservices ที่ทำงานร่วมกับ **Kong API Gateway**

---

## 🏛️ Architectural Design & Responsibility

* เป็น **เจ้าของข้อมูล (Data Owner)** สำหรับตารางผู้ใช้: `users`, `password_reset_tokens`
* ข้อมูลผู้ใช้ **ไม่ถูกเข้าถึงโดยตรงจาก service อื่น** — ต้องเรียกผ่าน API ของ Users Service เท่านั้น
* บริการอื่น (Shop, Admin) ถ้าต้องการข้อมูลผู้ใช้ จะต้องเรียก API นี้ผ่าน Gateway

---

## ✨ Features & Endpoints

| Feature            | HTTP Method | Path                                                    | Description                           |
| ------------------ | ----------- | ------------------------------------------------------- | ------------------------------------- |
| User Registration  | POST        | `/api/auth/register`                                    | ลงทะเบียนผู้ใช้ใหม่                   |
| Login / Logout     | POST        | `/api/auth/login`, `/api/auth/logout`                   | ส่ง JWT / ยกเลิก session              |
| Profile Management | GET / PUT   | `/api/user/profile`                                     | ดูและอัปเดตข้อมูลผู้ใช้ (ต้องล็อกอิน) |
| Password Reset     | POST        | `/api/auth/forgot-password`, `/api/auth/reset-password` | Flow รีเซ็ตรหัสผ่าน                   |

นอกจากนี้ ควรมี endpoint สำหรับ **Healthcheck**:

```
GET /healthz → 200 OK
```

---

## 📂 Project Structure (Standard Go Layout)

```
.
├── cmd/
│   └── api/
│       └── main.go
├── internal/
│   ├── handlers/
│   │   └── user_handler.go
│   ├── services/
│   │   └── user_service.go
│   ├── repositories/
│   │   └── user_repository.go
│   └── models/
│       └── user_model.go
├── docs/
│   └── swagger (ไฟล์ที่สร้างโดย swag)
├── .env.example
├── go.mod
├── go.sum
├── Dockerfile.dev
├── docker-compose.override.yml
├── docker-compose.kong.yml
├── .dockerignore
└── README.md
```

คำอธิบาย:

* **cmd/api** — จุดเริ่มต้นของโปรแกรมหลัก (main entry point) สำหรับรันเซิร์ฟเวอร์ API
* **internal/handlers** — ชั้นรับคำขอ (HTTP Request) และส่งคำตอบ (Response)
* **internal/services** — ชั้นของ Business Logic ที่ประมวลผลคำขอและควบคุมการทำงานของระบบ
* **internal/repositories** — ชั้นติดต่อกับฐานข้อมูล เช่น การ Query, Insert, Update, Delete
* **internal/models** — กำหนดโครงสร้างข้อมูล (Struct) ที่ใช้ภายในระบบและ ORM (GORM Models)
* **docs/swagger** — เก็บไฟล์เอกสาร API ที่สร้างโดย `swag init` (OpenAPI/Swagger)
* **.env.example** — ตัวอย่างไฟล์สำหรับตั้งค่าคอนฟิก เช่น Database URL, JWT, Email
* **Dockerfile.dev** — ไฟล์ Docker สำหรับโหมดพัฒนา (Dev Mode) ใช้ `air` สำหรับ hot-reload
* **docker-compose.override.yml** — ใช้รัน service พร้อม PostgreSQL ในโหมดพัฒนา
* **docker-compose.kong.yml** — ใช้เป็นแนวทางการเชื่อมต่อกับ Kong Gateway ตามสถาปัตยกรรมของโปรเจกต์หลัก (Mini-Project-Golang)
* **.dockerignore** — รายการไฟล์/โฟลเดอร์ที่ไม่ต้องการนำเข้าเวลาสร้าง Docker image
* **README.md** — เอกสารอธิบายรายละเอียดการติดตั้งและใช้งาน Service

---

## 🚀 Getting Started (Local Development)

### 1. Clone & เข้าโฟลเดอร์

```bash
git clone https://github.com/Wattanaroj2567/users-service.git
cd users-service
```

### 2. ติดตั้ง dependencies

```bash
go mod tidy
```

### 3. ตั้งค่า Database

ตรวจสอบว่า PostgreSQL ทำงานอยู่ แล้วสร้างฐานข้อมูล:

```sql
CREATE DATABASE gamegear_users_db;
```

### 4. สร้าง `.env` จาก `.env.example`

ปรับค่าสำหรับเครื่องคุณ เช่น user, password:

```env
APPLICATION_PORT=8080
DATABASE_URL="host=localhost user=xxx password=yyy dbname=gamegear_users_db sslmode=disable"
JWT_SECRET_KEY="your_super_secret"
EMAIL_HOST="smtp.example.com"
EMAIL_PORT=587
EMAIL_USERNAME="you@example.com"
EMAIL_PASSWORD="password"
FRONTEND_URL="http://localhost:3000"
```

### 5. รันโปรเจกต์

```bash
go run cmd/api/main.go
```

ตอนเริ่มต้นจะทำการ migrate ตารางอัตโนมัติ และรันที่ `http://localhost:8080`

---

## 🐋 Run with Docker (Dev)

### Dockerfile.dev

```dockerfile
FROM golang:1.22-alpine

RUN apk add --no-cache git bash build-base tzdata ca-certificates \
    && update-ca-certificates \
    && go install github.com/cosmtrek/air@latest

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download
COPY . .

EXPOSE 8080
CMD ["air"]
```

### docker-compose.override.yml

```yaml
version: "3.9"
services:
  users-db:
    image: postgres:15-alpine
    environment:
      POSTGRES_DB: gamegear_users_db
      POSTGRES_USER: dev
      POSTGRES_PASSWORD: dev
    ports:
      - "5432:5432"

  users-service:
    build:
      context: .
      dockerfile: Dockerfile.dev
    environment:
      DATABASE_URL: postgres://dev:dev@users-db:5432/gamegear_users_db?sslmode=disable
      APPLICATION_PORT: 8080
      JWT_SECRET_KEY: "supersecretkey"
    ports:
      - "8080:8080"
    depends_on:
      - users-db
    volumes:
      - .:/app
```

### docker-compose.kong.yml

```yaml
version: "3.9"
services:
  users-service:
    build:
      context: .
      dockerfile: Dockerfile.dev
    container_name: users-service
    environment:
      APPLICATION_PORT: 8080
      DATABASE_URL: postgres://dev:dev@users-db:5432/gamegear_users_db?sslmode=disable
    ports:
      - "8080:8080"
    networks:
      - gamegear-network
    depends_on:
      - users-db

  users-db:
    image: postgres:15-alpine
    environment:
      POSTGRES_DB: gamegear_users_db
      POSTGRES_USER: dev
      POSTGRES_PASSWORD: dev
    networks:
      - gamegear-network

networks:
  gamegear-network:
    external: true  # ใช้ network เดียวกับ Kong จากโปรเจกต์หลัก
```

> ไฟล์นี้ใช้เป็นแนวทางการเชื่อมต่อกับ Kong Gateway ตามสถาปัตยกรรมของโปรเจกต์หลัก (Mini-Project-Golang)

---

## 📝 API Documentation (Swagger / OpenAPI)

### ติดตั้ง Swag

```bash
go install github.com/swaggo/swag/cmd/swag@latest
```

### สร้างเอกสาร Swagger

```bash
swag init
```

จะสร้างโฟลเดอร์ `docs` และไฟล์ Swagger JSON/YAML ให้อัตโนมัติ

### เปิดใช้งาน Swagger UI

```
http://localhost:8080/swagger/index.html
```

---

## 🔁 Remote Development (ngrok หรือ Tunnel)

ถ้าคุณอยากให้เพื่อนร่วมทีมหรือฝั่ง admin-service เข้าถึง users-service ผ่าน tunnel:

```bash
ngrok http 8080
```

แล้วใช้ URL ที่ ngrok ให้ใน `.env` ของ service อื่น (เช่น `USER_SERVICE_URL=<ngrok-url>`)

---

## 🧭 ทำไมต้องใช้ ngrok ร่วมกับ Kong Gateway (กรณีพัฒนาแบบทีม)

ในการพัฒนาแบบทีมที่สมาชิกแต่ละคนรัน service บนเครื่องของตัวเอง เช่น `users-service`, `shop-service`, หรือ `admin-service` ไม่ได้อยู่ใน Docker network เดียวกัน — การสื่อสารตรงผ่านชื่อ service จะทำไม่ได้ เพราะอยู่คนละเครือข่ายกัน

ดังนั้น **ngrok** จึงถูกใช้เพื่อสร้าง tunnel ระหว่างเครื่องนักพัฒนาแต่ละคนให้สื่อสารกันได้ผ่านอินเทอร์เน็ต โดยมีเหตุผลหลักดังนี้:

| เหตุผล                         | รายละเอียด                                                                                                                                              |
| ------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------------- |
| 🌍 การทำงานคนละเครื่อง         | เมื่อสมาชิกทีมรัน service ของตนเอง (เช่น คุณรัน users-service แต่เพื่อนรัน admin-service) ต้องใช้ ngrok แชร์ URL ให้กันเพื่อให้เชื่อมต่อได้ผ่าน Gateway |
| ⚙️ Integration กับ Kong        | Kong สามารถใช้ URL จาก ngrok เป็นปลายทาง service ได้โดยตรง เช่น `USER_SERVICE_URL=https://abc1234.ngrok.io` เพื่อ proxy ไปยังเครื่องเพื่อน              |
| 🚀 สะดวกต่อการทดสอบ            | ใช้ทดสอบระบบรวม (Integration Test) ระหว่าง service จริง โดยไม่ต้อง deploy ขึ้นเซิร์ฟเวอร์กลาง                                                           |
| 🎓 ใช้สาธิต/ส่งอาจารย์ได้ทันที | สามารถเปิดระบบในเครื่องแล้วแชร์ให้ผู้สอนหรือผู้ทดสอบเข้ามาเรียก API ได้ผ่านอินเทอร์เน็ต                                                                 |

> 🔸 สรุป: ถ้าพัฒนาใน Docker Compose เดียวกัน → ไม่ต้องใช้ ngrok
> 🔹 แต่ถ้าอยู่คนละเครื่อง → ใช้ ngrok เพื่อให้ Kong เชื่อมต่อได้ครบทุก service

---

## ✅ Summary

* README นี้อัปเดตให้สอดคล้องกับ **แนวทางหลักของโปรเจกต์ Mini-Project-Golang**
* รองรับทั้งการพัฒนา, ทดสอบ และเชื่อมต่อกับ Kong Gateway
* มีวิธีรันแบบ local, Docker, Swagger, Kong Integration และ Remote Dev พร้อมใช้งานจริง
