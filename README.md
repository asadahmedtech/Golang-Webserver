# AA Webserver

Simple Webserver implemented as the part of onboarding process for summer interns 2020. This project implements following features.

  - Simple webserver for ToDo list management
  - Authentication : User registration, login and logout
  - JWT Authentication with Middlewares 
  - MongoDB as database
  - Redis as cache server
  - v1 : Static API
  - v2 : Dynamic API with Database
  - Dashboard urls for http server
  - Blue-Green Load Balancer between v1 and v2


Launch the app using

```sh
go run main.go
```

Make sure redis server and MongoDB server is available in the background

# API Usage 

API Server can be accessed through  [http://localhost:9000/]
 
 - V1 API Server is hosted at [http://localhost:8001/]
 - V2 API Server is hosted at [http://localhost:8002/]

Using Blue-Green load balancer, uses reverse proxy to direct the load to different versions

| Task | URL |
| ------ | ------ |
| Login | [/api/auth/login] |
| Register | [/api/auth/register] |
| Logout | [/api/auth/logout] | 
| Profile Page | [/api/profile] |
| Add new ToDo | [/api/todo/insert]  |
| Fetch All the ToDos | [/api/todo/fetch]  |

# DashBoard Usage

Dashboard Server can be accessed through  [http://localhost:8000/]

| Page | URL |
| ------ | ------ |
| Login | [/login]  |
| Register | [/register] |
| Logout | [/logout][PlGa] 
| Profile Page | [/profile] |
| Add new ToDo | [/todopost]  |
| Fetch All the ToDos | [/todoview] |