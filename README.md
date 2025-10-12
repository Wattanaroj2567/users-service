# 👤 Users Service — GameGear E-commerce

บริการสำหรับจัดการผู้ใช้และการยืนยันตัวตนทั้งหมดภายในระบบ **GameGear E-commerce**
เป็นหนึ่งใน microservices ที่ทำงานร่วมกับ **Kong API Gateway**

> 📖 **ดูเอกสารหลักของระบบ**: สำหรับ Kong Gateway setup, Architecture overview และการ integrate ทั้งระบบ → [Main README](../Mini-Project-Golang/README.md)
---

## 🏛️ Architectural Design & Responsibility

- เป็น **เจ้าของข้อมูล (Data Owner)** สำหรับตารางผู้ใช้: `users`, `password_reset_tokens`
- ข้อมูลผู้ใช้ **ไม่ถูกเข้าถึงโดยตรงจาก service อื่น** — ต้องเรียกผ่าน API ของ Users Service เท่านั้น
- บริการอื่น (Shop, Admin) ถ้าต้องการข้อมูลผู้ใช้ จะต้องเรียก API นี้ผ่าน Gateway

---

## ✅ สิ่งที่ต้องทำ (สำหรับผู้พัฒนา Service นี้)

> 👨‍💻 **Developer**: ณิชพน มานิตย์

### 📝 ไฟล์ที่ต้องแก้ไข/เขียนโค้ด

คุณต้องเขียนโค้ดเฉพาะใน **2 จุดหลัก**:

#### 1. **`cmd/api/main.go`**

- เริ่มต้น Gin server
- Setup routes
- Connect database
- Middleware configuration

#### 2. **โฟลเดอร์ `internal/`**

