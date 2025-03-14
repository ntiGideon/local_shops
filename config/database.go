package config

import (
	"baseProject/ent"
	"context"
	"fmt"
	_ "github.com/lib/pq"
)

func ConnectDb(connectionString string) (*ent.Client, error) {
	client, err := ent.Open("postgres", connectionString)

	if err != nil {
		return nil, err
	}
	defer func(client *ent.Client) {
		err := client.Close()
		if err != nil {
			return
		}
	}(client)
	ctx := context.Background()
	err = client.Schema.Create(ctx)
	if err != nil {
		return nil, err
	}
	fmt.Println("successfully connected to the database!")
	return client, err
}
