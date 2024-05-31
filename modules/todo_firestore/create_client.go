package todo_firestore

import (
	"context"
	"errors"
	"fmt"

	firestore "cloud.google.com/go/firestore"
	firebase "firebase.google.com/go"
	option "google.golang.org/api/option"
)

func CreateFirestoreClient(ctx context.Context, serviceAccountKeyPath string) (*firestore.Client, error) {
	app, err := initialiseFirebaseApp(ctx, serviceAccountKeyPath)
	if err != nil {
		errorMessage := fmt.Sprintf("Error initialising Firebase app: %v", err)
		return nil, errors.New(errorMessage)
	}

	client, err := app.Firestore(ctx)
	if err != nil {
		errorMessage := fmt.Sprintf("Error initialising Firestore client: %v", err)
		return nil, errors.New(errorMessage)
	}

	return client, nil
}

func initialiseFirebaseApp(ctx context.Context, serviceAccountKeyPath string) (*firebase.App, error) {
	opt := option.WithCredentialsFile(serviceAccountKeyPath)
	app, err := firebase.NewApp(ctx, nil, opt)
	if err != nil {
		return nil, err
	}

	return app, nil
}
