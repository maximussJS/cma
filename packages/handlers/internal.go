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

type InternalHandler struct {
	state *services.State
}

func NewInternalHandler() *InternalHandler {
	return &InternalHandler{
		state: services.StateSingletonInstance(),
	}
}

func (h *InternalHandler) getLastBlock(w http.ResponseWriter, r *http.Request) {
	blockchain, err := utils.GetBlockchainTypeParam(r)

	if err != nil {
		responses.ErrorJson(w, "Invalid blockchain type", http.StatusBadRequest)
		return
	}

	content, err := h.state.Get(blockchain)

	if err != nil {
		log.Println("Error reading file:", err)
		responses.ErrorJson(w, "Error reading block data", http.StatusInternalServerError)
		return
	}

	if content == nil {
		responses.SuccessNull(w)
		return
	}

	responses.SuccessJson(w, content.String())
}

func (h *InternalHandler) postLastBlock(w http.ResponseWriter, r *http.Request) {
	blockchain, err := utils.GetBlockchainTypeParam(r)

	if err != nil {
		responses.ErrorJson(w, "Invalid blockchain type", http.StatusBadRequest)
		return
	}

	var block models.StateBlock

	if err := json.NewDecoder(r.Body).Decode(&block); err != nil {
		responses.ErrorJson(w, "Invalid JSON data"+err.Error(), http.StatusBadRequest)
		return
	}

	if err := utils.ValidateStruct(block); err != nil {
		responses.ErrorJson(w, "Validation error: "+err.Error(), http.StatusBadRequest)
		return
	}

	h.state.Set(blockchain, &block)

	responses.SuccessJson(w, block.String())
}

func (h *InternalHandler) RegisterRoutes(router *mux.Router) {
	router.HandleFunc("/internal/{blockchain}/last-block", h.getLastBlock).Methods("GET")
	router.HandleFunc("/internal/{blockchain}/last-block", h.postLastBlock).Methods("POST")
}
