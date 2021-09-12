package tests

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"math/rand"
	"net/http/httptest"
	"strconv"
	"testing"
	"time"
	"todolist_backend/database"
	"todolist_backend/models"
	"todolist_backend/routes"
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
	req := httptest.NewRequest("GET","/api/get_todos",nil)
	resp, _ := app.Test(req,100)

	// Parsing the response to a raw string
	bodyBytes, _ := ioutil.ReadAll(resp.Body)
	bodyString := string(bodyBytes)

	// Getting the result from the database
	dbResultString := GetResultFromDB()

	// Comparing two text
	assert.Equal(t,bodyString,dbResultString,"They are not equal" )
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
		"text": text,
		"completed": "false"}
	body, _ := json.Marshal(postBody)
	req := httptest.NewRequest("POST", "/api/add_todos",  bytes.NewReader(body))
	req.Header.Set("Content-Type","application/json")
	_, err := app.Test(req,100)
	dbResult_After_Add := GetResultFromDB()

	if err != nil {
		fmt.Println(err.Error())
	}
	// comparing the table status before and after adding new todo_object and if they are not equal we pass the test
	check := dbResult_Before_Add==dbResult_After_Add

	assert.False(t, check, "They are not equal " )
	// truncating the todos in order to not save test data
	database.DB.Exec("Truncate todos")
}

// Integration testing between add and get Todos

func TestIntegration(t *testing.T)  {
	app := SetupApp()
	dbResultBeforeAdd := GetResultFromDB()
	// creating a random int to add as a text in our database
	rand.Seed(time.Now().UnixNano())
	randInt := rand.Intn(10000)
	text := strconv.Itoa(randInt)

	// adding new Todo_object with random text with add_todos
	postBody := map[string]interface{}{
		"text": text,
		"completed": "false"}
	body, _ := json.Marshal(postBody)
	req := httptest.NewRequest("POST", "/api/add_todos",  bytes.NewReader(body))
	req.Header.Set("Content-Type","application/json")
	_, err := app.Test(req,100)
	// Checking if there is any error
	if err != nil {
		fmt.Println(err.Error())
	}

	// Getting the latest status of the database with get_todos function
	reqG := httptest.NewRequest("GET","/api/get_todos",nil)
	resp, _ := app.Test(reqG,100)

	// Parsing the response to a raw string
	bodyBytes, _ := ioutil.ReadAll(resp.Body)
	bodyString := string(bodyBytes)

	// Making sure that added todo_object can be seen with get_todo

	db_result_normal := GetResultFromDB()
	check := dbResultBeforeAdd == bodyString

	assert.Equal(t, bodyString,db_result_normal, "add function and get function are not working together")
	assert.False(t, check)
	// truncating the todos in order to not save test data
	database.DB.Exec("Truncate todos")
}
