package controllers

import (
	"net/http"

	"github.com/Polengolas/gin-api-rest/database"
	"github.com/Polengolas/gin-api-rest/models"
	"github.com/gin-gonic/gin"
)

func ShowAllStudents(c *gin.Context) {
	var students []models.Student

	database.DB.Find(&students)

	c.JSON(200, students)
}

func ShowUniqueStudent(c *gin.Context) {
	id := c.Param("id")

	var student models.Student

	database.DB.Find(&student, id)

	if student.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"Not found": "Student not found"})
		return
	}

	c.JSON(200, student)
}

func TestQuaryParameter(c *gin.Context) {
	name := c.Params.ByName("name")
	c.JSON(200, gin.H{
		"API say:": "Hello " + name + ", how are you?",
	})
}

func CreateNewStudent(c *gin.Context) {
	var student models.Student

	if err := c.ShouldBindJSON(&student); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	database.DB.Create(&student)
	c.JSON(http.StatusCreated, student)
}

func DeleteStudent(c *gin.Context) {
	id := c.Param("id")

	var student models.Student

	database.DB.Delete(&student, id)

	c.JSON(http.StatusOK, gin.H{
		"data": "Student deleted successfully"})
}

func EditStudent(c *gin.Context) {
	id := c.Param("id")
	var student models.Student

	database.DB.First(&student, id)

	if err := c.ShouldBindJSON(&student); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	database.DB.Model(&student).UpdateColumns(student)

	c.JSON(http.StatusOK, student)
}

func SearchStudentByCPF(c *gin.Context) {
	var student models.Student

	cpf := c.Param("cpf")

	database.DB.Where(&models.Student{CPF: cpf}).First(&student)

	if student.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"Not found": "Student not found"})
		return
	}

	c.JSON(200, student)
}
