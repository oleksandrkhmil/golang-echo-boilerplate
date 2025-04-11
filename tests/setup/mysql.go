package setup

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/modules/mysql"
)

const (
	mysqlImage         = "mysql:9.2.0"
	mysqlDatabase      = "db_name"
	mysqlUsername      = "username"
	mysqlPassword      = "password"
	mysqlPort          = "3306"
	mysqlContainerName = "app_db_mysql"
)

type DBConfig struct {
	ContainerName string
	User          string
	Password      string
	Host          string
	ExposedPort   string
	LocalPort     string
	Name          string
}

func SetupMYSQL(ctx context.Context) (DBConfig, func(ctx context.Context) error, error) {
	container, err := mysql.Run(
		ctx,
		mysqlImage,
		mysql.WithDatabase(mysqlDatabase),
		mysql.WithUsername(mysqlUsername),
		mysql.WithPassword(mysqlPassword),
		testcontainers.CustomizeRequest(testcontainers.GenericContainerRequest{
			ContainerRequest: testcontainers.ContainerRequest{
				Name: mysqlContainerName,
			},
		}),
	)
	if err != nil {
		return DBConfig{}, nil, fmt.Errorf("run mysql container: %w", err)
	}

	teardown := func(ctx context.Context) error {
		if err := container.Terminate(ctx); err != nil {
			return fmt.Errorf("terminate mysql container: %w", err)
		}

		return nil
	}

	port, err := container.MappedPort(ctx, mysqlPort+"/tcp")
	if err != nil {
		if errTeardown := teardown(ctx); errTeardown != nil {
			err = errors.Join(err, errTeardown)
		}

		return DBConfig{}, nil, fmt.Errorf("get mysql exposed port: %w", err)
	}

	config := DBConfig{
		ContainerName: mysqlContainerName,
		User:          mysqlUsername,
		Password:      mysqlPassword,
		Host:          "localhost",
		ExposedPort:   port.Port(),
		LocalPort:     mysqlPort,
		Name:          mysqlDatabase,
	}

	time.Sleep(30 * time.Second)

	return config, teardown, nil
}
