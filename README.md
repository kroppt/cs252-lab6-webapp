# cs252-lab6-webapp
Implementation of a web application using HTML5 and cloud services

### Prerequisites

[golang](https://golang.org/project/)

### Running

Fetch the code
```
git clone https://github.com/kroppt/cs252-lab6-webapp
cd cs252-lab6-webapp
```

Run the code:

```
go run cmd/webapp/main.go
```
or
 
```bash
go build cmd/webapp/main.go
./main
```
 
To test the program, `curl -X POST -d '{"id": "test"}' http://localhost:8080/`, open localhost:8080/ to see "test".
