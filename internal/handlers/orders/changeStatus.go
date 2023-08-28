package orders

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strconv"

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
		orderId, err := strconv.Atoi(params["orderId"])
		if err != nil {
			utils.SendHttpResponseError(w, constants.ErrBadRequest, http.StatusBadRequest)
			return
		}

		if !existsOrder(orderId) {
			utils.SendHttpResponseError(w, errors.New("orden no existe"), http.StatusNotFound)
			return
		}

		err = json.NewDecoder(r.Body).Decode(&request)
		if err != nil {
			utils.SendHttpResponseError(w, constants.ErrBadRequest, http.StatusBadRequest)
			return
		}

		fmt.Println(request.StatusId)
		err = changeStatusOrder(orderId, request.StatusId)
		if err != nil {
			utils.SendHttpResponseError(w, constants.ErrBadRequest, http.StatusInternalServerError)
			return
		}

		utils.SendHttpResponse(w, ChangeStatusResponse{
			Status:  constants.Success,
			Message: constants.ChangeStatusSuccessfully,
		}, http.StatusOK)
	}
}

func existsOrder(orderId int) bool {
	return true
}

func changeStatusOrder(orderId int, statusId int) error {
	return nil
}
