# Go-Lang-RESTfulAPI

This repository implements a simple category REST API using the Go programming language.

## Features

- **CRUD Operations:**
    - Create new categories (`POST /api/categories`)
    - Retrieve all categories (`GET /api/categories`)
    - Get a specific category by ID (`GET /api/categories/:category_id`)
    - Update a category (`PUT /api/categories/:category_id`)
    - Delete a category (`DELETE /api/categories/:category_id`)
- **Authentication:**
    - Requires a valid `X-API-Key` header set to `TESTING` for authorization (**Note:** This is only for practice purpose , not secure for production).

## Dependencies

- `github.com/go-playground/validator` (v9.31.0 or later): Data validation
- `github.com/go-sql-driver/mysql` (v1.7.1 or later): MySQL database driver (if using a MySQL database)
- `github.com/joho/godotenv` (v1.5.1 or later): Environment variable management
- `github.com/julienschmidt/httprouter` (v1.3.0 or later): Routing library

## Installation

1. Clone the repository:

   ```bash
   git clone https://github.com/kevinjuliow/Go-Lang-RESTfulAPI

2. Open Project Directory

   ```
   cd Go-Lang-RESTfulAPI

## Environment
Remember to replace `DB_USERNAME` and `DB_PASSWORD` with your actual database credentials if you're using a MySQL database.

    DB_USERNAME=your_username
    DB_PASSWORD=your_password
    
