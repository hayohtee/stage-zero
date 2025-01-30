# stage-zero
The first task for HNG internship for backend track. 
A public API that returns the following information in JSON format:
- Email address
- The current datetime as an ISO 8601 formatted timestamp.
- The GitHub URL of the project's codebase.

## How to run
1. Ensure the GO SDK is installed
2. Clone the repository
   ```bash
   git clone https://github.com/hayohtee/stage-zero.git
   ```
3. Change into the directory
   ```bash
   cd stage-zero
   ```
4. Build the project
   ```bash
   go build -o api ./cmd/api
   ```
5. Run the program\
   The server should start on http://localhost:4000, send a GET request to the endpoint to retrieve the information
   ```bash
   ./api
   ```

## API Documentation
Endpoint URL
```bash
GET http://localhost:4000
```
Response
```json
{
  "email": "your-email@example.com",
  "current_datetime": "2025-01-30T09:30:00Z",
  "github_url": "<https://github.com/yourusername/your-repo>"
}
```

[Hire GO developer](https://hng.tech/hire/golang-developers)
