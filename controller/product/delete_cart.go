package product

import (
	"encoding/json"
	"go-learn/entities"
	"go-learn/library/response"
	"net/http"
)

func (c *_ControllerProduct) DeleteCart(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var (
		payload     entities.CartDeletePayload
		errResponse = response.NewResponse().
				WithCode(http.StatusUnprocessableEntity).
				WithStatus("Failed").
				WithMessage("Failed")
		succResponse = response.NewResponse().
				WithStatus("Success").
				WithMessage("Success")
		bearer = r.Header.Get("Authorization")
	)

	if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
		response := *errResponse.WithError(err.Error())
		output, _ := json.Marshal(response)
		w.WriteHeader(http.StatusBadRequest)
		w.Write(output)
		return
	}

	if err := c.service.ProductService.DeleteCart(bearer, payload.CartID); err != nil {
		response := *errResponse.WithError(err.Error())
		output, _ := json.Marshal(response)
		w.WriteHeader(http.StatusBadRequest)
		w.Write(output)
		return
	}

	object, err := json.Marshal(succResponse)
	if err != nil {
		response := *errResponse.WithError(err)
		output, _ := json.Marshal(response)
		w.WriteHeader(http.StatusBadRequest)
		w.Write(output)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(object)
}
