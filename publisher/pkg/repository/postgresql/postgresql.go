package postgresql

import (
	"fmt"
	"github.com/jmoiron/sqlx"
)

const (
	usersTable           = "tbl_user"
	storiesTable         = "tbl_story"
	commentsTable        = "tbl_comment"
	stickersTable        = "tbl_sticker"
	storiesStickersTable = "tbl_story_sticker"
)

type Config struct {
	Host     string
	Port     string
	Username string
	DBName   string
	Password string
	SSLMode  string
}

func New(cfg Config) (*sqlx.DB, error) {
	db, err := sqlx.Open("postgres", fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=%s",
		cfg.Host, cfg.Port, cfg.Username, cfg.DBName, cfg.Password, cfg.SSLMode))
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return db, nil
}
