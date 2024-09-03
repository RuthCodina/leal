package repository

import (
	"context"
	"log"

	"github.com/leal/pkg/users/domain"
)

func (r *Repository) UpdateUser(ctx context.Context, id int, name string, acc *domain.UserAccumulated) (*domain.User, error) {
	userDB := domain.UserDB{}
	err := r.db.QueryRow(userQuery, id, name).Scan(&userDB.Id, &userDB.Name, &userDB.LastName, &userDB.DNI, &userDB.State, &userDB.Points, &userDB.LealCoins, &userDB.CreatedAt, &userDB.UpdatedAt)
	if err != nil {
		log.Println("cannot fetch db user name")
		return nil, err
	}
	usr := domain.User{
		Id:        userDB.Id,
		Name:      userDB.Name,
		LastName:  userDB.LastName,
		DNI:       userDB.DNI,
		State:     userDB.State,
		Points:    userDB.Points.Add(acc.Points),
		LealCoins: userDB.LealCoins.Add(acc.Cashback),
		CreatedAt: userDB.CreatedAt.Time,
		UpdatedAt: userDB.UpdatedAt.Time,
	}
	return &usr, nil
}
