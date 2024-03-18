package validation

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/go-playground/validator"
)

// CustomError represents the error structure
type CustomError struct {
	Status       string `json:"status"`
	ErrorMessage string `json:"error"`
}

// Implementing the error interface by adding an Error() method
func (c *CustomError) Error() string {
	return c.ErrorMessage
}

// RespondWithError sends a JSON error response with the provided status code and message
func respondWithError(w http.ResponseWriter, statusCode int, message string) {
	ce := CustomError{
		Status:       fmt.Sprintf("%d", statusCode),
		ErrorMessage: message,
	}
	ceJSON, err := json.Marshal(ce)
	if err != nil {
		http.Error(w, "Error encoding response to JSON", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	w.Write(ceJSON)
}

// Validate validates the request data and writes error response if validation fails
func (c *CustomError) Validate(w http.ResponseWriter, r *http.Request, data interface{}) bool {
	if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
		respondWithError(w, http.StatusBadRequest, "Error decoding JSON")
		return true
	}

	validate := validator.New()
	if err := validate.Struct(data); err != nil {
		return handleValidationErrors(w, err)
	}

	return false
}

func handleValidationErrors(w http.ResponseWriter, err error) bool {
	switch e := err.(type) {
	case validator.ValidationErrors:
		var errorMessages []string
		for _, err := range e {
			errorMessages = append(errorMessages, getErrorMessage(err))
		}
		respondWithError(w, http.StatusBadRequest, errorMessages[0])
		return true
	default:
		http.Error(w, fmt.Sprintf("Unexpected validation error: %v", err), http.StatusInternalServerError)
		return true
	}
}

func getErrorMessage(e validator.FieldError) string {
	switch e.Tag() {
	case "required":
		return fmt.Sprintf("The %s field is required", strings.Replace(e.Field(), "/n", " ", -1))
	case "gte":
		return fmt.Sprintf("The %s field must be greater than or equal to the required value", strings.Replace(e.Field(), "/n", " ", -1))
	case "min":
		return fmt.Sprintf("The %s field must be at least 5 characters long", strings.Replace(e.Field(), "/n", " ", -1))
	default:
		return fmt.Sprintf("Validation error for %s: %s", e.Field(), e.Tag())
	}
}
