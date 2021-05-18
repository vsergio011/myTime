package database

import (
	"context"
	"log"

	firebase "firebase.google.com/go/v4"

	"google.golang.org/api/option"
)

func InitializeAppWithServiceAccount() *firebase.App {
	// [START initialize_app_service_account_golang]
	//C:\\Users\\sergi\\Documents\\myTime\\apitasks\\database\\private\\apikey.json
	opt := option.WithCredentialsFile("C:\\Users\\sergi\\Documents\\myTime\\myTime\\apitasks\\database\\private\\apikey.json" /*"/app/database/private/apikey.json"*/)
	app, err := firebase.NewApp(context.Background(), nil, opt)
	if err != nil {
		log.Fatalf("error initializing app: %v\n", err)
	}
	// [END initialize_app_service_account_golang]

	return app
}
