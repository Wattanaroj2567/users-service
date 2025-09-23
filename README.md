# 👤 User Service

Service สำหรับจัดการผู้ใช้และการยืนยันตัวตนทั้งหมดของโปรเจกต์ **GameGear E-commerce**

---

## 🏛️ Architectural Design

Service นี้ถูกออกแบบให้เป็น **"เจ้าของข้อมูล" (Data Owner)** โดยมีหน้าที่รับผิดชอบในการจัดการข้อมูลและตารางในฐานข้อมูลที่เกี่ยวกับผู้ใช้ (`users`, `password_reset_tokens`) โดยตรง
Service อื่นๆ ที่ต้องการเข้าถึงข้อมูลเหล่านี้ จะต้องเรียกใช้งานผ่าน API ที่ Service นี้มีให้เท่านั้น เพื่อรักษาความถูกต้องของข้อมูลและแบ่งหน้าที่ความรับผิดชอบอย่างชัดเจน

---

## ✨ Features & Responsibilities

Service นี้มีหน้าที่รับผิดชอบฟีเจอร์หลัก 4 ส่วน:

* **User Registration** (`POST /api/auth/register`):

  * สร้างบัญชีผู้ใช้ใหม่ด้วย `username`, `email`, และ `password`
  * ตรวจสอบข้อมูลซ้ำซ้อนและเข้ารหัสรหัสผ่านก่อนบันทึกลงฐานข้อมูล

* **Authentication** (`POST /api/auth/login`, `POST /api/auth/logout`):

  * ตรวจสอบ `username/email` และ `password` เพื่อยืนยันตัวตน
  * สร้างและส่ง JSON Web Token (JWT) กลับไปให้ Client เพื่อใช้ในการยืนยันตัวตนครั้งถัดไป

* **Profile Management** (`GET /api/user/profile`, `PUT /api/user/profile`):

  * อนุญาตให้ผู้ใช้ที่ล็อกอินแล้วสามารถดูและอัปเดตข้อมูลส่วนตัวได้ (เช่น `display_name`, `email`)

* **Password Reset** (`POST /api/auth/forgot-password`, `POST /api/auth/reset-password`):

  * จัดการ Flow การรีเซ็ตรหัสผ่านผ่านอีเมล โดยใช้ Token ที่มีเวลาหมดอายุ

---

## 📂 Project Structure

โครงสร้างของโปรเจกต์จัดเรียงตามหลัก **Standard Go Layout** โดยแต่ละส่วนมีหน้าที่ดังนี้

<table>
<tr>
<td width="60%">
<pre>
.
├── cmd/api/
│   └── main.go
├── internal/
│   ├── handlers/
│   │   └── user_handler.go
│   ├── services/
│   │   └── user_service.go
│   ├── repositories/
│   │   └── user_repository.go
│   └── models/
│       └── user_model.go
├── .env.example
├── .gitignore
├── go.mod
├── Dockerfile.dev
├── docker-compose.override.yml
├── .dockerignore
└── README.md
</pre>
</td>
<td>
<ul>
<li><b>cmd/api</b>: จุดเริ่มต้นในการรันโปรแกรมและตั้งค่าเริ่มต้น</li>
<li><b>internal</b>: โค้ดหลักทั้งหมดของ Service</li>
<ul>
<li><b>handlers</b>: ส่วนที่รับ Request และส่ง Response</li>
<li><b>services</b>: ส่วนของ Logic หลักทางธุรกิจ</li>
<li><b>repositories</b>: ส่วนที่ใช้สื่อสารกับฐานข้อมูล</li>
<li><b>models</b>: ส่วนกำหนดโครงสร้างข้อมูล (Structs)</li>
</ul>
<li><b>.env.example</b>: ไฟล์ตัวอย่างสำหรับ Configuration</li>
<li><b>.gitignore</b>: ไฟล์กำหนดรายการที่ไม่ต้องนำขึ้น Git Repository</li>
<li><b>Dockerfile.dev</b>: Dockerfile สำหรับ Dev Mode (ใช้ air เพื่อ hot-reload)</li>
<li><b>docker-compose.override.yml</b>: Compose ไฟล์สำหรับรัน Users Service + PostgreSQL</li>
<li><b>.dockerignore</b>: ไฟล์ ignore เพื่อลด context ขณะ build image</li>
</ul>
</td>
</tr>
</table>


---

## 🚀 Getting Started (Step-by-step)

