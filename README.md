# Clean Architecture API

A Go-based REST API built with Clean Architecture principles, featuring Gin framework, GORM, MySQL, and JWT authentication.

## 🏗️ Architecture

This project follows **Clean Architecture** principles with clear separation of concerns:

- **Domain Layer**: Core business entities, repositories interfaces, and domain services
- **Application Layer**: Use cases that orchestrate business logic
- **Interface Layer**: HTTP handlers, controllers, and middleware
- **Infrastructure Layer**: Database implementations, external services, and utilities

### Architecture Flow
```
HTTP Request → Interface Layer → Application Layer → Domain Layer
                                    ↓
Infrastructure Layer ← Domain Layer ← Application Layer
```

## 📁 Directory Structure

```
├── cmd/                    # Application entry points
├── internal/              # Private application code
│   ├── application/       # Application layer (use cases)
│   │   └── usecase/       # Business logic orchestration
│   ├── config/           # Configuration management
│   ├── domain/           # Domain layer (core business)
│   │   ├── entity/       # Business entities
│   │   ├── repository/   # Repository interfaces
│   │   ├── service/      # Domain services
│   │   ├── dto/          # Data Transfer Objects
│   │   └── error/        # Domain errors
│   ├── infrastructure/   # Infrastructure layer
│   │   ├── database/     # Database connection
│   │   ├── repository/   # Repository implementations
│   │   ├── auth/         # Authentication services
│   │   ├── logger/       # Logging implementation
│   │   └── model/        # Database models
│   ├── interfaces/       # Interface layer
│   │   └── http/         # HTTP handlers, controllers, middleware
│   └── util/             # Utility functions
├── docs/                 # API documentation (Swagger)
├── db/                   # Database migrations and data
├── main.go              # Application entry point
├── go.mod               # Go module dependencies
├── Dockerfile           # Production Docker image
├── docker-compose.yml   # Development environment
└── Makefile             # Build and development commands
```

## 🚀 Quick Start

### Prerequisites

- Go 1.24.2+
- Docker & Docker Compose
- MySQL 8.0+

### Option 1: Using Docker (Recommended)

1. **Clone the repository**
   ```bash
   git clone <repository-url>
   cd clean-architecture
   ```

2. **Start the development environment**
   ```bash
   docker-compose up -d
   ```

3. **Run database migrations**
   ```bash
   make migrate-dev
   ```

4. **Access the API**
   - API: http://localhost:8080
   - Swagger Docs: http://localhost:8080/swagger/index.html

### Option 2: Local Development

1. **Install dependencies**
   ```bash
   go mod download
   ```

2. **Setup environment variables**
   ```bash
   cp .env.example .env
   # Edit .env with your configuration
   ```

3. **Start MySQL database**
   ```bash
   docker-compose up mysql -d
   ```

4. **Run migrations**
   ```bash
   make migrate-dev
   ```

5. **Start development server**
   ```bash
   make dev
   ```

## 🛠️ Available Commands

```bash
# Build the application
make build

# Start development server with hot reload
make dev

# Run database migrations
make migrate-dev

# Generate API documentation
make docs
```

## 🔧 Configuration

The application uses environment variables for configuration:

| Variable | Default | Description |
|----------|---------|-------------|
| `PORT` | `8080` | Server port |
| `DB_HOST` | `127.0.0.1` | Database host |
| `DB_PORT` | `3306` | Database port |
| `DB_DATABASE` | `` | Database name |
| `DB_USERNAME` | `root` | Database username |
| `DB_PASSWORD` | `` | Database password |
| `JWT_SECRET` | `secret` | JWT signing secret |
| `LOG_LEVEL` | `info` | Logging level |

## 📚 API Documentation

Once the server is running, you can access:
- **Swagger UI**: http://localhost:8080/swagger/index.html
- **API Base Path**: `/api/v1`

## 🧪 Testing

```bash
# Run all tests
go test ./...

# Run tests with coverage
go test -cover ./...
```

## 🐳 Docker

### Production Build
```bash
docker build -t clean-architecture .
docker run -p 8080:8080 clean-architecture
```

### Development Build
```bash
docker-compose up
```

## 📝 Project Features

- ✅ Clean Architecture implementation
- ✅ RESTful API with Gin framework
- ✅ JWT Authentication
- ✅ MySQL database with GORM
- ✅ Structured logging with Zap
- ✅ API documentation with Swagger
- ✅ Docker support
- ✅ Hot reload development
- ✅ Graceful shutdown
- ✅ CORS middleware
- ✅ Environment-based configuration
