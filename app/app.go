package app

import (
	"net/http"
	"strings"

	"github.com/makasim/sentryhook"

	"github.com/getsentry/sentry-go"

	"github.com/sirupsen/logrus"

	"github.com/gorilla/mux"
)

// Args holds the arguments for App.
type Args struct {
	Router *mux.Router
	Log    *logrus.Logger
	DNS    string
}

// App will hold all the dependencies the application needs.
type App struct {
	db     interface{}
	router *mux.Router
	log    *logrus.Logger
}

// NewApp creates a new instance of the App, registers the routes, and returns the instance.
func NewApp(args Args) App {
	app := App{
		router: args.Router,
		log:    args.Log,
	}

	app.routes()

	if strings.TrimSpace(args.DNS) != "" {
		if err := sentry.Init(sentry.ClientOptions{Dsn: args.DNS}); err != nil {
			app.log.Errorln("failed to connect to Sentry:", err)

			return app
		}

		hook := sentryhook.New([]logrus.Level{
			logrus.PanicLevel,
			logrus.FatalLevel,
			logrus.ErrorLevel,
		})

		app.log.AddHook(hook)

		app.log.Infoln("connected to Sentry")
	}

	return app
}

// StartServer starts the HTTP server on the specified port. Any errors will be returned on the specified channel.
func (app App) StartServer(errorChan chan error, port string) {
	app.log.Infof("Starting server on port %s", port)
	errorChan <- http.ListenAndServe(":"+port, app.router)
}
