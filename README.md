# GoChat

![Go](https://img.shields.io/badge/Go-1.20+-blue?logo=go&logoColor=white)
![Gin](https://img.shields.io/badge/Framework-Gin-green?logo=go&logoColor=white)
![WebSockets](https://img.shields.io/badge/RealTime-WebSockets-orange?logo=socket.io&logoColor=white)
![License](https://img.shields.io/badge/License-MIT-yellow)

GoChat is a real-time chat system built with **Go** and **WebSockets**, enabling instant peer-to-peer messaging between users.
---
## 🛠 Tech Stack

- **Go** – Core backend language  
- **Gin** – HTTP web framework  
- **WebSockets** – Real-time communication for peer-to-peer messaging  
---
## ✨ Features

- 🔗 Real-time peer-to-peer messaging using WebSockets  
- ⚡ Built with Go + Gin for high performance  
- 🖥 Simple `index.html` client for testing messages  
- 🗑 Ephemeral: messages are not stored (cleared once you close the page)  
---
## 🚀 Getting Started

### Prerequisites
- Go 1.20+  

### Installation & Run
```bash
git clone https://github.com/yourusername/gochat.git
cd gochat
go run main.go
```
The server will start at:
```bash
👉 http://localhost:8080
```
#### Client
**Open index.html in your browser and start chatting in real time!
Messages are ephemeral — once you close the page, the chat is gone.**
---
## ☁️ Deployment

You can deploy GoChat in two ways:

### 1. Docker
```bash
# Build the Docker image
docker build -t gochat .

# Run the container
docker run -p 8080:8080 gochat
```
### 2. Render / Railway

1.Directly deploy the Go app using native Go runtime support.
2.Set the PORT environment variable (default: 8080).

## 💬 Demo / Usage

1. Start the server:
   ```bash
   go run main.go
   ```
**2.Open index.html in a browser.
3.Open another tab or browser window with the same index.html.
4.Type a message in one tab → it instantly appears in the other tab.
5. 🔄 All communication happens over WebSockets.
6.⚠️ Messages are not stored — they vanish once you close the page.**
---
## 🛣 Roadmap

Planned future improvements:

- 🔑 Add authentication (JWT or OAuth)  
- 🗄 Persist messages with a database (Postgres or Redis)  
- 👥 Enable group chats and chat rooms  
- 📱 Add typing indicators and read receipts  
- 📊 Add analytics (active users, message volume, etc.)  
- ☁️ Deploy with WebSocket scaling support (Redis Pub/Sub or NATS)  
---
## 🤝 Contributing

Contributions, issues, and feature requests are welcome!  
Feel free to fork this repo and submit pull requests.

## 📜 License

This project is licensed under the **MIT License** – see the [LICENSE](LICENSE) file for details.


