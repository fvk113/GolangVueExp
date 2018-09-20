//acount.go
package dominio

import "time"

type Account struct {
	ID         string     `json:"id"`
	EMAIL      string     `json:"email"`
	FIRST_NAME string     `json:"firstName"`
	LAST_NAME  string     `json:"lastName"`
	LOGIN      string     `json:"login"`
	PASSWORD   string     `json:"password"`
	ROLE_NAME  string     `json:"role"`
	IS_ENABLED bool       `json:"enabled"`
	CreatedAt  time.Time  `json:"created"`
	UpdatedAt  time.Time  `json:"updated"`
	DeletedAt  *time.Time `sql:"index" json:"deleted"`
}

type AccountCollection struct {
	Accounts []Account `json:"items"`
}
