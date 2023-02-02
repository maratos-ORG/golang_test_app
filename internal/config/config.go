package config

import (
	"os"

	"gopkg.in/alecthomas/kingpin.v2"
)

var (
	params BackendParameters
	err    error
)

// GetBackendParameters uses kingpin library for providing parameters via flags or it's env analogues.
// Functions fills BackendParameters struct (https://github.com/alecthomas/kingpin).
func GetBackendParameters() (*BackendParameters, error) {
	app := kingpin.New("backend", "DBaaS backend API.")
	params.Port = app.Flag("port", "Port to run API server. Default: 8080. Env: DBAAS_API_PORT.").Default("8080").Envar("DBAAS_API_PORT").Short('P').String()
	params.DBURL = app.Flag("db", "Database connection string. Ex: postgres://<user>:<password>@<hostname>:<port>/<database>. Env: DBAAS_DB_URL.").Envar("DBAAS_DB_URL").Short('D').Required().String()
	app.HelpFlag.Short('h')
	kingpin.MustParse(app.Parse(os.Args[1:]))
	return &params, nil
}
