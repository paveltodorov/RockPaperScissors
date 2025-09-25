# Rock Paper Scissors Betting Game

A RESTful API for a multiplayer Rock Paper Scissors game with a betting system, built in Go using Gin framework.

## Features

- **User Management**: Login/register with username validation
- **Betting System**: Wager coins on Rock Paper Scissors matches
- **Challenge System**: Create, accept, and decline game challenges
- **Funds Management**: Deposit and withdraw coins
- **Real-time Game Logic**: Automatic winner determination and money transfer
- **AI Players**: Play against AI opponents with different strategies

## API Endpoints

### Users
- `POST /login` - Login or register a new user
- `POST /logout` - Logout confirmation
- `GET /users` - List all users

### Funds
- `POST /deposit` - Add coins to user balance
- `POST /withdraw` - Remove coins from user balance

### Challenges
- `POST /challenges` - Create a new challenge
- `GET /challenges` - List all challenges
- `POST /challenges/accept` - Accept a challenge
- `POST /challenges/decline` - Decline a challenge

### AI Players
- `POST /ai/create` - Create a new AI player
- `GET /ai/list` - List all AI players
- `POST /ai/challenge` - Challenge an AI player

## Game Rules

- Both players bet the same amount
- Winner takes the total pot (bet × 2)
- Ties result in no money transfer
- Valid moves: rock, paper, scissors

## Tech Stack

- **Language**: Go 1.24+
- **Framework**: Gin (HTTP web framework)
- **Architecture**: Clean Architecture with Repository pattern
- **Storage**: In-memory (development) with interface for easy database integration

## Quick Start

```bash
# Install dependencies
go mod tidy

# Run the server
go run main.go

# Server runs on http://localhost:8080
```

## Example Usage

```bash
# Login/register a user
curl -X POST http://localhost:8080/login \
  -H "Content-Type: application/json" \
  -d '{"username":"alice","password":"secret"}'

# list users
 curl -X GET http://localhost:8080/users   -H "Content-Type: application/json"

# get stats for a specific user
curl GET  http://localhost:8080/users/stats -H "X-User-ID: 2"

# Create a challenge
curl -X POST http://localhost:8080/challenges \
  -H "Content-Type: application/json" \
  -d '{"challenger_id":1,"opponent_id":2,"bet":50,"move":"rock"}'

# Accept the challenge
curl -X POST http://localhost:8080/challenges/accept \
  -H "Content-Type: application/json" \
  -d '{"challenge_id":1,"opponent_id":2,"move":"paper"}'

curl -X GET http://localhost:8080/challenges

curl -X GET http://localhost:8080/challenges/pending -H "X-User-ID: 1"

curl -X POST http://localhost:8080/deposit   -H "Content-Type: application/json"   -d '{"user_id":1, "amount": 30}'

curl -X POST http://localhost:8080/withdraw   -H "Content-Type: application/json"   -d '{"user_id":1, "amount": 30}'

# Create an AI player
curl -X POST http://localhost:8080/ai/create \
  -H "Content-Type: application/json" \
  -d '{"strategy":"smart"}'

# Challenge an AI player
curl -X POST http://localhost:8080/ai/challenge \
  -H "Content-Type: application/json" \
  -d '{"ai_id":3,"player_id":1,"bet":25,"strategy":"smart"}'
```
