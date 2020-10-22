package probe

import (
	"net/http"
	"rest-api/utils"
)

// Index function
func Index(w http.ResponseWriter, r *http.Request) {
	response := utils.DTResponse{
		Message: "Probe",
		Data:    []struct{}{},
	}
	utils.Response(w, 200, response)
}
