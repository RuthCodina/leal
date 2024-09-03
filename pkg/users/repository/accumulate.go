package repository

import (
	"context"
	"database/sql"
	"log"

	"github.com/leal/pkg/helpers"
	"github.com/leal/pkg/users/domain"
	"github.com/shopspring/decimal"
)

func (r *Repository) AccumulatePoints(ctx context.Context, name string, userId int, promotion float32) (acc *domain.UserAccumulated, err error) {
	var (
		exists   bool
		userName string
	)

	if err = r.db.QueryRow(queryExistsUserId, userId).Scan(&exists); err != nil {
		if err == sql.ErrNoRows {
			return nil, helpers.ErrNoExist
		}
		return nil, err
	} else if exists {
		err = r.db.QueryRow(nameQuery, userId).Scan(&userName)
		if err != nil {
			log.Println("cannot fetch db user name")
			return nil, err
		}
	}
	if userName == name {
		userDB := domain.UserDB{}
		err := r.db.QueryRow(userQuery, userId, name).Scan(&userDB.Id, &userDB.Name, &userDB.LastName, &userDB.DNI, &userDB.State, &userDB.Points, &userDB.LealCoins, &userDB.CreatedAt, &userDB.UpdatedAt)
		if err != nil {
			log.Println("cannot fetch db user name")
			return nil, err
		}
		usr := domain.User{
			Points:    userDB.Points,
			LealCoins: userDB.LealCoins,
			CreatedAt: userDB.CreatedAt.Time,
			UpdatedAt: userDB.UpdatedAt.Time,
		}
		log.Println(usr)

		userCashback := decimal.NewFromFloat32(promotion).Mul(usr.LealCoins)
		userPoints := decimal.NewFromFloat32(promotion).Mul(usr.Points)

		res, err := r.db.Exec(updateUserCashbackPoints, userPoints, userCashback, userId, userDB.Name)
		if err != nil {
			log.Println("cannot fetch db user name")
			return nil, err
		}
		log.Println(res.RowsAffected())
		acc = &domain.UserAccumulated{
			Cashback: userCashback,
			Points:   userPoints,
		}
	}

	return acc, nil
}

func (r *Repository) GetCampaignName(ctx context.Context, branchId int) (string, error) {
	var (
		branchActivated bool
		campaignName    string
	)
	if err := r.db.QueryRow(branchActivatedQuery, branchId).Scan(&branchActivated); err != nil {
		if err == sql.ErrNoRows {
			return "", helpers.ErrActiveBranchCampign
		}
		return "", err
	} else if branchActivated {
		if err := r.db.QueryRow(campaignNameQuery, branchId).Scan(campaignName); err != nil {
			return "", err
		}
	}

	return campaignName, nil
}
