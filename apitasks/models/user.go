package models

import (
	"context"
	"errors"
	"log"

	"firebase.google.com/go/v4/auth"
	"github.com/vsergio011/apitasks/database"
	"google.golang.org/api/iterator"
)

type User struct {
	Id          string
	Name        string
	Surname     string
	Email       string `json:"Email"`
	PhoneNumber string
	Password    string
	DisplayName string
	Token       string
	RawID       string `json:"rawId"`
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

//*** CRUD USER **//
func GetUser(ctx context.Context, uid string) (User, error) {

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

	usr := User{
		Id:          u.UID,
		Surname:     "",
		Name:        u.DisplayName,
		Email:       u.Email,
		PhoneNumber: u.PhoneNumber,
	}

	//realizo la consulta en la base de datos propia
	db, err := database.Open()
	if err != nil {
		//error stament
	}
	defer db.Close()

	rows, err := db.Query("SELECT uid,name,surname FROM myTime.users where uid = 'nnaGlkOw9uPJVrqf1pwppxOSXOc2'")
	if err != nil {
		err = errors.New("CANNOT PREPARE STATMENT")
	}
	defer rows.Close()
	for rows.Next() {
		err := rows.Scan(&usr.Id, &usr.Name, &usr.Surname)
		if err != nil {
			log.Fatal(err)
		}
		log.Println(usr.Id, usr.Name, usr.Surname)
	}
	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Successfully fetched user data: %v\n", u)
	// [END get_user_golang]

	return usr, err
}

func GetUsers(ctx context.Context) ([]User, error) {
	var users []*auth.ExportedUserRecord
	var userq []User
	var app = database.InitializeAppWithServiceAccount()
	client, err := app.Auth(ctx)
	// Note, behind the scenes, the Users() iterator will retrive 1000 Users at a time through the API
	/*iter := client.Users(ctx, "")
	for {
		user, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			log.Fatalf("error listing users: %s\n", err)
		}
		log.Printf("read user user: %v\n", user)
	}*/

	// Iterating by pages 100 users at a time.
	// Note that using both the Next() function on an iterator and the NextPage()
	// on a Pager wrapping that same iterator will result in an error.
	pager := iterator.NewPager(client.Users(ctx, ""), 100, "")
	for {

		nextPageToken, err := pager.NextPage(&users)
		if err != nil {
			log.Fatalf("paging error %v\n", err)
		}
		for _, u := range users {
			usr := User{
				Id:          u.UID,
				Surname:     "",
				Name:        u.DisplayName,
				Email:       u.Email,
				PhoneNumber: u.PhoneNumber,
			}
			log.Printf("read user user: %v\n", u)
			//realizo la consulta en la base de datos propia
			db, err := database.Open()
			if err != nil {
				//error stament
			}
			defer db.Close()

			rows, err := db.Query("SELECT uid,name,surname FROM users where uid = ?", u.UID)
			if err != nil {
				err = errors.New("CANNOT PREPARE STATMENT")
			}
			defer rows.Close()
			for rows.Next() {
				err := rows.Scan(&usr.Id, &usr.Name, &usr.Surname)
				if err != nil {
					log.Fatal(err)
				}
				log.Println(usr.Id, usr.Name, usr.Surname)
			}
			err = rows.Err()
			if err != nil {
				log.Fatal(err)
			}
			userq = append(userq, usr)

		}
		if nextPageToken == "" {
			break
		}
	}

	return userq, err
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
	//una vez creado el usuario de autentificacion en firebase creamos una referencia del mismo en nuestra base de datos
	db, err := database.Open()
	if err != nil {
		//error stament
	}
	defer db.Close()

	sqlInsertUser := `INSERT INTO users (uid,name,surname,email) VALUES (?, ?, ?, ?);`
	stmt, err := db.Prepare(sqlInsertUser)
	if err != nil {
		err = errors.New("CANNOT PREPARE STATMENT")
	}

	data, err := stmt.Exec(u.UID, u.DisplayName, "ss", u.Email)
	if err != nil {
		err = errors.New("CANNOT PREPARE STATMENT")
	}

	log.Printf("Successfully created user: %v\n %v", u, data)

	return u, err
}

func UpdateUser(ctx context.Context, user User, uid string) (*auth.UserRecord, error) {
	var app = database.InitializeAppWithServiceAccount()
	client, err := app.Auth(ctx)
	params := (&auth.UserToUpdate{}).
		Email(user.Email).
		EmailVerified(true).
		PhoneNumber(user.PhoneNumber).
		Password(user.Password).
		DisplayName(user.DisplayName).
		PhotoURL("http://www.example.com/12345678/photo.png").
		Disabled(false)
	u, err := client.UpdateUser(ctx, uid, params)
	if err != nil {
		log.Fatalf("error updating user: %v\n", err)
	}
	log.Printf("Successfully updated user: %v\n", u)

	return u, err
}

func RemoveUser(ctx context.Context, uid string) error {
	var app = database.InitializeAppWithServiceAccount()
	client, err := app.Auth(ctx)
	err = client.DeleteUser(ctx, uid)
	if err != nil {
		log.Fatalf("error deleting user: %v\n", err)
	}
	//Eliminamos tambien la referencia en la base de datos propia
	db, err := database.Open()
	if err != nil {
		//error stament
	}
	defer db.Close()

	sqlInsertUser := `DELETE FROM users WHERE uid = ?;`
	stmt, err := db.Prepare(sqlInsertUser)
	if err != nil {
		err = errors.New("CANNOT PREPARE STATMENT")
	}
	stmt.Exec(uid)

	log.Printf("Successfully deleted user: %s\n ", uid)
	return err

}
func RemoveUsers(ctx context.Context, uids []string) error {
	var app = database.InitializeAppWithServiceAccount()
	client, err := app.Auth(ctx)
	deleteUsersResult, err := client.DeleteUsers(ctx, uids)
	if err != nil {
		log.Fatalf("error deleting users: %v\n", err)
	}

	log.Printf("Successfully deleted %d users", deleteUsersResult.SuccessCount)
	log.Printf("Failed to delete %d users", deleteUsersResult.FailureCount)
	for _, err := range deleteUsersResult.Errors {
		log.Printf("%v", err)
	}
	return err

}
