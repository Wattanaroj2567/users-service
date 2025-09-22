# 👤 User Service

![Go](https://img.shields.io/badge/Go-1.24.6-00ADD8?style=for-the-badge\&logo=go)
![PostgreSQL](https://img.shields.io/badge/PostgreSQL-4169E1?style=for-the-badge\&logo=postgresql)
![GORM](https://img.shields.io/badge/GORM-B93527?style=for-the-badge)

Service สำหรับจัดการผู้ใช้และการยืนยันตัวตนทั้งหมดของโปรเจกต์ **GameGear E-commerce**

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

โครงสร้างของโปรเจกต์จัดเรียงตามหลัก Standard Go Layout โดยแต่ละส่วนมีหน้าที่ดังนี้

<table>
  <tr>
    <td width="50%">
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
```
```bash
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

สร้างไฟล์ชื่อ `.env` ในระดับบนสุดของโปรเจกต์ แล้วคัดลอกเนื้อหาข้างล่างนี้ไปวาง **(อย่าลืมแก้ `your_user` และ `your_password` ให้ถูกต้อง)**

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