ทำตามขั้นตอนทีละสเต็ปเพื่อตั้งค่าและรัน Service ในเครื่องของคุณ

### Step 1 — Clone the Repository (Standard)

```bash
git clone https://github.com/Wattanaroj2567/users-service.git
cd users-service
```

### Step 1 (Alt) — Direct Branch Clone (Develop Branch)

```bash
git clone -b develop https://github.com/Wattanaroj2567/users-service.git
cd users-service
```

### Step 2 — Install Dependencies

```bash
go mod tidy
```

### Step 3 — Setup Database

ตรวจสอบให้แน่ใจว่า PostgreSQL Server ทำงานอยู่ แล้วสร้างฐานข้อมูล (ถ้ายังไม่มี)

**(A) ใช้ SQL โดยตรง**

```sql
CREATE DATABASE gamegear_users_db;
```

**(B) ใช้ psql ผ่าน bash (one-liner)**

```bash
psql -U your_user -h localhost -p 5432 -c "CREATE DATABASE gamegear_users_db;"
```

### Step 4 — Configure Environment Variables

สร้างไฟล์ `.env` ที่ root ของโปรเจกต์ และใส่ค่าตามตัวอย่าง (แก้ `your_user`, `your_password`, และค่าอื่น ๆ ให้ถูกต้อง)

```env
# Core Configuration
APPLICATION_PORT=8080

# PostgreSQL Database Connection URL
DATABASE_URL="host=localhost user=your_user password=your_password dbname=gamegear_users_db port=5432 sslmode=disable"

# JWT Authentication
JWT_SECRET_KEY="your_super_secret_key"

# Email Service
EMAIL_HOST="smtp.gmail.com"
EMAIL_PORT=587
EMAIL_USERNAME="your.email@example.com"
EMAIL_PASSWORD="your_email_password"

# Frontend/Client URL
FRONTEND_URL="http://localhost:3000"
```

### Step 5 — Run the Service

```bash
go run cmd/api/main.go
```

> เมื่อรันคำสั่งนี้ ระบบจะทำการ migrate ตารางที่จำเป็นทั้งหมด และเซิร์ฟเวอร์จะเริ่มที่ `http://localhost:8080`

---

## 🐋 Run with Docker (Recommended)

สำหรับการพัฒนาแบบทีม ควรใช้ Docker เพื่อให้ทุกคนทำงานในสภาพแวดล้อมเดียวกัน โดยไม่ต้องติดตั้ง Go หรือ PostgreSQL ในเครื่อง

### Step 1 — สร้างไฟล์ที่เกี่ยวข้อง

```bash
touch Dockerfile.dev
touch docker-compose.override.yml
touch .dockerignore
```

### Step 2 — Dockerfile.dev

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

### Step 3 — docker-compose.override.yml

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

### Step 4 — .dockerignore

```bash
.git
.env
/tmp
/docs
/vendor
**/*.log
```

### Step 5 — Run with Docker Compose

```bash
docker compose -f docker-compose.override.yml up --build
```

หยุดการทำงาน:

```bash
docker compose -f docker-compose.override.yml down
```

---

## 📝 API Documentation: Swagger (OpenAPI)

### Step 1 — ติดตั้ง `swag`

```bash
go install github.com/swaggo/swag/cmd/swag@latest
```

### Step 2 — สร้างไฟล์เอกสาร Swagger

```bash
swag init
```

> คำสั่งนี้จะสร้างโฟลเดอร์ `docs` และไฟล์ที่จำเป็นขึ้นมาโดยอัตโนมัติ

### Step 3 — เปิดดู API Docs

```
http://localhost:8080/swagger/index.html
```

---

## 🤝 Remote Development (ngrok)

### Step 1 — ติดตั้ง ngrok

ดาวน์โหลดได้ที่ [ngrok.com](https://ngrok.com)

### Step 2 — รัน User Service

```bash
go run cmd/api/main.go
```

### Step 3 — เปิดอุโมงค์ไปยังพอร์ต 8080

```bash
ngrok http 8080
```

### Step 4 — แชร์ URL ให้ทีม

คัดลอก URL ที่ขึ้นต้นด้วย `https://...` ส่งให้เพื่อนร่วมทีม

### Step 5 — ตั้งค่าใน .env ของ admin-service

```env
USER_SERVICE_URL="<THE_NGROK_URL_YOU_SENT>"
```
