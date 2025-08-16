# Go URL Shortener

A lightweight, in-memory URL shortener service written in Go.

## Features

- Generates secure, random 6-character short codes
- In-memory storage with mutex-protected concurrent access
- Simple REST API interface
- Automatic redirection from short URLs
- Minimal dependencies (only standard library)

## Quick Start

1. Clone the repository:
   ```bash
   git clone https://github.com/Sushanta175/URL_Shortener.git
   cd URL_Shortener
   ```

2. Run the service:
   ```bash
   go run main.go
   ```

3. Test the API:
   ```bash
   # Create short URL
   curl "http://localhost:8080/shorten?url=https://example.com/very/long/url"
   
   # Use short URL (replace {code} with actual code)
   curl -v "http://localhost:8080/{code}"
   ```

## API Documentation

### Create Short URL
```
GET /shorten?url=<your-long-url>
```

**Response:**
```
Shortened URL: http://localhost:8080/AbCdEf
```

### Redirect Endpoint
```
GET /{short-code}
```

**Responses:**
- `302 Found` with Location header (success)
- `404 Not Found` (invalid code)

## Configuration

| Aspect           | Default         | Customization                |
|-------------------|------------------|------------------------------|
| Port               | 8080           | Change in `main.go`          |
| Short code length | 6 characters     | Modify in `generateShortCode`|
| Base URL          | localhost:8080   | Update response formatting   |

## Development

### Dependencies
- Go 1.16+
- No external packages required

### Project Structure
```
.
├── main.go         # Server setup and routes
├── README.md         # This document
└── (future expansion)
```

### Testing
Basic manual testing:
```bash
# Test valid URL
curl "http://localhost:8080/shorten?url=https://google.com"

# Test missing URL parameter
curl "http://localhost:8080/shorten"

# Test redirect
curl -v "http://localhost:8080/abc123"
```

## License

MIT License - See [LICENSE](LICENSE) file for details.
