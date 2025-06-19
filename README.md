# BroadCastServer
https://github.com/kalyanipd/BroadCastServer.git
# 📡 Go WebSocket Broadcast Server

A simple CLI-based WebSocket broadcast server and client written in Go. This project demonstrates real-time communication between multiple clients using WebSockets—ideal for learning chat server mechanics, pub-sub, or live notifications.

---

 Project Page

👉 [Project Page URL](https://roadmap.sh/projects/broadcast-server) 
 

## 🚀 Features

- Start a WebSocket broadcast server via CLI.
- Connect multiple clients via CLI.
- Messages sent by one client are broadcasted to all other connected clients.
- Clean and modular code using `cobra` CLI and `gorilla/websocket`.

---

## 📦 Requirements

- Go 1.17+
- Internet access to fetch dependencies
- Terminal (to run CLI commands)

---

## 🔧 Installation

1. **Clone the repository**:

   ```bash
   git clone https://github.com/kalyanipd/BroadCastServer.git
   cd broadcast-server

2. **Initialize Go modules**:
    go mod tidy

3. **How to Run**:
    1. Start the Broadcast Server
        go run main.go start

    2. Connect as a Client
        // for multiple clients
        Open one or more terminals and run:
        go run main.go connect

    3. To exit the client:
        Type "exit" and press Enter.