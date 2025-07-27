# Go E-commerce Backend API

[![Go Version](https://img.shields.io/badge/Go-1.24.5-blue.svg)](https://golang.org/)
[![Gin Framework](https://img.shields.io/badge/Gin-1.10.1-green.svg)](https://gin-gonic.com/)
[![GORM](https://img.shields.io/badge/GORM-1.30.1-orange.svg)](https://gorm.io/)
[![License](https://img.shields.io/badge/License-MIT-yellow.svg)](LICENSE)

A robust and scalable e-commerce backend API built with Go, featuring clean architecture, comprehensive middleware, and modern development practices.

## 📋 Table of Contents

- [Features](#-features)
- [Architecture](#-architecture)
- [Tech Stack](#-tech-stack)
- [Project Structure](#-project-structure)
- [Prerequisites](#-prerequisites)
- [Installation](#-installation)
- [Configuration](#-configuration)
- [Running the Application](#-running-the-application)
- [API Endpoints](#-api-endpoints)
- [Database](#-database)
- [Development](#-development)
- [Testing](#-testing)
- [Docker](#-docker)
- [Contributing](#-contributing)
- [License](#-license)

## ✨ Features

- **Clean Architecture**: Follows clean architecture principles with clear separation of concerns
- **RESTful API**: RESTful endpoints with proper HTTP status codes
- **Authentication & Authorization**: JWT-based authentication with role-based access control
- **Database Integration**: MySQL with GORM ORM for data persistence
- **Caching**: Redis integration for improved performance
- **Logging**: Structured logging with Zap logger and log rotation
- **Configuration Management**: Viper for flexible configuration management
- **Middleware**: CORS, rate limiting, error handling, and authentication middleware
- **Docker Support**: Complete Docker setup with docker-compose
- **API Documentation**: Comprehensive API documentation
- **Testing**: Unit and integration tests
- **Monitoring**: Health checks and performance monitoring

## 🏗️ Architecture

The project follows a clean architecture pattern with the following layers:

```
┌─────────────────────────────────────────────────────────────┐
│                    Presentation Layer                      │
│  ┌─────────────┐  ┌─────────────┐  ┌─────────────┐      │
│  │ Controllers │  │ Middlewares │  │   Routers   │      │
│  └─────────────┘  └─────────────┘  └─────────────┘      │
└─────────────────────────────────────────────────────────────┘
┌─────────────────────────────────────────────────────────────┐
│                     Business Layer                         │
│  ┌─────────────┐  ┌─────────────┐  ┌─────────────┐      │
│  │   Services  │  │   Business  │  │   Domain    │      │
│  │             │  │   Logic     │  │   Models    │      │
│  └─────────────┘  └─────────────┘  └─────────────┘      │
└─────────────────────────────────────────────────────────────┘
┌─────────────────────────────────────────────────────────────┐
│                    Data Access Layer                       │
│  ┌─────────────┐  ┌─────────────┐  ┌─────────────┐      │
│  │ Repository  │  │     PO      │  │   Database  │      │
│  │   Pattern   │  │ (Persistence│  │   (MySQL)   │      │
│  │             │  │   Objects)  │  │             │      │
│  └─────────────┘  └─────────────┘  └─────────────┘      │
└─────────────────────────────────────────────────────────────┘
```

### Layer Responsibilities:

- **Controllers**: Handle HTTP requests and responses
- **Services**: Contain business logic and orchestrate operations
- **Repository**: Abstract data access layer
- **PO (Persistence Objects)**: Database models and entities
- **Middlewares**: Cross-cutting concerns like authentication, logging, CORS

## 🛠️ Tech Stack

### Core Framework & Libraries
- **Go 1.24.5**: Latest stable version of Go
- **Gin 1.10.1**: High-performance HTTP web framework
- **GORM 1.30.1**: Full-featured ORM for Go
- **Viper 1.20.1**: Configuration management
- **Zap 1.27.0**: Structured logging

### Database & Caching
- **MySQL 8.0**: Primary database
- **Redis 7.0**: Caching and session storage
- **GORM MySQL Driver**: Database driver

### Development Tools
- **Docker & Docker Compose**: Containerization
- **Make**: Build automation
- **UUID**: Unique identifier generation

### Testing
- **Testify**: Testing framework
- **Coverage**: Code coverage reporting

## 📁 Project Structure

```
go-ecommerce-backend-api/
├── cmd/                    # Application entry points
│   ├── cli/               # CLI commands
│   ├── cronjob/           # Scheduled jobs
│   └── server/            # Main server application
├── config/                # Configuration files
│   └── local.yml          # Local environment config
├── database_docker/        # Docker database volumes
├── docker-compose.yml      # Docker services configuration
├── global/                 # Global variables and state
├── init/                   # Database initialization scripts
├── internal/               # Private application code
│   ├── controller/         # HTTP request handlers
│   ├── initialize/         # Application initialization
│   ├── middlewares/        # HTTP middlewares
│   ├── models/             # Domain models
│   ├── po/                 # Persistence objects
│   ├── repo/               # Data access layer
│   ├── routers/            # Route definitions
│   └── service/            # Business logic layer
├── log/                    # Application logs
├── migrations/             # Database migrations
├── pkg/                    # Public packages
│   ├── logger/             # Logging utilities
│   ├── response/           # HTTP response utilities
│   ├── setting/            # Configuration utilities
│   └── utils/              # Common utilities
├── scripts/                # Build and deployment scripts
├── storages/               # File storage
├── tests/                  # Test files
├── third_party/            # Third-party dependencies
├── go.mod                  # Go module definition
├── go.sum                  # Go module checksums
├── Makefile                # Build automation
└── README.md               # Project documentation
```

## 📋 Prerequisites

Before running this project, ensure you have the following installed:

- **Go 1.24.5** or higher
- **Docker** and **Docker Compose**
- **MySQL 8.0** (or use Docker)
- **Redis 7.0** (or use Docker)
- **Make** (optional, for build automation)

## 🚀 Installation

### 1. Clone the Repository

```bash
git clone https://github.com/quanminhvu/go-ecommerce-backend-api.git
cd go-ecommerce-backend-api
```

### 2. Install Dependencies

```bash
go mod download
```

### 3. Set Up Environment

Copy the configuration file and modify as needed:

```bash
cp config/local.yml config/local.yml.example
```

## ⚙️ Configuration

The application uses Viper for configuration management. Configuration files are located in the `config/` directory.

### Configuration Structure

```yaml
server:
  port: 8080

mysql:
  username: quanvm
  password: 098poiA@
  host: localhost
  port: 3306
  dbname: golang-learn
  maxIdleConns: 10
  maxOpenConns: 100
  connMaxLifeTime: 3600

logger:
  log_level: debug
  file_log_name: "./storages/log/dev_001.log"
  max_size: 500
  max_backups: 3
  max_age: 28
  compress: true
```

### Environment Variables

You can override configuration using environment variables:

```bash
export MYSQL_HOST=localhost
export MYSQL_PORT=3306
export MYSQL_USERNAME=quanvm
export MYSQL_PASSWORD=098poiA@
export MYSQL_DBNAME=golang-learn
```

## 🏃‍♂️ Running the Application

### Option 1: Using Make (Recommended)

```bash
# Run the application
make run

# Or directly
go run ./cmd/server/
```

### Option 2: Using Docker Compose

```bash
# Start all services (MySQL, Redis, Application)
docker-compose up -d

# View logs
docker-compose logs -f

# Stop all services
docker-compose down
```

### Option 3: Manual Setup

1. **Start Database Services**

```bash
# Start MySQL and Redis
docker-compose up mysql redis -d
```

2. **Run Migrations**

```bash
# Apply database migrations
# (Migrations are automatically applied on startup)
```

3. **Start the Application**

```bash
go run ./cmd/server/
```

The application will be available at `http://localhost:8002`

## 📡 API Endpoints

### Base URL
```
http://localhost:8002/v1
```

### Available Endpoints

| Method | Endpoint | Description | Authentication |
|--------|----------|-------------|----------------|
| GET | `/ping/:name` | Health check endpoint | No |
| GET | `/user/1` | Get user information | Yes |

### Example Requests

```bash
# Health check
curl http://localhost:8002/v1/ping/test

# Get user (requires authentication)
curl -H "Authorization: Bearer YOUR_TOKEN" \
     http://localhost:8002/v1/user/1
```

## 🗄️ Database

### Database Schema

The application uses MySQL with the following main tables:

#### Users Table
```sql
CREATE TABLE users (
    id BIGINT PRIMARY KEY AUTO_INCREMENT,
    uuid VARCHAR(255) NOT NULL UNIQUE,
    name VARCHAR(255) NOT NULL,
    email VARCHAR(255) NOT NULL UNIQUE,
    password VARCHAR(255) NOT NULL,
    avatar VARCHAR(255),
    is_active BOOLEAN DEFAULT TRUE,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP NULL
);
```

#### Roles Table
```sql
CREATE TABLE roles (
    id BIGINT PRIMARY KEY AUTO_INCREMENT,
    name VARCHAR(255) NOT NULL UNIQUE,
    description TEXT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP NULL
);
```

#### User Roles (Many-to-Many)
```sql
CREATE TABLE user_roles (
    user_id BIGINT,
    role_id BIGINT,
    PRIMARY KEY (user_id, role_id),
    FOREIGN KEY (user_id) REFERENCES users(id),
    FOREIGN KEY (role_id) REFERENCES roles(id)
);
```

### Database Connection

The application uses GORM with the following connection settings:

- **Max Idle Connections**: 10
- **Max Open Connections**: 100
- **Connection Max Lifetime**: 3600 seconds

## 🧪 Development

### Code Structure

The project follows Go best practices and clean architecture:

- **Controllers**: Handle HTTP requests/responses
- **Services**: Business logic implementation
- **Repository**: Data access abstraction
- **PO (Persistence Objects)**: Database models
- **Middlewares**: Cross-cutting concerns

### Adding New Features

1. **Create PO (Persistence Object)**
```go
// internal/po/example.po.go
type Example struct {
    gorm.Model
    Name string `gorm:"column:name; type:varchar(255);not null"`
}
```

2. **Create Repository**
```go
// internal/repo/example.repo.go
type ExampleRepository struct {
    db *gorm.DB
}
```

3. **Create Service**
```go
// internal/service/example.service.go
type ExampleService struct {
    repo *repo.ExampleRepository
}
```

4. **Create Controller**
```go
// internal/controller/example.controller.go
type ExampleController struct {
    service *service.ExampleService
}
```

5. **Add Routes**
```go
// internal/routers/router.go
v1.GET("/example", c.NewExampleController().GetExample)
```

### Code Style

- Follow Go formatting standards (`gofmt`)
- Use meaningful variable and function names
- Add comments for complex logic
- Keep functions small and focused
- Use proper error handling

## 🧪 Testing

### Running Tests

```bash
# Run all tests
go test ./...

# Run tests with coverage
go test -cover ./...

# Run specific test
go test ./internal/controller

# Generate coverage report
go test -coverprofile=coverage.out ./...
go tool cover -html=coverage.out -o coverage.html
```

### Test Structure

Tests are organized alongside the code they test:

```
internal/
├── controller/
│   ├── user.controller.go
│   └── user.controller_test.go
├── service/
│   ├── user.service.go
│   └── user.service_test.go
└── repo/
    ├── user.repo.go
    └── user.repo_test.go
```

## 🐳 Docker

### Docker Services

The `docker-compose.yml` file includes:

- **MySQL 8.0**: Primary database
- **Redis 7.0**: Caching layer
- **Application**: Go API server

### Docker Commands

```bash
# Build and start all services
docker-compose up -d

# View logs
docker-compose logs -f

# Stop all services
docker-compose down

# Rebuild and start
docker-compose up --build -d

# Access MySQL
docker-compose exec mysql mysql -u quanvm -p

# Access Redis
docker-compose exec redis redis-cli
```

### Docker Environment Variables

Create a `.env` file for Docker environment variables:

```env
MYSQL_ROOT_PASSWORD=098poiA@
MYSQL_DATABASE=golang-learn
MYSQL_USER=quanvm
MYSQL_PASSWORD=098poiA@
```

## 🔧 Development Tools

### Make Commands

```bash
# Run the application
make run

# Build the application
make build

# Run tests
make test

# Clean build artifacts
make clean
```

### Logging

The application uses structured logging with Zap:

- **Log Level**: Configurable (debug, info, warn, error)
- **Log Rotation**: Automatic log rotation with compression
- **Log Location**: `./storages/log/dev_001.log`

### Monitoring

- **Health Checks**: Built-in health check endpoints
- **Metrics**: Application metrics collection
- **Performance**: Request/response timing

## 🤝 Contributing

1. Fork the repository
2. Create a feature branch (`git checkout -b feature/amazing-feature`)
3. Commit your changes (`git commit -m 'Add some amazing feature'`)
4. Push to the branch (`git push origin feature/amazing-feature`)
5. Open a Pull Request

### Development Guidelines

- Follow Go coding standards
- Add tests for new features
- Update documentation
- Use conventional commit messages
- Ensure all tests pass before submitting

## 📝 License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## 🙏 Acknowledgments

- [Gin Web Framework](https://gin-gonic.com/)
- [GORM](https://gorm.io/)
- [Viper](https://github.com/spf13/viper)
- [Zap Logger](https://github.com/uber-go/zap)

## 📞 Support

For support and questions:

- Create an issue in the GitHub repository
- Contact the development team
- Check the documentation

---

**Happy Coding! 🚀**