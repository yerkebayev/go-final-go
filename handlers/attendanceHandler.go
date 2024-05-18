package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/yerkebayev/go-final-go/models"
	"net/http"
	"time"
)

func (h Handler) AddAttendance(c *gin.Context) {
	var attendance models.Attendance
	if err := c.BindJSON(&attendance); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	attendance.Time = time.Now().Format(time.RFC3339)
	h.DB.Create(&attendance)
	c.JSON(http.StatusOK, &attendance)
}
