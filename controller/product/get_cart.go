package product

import (
	"encoding/json"
	"go-learn/library/response"
	"net/http"
)

func (c *_ControllerProduct) GetCart(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var (
		errResponse = response.NewResponse().
				WithCode(http.StatusUnprocessableEntity).
				WithStatus("Failed").
				WithMessage("Failed")
		succResponse = response.NewResponse().
				WithStatus("Success").
				WithMessage("Success")
		bearer = r.Header.Get("Authorization")
	)

	data, err := c.service.ProductService.GetCart(bearer)
	if err != nil {
		response := *errResponse.WithError(err)
		output, _ := json.Marshal(response)
		w.WriteHeader(http.StatusBadRequest)
		w.Write(output)
		return
	}
	response := *succResponse.WithData(data)
	object, err := json.Marshal(response)
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
