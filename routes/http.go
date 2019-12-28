package routes

import (
	"database/sql"
	"log"

	"github.com/bradj/iridium/auth"
	"github.com/bradj/iridium/config"
)

// App gets instantiated once at the beginning of the app and then
// the other types are constructed with it. It is read-only and never changes
// during the running of the application. It must contain pointer-like
// types so everything is accessing the same stuff. Its underlying types
// must be safe for access by multiple goroutines.
type App struct {
	DB     *sql.DB
	Logger *log.Logger
	Config config.TomlConfig
}

// HTTP deals with http responses, the reason to do this is that you have
// a tiny layer that deals with the HTTP protocol, and responding in JSON
// or whatever. This makes testing very separate.
type HTTP struct {
	App // embeds app for easy access
	JWT auth.JWT

	// Users controllers.Users
}
