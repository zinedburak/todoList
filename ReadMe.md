# TodoList Application



This is a simple TodoList application that is created by React and Go

# Acceptance Test
- Create a user interface for only adding ToDo's -- Completed
- Create a backend service to store persistent state of ToDo List -- Completed
- Writing deployment files for the front-end and back-end -- Completed
- Automating build, test and deployment of the application via CI/CD

# Integration Testing Roadmap
- Create a test that adds a new todo and check the comminication between functions (completed)
- Check the communication between API and frontend

# Unit Testing Roadmap
- Create a test that checks the bussiness logics one by one in API (completed)

# Current Status of the Application
- A user interface for adding different todos has been created
- Created six different tests that checks the functionalty and the correctness of the render (All six tests are green)
- Created a backend API for adding new todos to a database
- Created a backend API for getting the todos from the database
- Created a test that checks the routes of the API
- Created unit tests for the  backend functions
- Created an integration test that checks the communication between add_todo and get_todo
- Created docker files for frontend and backend
- Created docker-compose file to dockerize the database, backend and frontend
