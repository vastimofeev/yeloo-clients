package controllers

import (
	"net/http"
	"strconv"

	"yeloo-clients/internal/models"
	"yeloo-clients/internal/services"

	"github.com/gin-gonic/gin"
)

type ProfileController struct {
	Service services.ProfileServiceInterface
}

// GetProfile godoc
// @Summary Get a profile
// @Description Retrieve a user profile by its ID
// @Tags profiles
// @Param id path int true "Profile ID"
// @Success 200 {object} models.Profile
// @Failure 400 {object} models.ErrorResponse "Invalid ID"
// @Failure 404 {object} models.ErrorResponse "Profile not found"
// @Router /profile/{id} [get]
func (c *ProfileController) GetProfile(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	profile, err := c.Service.GetProfile(uint(id))
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Profile not found"})
		return
	}

	ctx.JSON(http.StatusOK, profile)
}

// CreateProfile godoc
// @Summary Create a new profile
// @Description Add a new user profile to the database
// @Tags profiles
// @Accept json
// @Produce json
// @Param profile body models.Profile true "Profile data"
// @Success 201 {object} models.Profile
// @Failure 400 {object} models.ErrorResponse "Invalid input"
// @Failure 409 {object} models.ErrorResponse "Profile already exists"
// @Router /profile [post]
func (c *ProfileController) CreateProfile(ctx *gin.Context) {
	var profile models.Profile
	if err := ctx.ShouldBindJSON(&profile); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := c.Service.CreateProfile(&profile); err != nil {
		ctx.JSON(http.StatusConflict, gin.H{"error": "Profile already exists"})
		return
	}

	ctx.JSON(http.StatusCreated, profile)
}

// DeleteProfile godoc
// @Summary Delete a profile
// @Description Remove a user profile by its ID
// @Tags profiles
// @Param id path int true "Profile ID"
// @Success 200 {object} models.SuccessResponse "Profile deleted"
// @Failure 400 {object} models.ErrorResponse "Invalid ID"
// @Failure 404 {object} models.ErrorResponse "Profile not found"
// @Router /profile/{id} [delete]
func (c *ProfileController) DeleteProfile(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	if err := c.Service.DeleteProfile(uint(id)); err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Profile not found"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Profile deleted"})
}
