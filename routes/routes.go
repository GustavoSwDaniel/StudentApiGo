package routes

import (
	"api-go-gin/controllers"

	"github.com/gin-gonic/gin"
)

func HandleRequests() {
	r := gin.Default()
	r.GET("/student", controllers.GetAllStudents)
	r.GET("/student/cpf/:cpf", controllers.GetStudentByCpf)
	r.GET("/student/:id", controllers.GetStudentById)
	r.POST("/student", controllers.CreateNewStudent)
	r.PATCH("/student/:id", controllers.UpdateStudent)
	r.DELETE("/student/:id", controllers.DeleteStudent)

	r.Run()
}
