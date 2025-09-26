# Job Tracker API Documentation

This project demonstrates a simple job applicant tracking system built with modern technologies.

## Technologies Used

1. **Golang** - Backend programming language
2. **Fiber Framework** - Fast HTTP web framework for Go
3. **KrakenD** - API Gateway for rate limiting and routing
4. **Docker** - Containerization with multi-stage builds
5. **Redis** - Caching layer for improved performance
6. **Apidog** - API documentation and testing
7. **PostgreSQL** - Primary database with optimized queries

## API Endpoints

### Health Check
- `GET /health` - Check service status

### Applicants
- `GET /applicants` - Get paginated list of applicants
- `POST /applicants` - Create new applicant
- `GET /applicants/{id}` - Get specific applicant
- `PUT /applicants/{id}` - Update applicant
- `DELETE /applicants/{id}` - Delete applicant

## Features

- **Input Validation**: Email format, phone number validation
- **Caching**: Redis-based caching with pagination support
- **Rate Limiting**: KrakenD gateway with per-endpoint limits
- **Database Optimization**: PostgreSQL with indexes and connection pooling
- **Security**: Basic authentication middleware
- **Logging**: Request logging and error tracking
- **Containerization**: Multi-stage Docker builds with security best practices

## Running the Application

```bash
# Using Docker Compose
docker-compose up -d

# Direct access to API
curl http://localhost:3000/health

# Through KrakenD Gateway
curl http://localhost:8081/api/health
```

## API Testing

Import the `job-tracker-api.json` file into Apidog to test all endpoints with sample data.
