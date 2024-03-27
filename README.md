# CRUD
Showcasing a Go project utilizing GORM and Gin frameworks, featuring RESTful endpoints for CRUD operations including user creation, retrieval, updating, and deletion.

## Features

- User Profile Creation
- Get All Users
- Get User By ID
- Update User By ID
- Delete User by ID

## End Points

- `GET /health/liveness`: Endpoint to check the liveness of the application.
- `GET /health/readiness`: Endpoint to check the readiness of the application.
- `GET /metrics`: Endpoint to expose application metrics.
- `POST /createUser`: Endpoint to create a new user profile.
- `GET /getAllUsers`: Endpoint to retrieve all user profiles.
- `GET /getUserById/:id`: Endpoint to retrieve a user profile by ID.
- `PUT /updateUserById/:id`: Endpoint to update a user profile by ID.
- `DELETE /deleteUserById/:id`: Endpoint to delete a user profile by ID.

## Tech Stack

**Server:** Go, gin, gorm

**Database:** MySql


## Usage

To use this API, you can send HTTP requests to the specified routes using tools like cURL, Postman, or any HTTP client library.


