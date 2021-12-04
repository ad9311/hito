package app

import (
	"html/template"
	"net/http"

	"github.com/ad9311/hito/internal/console"

	"github.com/ad9311/hito/internal/driverdb"
	"github.com/alexedwards/scs/v2"
)

// InitConfig provides initial configuration fields for the app.
type InitConfig struct {
	ConnDB        *driverdb.DB
	PortNumber    string
	UseCache      bool
	Production    bool
	TemplateCache map[string]*template.Template
	Session       *scs.SessionManager
}

// InitData provides initial data fields that are passed on to templates.
type InitData struct {
	CurrentUser driverdb.User
	UserSlice   []driverdb.User
	StringMap   map[string]string
	BoolMap     map[string]bool
	CSRFToken   string
}

// Config holds the current configuration of the app.
var Config InitConfig

// Data holds the current data that are passed on to templates.
var Data = InitData{
	StringMap: make(map[string]string),
	BoolMap:   make(map[string]bool),
}

// New sets the initial values for the app
// and returns a pointer for the config and data structs.
func New() (*InitConfig, *InitData) {
	console.Log("Generating application configuration")
	conn, err := driverdb.ConnectSQL("host=localhost port=5432 dbname=hito_devel user=ad9311 password=")
	console.AssertPanic(err)

	tmplCache, err := Gen()
	console.AssertPanic(err)

	Config.ConnDB = conn
	Config.PortNumber = ":3000"
	Config.UseCache = false
	Config.Production = false
	Config.TemplateCache = tmplCache
	Config.Session = scs.New()
	Config.Session.Cookie.Persist = true
	Config.Session.Cookie.SameSite = http.SameSiteLaxMode
	Config.Session.Cookie.Secure = true

	return &Config, &Data
}
