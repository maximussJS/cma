package utils

import (
	"cma/packages/constants"
	"github.com/gorilla/mux"
	"net/http"
)

func GetBlockchainTypeParam(r *http.Request) (constants.BlockchainType, error) {
	vars := mux.Vars(r)
	blockchainStr := vars["blockchain"]

	return constants.NewBlockchainType(blockchainStr)
}
