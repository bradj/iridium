package persistence

import (
	"context"
	"database/sql"
	"sync"

	"github.com/bradj/iridium/models"
)

var activeUsers = struct {
	sync.RWMutex
	m map[string]*models.User
}{m: make(map[string]*models.User)}

func GetUser(userID string, ctx context.Context, db *sql.DB) (*models.User, error) {
	var (
		user *models.User
		err  error
	)

	activeUsers.RLock()
	user, ok := activeUsers.m[userID]
	activeUsers.RUnlock()

	if ok == false {
		user, err = models.Users(models.UserWhere.ID.EQ(userID)).One(ctx, db)
		activeUsers.Lock()
		activeUsers.m[userID] = user
		activeUsers.Unlock()
	}

	if err != nil {
		return nil, err
	}

	return user, nil
}

func PurgeUser(userID string) {
	activeUsers.Lock()
	delete(activeUsers.m, userID)
	activeUsers.Unlock()
}
