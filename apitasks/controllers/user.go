package controllers

import (
	"context"
	"log"

	"firebase.google.com/go/v4/auth"
	"github.com/vsergio011/apitasks/database"
	"github.com/vsergio011/apitasks/models"
)

func GetUser(ctx context.Context, uid string) (models.User, error) {

	return models.GetUser(ctx, uid)
}

func GetUsers(ctx context.Context) ([]models.User, error) {

	return models.GetUsers(ctx)
}
func AddUser(ctx context.Context, user models.User) (*auth.UserRecord, error) {
	return models.AddUser(ctx, user)
}
func RemoveUser(ctx context.Context, uid string) error {
	return models.RemoveUser(ctx, uid)
}
func RemoveUsers(ctx context.Context, uids []string) error {
	return models.RemoveUsers(ctx, uids)
}
func UpdateUser(ctx context.Context, user models.User, uid string) (*auth.UserRecord, error) {
	return models.UpdateUser(ctx, user, uid)
}

func CreateToken(ctx context.Context) (string, error) {
	var uid = "M1VXKsoMHqO9vMFZaREVSVKJxQX2"
	var app = database.InitializeAppWithServiceAccount()
	client, err := app.Auth(context.Background())
	if err != nil {
		log.Fatalf("error getting Auth client: %v\n", err)
	}

	claims := map[string]interface{}{
		"premiumAccount": true,
	}

	token, err := client.CustomTokenWithClaims(ctx, uid, claims)
	if err != nil {
		log.Fatalf("error minting custom token: %v\n", err)
	}

	return token, err
}

func VerifyToken(idToken string, ctx context.Context) (*auth.Token, error) {

	var app = database.InitializeAppWithServiceAccount()

	client, err := app.Auth(ctx)
	if err != nil {
		log.Fatalf("error getting Auth client: %v\n", err)
	}

	token, err := client.VerifyIDToken(ctx, idToken)
	if err != nil {
		log.Fatalf("error verifying ID token: %v\n", err)
	}

	log.Printf("Verified ID token: %v\n", token)

	return token, err
}
