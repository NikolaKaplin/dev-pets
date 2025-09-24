package handlers

import (
	"dev-pets-backend/internal/config"
	"dev-pets-backend/internal/models"
	"dev-pets-backend/internal/sevices"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"

	"github.com/gofiber/fiber/v2"
)

func GoogleLogin(c *fiber.Ctx) error {
	path := config.AppConfig.GoogleLoginConfig
	url := path.AuthCodeURL("state")
	return c.Redirect(url)
}

func GoogleCallback(c *fiber.Ctx) error {
	token, error := config.AppConfig.GoogleLoginConfig.Exchange(c.Context(), c.FormValue("code"))
	if error != nil {
		panic(error)
	}
	email := sevices.GetEmail(token.AccessToken)
	return c.Status(200).JSON(fiber.Map{"email": email, "login": true})
}

func GetEmail(token string) string {
	reqURL, err := url.Parse("https://www.googleapis.com/oauth2/v1/userinfo")
	if err != nil {
		panic(err)
	}
	ptoken := fmt.Sprintf("Bearer %s", token)
	res := &http.Request{
		Method: "GET",
		URL:    reqURL,
		Header: map[string][]string{
			"Authorization": {ptoken},
		},
	}
	req, err := http.DefaultClient.Do(res)
	if err != nil {
		panic(err)
	}
	defer req.Body.Close()
	body, err := io.ReadAll(req.Body)
	if err != nil {
		panic(err)
	}
	var data models.GoogleResponse
	err = json.Unmarshal(body, &data)
	if err != nil {
		panic(err)
	}
	return data.Email
}
