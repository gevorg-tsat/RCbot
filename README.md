# Random Coffee Telegram Bot

## Requirements:
- Docker

or

- Go v1.21
- Postgres v15+

## Configuration:
1) Open [secrets/.env](secrets/.env)
2) Add Bot-Token info
3) Add your TelegramID (not username!!) to be admin
4) Also, you can change other field, if you want

## How to run:
If you have docker:
```bash
docker-compose up
```
Else:
1) Run your postgres server
2) 
```bash
go run cmd/main.go
```

## Commands for bot
### User commands
- /start -- to participate in RandomCoffee
- /stop -- stop participating in RandomCoffee
### Admin commands
- /start_event -- create pairs for RC and send every member a message with info about partner