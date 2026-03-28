# 🚀 Go User API (Gin + Clean Architecture)

一個使用 Go + Gin 實作的簡單 User CRUD API，展示後端基礎能力與系統分層設計。

---

## 📌 專案特色

* 使用 **Gin** 建立 RESTful API
* 採用 **分層架構（Handler / Service / Model）**
* 使用 **Interface + Dependency Injection（DI）** 降低耦合
* 實作 **Custom Error + Middleware 統一錯誤處理**
* 符合基本 RESTful 設計原則

---

## 🧱 專案架構

```
project/
├── main.go              # 入口（組裝依賴）
├── handler/             # HTTP 層（處理 request/response）
│   └── user_handler.go
├── service/             # 商業邏輯層
│   └── user_service.go
├── model/               # 資料結構與錯誤定義
│   ├── user.go
│   └── error.go
├── middleware/          # 中介層（統一錯誤處理）
│   └── error_middleware.go
```

---

## 🧠 架構說明

### 🔹 Handler

* 負責接收 HTTP Request
* 呼叫 Service
* 回傳 HTTP Response

### 🔹 Service

* 負責商業邏輯
* 不依賴 Gin（與 HTTP 解耦）
* 透過 Interface 設計提高可測試性

### 🔹 Model

* 定義資料結構（User）
* 定義自訂錯誤（AppError）

### 🔹 Middleware

* 統一處理錯誤回傳
* 避免在每個 handler 重複寫 error handling

---

## 🔧 技術重點

### ✅ Interface + DI（依賴注入）

```go
type UserService interface {
    GetAllUsers() []model.User
}
```

* Handler 不依賴具體實作
* 可替換為 mock / DB 實作
* 提升可測試性與擴展性

---

### ✅ Custom Error

```go
type AppError struct {
    Code    int
    Message string
}
```

* 統一錯誤格式
* 搭配 middleware 處理

---

### ✅ Middleware Error Handling

```go
c.Error(err)
```

* handler 不直接回錯誤
* middleware 統一處理 response

---

## 📡 API 規格

### 🔹 取得所有使用者

```
GET /users
```

Response:

```json
[
  { "id": 1, "name": "John" },
  { "id": 2, "name": "Jane" }
]
```

---

### 🔹 建立使用者

```
POST /users
```

Request:

```json
{
  "name": "Alice"
}
```

Response:

```json
{
  "id": 3,
  "name": "Alice"
}
```

---

### 🔹 查詢單一使用者

```
GET /users/:id
```

Response:

```json
{
  "id": 1,
  "name": "John"
}
```

Error:

```json
{
  "error": "user not found"
}
```

---

## ▶️ 如何執行

```bash
go mod tidy
go run main.go
```

預設啟動於：

```
http://localhost:8080
```

---

## 🧪 測試 API（範例）

```bash
curl -X POST http://localhost:8080/users \
-H "Content-Type: application/json" \
-d '{"name":"Alice"}'
```

---

## 🎯 學習重點

本專案主要練習：

* Go 基本語法（struct / slice / pointer）
* RESTful API 設計
* Gin framework 使用
* 分層架構設計
* Interface 與 Dependency Injection
* Middleware 與錯誤處理

---

## 🚀 未來可擴充

* 串接資料庫（MySQL / PostgreSQL）
* 使用 ORM（GORM）
* 加入 JWT 驗證
* 加入 logging（ELK / Zap）
* 單元測試（mock service）

---

## 👩‍💻 Author

Jessie Chen
Backend Developer (Java → Go learning journey)

---
