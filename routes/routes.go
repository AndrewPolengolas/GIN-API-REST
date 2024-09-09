package routes

import (
	"github.com/Polengolas/gin-api-rest/controllers"
	"github.com/gin-gonic/gin"
)

func HandleRequests() {
	r := gin.Default()

	r.GET("/students", controllers.ShowAllStudents)
	r.GET("/students/:id", controllers.ShowUniqueStudent)
	r.POST("/students", controllers.CreateNewStudent)
	r.GET("/:name", controllers.TestQuaryParameter)
	r.DELETE("/students/:id", controllers.DeleteStudent)
	r.PATCH("/students/:id", controllers.EditStudent)
	r.GET("/students/cpf/:cpf", controllers.SearchStudentByCPF)
	r.Run()
}
