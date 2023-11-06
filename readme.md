# Crud Golang

## Requirement Install
* [GoLang](https://golang.org/)
* [Docker](https://www.docker.com/)
* [Postman](https://www.postman.com/) or [Insomnia](https://insomnia.rest/download) to test API

## Project Description

* What does the application do?
    - Management of a User CRUD

* Why did I use these technologies?
    - I used Golang because I had no knowledge of the technology and wanted to learn it. I used the Gin framework and GORM as an orm. I saw that Gin was a framework very similar to Django (Python, which I am familiar with). I saw that there are others like Fiber, Iris, Beego among others.

* Some of the challenges I faced and features I wanted to implement
    - Search for the most used language and frameworks and understand each one and create the application. I added Swagger but I didn't configure all possible responses. I read about unit tests and saw some examples. I saw that Go itself has a testing framework, Go Testing Package, I read about Testify too but I didn't feel comfortable creating them.


## Installation

Clone project
```sh
$ git clone https://github.com/vlfreitas/golang_crud.git
$ cd golang_crud
```
Running Docker to up database
```sh
$ docker-compose up -d
```
Running the application
```sh
$ go run main.go
```

See the API document and execute the calls. If you prefer, you can make calls via Postman or Insomnia.
```
Access the address in the browser: http://localhost:8080/swagger/index.html
```