# Low-Level Design (LLD) for Mini FPL Clone Backend

---

## 1. Overview

This document describes the design, entities, routes, and data flow for the backend service implementing a Fantasy Premier League clone.

---

## 2. Project Structure

- `cmd/` — application entry point  
- `internal/` — core business logic and services  
- `pkg/` — reusable packages (e.g., middleware, utils)  
- `docs/` — design and API documentation

---

## 3. DB & Schema Structure

### User
- `ID` (UUID)  
- `Username` (string)  
- `Email` (string)  
- `PasswordHash` (string)  
- `Role` (string) — e.g., "user", "admin"  
- `CreatedAt`, `UpdatedAt`

### Player
- `ID` (UUID)  
- `Name` (string)  
- `Position` (string) — e.g., "GK", "DEF"  
- `Team` (string)  
- `Price` (float64)  
- `IsActive` (bool)  

### Gameweek
- `ID` (UUID)  
- `Number` (int)  
- `StartDate`, `EndDate` (timestamp)  
- `IsActive` (bool)

### UserTeam
- `ID` (UUID)  
- `UserID` (UUID)  
- `GameweekID` (UUID)  
- `PlayerIDs` ([]UUID) — selected players for that GW  
- `CreatedAt`, `UpdatedAt`

### PlayerPoints
- `ID` (UUID)  
- `PlayerID` (UUID)  
- `GameweekID` (UUID)  
- `Points` (int)

---

## 4. Routes and Access Control

| Route                       | Method | Access      | Description                        |
|-----------------------------|--------|-------------|------------------------------------|
| `/auth/register`            | POST   | Public      | Register new user                  |
| `/auth/login`               | POST   | Public      | Login and get JWT                  |
| `/user/me`                  | GET    | Protected   | Get logged-in user profile         |
| `/user/team`                | GET    | Protected   | Get user's team for current GW     |
| `/user/team`                | POST   | Protected   | Create/update user's team          |
| `/user/points/:gameweek_id` | GET    | Protected   | Get user's points for gameweek     |
| `/user/gameweeks`           | GET    | Protected   | List user's teams for all gameweeks|
| `/players`                  | GET    | Public      | Get all active players             |
| `/players/:id`              | GET    | Public      | Get details of a specific player   |
| `/gameweeks`                | GET    | Public      | List all gameweeks                 |
| `/gameweeks/live`           | GET    | Public      | Get current active gameweek        |
| `/points/:gameweek_id`      | GET    | Public      | Get all players' points for GW     |
| `/admin/players`            | POST   | Admin       | Add/update player data (future)    |
| `/admin/points/:gameweek_id`| POST   | Admin       | Upload points for gameweek (future)|
| `/admin/gameweeks`          | POST   | Admin       | Create/update gameweek (future)    |

---

## 5. Data Flow Summary

1. **User registration & authentication:**  
   - User registers → password is hashed → JWT issued on login.

2. **Player data:**  
   - Publicly accessible for building teams.

3. **Team selection:**  
   - Users select 15 players per gameweek; saved under `UserTeam`.

4. **Points:**  
   - Real player points stored in `PlayerPoints`.  
   - User points calculated by summing player points from their `UserTeam` for the gameweek.

5. **Gameweeks:**  
   - Managed centrally; users can view current and past gameweeks.

---

## 6. Security & Middleware

- JWT authentication middleware for `/user/*` routes.  
- Role-based middleware for `/admin/*` routes (to be implemented).  
- Input validation and rate limiting applied globally.

---

## 7. Extensibility

- Designed for adding mini-leagues, transfers, and live scoring.  
- Modular structure supports adding new features with minimal disruption.

---

End of document.
