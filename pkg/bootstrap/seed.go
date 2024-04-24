package bootstrap

import (
	"github.com/kyloReneo/go-blog/internal/database/seeder"
	"github.com/kyloReneo/go-blog/pkg/config"
	"github.com/kyloReneo/go-blog/pkg/database"

)

// Bootstraps the Database Migrations
func Seed() {
	config.Set()
	database.Connect()

	seeder.Seed()
}
