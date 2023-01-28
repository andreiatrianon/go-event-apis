Server Instructions
=================

## Run server
Go to the server directory and run:
```terminal
$ cd server
$ go run .
```

### Make Requests

#### Via curl
Example:
```terminal
curl -H "Content-Type: application/json" http://localhost:8080/ \
    --request "POST" \
    --data '{"eventType": "copyAndPaste",
  "websiteUrl": "https://ravelin.com",
  "sessionId": "123123-123123-123123123",
  "pasted": true,
  "formId": "inputCardNumber"}'
```
#### Via frontend
Open the HTML file `client/index.html` in a web browser and send events by:
- Resizing the screen using the browser inspector.
- Pasting a value inside each field of the form.
- Filling the fields and submit the form.

### Run tests
Inside the server directory, you can run all the tests:
```terminal
$ go test ./...
```

Or you can run the tests from each package:

- `package main`
```terminal
$ go test
```

- `package apis`
```terminal
$ go test ./apis/
```
