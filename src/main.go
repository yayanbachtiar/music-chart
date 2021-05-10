package main

import (
	_ "github.com/go-sql-driver/mysql"

	"database/sql"
	"github.com/go-chi/chi"
	"github.com/yayanbachtiar/music-chart/src/bussiness/domain"
	service2 "github.com/yayanbachtiar/music-chart/src/bussiness/service"
	"github.com/yayanbachtiar/music-chart/src/rest"
	"log"
	"net/http"
	"time"
)

var (
	a                  App
	timeout            = 5 * time.Second
)

type App struct {
	Router *chi.Mux
	DB     *sql.DB
}

func (a *App) Initialize() {
	route:= chi.NewRouter()
	dbURI := "sql6411390:Bz426v1YUC@tcp(sql6.freemysqlhosting.net:3306)/sql6411390"
	db, err := sql.Open("mysql", dbURI)
	if err != nil {
		log.Fatal(err)
	}
	a.DB = db
	for i := 0; i < 5; i++ {
		time.Sleep(time.Duration(i) * time.Second)

		if err = a.DB.Ping(); err == nil {
			break
		}
		log.Println(err)
	}
	if err != nil {
		log.Fatal(err)
	}
	a.DB.SetMaxOpenConns(10)
	dom := domain.InitDomain(a.DB)
	service := service2.Init(dom)
	//a.InitRoutes()
	a.Router = rest.InitRoutes(route, service)
}

// @title APIs with chi swagger and jwt
// @version 1.0
// @description APIs with chi swagger and jwt
// @BasePath /
//
//func (a *App) InitRoutes() {
//	a.Router.Use(middleware.Logger)
//	a.Router.Use(middleware.Timeout(10 * time.Second))
//	a.Router.Use(render.SetContentType(render.ContentTypeJSON))
//
//	a.Router.Get("/ping", func(w http.ResponseWriter, r *http.Request) {
//		w.Write([]byte("pong"))
//	})
//	a.Router.Mount("/swagger", httpSwagger.Handler(
//		httpSwagger.URL("http://localhost:8081/swagger/doc.json")))
//
//	a.Router.Post("/register", a.RegisterUser)
//	a.Router.Post("/login", a.Login)
//	a.Router.Route("/secret", func(r chi.Router) {
//		r.Use(MyMiddleware)
//		r.Use(IsAdmin)
//		r.Get("/claims", a.Claims)
//	})
//	//a.Router.Post("/account/{from_account_number}/transfer", a.RestTransferBalance)
//}

// Run simply starts the application.
func (a *App) Run(addr string) {
	server := http.Server{
		ReadTimeout:  timeout,
		WriteTimeout: timeout,
		Handler:      a.Router,
		Addr:         addr,
	}
	log.Fatal(server.ListenAndServe())
}

func main() {
	a = App{}
	a.Initialize()
	a.Run(":8081")
}
