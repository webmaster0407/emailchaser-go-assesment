package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"emailchaser.com/backend-go/initializers"
	"emailchaser.com/backend-go/models"
	"emailchaser.com/backend-go/serializers"
)

// Helper function to send a JSON error response
func sendErrorJSON(ctx *gin.Context, statusCode int, message string) {
	ctx.JSON(statusCode, gin.H{"error": message})
}

// Helper function to send a JSON success response
func sendSuccessJSON(ctx *gin.Context, statusCode int, data interface{}) {
	ctx.JSON(statusCode, data)
}

// Helper function to map LeadSerializer data to models.Lead
func mapLeadBodyToModel(inputLead serializers.LeadSerializer) models.Lead {
	return models.Lead{
		FirstName:      inputLead.FirstName,
		MiddleName:     inputLead.MiddleName,
		LastName:       inputLead.LastName,
		JobTitle:       inputLead.JobTitle,
		Email:          inputLead.Email,
		PhoneNumber:    inputLead.PhoneNumber,
		City:           inputLead.City,
		CurrentCompany: inputLead.CurrentCompany,
		CompanyWebsite: inputLead.CompanyWebsite,
		LinkedIn:       inputLead.LinkedIn,
		Status:         inputLead.Status,
	}
}

func CreateLead(ctx *gin.Context) {
	var requestBody serializers.LeadSerializer
	if err := ctx.ShouldBindJSON(&requestBody); err != nil {
		sendErrorJSON(ctx, http.StatusBadRequest, "Invalid request body")
		return
	}

	lead := mapLeadBodyToModel(requestBody)

	if err := createLeadInDB(&lead); err != nil {
		sendErrorJSON(ctx, http.StatusInternalServerError, "Failed to create lead in the database")
		return
	}

	sendSuccessJSON(ctx, http.StatusOK, lead)
}

func GetLead(ctx *gin.Context) {
	leadID, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		sendErrorJSON(ctx, http.StatusBadRequest, "Invalid lead ID")
		return
	}

	lead, err := getLeadFromDB(leadID)
	if err != nil {
		sendErrorJSON(ctx, http.StatusNotFound, "Lead not found in the database")
		return
	}

	sendSuccessJSON(ctx, http.StatusOK, lead)
}

func DeleteLead(ctx *gin.Context) {
	leadID, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		sendErrorJSON(ctx, http.StatusBadRequest, "Invalid lead ID")
		return
	}

	if err := deleteLeadFromDB(leadID); err != nil {
		sendErrorJSON(ctx, http.StatusNotFound, "Lead not found in the database")
		return
	}

	sendSuccessJSON(ctx, http.StatusOK, gin.H{"message": "Lead deleted successfully"})
}

func UpdateLead(ctx *gin.Context) {
	leadID, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		sendErrorJSON(ctx, http.StatusBadRequest, "Invalid lead ID")
		return
	}

	lead, err := getLeadFromDB(leadID)
	if err != nil {
		sendErrorJSON(ctx, http.StatusBadRequest, "Lead not found in the database")
		return
	}

	if err := ctx.ShouldBindJSON(&lead); err != nil {
		sendErrorJSON(ctx, http.StatusBadRequest, "Invalid request body")
		return
	}

	if err := updateLeadInDB(&lead); err != nil {
		sendErrorJSON(ctx, http.StatusInternalServerError, "Failed to update lead in the database")
		return
	}

	sendSuccessJSON(ctx, http.StatusOK, lead)
}

// createLeadInDB creates the lead in the database
func createLeadInDB(lead *models.Lead) error {
	return initializers.DB.Create(lead).Error
}

// getLeadFromDB retrieves the lead from the database by ID
func getLeadFromDB(leadID int) (models.Lead, error) {
	var lead models.Lead
	if result := initializers.DB.First(&lead, leadID); result.Error != nil {
		return models.Lead{}, result.Error
	}
	return lead, nil
}

// deleteLeadFromDB deletes the lead from the database by ID
func deleteLeadFromDB(leadID int) error {
	var lead models.Lead
	if result := initializers.DB.Delete(&lead, leadID); result.Error != nil {
		return result.Error
	}
	return nil
}

// updateLeadInDB updates the lead in the database
func updateLeadInDB(lead *models.Lead) error {
	return initializers.DB.Save(lead).Error
}
