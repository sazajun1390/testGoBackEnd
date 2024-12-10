package Infra

import (
	"context"
	"log"
	"net/http"

	"chatSystem/ent/user"

	//"github.com/golang-jwt/jwt"
	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
)

//var secretKey = []byte("your-secret-key")

// JWTのクレーム構造
type JwtCustomClaims struct {
	Email string `json:"email"`
	jwt.RegisteredClaims
}

type PostUser struct {
	Email    string `json:"email"`
	UserName string `json:"name"`
	Password string `json:"password"`
}

type SignInRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type CreateRoomRequest struct {
	RoomName string `json:"roomName"`
}

type JoinRoomRequest struct {
	RoomID string `json:"roomId"`
}

type SendMessageRequest struct {
	RoomID  string `json:"roomId"`
	Message string `json:"message"`
}

type GetChatRoomMessageRequest struct {
	RoomID string `json:"roomId"`
}

func CreateUser(c echo.Context) error {
	var userRequest PostUser
	client, err := getDBClient()
	if err != nil {
		return c.String(http.StatusInternalServerError, "Failed to connect to database")
	}
	defer client.Close()

	if err := c.Bind(&userRequest); err != nil {
		return c.String(http.StatusBadRequest, "Invalid request")
	}
	passwordHash, err := HashPassword(userRequest.Password)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Failed to hash password"})
	}
	users, err := client.User.
		Create().
		SetEmail(userRequest.Email).
		SetUsername(userRequest.UserName).
		SetPasswordHash(passwordHash).
		Save(context.Background())
	if err != nil {
		return c.String(http.StatusInternalServerError, "Failed to create user")
	}
	log.Println(users)

	//return c.String(http.StatusOK, "Success Create User")
	userQuery, err := client.User.
		Query().
		Where(
			user.EmailEQ(userRequest.Email),
		).
		Only(context.Background())
	if err != nil {
		return c.String(http.StatusInternalServerError, "Failed to get users")
	}

	return c.JSON(http.StatusOK, echo.Map{
		"message": "Success Create User",
		"user":    userQuery,
		"token":   "token",
	})
}

func SignIn(c echo.Context) error {
	var userRequest SignInRequest
	client, err := getDBClient()
	if err != nil {
		//return c.String(http.StatusInternalServerError, "Failed to connect to database")
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Failed to connect to database"})
	}
	defer client.Close()

	if err := c.Bind(&userRequest); err != nil {
		//return c.String(http.StatusBadRequest, "Invalid request")
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "Invalid request"})
	}
	log.Println(userRequest)

	userQuery, err := client.User.
		Query().
		Where(
			user.EmailEQ(userRequest.Email),
		).
		Only(context.Background())
	/*
		hashedPassword, err := HashPassword(user.Password)
			if err != nil {
				return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Failed to hash password"})
			}
	*/
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Failed to hash password"})
	}

	//return c.String(http.StatusOK, "Success SignIn")

	return c.JSON(http.StatusOK, echo.Map{"message": "Success SignIn", "user": echo.Map{"Account": userQuery}})
}

func CreateRoom(c echo.Context) error {
	return c.String(http.StatusOK, "Success Create Room")
}

func JoinRoom(c echo.Context) error {
	return c.String(http.StatusOK, "Success Join Room")
}

func GetChatRoomList(c echo.Context) error {
	return c.String(http.StatusOK, "Success Get Chat Room List")
}

func SendMessage(c echo.Context) error {
	return c.String(http.StatusOK, "Success Send Message")
}

func GetChatRoomMessage(c echo.Context) error {
	return c.String(http.StatusOK, "Success Get Chat Room Message")
}
