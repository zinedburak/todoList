package tests

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/stretchr/testify/assert"
	"math/rand"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"
	"time"
	"todolist_backend/database"
	"todolist_backend/routes"
)

func TestSetup(t *testing.T) {
	tests := []struct {
		description string
		method 		string
		route       string
		expectedCode int
	}{
		// First test
		{
			description:  "get todos test",
			method: 	  "GET",
			route:        "/api/get_todos",
			expectedCode: 200,
		},
		// Second Test
		{
			description:  "add todos test",
			method:       "POST",
			route:        "/api/add_todos",
			expectedCode: 200,
		},
		// Third Test
		{
			description:  "non existing page test",
			method:       "GET",
			route:        "/not_found",
			expectedCode: 404,
		},
	}
	database.Connect()
	app := fiber.New()
	app.Use(cors.New(cors.Config{
		AllowCredentials: true,
	}))
	routes.Setup(app)

	for _, test := range tests {
		//fmt.Println(test)
		var req *http.Request
		rand.Seed(time.Now().UnixNano())
		randInt := rand.Intn(10000)
		text := strconv.Itoa(randInt)

		if test.method == "POST"{
			postBody := map[string]interface{}{
				"text": text,
				"completed": "false"}
			body, _ := json.Marshal(postBody)
			req = httptest.NewRequest(test.method, test.route,  bytes.NewReader(body))

			req.Header.Set("Content-Type", "application/json")
		}else {
			req = httptest.NewRequest(test.method, test.route,  nil)
		}
		//fmt.Println(req)
	////
		resp, err := app.Test(req,100)
		//fmt.Println(err.Error())
		//fmt.Println(resp.Status)
	//
		if err != nil{
			fmt.Println(err.Error())
		}

		assert.Equalf(t, test.expectedCode, resp.StatusCode, test.description)
		// truncating the todos in order to not save test data
		database.DB.Exec("Truncate todos")
	}

}
