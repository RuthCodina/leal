package repository

import (
	"context"
	"database/sql"
	"log"

	c "github.com/leal/pkg/campaigns/domain"
	cs "github.com/leal/pkg/campaigns/services"
	"github.com/leal/pkg/helpers"
	"github.com/leal/pkg/users/domain"
	"github.com/shopspring/decimal"
)

func (r *Repository) AccumulatePoints(ctx context.Context, name string, userId int, branchId int, payment domain.UserPayment) (*decimal.Decimal, error) {
	var (
		exists          bool
		dbName          string
		usr             domain.User
		branchActivated bool
		campaignStatus  c.CampaignStatus
	)

	if err := r.db.QueryRow(queryExistsId, userId).Scan(&exists); err != nil {
		if err == sql.ErrNoRows {
			return nil, helpers.ErrNoExist
		}
		return nil, err
	} else if exists {
		err = r.db.QueryRow(queryExistsId, userId).Scan(&dbName)
		if err != nil {
			log.Println("cannot fetch db user name")
			return nil, err
		}
	}

	if dbName == name {
		userDB := domain.UserDB{}
		err := r.db.QueryRow(userQuery, userId, name).Scan(&userDB.Id, &userDB.Name, &userDB.LastName, &userDB.DNI, &userDB.State, &userDB.Points, &userDB.LealCoins, &userDB.CreatedAt, &userDB.UpdatedAt)
		if err != nil {
			log.Println("cannot fetch db user name")
			return nil, err
		}
		usr = domain.User{
			Id:        userDB.Id,
			Name:      userDB.Name,
			LastName:  userDB.LastName,
			DNI:       userDB.DNI,
			State:     userDB.State,
			Points:    userDB.LealCoins,
			LealCoins: userDB.LealCoins,
			CreatedAt: userDB.CreatedAt.Time,
			UpdatedAt: userDB.UpdatedAt.Time,
		}
		log.Println(usr)
	}

	if err := r.db.QueryRow(branchActivatedQuery, branchId).Scan(&branchActivated); err != nil {
		if err == sql.ErrNoRows {
			return nil, helpers.ErrAccumulation
		}
		return nil, err
	} else if branchActivated {
		status, err := cs.Status(ctx, branchId)
		if err != nil {
			return nil, err
		}
		campaignStatus = status
	}

	return nil, nil
}
