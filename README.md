# Job Applicant Tracking System

A modern, scalable job applicant management system built with Go, demonstrating enterprise-level architecture patterns and best practices.

## 🏗️ Architecture Overview

This system implements a microservices architecture with API Gateway pattern, featuring:

- **Backend API**: Go with Fiber framework for high-performance HTTP handling
- **API Gateway**: KrakenD for rate limiting, caching, and request routing
- **Database**: PostgreSQL with GORM for data persistence and relationships
- **Caching**: Redis for performance optimization and session management
- **Containerization**: Docker with multi-stage builds for production deployment
- **Documentation**: Apidog integration for API testing and documentation

## 📁 Project Structure

```
job_tracker/
├── 📁 apidog/                          # API Documentation & Testing
│   ├── 📄 job-tracker-api.json         # OpenAPI specification
│   └── 📄 README.md                    # API documentation guide
├── 📁 controllers/                     # HTTP Request Handlers
│   └── 📄 applicantController.go       # Applicant CRUD operations
├── 📁 database/                        # Database Configuration
│   └── 📄 db.go                        # PostgreSQL connection & setup
├── 📁 krakend/                         # API Gateway Configuration
│   └── 📄 krakend.json                 # KrakenD routing & rate limiting
├── 📁 middleware/                      # Custom Middleware
│   ├── 📄 auth.go                      # Authentication middleware
│   └── 📄 request_logger.go            # Request logging middleware
├── 📁 models/                          # Data Models
│   └── 📄 applicant.go                 # Applicant struct definition
├── 📁 routes/                          # Route Configuration
│   └── 📄 applicant.go                 # API route definitions
├── 📁 utils/                           # Utility Functions
│   └── 📄 validation.go                # Input validation helpers
├── 📄 .env.example                     # Environment variables template
├── 📄 docker-compose.yml               # Multi-service orchestration
├── 📄 Dockerfile                       # Container build configuration
├── 📄 go.mod                           # Go module dependencies
├── 📄 go.sum                           # Go module checksums
├── 📄 main.go                          # Application entry point
└── 📄 README.md                        # Project documentation
```

### 📋 Directory Breakdown

| Directory | Purpose | Key Files |
|-----------|---------|-----------|
| **`/apidog/`** | API Documentation | OpenAPI spec, testing collection |
| **`/controllers/`** | Business Logic | HTTP handlers, validation, caching |
| **`/database/`** | Data Layer | Connection setup, migrations, indexes |
| **`/krakend/`** | API Gateway | Rate limiting, routing, caching config |
| **`/middleware/`** | Cross-cutting Concerns | Auth, logging, error handling |
| **`/models/`** | Data Models | Struct definitions, GORM tags |
| **`/routes/`** | URL Routing | Endpoint definitions, middleware setup |
| **`/utils/`** | Helper Functions | Validation, sanitization, utilities |

### 🔧 Key Configuration Files

- **`docker-compose.yml`**: Orchestrates all services (app, db, redis, krakend)
- **`Dockerfile`**: Multi-stage build for optimized Go application
- **`krakend.json`**: API Gateway configuration with rate limiting
- **`go.mod`**: Go module dependencies and versions
- **`.env.example`**: Environment variable templates

## 🚀 Technologies Used

### 1. **Golang** - Backend Programming Language
- **Usage**: Core application logic, HTTP server, business logic implementation
- **Features Demonstrated**:
  - Struct-based data modeling with GORM
  - Context-based request handling
  - Error handling with custom error types
  - Environment-based configuration
  - Connection pooling and database optimization

### 2. **Fiber Framework** - High-Performance Web Framework
- **Usage**: HTTP server, middleware, routing, and request/response handling
- **Features Demonstrated**:
  - RESTful API endpoints with proper HTTP methods
  - Custom middleware for logging and authentication
  - CORS configuration for cross-origin requests
  - Panic recovery and error handling
  - Request validation and sanitization

### 3. **KrakenD** - API Gateway
- **Usage**: Request routing, rate limiting, caching, and API management
- **Features Demonstrated**:
  - Endpoint routing with different rate limits per operation
  - Response caching with TTL configuration
  - CORS handling at gateway level
  - Request/response transformation
  - Health check endpoint routing

### 4. **Docker** - Containerization
- **Usage**: Application packaging, deployment, and orchestration
- **Features Demonstrated**:
  - Multi-stage builds for optimized image size
  - Docker Compose for multi-service orchestration
  - Health checks and restart policies
  - Environment variable configuration
  - Volume management for data persistence

### 5. **Redis** - In-Memory Caching
- **Usage**: Application-level caching, session storage, and performance optimization
- **Features Demonstrated**:
  - Pagination-aware caching strategies
  - Cache invalidation on data mutations
  - Connection pooling and error handling
  - Fallback mechanisms when Redis is unavailable
  - TTL-based cache expiration

### 6. **Apidog** - API Documentation & Testing
- **Usage**: API documentation, testing, and client integration
- **Features Demonstrated**:
  - OpenAPI specification with detailed schemas
  - Example requests and responses
  - Endpoint documentation with parameters
  - Error response documentation
  - Collection for automated testing

### 7. **PostgreSQL Database** - Primary Data Storage
- **Usage**: Data persistence, relationships, and ACID compliance
- **Features Demonstrated**:
  - GORM ORM with auto-migration
  - Database indexes for performance optimization
  - Connection pooling and timeout configuration
  - Soft deletes with GORM
  - Transaction handling and data integrity

## 📋 Business Logic

### Core Functionality
The system manages job applicants through a complete CRUD lifecycle:

