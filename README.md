# GoLogin

This repository is for learning the Go programming language.

Let's start with implementing a web server with connection to a MySQL database including login functionality.

## Preparation

1. Install go -> [https://go.dev/doc/install](https://go.dev/doc/install/)
2. create .env File in project root directory
3. Start MySQL Database
4. insert MySQL Credentials into .env

### .env Example

```
DB_USERNAME=username
DB_PASSWORD=password
DB_HOST=localhost
DB_PORT=3306
DB_NAME=gologin
DB_OPTIONS="?charset=utf8mb4&parseTime=True&loc=Local"
```

## Installation

```
// go into project root directory
$ go install
$ go run .
```
