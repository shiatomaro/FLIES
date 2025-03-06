# User Management System

## Overview

**FLIES** is a full-stack web application built with **React (Material UI)** for the frontend and **Go (Gin)** for the backend. It provides authentication, role-based access control, and user management capabilities.

## Features

c

## Technologies Used

### **Frontend (React.js)**

- React.js with Material UI
- React Router for navigation
- Axios for API requests
- JWT Authentication for secure access

### **Backend (Go)**

- Go (Golang) for the backend API
- Gin Framework for routing and middleware
- GORM for database management
- PostgreSQL as the database
- JWT Authentication for secure user sessions

## User Roles & Permissions

| Role      | Permissions                                              |
| --------- | -------------------------------------------------------- |
| **User**  | View Profile                                             |
| **Admin** | Manage Users, Change Roles, Delete Users (except Admins) |

## Installation and Setup

### **Backend (Go)**

#### 1. Clone the repository:

```sh
git clone https://github.com/your-repo/user-management.git
cd user-management/backend
```

#### 2. Install dependencies:

```sh
go mod tidy
```

#### 3. Configure environment variables in `.env`:

```ini
DATABASE_URL=your_postgres_connection_string
JWT_SECRET=your_secret_key
```

#### 4. Run the server:

```sh
go run main.go
```

### **Frontend (React)**

#### 1. Navigate to the frontend directory:

```sh
cd ../frontend
```

#### 2. Install dependencies:

```sh
npm install
```

#### 3. Start the development server:

```sh
npm start
```

## API Endpoints

### **Authentication**

- `POST /login` - User login (returns JWT token)
- `POST /register` - Register a new user

### **Users (Admin Only)**

- `GET /users` - Get all users
- `PATCH /users/:id/role` - Update user role (Admin cannot change their own role)
- `DELETE /users/:id` - Delete a user (Admins cannot be deleted)

### **User Profile**

- `GET /profile` - Get current user profile
- `PATCH /profile` - Update user details

## Security Considerations

- **JWT Storage**: Stored in `localStorage` for persistence.
- **Token Expiry**: Tokens expire after **X hours** (define if necessary).
- **Role Protection**: Admin roles cannot be modified or removed.

## Contributing

Contributions are welcome! Feel free to submit **issues** and **pull requests**.

### Steps to Contribute

1. Fork the repository.
2. Create a new branch (`git checkout -b feature-name`).
3. Make your changes and commit (`git commit -m "Added new feature"`).
4. Push to your branch (`git push origin feature-name`).
5. Open a pull request.

## License

📜 This project is licensed under the **MIT License**.


