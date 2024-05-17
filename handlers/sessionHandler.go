package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/yerkebayev/go-final-go/models"
	"net/http"
)

func (h Handler) AllSession(c *gin.Context) {
	var sessions []models.Session
	h.DB.Preload("Students").Preload("Attendances").Find(&sessions)
	c.IndentedJSON(http.StatusOK, sessions)
}
func (h Handler) AddSession(c *gin.Context) {
	var session models.Session
	if err := c.BindJSON(&session); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	h.DB.Create(&session)
	c.JSON(http.StatusOK, &session)
}
func (h Handler) UpdateSession(c *gin.Context) {
	id := c.Param("id")
	var session models.Session
	if err := h.DB.Where("id = ?", id).First(&session).Error; err != nil {
		c.AbortWithStatus(http.StatusNotFound)
		return
	}
	if err := c.BindJSON(&session); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	h.DB.Save(&session)
	c.JSON(http.StatusOK, &session)
}
func (h Handler) DeleteSession(c *gin.Context) {
	id := c.Param("id")
	var session models.Session

	if err := h.DB.Where("id = ?", id).First(&session).Error; err != nil {
		c.AbortWithStatus(http.StatusNotFound)
		return
	}

	h.DB.Delete(&session)
	c.Status(http.StatusOK)
}
