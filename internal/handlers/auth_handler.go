package handlers

import (
	"encoding/json"
	"net/http"
	"rest-api/internal/models"
	"rest-api/internal/services"

	"github.com/go-playground/validator/v10"
)

var authService services.AuthService

func InitAuthHandler(service services.AuthService) {
	authService = service
}

func Register(w http.ResponseWriter, r *http.Request) {
	var user models.User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		response := models.APIResponse{
			Type:       "error",
			StatusCode: http.StatusBadRequest,
			Message:    "Invalid request payload",
			Data:       nil,
		}
		SendApiResponse(w, response)
		return
	}

	// Validate the user struct
	if err := models.Validate.Struct(user); err != nil {
		var validationErrors []models.ValidationErrors

		for _, err := range err.(validator.ValidationErrors) {
			validationErrors = append(validationErrors, models.ValidationErrors{
				Field:   err.Field(),
				Message: err.Tag(),
			})
		}

		response := models.APIResponse{
			Type:       "error",
			StatusCode: http.StatusBadRequest,
			Message:    "Validation failed",
			Data: map[string][]models.ValidationErrors{
				"error": validationErrors,
			},
		}
		SendApiResponse(w, response)
		return
	}

	// Register the user
	token, err := authService.Register(&user)
	if err != nil {
		response := models.APIResponse{
			Type:       "error",
			StatusCode: http.StatusInternalServerError,
			Message:    err.Error(),
			Data:       nil,
		}
		SendApiResponse(w, response)
		return
	}

	response := models.APIResponse{
		Type:       "success",
		StatusCode: http.StatusCreated,
		Message:    "User registered successfully",
		Data: map[string]interface{}{
			"user": map[string]string{
				"first_name": user.FirstName,
				"last_name":  user.LastName,
				"email":      user.Email,
			},
			"token": token,
		},
	}
	SendApiResponse(w, response)
}

func Login(w http.ResponseWriter, r *http.Request) {
	var user models.User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		response := models.APIResponse{
			Type:       "error",
			StatusCode: http.StatusBadRequest,
			Message:    "Invalid request payload",
			Data:       nil,
		}
		SendApiResponse(w, response)
		return
	}

	token, err := authService.Login(user.Email, user.Password)
	if err != nil {
		response := models.APIResponse{
			Type:       "error",
			StatusCode: http.StatusBadRequest,
			Message:    err.Error(),
			Data:       nil,
		}
		SendApiResponse(w, response)
		return
	}
	response := models.APIResponse{
		Type:       "success",
		StatusCode: http.StatusOK,
		Message:    "User logged in successfully",
		Data: map[string]interface{}{
			"token": token,
		},
	}
	SendApiResponse(w, response)
}
