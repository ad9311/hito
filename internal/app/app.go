package app

import (
	"fmt"
	"html/template"
	"net/http"
	"os"
	"path/filepath"

	"github.com/BurntSushi/toml"
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
	loadCfg := loadConfigFile()

	console.Log("Generating application configuration")

	connStr := fmt.Sprintf(
		"host=%s port=%s dbname=%s user=%s password=%s",
		loadCfg.Host,
		loadCfg.PGPort,
		loadCfg.DBName,
		loadCfg.User,
		loadCfg.Password,
	)

	conn, err := driverdb.ConnectSQL(connStr)
	console.AssertPanic(err)

	tmplCache, err := Gen()
	console.AssertPanic(err)

	Config.ConnDB = conn
	Config.PortNumber = loadCfg.ServerPort
	Config.UseCache = loadCfg.UseCache
	Config.Production = loadCfg.Production
	Config.TemplateCache = tmplCache
	Config.Session = scs.New()
	Config.Session.Cookie.Persist = true
	Config.Session.Cookie.SameSite = http.SameSiteLaxMode
	Config.Session.Cookie.Secure = true

	return &Config, &Data
}

type loadConfig struct {
	Host       string `toml:"host"`
	PGPort     string `toml:"pg_port"`
	DBName     string `toml:"dbname"`
	User       string `toml:"user"`
	Password   string `toml:"password"`
	ServerPort string `toml:"server_port"`
	UseCache   bool   `toml:"use_cache"`
	Production bool   `toml:"production"`
}

func loadConfigFile() loadConfig {
	configFile := "./config/config.toml"
	console.Log(fmt.Sprintf("Reading %s", filepath.Base(configFile)))
	rPath, err := filepath.Abs(configFile)
	console.AssertPanic(err)

	data, err := os.ReadFile(rPath)
	console.AssertError(err)

	var loadCfg loadConfig
	_, err = toml.Decode(string(data), &loadCfg)
	console.AssertPanic(err)

	return loadCfg
}
