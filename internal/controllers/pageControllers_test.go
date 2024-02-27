package controllers

import (
	"github.com/stretchr/testify/suite"
	"github.com/terryhycheng/go-todo-list/internal/models"
)

type UnitTestSuite struct {
	suite.Suite
	models      *models.MockTodoRepository
	controllers PageControllers
}

func (suite *UnitTestSuite) SetupTest() {
	mockModels := new(models.MockTodoRepository)
	controllers := NewHomepageControllers(mockModels)

	suite.models = mockModels
	suite.controllers = controllers
}

func (suite *UnitTestSuite) TestHomePageController() {

	todosRes := []models.TodoGorm{
		{
			Title:    "Test Title",
			Content:  "Test Content",
			Priority: "normal",
			IsDone:   false,
		},
	}

	suite.models.On("GetTodos").Return(todosRes)

	suite.Assert().NoError(suite.controllers.HomePageController(nil))
}
