# Users API

## Description
The Users API is a RESTful web service built with Go that allows users to manage user data. It provides endpoints for creating, retrieving, updating, and deleting user information. The API utilizes MongoDB for data storage and supports environment variable configuration for sensitive information.

## Table of Contents
- [Features](#features)
- [Technologies Used](#technologies-used)
- [Installation](#installation)
- [Configuration](#configuration)
- [Usage](#usage)
- [API Endpoints](#api-endpoints)
- [Contributing](#contributing)
- [License](#license)

## Features
- Create, Read, Update, and Delete (CRUD) operations for user data.
- Environment variable configuration using dotenv for sensitive information.
- MongoDB as the database for storing user data.
- Built with Go and utilizes the Gorilla Mux router for handling HTTP requests.

## Technologies Used
- Go (Golang)
- MongoDB
- Gorilla Mux
- dotenv (for environment variables)

## Installation
1. Clone the repository:
   ```bash
   git clone https://github.com/paulaneesh7/User-API-Golang.git
   ```
2. Navigate to the project directory:
   ```bash
   cd User-API-Golang
   ```
3. Install the required Go packages:
   ```bash
   go mod tidy
   ```

## Configuration
1. Create a `.env` file in the root directory of the project.
2. Add your MongoDB connection string:
   ```plaintext
   MONGODB_URI=mongodb://username:password@host:port/database
   ```

## Usage
1. Start the server:
   ```bash
   go run main.go
   ```
2. The API will be available at `http://localhost:8080`.

## API Endpoints
- `POST /api/user` - Create a new user.
- `GET /api/users` - Retrieve all users.
- `GET /api/user/{id}` - Retrieve a user by ID.
- `PUT /api/user/{id}` - Update a user by ID.
- `DELETE /api/user/{id}` - Delete a user by ID.

## Contributing
Contributions are welcome! Please feel free to submit a pull request or open an issue for discussion.

## License
This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.
