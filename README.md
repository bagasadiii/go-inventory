# Inventory Management System (Under Development)

This project is a work in progress and aims to showcase my skills in Go (Golang), Gin Gonic framework, PostgreSQL, and JWT authentication. The system is built using a layered architecture, with the goal of implementing a robust and scalable solution. While currently featuring basic CRUD operations for user management, additional functionalities are planned for future development.

## Project Features
- **CRUD Operations**: Currently, the project supports basic Create, Read, Update, and Delete (CRUD) operations for user management.
- **JWT Authentication**: JSON Web Token (JWT) is used for secure authentication and authorization.
- **Gin Gonic Framework**: The web server is built using the Gin Gonic framework, ensuring fast and efficient routing.
- **PostgreSQL Database**: PostgreSQL is used as the database for storing user data and other information.
- **Layered Architecture**: The application is structured using a layered architecture pattern to ensure maintainability and scalability.
- **Middleware**: Custom middleware is used for logging, authentication, and other essential tasks.
  
## Planned Features
- More features related to **inventory management** (CRUD for products, orders, etc.)
- Advanced **role-based access control (RBAC)**
- **Rate-limiting** middleware to handle traffic more efficiently
- Integration with **external APIs** for enhanced functionality

## Technologies Used
- **Go (Golang)**: The primary programming language used to build the application.
- **Gin Gonic**: A fast and lightweight web framework for Go, providing a simple API for handling HTTP requests and responses.
- **PostgreSQL**: A powerful, open-source relational database used to store application data.
- **JWT (JSON Web Token)**: For secure user authentication and session management.
- **Gin Middleware**: Custom middleware for logging, authentication, and error handling.

## Setup & Installation

To run this project locally, follow these steps:

1. **Clone the Repository**:
   ```bash
   git clone https://github.com/bagasadiii/inventory-app.git
   cd inventory-app
