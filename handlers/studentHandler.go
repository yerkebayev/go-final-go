package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/yerkebayev/go-final-go/models"
	"net/http"
)

func (h Handler) AddStudent(c *gin.Context) {
	var student models.Student
	err := c.BindJSON(&student)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	h.DB.Create(&student)
	c.JSON(http.StatusOK, &student)
}
