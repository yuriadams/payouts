package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/julienschmidt/httprouter"
  "github.com/yuriadams/payout/services"
  "github.com/yuriadams/payout/models"
)

func extractRankings(r *http.Request) string {
	rawBody := make([]byte, r.ContentLength, r.ContentLength)
	r.Body.Read(rawBody)
	return string(rawBody)
}

func GamePayoutHandler(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
  bytes := []byte(extractRankings(r))
  var rankings []models.Result
  json.Unmarshal(bytes, &rankings)

	result := services.Finalize(p.ByName("id"), rankings)
  // var status int

	// if result == "" {
	// 	status = http.StatusOK
	// } else {
	// 	status = http.StatusInternalServerError
	// }

	Logging("Game finalized succesfully %s.", result)

	RespondWithJSON(w, http.StatusOK, string(result))
}
