package handlers

import (
	"context"
	"encoding/json"
	"net/http"
	"time"

	"github.com/leal/pkg/helpers"
	"github.com/leal/pkg/users/domain"
	"github.com/leal/pkg/users/handlers/dtos"
)

func (h *Handler) AccumulatePoints(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()
	accDto := dtos.AccumulatePointsDTO{}

	err := json.NewDecoder(r.Body).Decode(&accDto)
	if err != nil {
		helpers.ErrInvalidJSON.Send(w)
		return
	}
	payDate, err := time.Parse("2006-01-02 15:04:05", accDto.Payment.PayDate)
	if err != nil {
		helpers.ErrParsingDate.Send(w)
		return
	}

	payment := domain.UserPayment{
		Amount:   accDto.Payment.Amount,
		PayDate:  payDate,
		BranchId: accDto.Payment.BranchId,
	}
	acc, err := h.UserService.AccumulatePoints(ctx, accDto.Name, accDto.Id, payment)

	if err != nil {
		helpers.ErrBadRequest.Send(w)
		return
	}

	helpers.Success(&acc, http.StatusOK).Send(w)
}
