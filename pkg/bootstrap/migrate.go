package bootstrap

import (
	"github.com/kyloReneo/go-blog/pkg/config"
	"github.com/kyloReneo/go-blog/pkg/database"
)

// Bootstraps the Database Migrations
func Migrate() {
	config.Set()
	database.Connect()
}
