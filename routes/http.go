package routes

import (
	"database/sql"
	"encoding/json"
	"io"
	"io/ioutil"
	"log"

	"github.com/bradj/iridium/config"
)

// App gets instantiated once at the beginning of the app and then
// the other types are constructed with it. It is read-only and never changes
// during the running of the application. It must contain pointer-like
// types so everything is accessing the same stuff. Its underlying types
// must be safe for access by multiple goroutines.
type App struct {
	DB     *sql.DB
	Config config.TomlConfig
}

// HTTP deals with http responses, the reason to do this is that you have
// a tiny layer that deals with the HTTP protocol, and responding in JSON
// or whatever. This makes testing very separate.
type HTTP struct {
	App    // embeds app for easy access
	Logger *log.Logger
}

func (h HTTP) bodyDecode(body io.ReadCloser, t interface{}) error {
	buf, err := ioutil.ReadAll(body)

	if err != nil {
		return err
	}

	err = json.Unmarshal(buf, &t)

	if err != nil {
		return err
	}

	return nil
}
