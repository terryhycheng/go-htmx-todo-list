package controllers_test

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/terryhycheng/go-todo-list/internal/controllers"
)

var id = uuid.New()

func TestAddTodoController(t *testing.T) {
	// Initialize echo framework for testing
	e := echo.New()

	// Prepare request
	req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader("title=TestTitle&description=TestDescription&priority=normal&id="+id.String()))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationForm)

	// Prepare response recorder
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	// Execute controller function
	err := controllers.AddTodoController(c)

	// Assertions
	require.NoError(t, err)
	require.Equal(t, http.StatusCreated, rec.Code)
	// Add more assertions as needed
}

func TestChangeTodoStatusController(t *testing.T) {
	// Mock the necessary dependencies if required for testing
	// Initialize echo framework for testing
	e := echo.New()

	// Prepare request
	req := httptest.NewRequest(http.MethodGet, "/"+id.String(), nil)

	// Prepare response recorder
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetParamNames("id")
	c.SetParamValues(id.String())

	// Execute controller function
	err := controllers.ChangeTodoStatusController(c)

	// Assertions
	require.NoError(t, err)
	require.Equal(t, http.StatusOK, rec.Code)
	// Add more assertions as needed
}

func TestDeleteTodoController(t *testing.T) {
	asserts := assert.New(t)

	// Mock the necessary dependencies if required for testing
	// Initialize echo framework for testing
	e := echo.New()

	// Prepare request
	req := httptest.NewRequest(http.MethodDelete, "/"+id.String(), nil)

	// Prepare response recorder
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetParamNames("id")
	c.SetParamValues(id.String())

	// Execute controller function
	err := controllers.DeleteTodoController(c)

	// Assertions
	asserts.NoError(err)
	asserts.Equal(http.StatusOK, rec.Code)
	// Add more assertions as needed
}
