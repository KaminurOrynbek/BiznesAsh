# BiznesAsh

BiznesAsh is a digital platform built to empower aspiring entrepreneurs in Kazakhstan. It simplifies business registration by consolidating essential information, expert guidance, and step-by-step procedures into one intuitive and accessible platform.

> *Start smart, grow strong.*

---

## Project Overview  📌

In Kazakhstan, youth and returning migrants are increasingly interested in launching private businesses. However, navigating business registration (e.g., opening an ИП or ТОО) is often confusing and time-consuming due to fragmented information spread across:

- [egov.kz](https://egov.kz)
- [atameken.kz](https://atameken.kz)
- Social media
- Informal communities

BiznesAsh solves this by centralizing business-starting knowledge in a structured, localized, and community-driven platform.

---

## Technologies Used ⚙️
 
- Backend: Golang, gRPC 
- Database: PostgreSQL, Redis  
- DevOps: Docker 
- Testing: 
- Authentication: JWT, bcrypt  
- Message queue: NATS
- Architecture: Microservices

---

## How to Run Locally 🧪

```
git clone https://github.com/KaminurOrynbek/BiznesAsh.git
```
```
cd BiznesAsh
```
- Configure your database in .env
- Create docker container for redis and nats
```
cd UserService
```
```
go run cmd/user/main.go
```
```
cd ContentService
```
```
go run cmd/content/main.go
```
```
cd NotificationService.go
```
```
go run cmd/notification/main.go
```
```
cd APIGateway
```
```
go run cmd/gateway/main.go
```
![image](https://github.com/user-attachments/assets/1c228ba2-2fa0-4829-8a73-83de11ae6b0f)
![image](https://github.com/user-attachments/assets/8dcbfda2-b643-4216-843d-bdf16f0a144a)
![image](https://github.com/user-attachments/assets/6a35ff50-4579-4287-9cc0-000d96ffb6f8)
![image](https://github.com/user-attachments/assets/1675a63c-2347-474d-be0c-c41d9fd5adb4)

