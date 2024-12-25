# SignIn  Web Application

This is a simple web application built using **Go**, **Fiber**, and **MongoDB Atlas**. It allows users to sign up by entering a username and password, and the user data is stored in a MongoDB Atlas database.

## Table of Contents

- [Prerequisites](#prerequisites)
- [Installation](#installation)
- [Running the Application](#running-the-application)
- [How it Works](#how-it-works)
- [Folder Structure](#folder-structure)
- [MongoDB Atlas Configuration](#mongodb-atlas-configuration)
- [Contributing](#contributing)

## Prerequisites

Before you begin, ensure you have the following installed on your system:

- **Go** (Go 1.18 or higher): [Install Go](https://go.dev/dl/)
- **MongoDB Atlas** account: [Sign up for MongoDB Atlas](https://www.mongodb.com/cloud/atlas)

## Installation

1. **Clone the repository**:
    ```bash
    git clone https://github.com/your-username/signin-signout-go.git
    cd signin-signout-go
    ```

2. **Install dependencies**:
    You will need to install the required Go packages:
    ```bash
    go get github.com/gofiber/fiber/v2
    go get go.mongodb.org/mongo-driver/bson
    go get go.mongodb.org/mongo-driver/mongo
    go get go.mongodb.org/mongo-driver/mongo/options
    ```

3. **Set up MongoDB Atlas**:
    - Create a MongoDB Atlas account (if you don't have one already).
    - Create a cluster and get your connection string for the cluster.
    - Add your local machine’s IP address to the **Network Access** IP Whitelist in MongoDB Atlas.
    - Update the `main.go` file with your MongoDB Atlas connection string:
    ```go
    const MongoDBLink = "mongodb+srv://<username>:<password>@cluster0.lpuxv.mongodb.net/testdb?retryWrites=true&w=majority"
    ```

4. **Create `go.mod` file**:
    If you haven’t initialized a Go module, run:
    ```bash
    go mod init signin-signout-go
    ```

5. **Tidy up dependencies**:
    Run the following command to clean up and ensure all dependencies are correctly installed:
    ```bash
    go mod tidy
    ```

## Running the Application

1. To run the application locally, execute the following command:
    ```bash
    go run main.go
    ```

2. The application will start a web server on `http://localhost:4000`.

3. Open your browser and navigate to `http://localhost:4000` to view the sign-in form.

4. Fill out the form with a username and password, and click **Submit**. This will store the user data in the `testdb` database under the `users` collection.

## How it Works

- The application serves an HTML form at the root (`/`) endpoint where users can submit their username and password.
- When the form is submitted, the data is sent via **POST** to the `/submit` endpoint.
- The user data is then inserted into MongoDB Atlas under the `testdb` database in the `users` collection.

### Database Schema

The MongoDB schema for the `users` collection contains the following fields:

- `username` (string): The username provided by the user.
- `password` (string): The password provided by the user.

## Folder Structure

```bash
.
├── main.go              # Main Go application file
├── index.html           # HTML form for user sign-in
└── README.md            # This readme file
