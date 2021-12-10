# GoLogin

This repository is for learning the Go programming language.

Let's start with implementing a web server with connection to a MySQL database including login functionality.

## Preparation

1. create .env File in project root directory
2. Start MySQL Database
3. insert MySQL Credentials into .env

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
Install go -> https://go.dev/doc/install

$ go install
$ go run .
```
