package repo

import (
	"time"

	"github.com/gin-gonic/gin"
)

type UserStorageI interface {
	Create(ctx gin.Context, req *User) (*User, error)
	Update(ctx gin.Context, req *User) error
	Get(ctx gin.Context, req *User) (*User, error)
	Delete(ctx gin.Context, req *User) error 

}

type User struct {
	ID        string
	FirstName string
	LastName  string
	Email     string
	Password  string
	CreatedAt time.Time
}

type UpdateUser struct {
	ID        string
	FirstName string
	LastName  string
}
