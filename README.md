# Users API

## Description
The Users API is a RESTful web service built with Go that allows users to manage user data. It provides endpoints for creating, retrieving, updating, and deleting user information. The API utilizes MongoDB for data storage and supports environment variable configuration for sensitive information.

## Table of Contents
- [Users API](#users-api)
  - [Description](#description)
  - [Table of Contents](#table-of-contents)
  - [Features](#features)
  - [Technologies Used](#technologies-used)
  - [Installation](#installation)
  - [Configuration](#configuration)
  - [Usage](#usage)
  - [API Endpoints](#api-endpoints)
    - [Insert an User](#insert-an-user)
    - [Get All Users](#get-all-users)
    - [Get an User By Id](#get-an-user-by-id)
    - [Update an User](#update-an-user)
    - [Delete an User](#delete-an-user)
    - [Note](#note)
  - [Security Note](#security-note)
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

### Insert an User

- **Endpoint:** `POST /api/user`
- **Description:** Insert a new user record.
- **Request Body:**
  ```json
  {
    "name": "Jane Smith",
    "email": "janesmith@example.com",
    "gender": "Female",
    "education": "Master of Business Administration",
    "bio": "Experienced business strategist with a focus on digital transformation and innovation."
  }
  ```

### Get All Users

- **Endpoint:** `GET /api/users`
- **Description:** Retrieves a list of all users.

### Get an User By Id

- **Endpoint:** `GET /api/user/{id}`
- **Description:** Retrieves a user record by ID.

### Update an User

- **Endpoint:** `PUT /api/user/{key}/{id}`
- **Description:** Updates the user by Id.


### Delete an User

- **Endpoint:** `DELETE /api/user/{key}/{id}`
- **Description:** Deletes a user record by ID.


### Note
- `PUT` and `DELETE` require the `key` parameter to be passed in the URL. The `key` is the value of the `SECRET_KEY` environment variable. 

The API supports CORS for cross-origin requests.


## Security Note

For the **key** parameter, you can contact me to get the key for `updating` and `deleting` users.

## Contributing
Contributions are welcome! Please feel free to submit a pull request or open an issue for discussion.

## License
This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.
