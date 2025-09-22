# 👤 User Service

Service สำหรับจัดการผู้ใช้และการยืนยันตัวตนทั้งหมดของโปรเจกต์ **GameGear E-commerce**

---

## 🏛️ Architectural Design

Service นี้ถูกออกแบบให้เป็น **"เจ้าของข้อมูล" (Data Owner)** โดยมีหน้าที่รับผิดชอบในการจัดการข้อมูลและตารางในฐานข้อมูลที่เกี่ยวกับผู้ใช้ (`users`, `password_reset_tokens`) โดยตรง
Service อื่นๆ ที่ต้องการเข้าถึงข้อมูลเหล่านี้ จะต้องเรียกใช้งานผ่าน API ที่ Service นี้มีให้เท่านั้น เพื่อรักษาความถูกต้องของข้อมูลและแบ่งหน้าที่ความรับผิดชอบอย่างชัดเจน

---

## ✨ Features & Responsibilities

Service นี้มีหน้าที่รับผิดชอบฟีเจอร์หลัก 4 ส่วน:

* **User Registration** (`POST /api/auth/register`)

  * สร้างบัญชีผู้ใช้ใหม่ด้วย `username`, `email`, และ `password`
  * ตรวจสอบข้อมูลซ้ำซ้อนและเข้ารหัสรหัสผ่านก่อนบันทึกลงฐานข้อมูล

* **Authentication** (`POST /api/auth/login`, `POST /api/auth/logout`)

  * ตรวจสอบ `username/email` และ `password` เพื่อยืนยันตัวตน
  * สร้างและส่ง JSON Web Token (JWT) กลับไปให้ Client เพื่อใช้ในการยืนยันตัวตนครั้งถัดไป

* **Profile Management** (`GET /api/user/profile`, `PUT /api/user/profile`)

  * อนุญาตให้ผู้ใช้ที่ล็อกอินแล้วสามารถดูและอัปเดตข้อมูลส่วนตัวได้ (เช่น `display_name`, `email`)

* **Password Reset** (`POST /api/auth/forgot-password`, `POST /api/auth/reset-password`)

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
├── .env
├── go.mod
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
<li><b>.env</b>: ไฟล์เก็บ Configuration</li>
</ul>
</td>
</tr>
</table>

---

## 🚀 Getting Started (Step-by-step)

ทำตามขั้นตอนทีละสเต็ปเพื่อตั้งค่าและรัน Service ในเครื่องของคุณ

### Step 1 — Clone the Repository (Standard) สำหรับใช้งานได้จริง

```bash
git clone https://github.com/Wattanaroj2567/users-service.git
cd users-service
```

### Step 1 (Alt) — Direct Branch Clone สำหรับผู้พัฒนาให้เริ่มต้นที่นี่เท่านั้น
คำสั่งนี้จะ Clone โปรเจกต์ทั้งหมดมาที่เครื่องของคุณ โดยจะได้ Branch develop เป็นค่าเริ่มต้น
```bash
git clone -b develop https://github.com/Wattanaroj2567/users-service.git
cd users-service
```

### Step 2 — Install Dependencies

> คำสั่งนี้จะดาวน์โหลด dependencies ทั้งหมดตาม `go.mod`

```bash
go mod tidy
```

### Step 3 — Setup Database

ตรวจสอบให้แน่ใจว่า PostgreSQL Server ทำงานอยู่ แล้วสร้างฐานข้อมูล (ถ้ายังไม่มี)

**(A) ใช้ SQL โดยตรง**

```sql
CREATE DATABASE gamegear_db;
```

**(B) ใช้ psql ผ่าน bash (one‑liner)**

```bash
psql -U your_user -h localhost -p 5432 -c "CREATE DATABASE gamegear_db;"
```

### Step 4 — Configure Environment Variables

สร้างไฟล์ `.env` ที่ root ของโปรเจกต์ และใส่ค่าตามตัวอย่าง (แก้ `your_user` และ `your_password` ให้ถูกต้อง)

```env
# PostgreSQL Database Connection URL
DATABASE_URL="host=localhost user=your_user password=your_password dbname=gamegear_db port=5432 sslmode=disable"
```

### Step 5 — Run the Service

รันเซิร์ฟเวอร์

```bash
go run cmd/api/main.go
```

> เมื่อรันคำสั่งนี้ ระบบจะทำการ **migrate** ตารางที่จำเป็นทั้งหมด และเซิร์ฟเวอร์จะเริ่มที่ `http://localhost:8080`

## 🤝 Remote Development (Working from Different Locations)

ต้องการให้เพื่อนร่วมทีมเข้าถึง `users-service` ของคุณจากภายนอก? ใช้ **ngrok** ตามขั้นตอนนี้

### Step 1 — ติดตั้ง/ดาวน์โหลด ngrok

ไปที่ [ngrok.com](https://ngrok.com) และติดตั้งตามระบบปฏิบัติการของคุณ

### Step 2 — รัน User Service ในเครื่อง

```bash
go run cmd/api/main.go
```

### Step 3 — เปิดอุโมงค์ไปยังพอร์ต 8080

เปิด Terminal ใหม่แล้วรัน

```bash
ngrok http 8080
```

### Step 4 — แชร์ URL ให้ทีม

คัดลอก URL ที่ขึ้นต้นด้วย `https://...` ส่งให้เพื่อนร่วมทีม

### Step 5 — เพื่อนนำ URL ไปตั้งใน .env (ฝั่ง admin-service)

```env
# .env file on admin-service
USER_SERVICE_URL="<THE_NGROK_URL_YOU_SENT>"
```
