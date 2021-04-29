package rest

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/go-chi/render"
	"github.com/yayanbachtiar/music-chart/src/bussiness/model"
	"io/ioutil"
	"math/rand"
	"net/http"
	"time"
)

// Register godoc
// @Summary Show a account
// @Description Register User
// @Accept  json
// @Produce  json
// @Param body body model.UserInput true "Account"
// @Success 200 {object} model.User
// @Router /register [post]
func (e *rest) RegisterUser(writer http.ResponseWriter, request *http.Request) {
	body, err := ioutil.ReadAll(request.Body)
	if err != nil {
		render.Render(writer, request, ErrInvalidRequest(err))
		return
	}

	var user model.User
	err = json.Unmarshal(body, &user)
	if err != nil {
		render.Render(writer, request, ErrInvalidRequest(err))
		return
	}

	//err = e.service.saveUser(user)
	//if err != nil {
	//	render.Render(writer, request, ErrNotAlreadyRegisterd(err))
	//	return
	//}
	render.Status(request, http.StatusCreated)
	render.JSON(writer, request, user)
}

// @Summary Show a account
// @Description Login User
// @Accept  json
// @Produce  json
// @Param body body model.UserLogin true "Account"
// @Success 200 {object} model.User
// @Router /login [post]
func (a *rest) Login(writer http.ResponseWriter, request *http.Request) {
	body, err := ioutil.ReadAll(request.Body)
	if err != nil {
		render.Render(writer, request, ErrInvalidRequest(err))
		return
	}

	var user model.User
	err = json.Unmarshal(body, &user)
	if err != nil {
		render.Render(writer, request, ErrInvalidRequest(err))
		return
	}
	userFound := false
	//
	//for _, u := range GetUsers() {
	//	if u.Phone == user.Phone && u.Password == user.Password {
	//		userFound = true
	//		user = u
	//	}
	//}

	if !userFound {
		render.Render(writer, request, ErrUnauthorized(errors.New("Unauthorized")))
		return
	}
	// jwt claims
	var expired = time.Now().Add(time.Duration(1) * time.Hour).Unix()
	claims := model.Claims{
		StandardClaims: jwt.StandardClaims{
			Issuer:    "my-application",
			ExpiresAt: expired,
		},
		Phone: user.Phone,
		Name:  user.Name,
		Role:  user.Role,
	}
	token := jwt.NewWithClaims(
		JWT_SIGNING_METHOD,
		claims,
	)

	signedToken, err := token.SignedString(JWT_SIGNATURE_KEY)
	if err != nil {
		render.Render(writer, request, ErrUnauthorized(err))
		return
	}

	type M struct {
		AccessToken string `json:"access_token"`
		ExpiredAt   int64  `json:"expired_at"`
	}

	render.Status(request, http.StatusOK)
	render.JSON(writer, request, M{
		AccessToken: signedToken,
		ExpiredAt:   expired,
	})
	return
}

// Login godoc
// @Summary Show a account
// @Description Login User
// @Accept  json
// @Produce  json
// @Param access_token query string true "acceess token"
// @Success 200 {object} model.User
// @Router /claims [get]
func (a *rest) Claims(w http.ResponseWriter, r *http.Request) {
	accessToken := r.URL.Query().Get("access_token")
	if accessToken == "" {
		render.Render(w, r, ErrUnauthorized(errors.New("Unauthorized")))
	}
	token, err := jwt.Parse(accessToken, func(token *jwt.Token) (interface{}, error) {
		if method, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Signing method invalid")
		} else if method != JWT_SIGNING_METHOD {
			return nil, fmt.Errorf("Signing method invalid")
		}

		return JWT_SIGNATURE_KEY, nil
	})
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	render.Status(r, http.StatusOK)
	render.JSON(w, r, claims)
	return
}

var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func RandStringRunes(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}
