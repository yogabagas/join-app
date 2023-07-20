package rest

import (
	"database/sql"
	"fmt"
	"github/yogabagas/print-in/config"
	"github/yogabagas/print-in/registry"
	groupV1 "github/yogabagas/print-in/transport/rest/group/v1"
	"github/yogabagas/print-in/transport/rest/handler"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	_ "github/yogabagas/print-in/docs"

	httpSwagger "github.com/swaggo/http-swagger"

	"github.com/gorilla/mux"
)

type Option struct {
	Port         string
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
	Sql          *sql.DB
	Mux          *mux.Router
}

type Handler struct {
	option    *Option
	listenErr chan error
}

// NewRest
// @title Join App API
// @version 1.0
// @description Mentoring Service API
// @termsOfService http://swagger.io/terms/
// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @BasePath /
// @name Mentoring App
func NewRest(o *Option) *Handler {

	reg := registry.NewRegistry(
		registry.NewSQLConn(o.Sql),
	)

	appController := reg.NewAppController()

	handlerImpl := handler.HandlerImpl{
		Controller: appController,
	}

	r := mux.NewRouter()

	URI := fmt.Sprintf("%s%s", config.GlobalCfg.App.Host, config.GlobalCfg.App.Port)

	r.PathPrefix("/swagger/").Handler(httpSwagger.Handler(
		httpSwagger.URL(fmt.Sprintf("http://%s/swagger/doc.json", URI)),
		httpSwagger.DeepLinking(true),
		httpSwagger.DocExpansion("none"),
		httpSwagger.DomID("swagger-ui"),
	)).Methods(http.MethodGet)
	r.PathPrefix("/health").HandlerFunc(handlerImpl.Healthcheck)

	v1 := r.PathPrefix("/v1").Subrouter()

	groupV1.NewUsersV1(handlerImpl, v1)
	groupV1.NewRolesV1(handlerImpl, v1)

	o.Mux = r

	return &Handler{
		option: o,
	}

}

func (h *Handler) Serve() {

	log.Printf("HTTP serve at : %s%s", config.GlobalCfg.App.Host, config.GlobalCfg.App.Port)

	srv := &http.Server{
		Handler:      h.option.Mux,
		Addr:         config.GlobalCfg.App.Port,
		ReadTimeout:  h.option.ReadTimeout,
		WriteTimeout: h.option.WriteTimeout,
	}

	h.listenErr <- srv.ListenAndServe()
}

func (h *Handler) ListenError() <-chan error {
	return h.listenErr
}

func (h *Handler) SignalCheck() {
	term := make(chan os.Signal, 1)

	signal.Notify(term, os.Interrupt, syscall.SIGTERM)
	select {
	case <-term:
		log.Println("Exiting gracefully . . .")
	case err := <-h.ListenError():
		log.Println("Error starting web server, exiting gracefully:", err)
	}
}
