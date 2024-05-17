package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/yerkebayev/go-final-go/models"
	"net/http"
	"time"
)

func (h Handler) AllAttendance(c *gin.Context) {
	var attendances []models.Attendance
	h.DB.Find(&attendances)
	c.IndentedJSON(http.StatusOK, attendances)
}
func (h Handler) AddAttendance(c *gin.Context) {
	var attendance models.Attendance
	if err := c.BindJSON(&attendance); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	if attendance.Time.IsZero() {
		attendance.Time = time.Now()
	}

	h.DB.Create(&attendance)
	c.JSON(http.StatusOK, &attendance)
}
