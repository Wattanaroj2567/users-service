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

## 🚀 Getting Started

ทำตามขั้นตอนเหล่านี้เพื่อตั้งค่าโปรเจกต์และรัน Service ในเครื่องของคุณ

### 1. Clone the Repository

```bash
git clone https://github.com/Wattanaroj2567/users-service.git
cd users-service
```

### 2. Install Dependencies

คำสั่งนี้จะดาวน์โหลด Library ที่จำเป็นทั้งหมดที่ระบุไว้ใน `go.mod`

```bash
go mod tidy
```

### 3. Setup Database

* ตรวจสอบให้แน่ใจว่า PostgreSQL Server ของคุณทำงานอยู่
* สร้างฐานข้อมูลสำหรับโปรเจกต์ โดยใช้คำสั่ง SQL นี้:

```sql
CREATE DATABASE gamegear_db;
```

### 4. Configure Environment Variables

สร้างไฟล์ชื่อ `.env` ในระดับบนสุดของโปรเจกต์ แล้วคัดลอกเนื้อหาข้างล่างนี้ไปวาง *(อย่าลืมแก้ `your_user` และ `your_password` ให้ถูกต้อง)*

```env
# PostgreSQL Database Connection URL
DATABASE_URL="host=localhost user=your_user password=your_password dbname=gamegear_db port=5432 sslmode=disable"
```

### 5. Run the Service

รันคำสั่งนี้ใน Terminal เพื่อเริ่มการทำงานของเซิร์ฟเวอร์

```bash
go run cmd/api/main.go
```

* เมื่อรันคำสั่งนี้ โปรแกรมจะทำการ **Migrate** หรือสร้างตารางที่จำเป็นในฐานข้อมูลให้โดยอัตโนมัติ
* เซิร์ฟเวอร์จะเริ่มต้นทำงานที่ `http://localhost:8080`

---

## 🤝 Remote Development (Working from Different Locations)

เมื่อเพื่อนร่วมทีมต้องการเรียกใช้ Service นี้จากเครื่องของพวกเขา (เช่น `admin-service` ต้องการเรียกใช้ `users-service`) เราจะใช้ **ngrok** เพื่อสร้าง URL สาธารณะชั่วคราว

### How to Share Your Service

1. ดาวน์โหลด **ngrok** จาก [ngrok.com](https://ngrok.com)
2. รัน Service ของคุณตามปกติ

```bash
go run cmd/api/main.go
```

3. เปิด Terminal ใหม่ขึ้นมา แล้วรันคำสั่งนี้เพื่อสร้าง "อุโมงค์" มายัง Port 8080 ของคุณ:

```bash
ngrok http 8080
```

4. ngrok จะแสดง URL สาธารณะ (ขึ้นต้นด้วย `https://...`) ขึ้นมา ให้คัดลอก URL นี้แล้วส่งให้เพื่อนร่วมทีมของคุณ
5. เพื่อนของคุณจะนำ URL ที่ได้ไปใส่ในไฟล์ `.env` ของ Service ที่เขากำลังพัฒนาอยู่ (เช่น `admin-service`)

```env
# .env file on admin-service
USER_SERVICE_URL="<THE_NGROK_URL_YOU_SENT>"
```