| Folder               | ต้องทำ       | หมายเหตุ                          |
| -------------------- | ------------ | --------------------------------- |
| ✅ **handlers/**     | ✅ ต้องทำ    | เขียน HTTP handlers (controllers) |
| ❌ **models/**       | ❌ ไม่ต้องทำ | **PM ทำให้แล้ว** (วรรธนโรจน์)     |
| ✅ **repositories/** | ✅ ต้องทำ    | เขียน database operations (CRUD)  |
| ✅ **services/**     | ✅ ต้องทำ    | เขียน business logic              |

#### 3. **ไฟล์อื่นๆ**

- ✅ **`.env`** - ตั้งค่า environment variables
- ✅ **`go.mod`** - เพิ่ม dependencies ตามต้องการ

### 🚫 ไฟล์ที่ไม่ต้องแก้

- ❌ `internal/models/` - PM ทำให้แล้ว
- ❌ `README.md` - มีให้แล้ว (แต่สามารถเพิ่มเติมได้)

### 📋 Checklist

- [ ] เขียน `cmd/api/main.go`
- [ ] เขียน handlers ใน `internal/handlers/`
- [ ] เขียน repositories ใน `internal/repositories/`
- [ ] เขียน services ใน `internal/services/`
- [ ] ตั้งค่า `.env`
- [ ] ทดสอบ API ด้วย Swagger
- [ ] ทดสอบผ่าน Kong Gateway

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
└── README.md
```

คำอธิบาย:

- **cmd/api** — จุดเริ่มต้นของโปรแกรมหลัก (main entry point) สำหรับรันเซิร์ฟเวอร์ API
- **internal/handlers** — ชั้นรับคำขอ (HTTP Request) และส่งคำตอบ (Response)
- **internal/services** — ชั้นของ Business Logic ที่ประมวลผลคำขอและควบคุมการทำงานของระบบ
- **internal/repositories** — ชั้นติดต่อกับฐานข้อมูล เช่น การ Query, Insert, Update, Delete
- **internal/models** — กำหนดโครงสร้างข้อมูล (Struct) ที่ใช้ภายในระบบและ ORM (GORM Models)
- **docs/swagger** — เก็บไฟล์เอกสาร API ที่สร้างโดย `swag init` (OpenAPI/Swagger)
- **.env.example** — ตัวอย่างไฟล์สำหรับตั้งค่าคอนฟิก เช่น Database URL, JWT, Email
- **Kong Gateway** — จัดการโดย PM (วรรธนโรจน์) ใน admin-service
- **README.md** — เอกสารอธิบายรายละเอียดการติดตั้งและใช้งาน Service

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

## 🦍 Kong API Gateway Integration

Service นี้เป็นส่วนหนึ่งของระบบ Microservices ที่ใช้ **Kong Gateway** เป็นจุดเข้าถึงหลัก (API Gateway)

> 👤 **ผู้ดูแล Kong Gateway**: วรรธนโรจน์ บุตรดี (Project Manager)  
> 💡 **คำแนะนำ**: หากมีปัญหาเกี่ยวกับ Kong setup, Routes หรือ Plugins โปรดติดต่อผู้ดูแล

### 🚀 Quick Start Options

#### Option 1: รันพร้อม Kong + Konga (แนะนำสำหรับ Integration Testing)

```bash
# จาก root directory (GameGear-Ecommerce/)
# Kong Gateway จัดการโดย PM ใน admin-service
# สำหรับการทดสอบ service เดี่ยว ให้ใช้:
go run cmd/api/main.go
```

**Services ที่จะรัน:**

- 🦍 Kong Gateway (port 8000, 8001)
- 🖥️ Konga Admin UI (port 1337)
- 👤 Users Service (port 8080)
- 🗄️ PostgreSQL Databases (Kong, Konga, Users)

#### Option 2: รัน Service เดี่ยว (สำหรับ Local Development)

```bash
# จาก service directory
go run cmd/api/main.go
```

**Service ที่จะรัน:**

- 👤 Users Service (port 8080)

**หมายเหตุ:** ต้องมี PostgreSQL database รันอยู่แล้ว หรือใช้ cloud database

### 🌐 การใช้ ngrok สำหรับ External Access

```bash
# รัน service
go run cmd/api/main.go

# เปิด terminal ใหม่ และรัน ngrok
ngrok http 8080
```

**ผลลัพธ์:**

- Service รันที่: `http://localhost:8080`
- External URL: `https://abc123.ngrok.io` (สำหรับ PM เรียกใช้)

---

### 📍 API Endpoints

| Access Method             | URL                                              | Use Case       | Note                      |
| ------------------------- | ------------------------------------------------ | -------------- | ------------------------- |
| **Via Kong Gateway**      | `http://localhost:8000/users/*`                  | ✅ Production  | เรียกผ่าน Gateway (แนะนำ) |
| **Direct Access**         | `http://localhost:8080/*`                        | 🔧 Development | เรียกตรง (Dev/Test only)  |
| **Swagger UI (via Kong)** | `http://localhost:8000/users/swagger/index.html` | 📖 API Docs    | ผ่าน Gateway              |
| **Swagger UI (direct)**   | `http://localhost:8080/swagger/index.html`       | 📖 API Docs    | เรียกตรง                  |

### 🔧 Kong Configuration

หากต้องการตั้งค่า Service นี้ใน Kong Gateway:

1. **เปิด Konga UI**: http://localhost:1337
2. **เพิ่ม Service**:
   ```
   Name: users-service
   Protocol: https
   Host: abc123.ngrok.io  (URL จาก ngrok ของ ณิชพน)
   Port: 443
   Path: /
   ```
3. **เพิ่ม Route**:
   ```
   Name: users-route
   Paths: /users
   Strip Path: false
   ```

### 🧪 ทดสอบการเชื่อมต่อ

```bash
# ทดสอบผ่าน Kong Gateway
curl http://localhost:8000/users/healthz

# ทดสอบเรียกตรง (Dev only)
curl http://localhost:8080/healthz

# ทดสอบ Login via Kong
curl -X POST http://localhost:8000/users/login \
  -H "Content-Type: application/json" \
  -d '{"email":"user@example.com","password":"password123"}'
```

### 📖 เอกสารเพิ่มเติม

สำหรับข้อมูลเพิ่มเติมเกี่ยวกับ Kong Gateway และการตั้งค่าทั้งระบบ:

- **Kong + Konga Setup Guide**: [Main README - Kong Setup](../Mini-Project-Golang/README.md#-quick-start-ติดตั้งและรัน-kong--konga)
- **System Architecture**: [Main README - Architecture](../Mini-Project-Golang/README.md#%EF%B8%8F-system-architecture-overview)
- **Ports Summary**: [Main README - Ports](../Mini-Project-Golang/README.md#-ports-summary)
- **Troubleshooting**: [Main README - Troubleshooting](../Mini-Project-Golang/README.md#-troubleshooting)

> 💡 **หมายเหตุ**: สำหรับการ setup Kong, Konga และ Plugins (CORS, JWT, Rate Limiting) โปรดดูเอกสารหลักที่ [Main README](../Mini-Project-Golang/README.md)

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

- README นี้อัปเดตให้สอดคล้องกับ **แนวทางหลักของโปรเจกต์ Mini-Project-Golang**
- รองรับทั้งการพัฒนา, ทดสอบ และเชื่อมต่อกับ Kong Gateway
- มีวิธีรันแบบ local, Docker, Swagger, Kong Integration และ Remote Dev พร้อมใช้งานจริง
