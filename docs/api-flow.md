# 📡 API Flow: Mini FPL Clone

This document describes all backend API endpoints with their purpose and access level.

---

## 🔐 Auth (Public)

| Method | Endpoint       | Description            |
|--------|----------------|------------------------|
| POST   | /auth/register | Register a new user     |
| POST   | /auth/login    | Login and get JWT token |

---

## 👤 User (Protected - JWT required)

| Method | Endpoint                  | Description                                 |
|--------|---------------------------|---------------------------------------------|
| GET    | /user/me                  | Get logged-in user profile                  |
| GET    | /user/team                | Get user's selected team for current GW    |
| POST   | /user/team                | Create or update user's team for current GW|
| GET    | /user/points/:gameweek_id | Get user's team points for specified GW    |
| GET    | /user/gameweeks           | List user's teams for all gameweeks         |

---

## 🧑‍🤝‍🧑 Players (Public)

| Method | Endpoint     | Description                  |
|--------|--------------|------------------------------|
| GET    | /players     | Get all available players    |
| GET    | /players/:id | Get details for a player     |

---

## 📅 Gameweeks (Public)

| Method | Endpoint       | Description                     |
|--------|----------------|---------------------------------|
| GET    | /gameweeks     | Get all gameweeks               |
| GET    | /gameweeks/live| Get current active gameweek    |

---

## 📈 Points (Public)

| Method | Endpoint             | Description                          |
|--------|----------------------|------------------------------------|
| GET    | /points/:gameweek_id | Get all players' points for a gameweek |

---

## 🛡️ Admin (Protected, JWT + Role Check, Future)

| Method | Endpoint                   | Description                          |
|--------|----------------------------|------------------------------------|
| POST   | /admin/players             | Add or update player data           |
| POST   | /admin/points/:gameweek_id | Upload points data for a gameweek   |
| POST   | /admin/gameweeks           | Create or update gameweek details   |

---

## 🧪 Example User Flow:

1. `POST /auth/register` → register & get JWT  
2. `GET /players` → fetch all players  
3. `POST /user/team` → create/update user team  
4. After GW ends: `GET /user/points/:gameweek_id` → get user's points  

---

## ⚠️ Notes:

- `/user/*` routes require valid JWT authentication.
- Admin routes require JWT + role-based authorization (middleware to be implemented).
- Rate limiting, input validation, and logging handled via middleware.
