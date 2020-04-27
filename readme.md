# Grader Ranker Server

Simple ranker and grader API that use Go, MySQL
## Environment
 - Go 1.14
 - MySQL 8.0

## How To Run
 - make sure your computer has Go and Mysql installed
 - set environment variable in `.env` file <br>
 .env file example <a href="./code/.env-example">.env-example</a>
 - run it by typing <br> 
    ```go run cmd/api/main.go``` in the terminal or <br>
    ```make run``` if Makefile is installed on your computer

## API
### Grader 
grader server run in <br> POST http://localhost:8080/grader/<br>
json example <a href="./code/scripts/examsheets_dummy.json">click this</a>

### Grader
grader server run in <br> GET http://localhost:8080/ranker/{idExam}
| Parameter | Type | Description |
| :--- | :--- | :--- |
| `idExam` | `unsigned integer` | **Required**.  |
### <a href="https://github.com/sy-anwar/grader-ranker-server">Github Repository</a> 