// using a very simple in memory cache now
// Future scaling can use a dedicated cache like redis or a pub-sub architecture
package cache

import (
	"github.com/Ahmed-Armaan/finance-dashboard-api.git/internal/database/models"
	"github.com/google/uuid"
	"sync"
)

var (
	store = make(map[uuid.UUID]string)
	mu    sync.RWMutex
)

func Set(userID uuid.UUID, role models.UserRole) {
	mu.Lock()
	defer mu.Unlock()
	store[userID] = string(role)
}

func Get(userID uuid.UUID) (string, bool) {
	mu.Lock()
	defer mu.Unlock()
	role, exists := store[userID]
	if exists {
		delete(store, userID)
	}
	return role, exists
}
