package controllers

import (
	"github.com/SpaceTesla/orbis-be/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

var deployments []models.Deployment // Temporary in-memory store

// GetDeployments Get all deployments
func GetDeployments(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"deployments": deployments})
}

// CreateDeployment Create a new deployment
func CreateDeployment(c *gin.Context) {
	var newDeployment models.Deployment
	if err := c.ShouldBindJSON(&newDeployment); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	deployments = append(deployments, newDeployment)
	c.JSON(http.StatusCreated, newDeployment)
}
