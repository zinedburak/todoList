# TodoList Application

[![Frontend CI](https://github.com/zinedburak/todoList/actions/workflows/node.js.yml/badge.svg)](https://github.com/zinedburak/todoList/actions/workflows/node.js.yml)
[![Backend CI](https://github.com/zinedburak/todoList/actions/workflows/go.yml/badge.svg)](https://github.com/zinedburak/todoList/actions/workflows/go.yml)

This is a simple TodoList application that is created by React and Go
# How To Run The Application
- Make sure that you have docker installed
- Run the "docker-compose up --build" command inside the main folder that you clone the project
- The application will be built and run with the above command
- If you want to run the tests in the docker please go to the docker-compose file and uncomment the tests_backend and tests_frontend parts of the code and run the "docker-compose up --build" command again
- The Acceptance test fails when it runs in docker image since it could not install chromium and mimic user. But it passes without any problems when the application built locally.
- Go to localhost:3000 and enjoy your todo application

# Acceptance Test
- Create a user interface for only adding ToDo's -- Completed
- Create a backend service to store the persistent state of ToDo List -- Completed
- Writing deployment files for the front-end and back-end -- Completed
- Automating build and test of the application (Completed)
- Automating deployment of the application via CI/CD (Not completed)

# Development Cycle Of The Application
## Frontend
The Frontend side of the application has been developed with React.js language and the
the first step of the development was to create a react application with "npx create-react-app" command.

After the initial application was created I started my TDD:

- First step was to write a test for my imaginary form component. In this test, I checked if the form component has the add button and text holder for writing the todos. Also, I have written a test for checking if the text inside of the text holder changes when an event is fired.
- After writing the tests for the Form I implemented the form component of the UI without any functionalities and refactored the tests
- Added tests for the button functionalities.
- Added button functionality and created a dummy business logic to pass the tests (Since there was no backend at this moment)
- Refactored the form test to make tests more compatible with future components
- I created tests for two more imaginary components which were Todolist and Todo. In these tests that I checked that both Todolist and todo components render correctly with given data
- I implemented the todo and todolist components. Basically, these todolist component creates Todo component for each todos in the application.
- Refactored the tests one last time for tidying up the code.
- After the backend was implemented changed the dummy business logic with the API calls to the get todo and add todo functions
- After the feedback I created an acceptance test using puppeteer that tries to add a new todo item mimicking a user

At this point, I had a fully functioning frontend application that was capable of adding todos and showing them from the memory

## Backend
The Backend side of the application has been developed with Go and for the database service I used the Postgres database

After completing the frontend part of the application I started my TDD for the backend services:

- The first test I wrote for the backend was to routes test. In this test, I aimed to check the status of the API routes by sending them requests and comparing the response's status with hardcoded status codes
- I created routes for the API and connect them to the empty functions.
- I created a unit test for the add todos function. In add todos unit test I checked if the function can add a todo object into the database without any problem
- I created a unit test for the get todos function. In this test, I checked if the function can successfully gets all the data from the database.
- I created an integration test in order to test the compatibility of the add and get todos
- After the feedback I changed the integration test from testing the compatibility of two functions with each other to compatibility of these functions with the database
- Implemented add and get todos functions to show and update the persistent state of todos
- Added database connection functionality to connect the Postgres DB
- Added Todo data structure to easily add and get todos with gorm

## Dockerization
- Created Dockerfile for the front end backend services
- Created docker-compose file to compose the frontend backend in the same container
- Added commands to run frontend and backend tests when the container first runs


## Build and Test Automation
- Created CI pipeline with github actions for frontend
    - In this pipeline the frontend part of the application checkedout to github servers
    - After the check out the fronted is built wiht npm run build
    - When the build is complete all the tests expect the Acceptance test runs and produces a result 
- Created CI pipeline with github actions for backend
    - In this pipeline the backend part of the application checkedout to github servers
    - In the github servers the backend is built with go build -v ./.. command
    - After build we ran tests in the tests directory with go test command
