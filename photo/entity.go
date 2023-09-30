package photo

import (
	"crowdfunding/user"
	"time"
)

type Photo struct {
	ID             int
	Title          string
	Caption  	   string
	PhotoURL       string
	UserID		   int
	CreatedAt      time.Time
	UpdatedAt      time.Time  
	User 		   user.User
}