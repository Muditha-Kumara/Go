package server

import (
	"context"
	"goapi/internal/api/handlers/data"
	"goapi/internal/api/handlers/setting"
	"goapi/internal/api/middleware"
	"goapi/internal/api/service"
	"log"
	"net/http"
)

type Server struct {
	ctx        context.Context
	HTTPServer *http.Server
	logger     *log.Logger
}

func NewServer(ctx context.Context, sf *service.ServiceFactory, logger *log.Logger) *Server {

	mux := http.NewServeMux()
	err := setupDataHandlers(mux, sf, logger)
	if err != nil {
		logger.Fatalf("Error setting up data handlers: %v", err)
	}

	middlewares := []middleware.Middleware{
		middleware.BasicAuthenticationMiddleware,
		middleware.CommonMiddleware,
	}

	return &Server{
		ctx:    ctx,
		logger: logger,
		HTTPServer: &http.Server{
			Handler: middleware.ChainMiddleware(mux, middlewares...),
		},
	}
}

func (api *Server) Shutdown() error {
	api.logger.Println("Gracefully shutting down server...")
	return api.HTTPServer.Shutdown(api.ctx)
}

func (api *Server) ListenAndServe(addr string) error {
	api.HTTPServer.Addr = addr
	return api.HTTPServer.ListenAndServe()
}

// * REST API handlers
func setupDataHandlers(mux *http.ServeMux, sf *service.ServiceFactory, logger *log.Logger) error {
       // Register GET /setting endpoint
       mux.HandleFunc("GET /setting", func(w http.ResponseWriter, r *http.Request) {
	       setting.GetSetting(w, r)
       })
       ds, err := sf.CreateDataService(service.SQLiteDataService)
       if err != nil {
	       return err
       }

       mux.HandleFunc("GET /data/{vehical_id}", func(w http.ResponseWriter, r *http.Request) {
	       data.GetByIDHandler(w, r, logger, ds)
       })

	mux.HandleFunc("OPTIONS /*", func(w http.ResponseWriter, r *http.Request) {
		data.OptionsHandler(w, r)
	})
	mux.HandleFunc("POST /data", func(w http.ResponseWriter, r *http.Request) {
		data.PostHandler(w, r, logger, ds)
	})
	mux.HandleFunc("PUT /data", func(w http.ResponseWriter, r *http.Request) {
		data.PutHandler(w, r, logger, ds)
	})
	mux.HandleFunc("GET /data", func(w http.ResponseWriter, r *http.Request) {
		data.GetHandler(w, r, logger, ds)
	})
	mux.HandleFunc("GET /data/{device_id}/{vehical_id}", func(w http.ResponseWriter, r *http.Request) {
		data.GetByIDHandler(w, r, logger, ds)
	})
	mux.HandleFunc("DELETE /data/{device_id}/{vehical_id}", func(w http.ResponseWriter, r *http.Request) {
		data.DeleteHandler(w, r, logger, ds)
	})
	return err
}
