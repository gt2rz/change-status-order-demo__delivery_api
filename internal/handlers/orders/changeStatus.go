package orders

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/gt2rz/change-status_order_demo/delivery_api/internal/constants"
	"github.com/gt2rz/change-status_order_demo/delivery_api/internal/servers"
	"github.com/gt2rz/change-status_order_demo/delivery_api/internal/utils"
)

type ChangeStatusRequest struct {
	StatusId int `json:"status_id"`
}

type ChangeStatusResponse struct {
	Status  bool   `json:"status"`
	Message string `json:"message"`
}

func ChangeStatusHandler(server *servers.HttpServer) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		var request = ChangeStatusRequest{}

		params := mux.Vars(r)
		orderId := params["orderId"]

		if !existsOrder(orderId) {
			utils.SendHttpResponseError(w, errors.New("orden no existe"), http.StatusNotFound)
			return
		}

		err := json.NewDecoder(r.Body).Decode(&request)
		if err != nil {
			utils.SendHttpResponseError(w, constants.ErrBadRequest, http.StatusBadRequest)
			return
		}

		fmt.Println(request.StatusId)

		utils.SendHttpResponse(w, ChangeStatusResponse{
			Status:  constants.Success,
			Message: constants.ChangeStatusSuccessfully,
		}, http.StatusOK)
	}
}

func existsOrder(orderId string) bool {
	if orderId != "1"{
		return false
	}
	return true
}
