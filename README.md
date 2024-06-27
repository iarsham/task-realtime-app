# 🚀 Go WebSocket Realtime App

![Chat Application](https://github.com/Marcel-MD/rooms-go-api/assets/87933412/0f9a521e-306d-4aed-8e30-4d35c2a669cb)
![Chat Application](https://github.com/Marcel-MD/rooms-go-api/assets/87933412/b943f3ee-6f37-41d7-8f07-6ba8c74b3916)

Welcome to the Go WebSocket WebSocket Realtime App! This is a simple chatroom application implemented using WebSockets
in the Go
language. The application allows multiple users to join a chatroom and exchange messages in real-time.

## ✨ Features

- Send & receive messages instantly in the chatroom. ➡️
- Secure login. Log in with your username for a personalized chat space.🔐
- Get push notifications for important updates.📢
- Caching stores frequently used data for a smoother experience.🌟
- Microservices power! Breaks down the app for better scaling and easier updates.📩
- RabbitMQ message queue! Delivers messages reliably between services.💬
- Rock-solid tests! Ensures everything works perfectly with TDD & unit testing. ✅

## ⚙️ Prerequisites

Before running this application, make sure you have the following installed:

- Go (1.22 or higher) 🐹 (https://go.dev/)
- Docker & Docker-Compose ☁️(https://www.docker.com/)

## 📥 Installation

1. **Clone this repository to your local machine:**

   ```shell
   git clone https://github.com/iarsham/task-realtime-app

2. **Change to the project directory:**
   ```shell
   cd ./task-realtime-app

3. **Create .env file and fill it based on .env-sample:**
   ```shell
   touch .env

4. **Build and Start the application:**
   ```shell
   make run-prod

## 🚀 Usage

- Auth Service: http://localhost:8000/docs. RestAPI
- Notification Service: http://localhost:8001/notification. Websocket
- Chat & Room Service: http://localhost:8002/docs. Websocket & RestAPI

![Message Example](https://github.com/Marcel-MD/rooms-go-api/assets/87933412/a78010f8-675f-4047-ac5e-1f6229256bd4)
![Notfication Example](https://github.com/Marcel-MD/rooms-go-api/assets/87933412/885ae8e6-0886-4bb6-8366-40ec65dea042)

## 🤝 Contributing

Contributions are welcome! If you find any issues or want to enhance the functionality of this application, feel free to
open an issue or submit a pull request. Please make sure to follow the Contributing Guidelines when contributing.

## 📄 License

This project is licensed under the MIT License.

## 📞 Contact

If you have any questions or need any assistance, feel free to reach out:

- Email: arshamdev2001@gmail.com