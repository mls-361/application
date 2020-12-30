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
		host      string
		devel     int
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

	return app
}

// ID AFAIRE.
func (a *Application) ID() string { return a.id }

// Name AFAIRE.
func (a *Application) Name() string { return a.name }

// Version AFAIRE.
func (a *Application) Version() string { return a.version }

// BuiltAt AFAIRE.
func (a *Application) BuiltAt() time.Time { return a.builtAt }

// StartedAt AFAIRE.
func (a *Application) StartedAt() time.Time { return a.startedAt }

// Host AFAIRE.
func (a *Application) Host() string { return a.host }

// Devel AFAIRE.
func (a *Application) Devel() int { return a.devel }

// Initialize AFAIRE.
func (a *Application) Initialize() error {
	s, ok := os.LookupEnv(strings.ToUpper(a.name) + "_DEVEL")
	if ok {
		devel, err := strconv.Atoi(s)
		if err != nil {
			return err
		}

		a.devel = devel
	}

	fqdn, err := fqdn.FQDN()
	if err != nil {
		return err
	}

	a.host = fqdn

	return nil
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

/*
######################################################################################################## @(°_°)@ #######
*/
