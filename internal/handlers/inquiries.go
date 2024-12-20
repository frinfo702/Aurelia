package handlers

import (
	"Aurelia/internal/models"
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
)

type inquiryHandler struct {
	db *sql.DB
}

func NewInquiryHandler(db *sql.DB) *inquiryHandler {
	return &inquiryHandler{db: db}
}

func (iH *inquiryHandler) PostInquiryHandler(w http.ResponseWriter, req *http.Request) {
	// receive json type inquiry
	var inquiry models.Inquiry // inquiry vertex to receive converted json

	// convert json to models.inquiry
	err := json.NewDecoder(req.Body).Decode(&inquiry)
	if err != nil {
		log.Printf("failed to decode json: %v", err)
		http.Error(w, "failed to receive the inquiry application", http.StatusInternalServerError)
	}

	// store data to inquiry database
	const query = `
		INSERT VALUES
	`

	// Encode back to JSON and return response
}
