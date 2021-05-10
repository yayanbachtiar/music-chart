package rest

import (
	"context"
	"errors"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/render"
	httpSwagger "github.com/swaggo/http-swagger"
	_ "github.com/swaggo/files" // swagger embed files
_ "github.com/yayanbachtiar/music-chart/docs"
	"github.com/yayanbachtiar/music-chart/src/bussiness/model"
	"github.com/yayanbachtiar/music-chart/src/bussiness/service"
	"net/http"
	"sync"
	"time"
)
type REST interface{}

var once = &sync.Once{}

type rest struct {
	Router *chi.Mux
	service *service.Services
}

var (
	JWT_SIGNATURE_KEY = []byte("my-application-app")
	JWT_SIGNING_METHOD = jwt.SigningMethodHS256
)

// @title APIs with chi swagger and jwt
// @version 1.0
// @description APIs with chi swagger and jwt
// @BasePath /

func InitRoutes(router *chi.Mux, s *service.Services) *chi.Mux {
	var e *rest
	once.Do(func() {
		e = &rest{
			Router: router,
			service: s,
		}
		e.Serve()
	})
	return e.Router
}

func (r *rest) Serve() {
	r.Router.Use(middleware.Logger)
	r.Router.Use(middleware.Timeout(10 * time.Second))
	r.Router.Use(render.SetContentType(render.ContentTypeJSON))

	r.Router.Get("/ping", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("pong"))
	})
	r.Router.Mount("/swagger", httpSwagger.Handler(
		httpSwagger.URL("http://localhost:8081/swagger/doc.json")))

	r.Router.Post("/register", r.RegisterUser)
	r.Router.Post("/login",  r.Login)
	//r.Router.Route("/secret", func(r chi.Router) {
	//	r.Use(MyMiddleware)
	//	r.Use(IsAdmin)
	//	r.Get("/claims", r.Claims)
	//})
}

// HTTP middleware setting a value on the request context
func MyMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authorizationHeader := r.URL.Query().Get("Authorization")
		//if !strings.Contains(authorizationHeader, "Bearer") {
		//	return
		//}
		if authorizationHeader == "" {
			http.Error(w, "Invalid token", http.StatusBadRequest)
			return
		}

		//tokenString := strings.Replace(authorizationHeader, "Bearer ", "", -1)
		token, err := jwt.ParseWithClaims(authorizationHeader, &model.Claims{}, func(token *jwt.Token) (interface{}, error) {
			return []byte(JWT_SIGNATURE_KEY), nil
		})
		if err != nil {
			http.Error(w, err.Error(), http.StatusUnauthorized)
			return
		}
		if claims, ok := token.Claims.(*model.Claims); ok && token.Valid {
			fmt.Printf("%v %v", claims.Name, claims.StandardClaims.ExpiresAt)
			ctx := context.WithValue(r.Context(), "role", claims.Role)
			ctx = context.WithValue(r.Context(), "phone", claims.Phone)
			ctx = context.WithValue(r.Context(), "name", claims.Name)
			next.ServeHTTP(w, r.WithContext(ctx))
			return
		}
		http.Error(w, "Unathorized", http.StatusUnauthorized)
		return
	})
}
func IsAdmin(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		role := r.Context().Value("role")
		if role == "admin" {
			next.ServeHTTP(w, r)
			return
		}
		render.Render(w, r, ErrInvalidRequest(errors.New("unauthenticated")))
		return
	})
}