1. **Applicant Registration**: Capture candidate information with validation
2. **Application Tracking**: Monitor application status through various stages
3. **Interview Management**: Update candidate progress and add notes
4. **Decision Making**: Track hiring decisions and maintain audit trail
5. **Data Analytics**: Pagination and filtering for large datasets

### Data Model
```go
type Applicant struct {
    ID        uint           `json:"id" gorm:"primarykey"`
    CreatedAt time.Time      `json:"created_at"`
    UpdatedAt time.Time      `json:"updated_at"`
    DeletedAt gorm.DeletedAt `json:"deleted_at,omitempty" gorm:"index"`
    
    Name     string `json:"name" gorm:"not null;size:100"`
    Email    string `json:"email" gorm:"unique;not null;size:150"`
    Position string `json:"position" gorm:"not null;size:100"`
    Status   string `json:"status" gorm:"default:'pending';size:20"`
    Phone    string `json:"phone,omitempty" gorm:"size:20"`
    Resume   string `json:"resume,omitempty" gorm:"type:text"`
    Notes    string `json:"notes,omitempty" gorm:"type:text"`
}
```

### Application Status Flow
```
pending → reviewed → interviewed → hired/rejected
```

## 🛠️ Installation & Setup

### Prerequisites
- Docker and Docker Compose
- Go 1.21+ (for local development)
- Git

### Quick Start
```bash
# Clone the repository
git clone <repository-url>
cd job_tracker

# Start all services
docker compose up --build -d

# Verify services are running
docker compose ps
```

### Service Endpoints
- **Direct API**: http://localhost:3000
- **KrakenD Gateway**: http://localhost:8081
- **PostgreSQL**: localhost:5432
- **Redis**: localhost:6379

## 📚 API Documentation

### Health Check
```bash
# Direct API
curl http://localhost:3000/health

# Through KrakenD Gateway
curl http://localhost:8081/api/health
```

### Applicant Management

#### Create New Applicant
```bash
curl -X POST http://localhost:8081/api/applicants \
  -H "Content-Type: application/json" \
  -d '{
    "name": "John Doe",
    "email": "john@example.com",
    "position": "Software Engineer",
    "phone": "+1234567890",
    "notes": "Experienced developer"
  }'
```

#### Get All Applicants (with pagination)
```bash
curl "http://localhost:8081/api/applicants?page=1&limit=10"
```

#### Get Specific Applicant
```bash
curl http://localhost:8081/api/applicants/1
```

#### Update Applicant
```bash
curl -X PUT http://localhost:8081/api/applicants/1 \
  -H "Content-Type: application/json" \
  -d '{
    "status": "interviewed",
    "notes": "Passed technical interview"
  }'
```

#### Delete Applicant
```bash
curl -X DELETE http://localhost:8081/api/applicants/1
```

## 🔧 Configuration

### Environment Variables
```bash
# Database Configuration
DB_HOST=localhost
DB_USER=postgres
DB_PASSWORD=password
DB_NAME=postgres
DB_PORT=5432

# Redis Configuration
REDIS_HOST=localhost
REDIS_PORT=6379
REDIS_PASSWORD=

# Application Configuration
PORT=3000
ENVIRONMENT=development
```

### KrakenD Configuration
The API Gateway is configured with:
- **Rate Limiting**: 100 requests/minute globally, 50 for GET, 20 for POST
- **Caching**: 300-second TTL for GET requests
- **CORS**: Configured for cross-origin requests
- **Timeout**: 3-5 second timeouts per endpoint

## 🧪 Testing

### Manual Testing
Use the provided Apidog collection (`apidog/job-tracker-api.json`) to test all endpoints with sample data.

### Automated Testing
```bash
# Run tests (if implemented)
go test ./...

# Test specific functionality
go test ./controllers
go test ./utils
```

## 📊 Performance Features

### Caching Strategy
- **Redis Caching**: Paginated results cached for 3 minutes
- **Cache Invalidation**: Automatic cache clearing on data mutations
- **Fallback**: Direct database access when Redis is unavailable

### Database Optimization
- **Indexes**: Created on email, status, and created_at fields
- **Connection Pooling**: 10 idle, 100 max connections
- **Query Optimization**: GORM with prepared statements

### API Gateway Features
- **Rate Limiting**: Per-endpoint rate limits
- **Response Caching**: KrakenD-level caching
- **Request Routing**: Load balancing and failover

## 🚀 Deployment

### Production Deployment
```bash
# Build production image
docker build -t job-tracker:latest .

# Deploy with production environment
docker compose -f docker-compose.prod.yml up -d
```

### Scaling
- **Horizontal Scaling**: Multiple app instances behind KrakenD
- **Database Scaling**: Read replicas for read-heavy operations
- **Cache Scaling**: Redis cluster for high availability

## 🔍 Monitoring & Logging

### Application Logs
- **Request Logging**: All HTTP requests with timing
- **Database Logging**: SQL queries with execution times
- **Error Logging**: Detailed error information with context
- **Cache Logging**: Cache hits/misses and performance metrics

### Health Monitoring
- **Health Checks**: Built-in health endpoints
- **Docker Health**: Container health monitoring
- **Database Health**: Connection status monitoring

## 🤝 Contributing

1. Fork the repository
2. Create a feature branch
3. Make your changes
4. Add tests if applicable
5. Submit a pull request

## 📄 License

This project is licensed under the MIT License - see the LICENSE file for details.

## 🎯 Future Enhancements

- [ ] Authentication and authorization
- [ ] File upload for resumes
- [ ] Email notifications
- [ ] Advanced search and filtering
- [ ] Analytics dashboard
- [ ] API versioning
- [ ] GraphQL endpoint
- [ ] WebSocket support for real-time updates

---

**Built with ❤️ using modern Go practices and enterprise architecture patterns.**
