package repository

import (
	"database/sql"
	"time"
)

type Client struct {
	ClientId   string
	Secret     string
	Active     bool
	StartWDate time.Time
	EndWDate   sql.NullTime
	ClaimIdref int
}

type Scope struct {
	Id string
}

type Grant struct {
	Id string
}
