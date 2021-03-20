package models

import (
	"context"
	"log"

	"firebase.google.com/go/v4/auth"
	"github.com/vsergio011/apitasks/database"
)

type User struct {
	Id          int
	Name        string
	Surname     string
	Email       string `json:"Email"`
	PhoneNumber string
	Password    string
	DisplayName string
	Token       string
}

/*type User struct {
	Email            string      `json:"email"`
	ProviderID       string      `json:"providerId"`
	RawID            string      `json:"rawId"`
	CustomClaims     interface{} `json:"CustomClaims"`
	Disabled         bool        `json:"Disabled"`
	EmailVerified    bool        `json:"EmailVerified"`
	ProviderUserInfo []struct {
		Email      string `json:"email"`
		ProviderID string `json:"providerId"`
		RawID      string `json:"rawId"`
	} `json:"ProviderUserInfo"`
	TokensValidAfterMillis int `json:"TokensValidAfterMillis"`
	UserMetadata           struct {
		CreationTimestamp    int64 `json:"CreationTimestamp"`
		LastLogInTimestamp   int   `json:"LastLogInTimestamp"`
		LastRefreshTimestamp int   `json:"LastRefreshTimestamp"`
	} `json:"UserMetadata"`
	TenantID string `json:"TenantID"`
}*/
//firebase function
func GetUser(ctx context.Context, uid string) (*auth.UserRecord, error) {

	// [START get_user_golang]
	// Get an auth client from the firebase.App
	var app = database.InitializeAppWithServiceAccount()

	client, err := app.Auth(ctx)
	if err != nil {
		log.Fatalf("error getting Auth client: %v\n", err)
	}

	u, err := client.GetUser(ctx, uid)
	if err != nil {
		log.Fatalf("error getting user %s: %v\n", uid, err)
	}
	log.Printf("Successfully fetched user data: %v\n", u)
	// [END get_user_golang]

	return u, err
}

func GetUsers(ctx context.Context) (*auth.UserRecord, error) {
	uid := "M1VXKsoMHqO9vMFZaREVSVKJxQX2"
	// [START get_user_golang]
	// Get an auth client from the firebase.App
	var app = database.InitializeAppWithServiceAccount()

	client, err := app.Auth(ctx)
	if err != nil {
		log.Fatalf("error getting Auth client: %v\n", err)
	}

	u, err := client.GetUser(ctx, uid)
	if err != nil {
		log.Fatalf("error getting user %s: %v\n", uid, err)
	}
	log.Printf("Successfully fetched user data: %v\n", u)
	// [END get_user_golang]

	return u, err
}

func AddUser(ctx context.Context, user User) (*auth.UserRecord, error) {
	var app = database.InitializeAppWithServiceAccount()
	client, err := app.Auth(ctx)
	params := (&auth.UserToCreate{}).
		Email(user.Email).
		EmailVerified(true).
		PhoneNumber(user.PhoneNumber).
		Password(user.Password).
		DisplayName(user.DisplayName).
		PhotoURL("http://www.example.com/12345678/photo.png").
		Disabled(false)
	u, err := client.CreateUser(ctx, params)
	if err != nil {
		log.Fatalf("error creating user: %v\n", err)
	}
	log.Printf("Successfully created user: %v\n", u)

	return u, err
}

/*func GetUsers() ([]User, error) {
	sql := `SELECT users.id, users.name,
									 users.surname, users.email FROM users`

	db, err := database.Open()
	if err != nil {
		return nil, err
	}
	defer db.Close()

	var users []User
	rows, err := db.Query(sql)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var b User
		if err := rows.Scan(&b.Id, &b.Name, &b.Surname, &b.Email); err != nil {
			continue
		}
		users = append(users, b)
	}

	return users, nil
}*/
