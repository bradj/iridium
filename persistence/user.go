package persistence

import (
	"context"
	"database/sql"

	"github.com/bradj/iridium/models"
)

var (
	activeUsers map[string]*models.User
)

func GetUser(userID string, ctx context.Context, db *sql.DB) (*models.User, error) {
	if activeUsers == nil {
		activeUsers = make(map[string]*models.User)
	}

	var (
		user *models.User
		err  error
	)

	user, ok := activeUsers[userID]

	if ok == false {
		user, err = models.Users(models.UserWhere.ID.EQ(userID)).One(ctx, db)
		activeUsers[userID] = user
	}

	if err != nil {
		return nil, err
	}

	return user, nil
}

func PurgeUser(userID string) {
	delete(activeUsers, userID)
}
