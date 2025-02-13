Social API Repository
=====================

Overview:
---------
This project is a sample Go API built with the Gin framework. It demonstrates a layered architecture where requests pass through multiple layers:
  • Router: Handles HTTP endpoints.
  • Provider: Processes and prepares data.
  • Depx: Fetches external data (from https://jsonplaceholder.typicode.com/posts).

The API exposes two endpoints:
  - /ready : A simple health-check endpoint.
  - /posts : Retrieves posts from an external API via the provider and depx layers.

Project Structure:
------------------
```
social/
├── go.mod           -- Go module file.
├── main.go          -- Application entry point.
├── routes/
│   ├── ready.go     -- Defines the /ready endpoint.
│   └── posts.go     -- Defines the /posts endpoint and invokes the provider.
├── providers/
│   └── posts.go     -- Contains business logic and data conversion for posts.
└── depx/
    └── posts.go     -- Handles the external API call to fetch posts data.
```

How It Works:
-------------
1. A client sends a request to the API.
2. The router (in routes/) processes the request:
   - For /ready, a simple response is returned.
   - For /posts, the router calls the provider.
3. The provider (in provider/posts.go) calls the depx layer.
4. The depx layer (in depx/posts.go) makes an HTTP GET request to:
   https://jsonplaceholder.typicode.com/posts
5. The depx layer returns the raw JSON data along with the HTTP status code.
6. The provider converts the JSON into a slice of Post structs, optionally processes it,
   and then marshals it back into JSON.
7. The router returns the final JSON data to the client.

Usage:
------
To run the application, ensure you have Go installed (v1.18+), then execute the following command
from the project root:

```
go run main.go
```

The server will start on port 8080. You can then access:
  • Health Check: http://localhost:8080/ready
  • Posts Data:  http://localhost:8080/posts

Dependencies:
-------------
- Go (v1.18 or later)
- Gin framework: github.com/gin-gonic/gin

Configuration & Improvements:
-----------------------------
- A custom HTTP client with a 10-second timeout is used in the depx layer to handle external API calls.
- Context propagation is implemented to support cancellation and timeouts.
- The layered architecture (router → provider → depx) ensures separation of concerns and makes the codebase more modular.
- Centralized error handling and logging can be further integrated for production readiness.

License:
--------
This project is licensed under the MIT License. See the LICENSE file for more details.