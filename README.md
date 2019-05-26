# Simple Go service
This is a simple Go service that listens for all requests on port :8080 and prints request path in response body after 200 millis, e.g.:
<br />GET http://localhost:8080/test
<br />Response: `{ "path": "/test" }`
<br />It also exposes prometheus metrics on port :9090.
