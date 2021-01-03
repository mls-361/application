/*
------------------------------------------------------------------------------------------------------------------------
####### application ####### (c) 2020-2021 mls-361 ################################################## MIT License #######
------------------------------------------------------------------------------------------------------------------------
*/

package application

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/mls-361/fqdn"
	"github.com/mls-361/util"
	"github.com/mls-361/uuid"
)

type (
	// Application AFAIRE.
	Application struct {
		id        string
		name      string
		version   string
		builtAt   time.Time
		startedAt time.Time
		debug     int
		host      string
	}
)

// New AFAIRE.
func New(name, version, builtAt string) *Application {
	t, _ := util.EpochStrToTime(builtAt, 0)

	app := &Application{
		id:        uuid.New(),
		name:      name,
		version:   version,
		builtAt:   t,
		startedAt: time.Now(),
	}

	s, ok := app.LookupEnv("DEBUG")
	if ok {
		if debug, err := strconv.Atoi(s); err == nil {
			app.debug = debug
		}
	}

	return app
}

// ID AFAIRE.
func (a *Application) ID() string {
	return a.id
}

// Name AFAIRE.
func (a *Application) Name() string {
	return a.name
}

// Version AFAIRE.
func (a *Application) Version() string {
	return a.version
}

// BuiltAt AFAIRE.
func (a *Application) BuiltAt() time.Time {
	return a.builtAt
}

// StartedAt AFAIRE.
func (a *Application) StartedAt() time.Time {
	return a.startedAt
}

// Debug AFAIRE.
func (a *Application) Debug() int {
	return a.debug
}

// LookupEnv AFAIRE.
func (a *Application) LookupEnv(suffix string) (string, bool) {
	return os.LookupEnv(strings.ToUpper(a.name) + "_" + suffix)
}

// OnError AFAIRE.
func (a *Application) OnError(err error) error {
	fmt.Fprintf( ///////////////////////////////////////////////////////////////////////////////////////////////////////
		os.Stderr,
		"Error: application=%s version=%s builtAt=%s >>> %s\n",
		a.name,
		a.version,
		a.builtAt.String(),
		err,
	)

	return err
}

// Initialize AFAIRE.
func (a *Application) Initialize() error {
	fqdn, err := fqdn.FQDN()
	if err != nil {
		return err
	}

	a.host = fqdn

	return nil
}

// Host AFAIRE.
func (a *Application) Host() string {
	return a.host
}

/*
######################################################################################################## @(°_°)@ #######
*/
