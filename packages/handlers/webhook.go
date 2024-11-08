package handlers

import (
	"cma/packages/models"
	"cma/packages/responses"
	"cma/packages/services"
	"cma/packages/utils"
	"encoding/json"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

type WebhookHandler struct {
	consistency *services.Consistency
	slack       *services.Slack
}

func NewWebhookHandler() *WebhookHandler {
	return &WebhookHandler{
		consistency: services.NewConsistency(),
		slack:       services.NewSlack(),
	}
}

func (h *WebhookHandler) sendErrorMessage(orig error) {
	if utils.IsErrorNeedToSend(orig) {
		if err := h.slack.SendMessage(orig.Error()); err != nil {
			log.Printf("Error sending error message %s", err)
		}
	} else {
		log.Printf("error %s", orig)
	}
}

func (h *WebhookHandler) handle(w http.ResponseWriter, r *http.Request) {
	blockchain, err := utils.GetBlockchainTypeParam(r)

	if err != nil {
		responses.ErrorJson(w, "Invalid blockchain type", http.StatusBadRequest)
		return
	}

	var blocks models.Blocks

	if err := json.NewDecoder(r.Body).Decode(&blocks); err != nil {
		responses.ErrorJson(w, "Invalid JSON data"+err.Error(), http.StatusBadRequest)
		return
	}

	if err := utils.ValidateStruct(blocks); err != nil {
		responses.ErrorJson(w, "Validation error: "+err.Error(), http.StatusBadRequest)
		return
	}

	if err := h.consistency.Check(blockchain, blocks.Data, blocks.DirectionString()); err != nil {
		h.sendErrorMessage(err)
		responses.ErrorJson(w, "Block Order error", http.StatusBadRequest)
		return
	}

	responses.SuccessEmpty(w)
	return
}

func (h *WebhookHandler) RegisterRoutes(router *mux.Router) {
	router.HandleFunc("/webhook/{blockchain}", h.handle).Methods("POST")
}
