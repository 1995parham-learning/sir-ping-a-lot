package service

import (
	"errors"
	"log"
	"net/http"
	"net/url"
	"user/authentication"
	"user/config"
	"user/model"
	"user/request"
	"user/store"

	"github.com/labstack/echo"
	"golang.org/x/crypto/bcrypt"
)

var ErrLoggedOut = errors.New("you are not logged in")

type API struct {
	User   store.User
	URL    store.URL
	Config config.JWT
}

func (a API) Run() {
	e := echo.New()

	e.POST("/register", a.Register)
	e.POST("/login", a.Login)
	e.POST("/url", a.Add)
	e.Logger.Fatal(e.Start(":8080"))
}

func (a API) Register(c echo.Context) error {
	var user model.User

	if err := c.Bind(&user); err != nil {
		return err
	}

	if user.Email == "" || user.Password == "" {
		return echo.NewHTTPError(http.StatusBadRequest, "Email and password cannot be empty")
	}

	hashPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), 14)
	if err != nil {
		log.Fatal(err)
	}

	user.Password = string(hashPassword)

	//nolint: errcheck
	if err := a.User.Insert(user); err != nil {
		c.JSON(http.StatusConflict, err.Error())
	}

	return c.JSON(http.StatusCreated, user)
}

func (a API) Login(c echo.Context) error {
	var user model.User

	if err := c.Bind(&user); err != nil {
		return err
	}

	if user.Email == "" || user.Password == "" {
		return echo.NewHTTPError(http.StatusBadRequest, "Email and password cannot be empty")
	}

	us, err := a.User.Retrieve(user)
	if err != nil {
		return c.JSON(http.StatusNotFound, err.Error())
	}

	token, err := authentication.CreateToken(us.ID, a.Config)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, token)
}

func (a API) Add(c echo.Context) error {
	var newURL request.URL

	token := c.Request().Header.Get("Authorization")

	if err := c.Bind(&newURL); err != nil {
		return err
	}

	in, id := authentication.ValidateToken(token, a.Config)

	if !in {
		return echo.NewHTTPError(http.StatusForbidden, ErrLoggedOut.Error())
	}

	_, err := url.ParseRequestURI(newURL.URL)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	var u model.URL

	u.UserID = id
	u.URL = newURL.URL
	u.Period = newURL.Period

	if err := a.URL.Insert(u); err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusCreated, u)
}
