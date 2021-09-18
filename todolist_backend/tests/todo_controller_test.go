package tests

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/rand"
	"net/http/httptest"
	"strconv"
	"testing"
	"time"
	"todolist_backend/database"
	"todolist_backend/models"
	"todolist_backend/routes"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/stretchr/testify/assert"
)

// Setting up the application to prevent code duplication
func SetupApp() *fiber.App {
	database.Connect()
	app := fiber.New()
	app.Use(cors.New(cors.Config{
		AllowCredentials: true,
	}))
	routes.Setup(app)
	return app
}

// Getting the result from DB and parsing it to a string // added to prevent code duplication
func GetResultFromDB() string {
	var todos []models.Todos
	database.DB.Find(&todos)

	// Parsing the result from database to a raw string
	dbResultJson, _ := json.Marshal(todos)
	dbResultString := string(dbResultJson)
	return dbResultString
}

// Unit testing the Get todos function
func TestGetTodos(t *testing.T) {

	app := SetupApp()

	// creating a http request to the get_todos
	req := httptest.NewRequest("GET", "/api/get_todos", nil)
	resp, _ := app.Test(req, 100)

	// Parsing the response to a raw string
	bodyBytes, _ := ioutil.ReadAll(resp.Body)
	bodyString := string(bodyBytes)

	// Getting the result from the database
	dbResultString := GetResultFromDB()

	// Comparing two text
	assert.Equal(t, bodyString, dbResultString, "They are not equal")
	// truncating the todos in order to not save test data
	database.DB.Exec("Truncate todos")

}

//Unit testing the Add todos function
func TestAddTodos(t *testing.T) {
	app := SetupApp()
	// creating a random int to add as a text in our database
	rand.Seed(time.Now().UnixNano())
	randInt := rand.Intn(10000)
	text := strconv.Itoa(randInt)

	dbResult_Before_Add := GetResultFromDB()

	postBody := map[string]interface{}{
		"text": text}
	body, _ := json.Marshal(postBody)
	req := httptest.NewRequest("POST", "/api/add_todos", bytes.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	_, err := app.Test(req, 100)
	dbResult_After_Add := GetResultFromDB()

	if err != nil {
		fmt.Println(err.Error())
	}
	// comparing the table status before and after adding new todo_object and if they are not equal we pass the test
	check := dbResult_Before_Add == dbResult_After_Add

	assert.False(t, check, "They are not equal ")
	// truncating the todos in order to not save test data
	database.DB.Exec("Truncate todos")
}

// Integration testing for add todos

func TestIntegrationAddTodos(t *testing.T) {
	app := SetupApp()
	// Getting todos from database before adding any todo_object
	dbResultBeforeAdd := GetResultFromDB()

	// Creating a random int to add as a text in our database
	rand.Seed(time.Now().UnixNano())
	randInt := rand.Intn(10000)
	text := strconv.Itoa(randInt)

	// adding new Todo_object with random text with add_todos
	postBody := map[string]interface{}{
		"text": text}
	body, _ := json.Marshal(postBody)
	req := httptest.NewRequest("POST", "/api/add_todos", bytes.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	_, err := app.Test(req, 100)
	// Checking if there is any error
	if err != nil {
		fmt.Println(err.Error())
	}
	// Getting todos after adding a todo_ with random text
	dbResultAfterAdd := GetResultFromDB()

	// Comparing the todos table before adding a new item and after
	check := dbResultBeforeAdd == dbResultAfterAdd
	assert.False(t, check)

	// truncating the todos in order to not save test data
	database.DB.Exec("Truncate todos")
}

// Integration testing for get todos
func TestIntegrationGetTodos(t *testing.T) {
	// Testing for empty todos table
	app := SetupApp()

	// Getting todos table with our GetResultFromDB function
	dbEmptyTodos := GetResultFromDB()

	// Getting todos table with our GetTodos API
	reqGetTodosEmpty := httptest.NewRequest("GET", "/api/get_todos", nil)
	respGetTodosEmpty, _ := app.Test(reqGetTodosEmpty, 100)

	// Parsing the response to a raw string
	bodyBytesEmptyTable, _ := ioutil.ReadAll(respGetTodosEmpty.Body)
	bodyStringEmptyTable := string(bodyBytesEmptyTable)

	// Testing for todos table with todos in it

	//Adding a new Todo_object to database
	rand.Seed(time.Now().UnixNano())
	randInt := rand.Intn(10000)
	text := strconv.Itoa(randInt)

	postBody := map[string]interface{}{
		"text": text}
	body, _ := json.Marshal(postBody)
	reqAddTodos := httptest.NewRequest("POST", "/api/add_todos", bytes.NewReader(body))
	reqAddTodos.Header.Set("Content-Type", "application/json")
	_, err := app.Test(reqAddTodos, 100)

	if err != nil {
		fmt.Println(err.Error())
	}

	// Getting todos table with our GetResultFrom DB function after adding a random todo_object with random text
	dbFilledTodos := GetResultFromDB()

	// Getting todos table wit our GetTodos API  after adding a todo_object
	reqGetTodosFilled := httptest.NewRequest("GET", "/api/get_todos", nil)
	respGetTodosFilled, _ := app.Test(reqGetTodosFilled, 100)

	// Parsing the response to a raw string
	bodyBytesFilledTable, _ := ioutil.ReadAll(respGetTodosFilled.Body)
	bodyStringFilledTable := string(bodyBytesFilledTable)

	// Comparing the results of db status with the api result for both empty and filled database
	assert.Equal(t, dbFilledTodos, bodyStringFilledTable, "Filled integration test fails")
	assert.Equal(t, bodyStringEmptyTable, dbEmptyTodos, "Empty Integration Test Fails")
	database.DB.Exec("Truncate todos")

}
