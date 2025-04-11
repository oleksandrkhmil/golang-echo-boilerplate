package integration_test

import (
	"context"
	"log/slog"
	"testing"

	"github.com/nix-united/golang-echo-boilerplate/internal/config"
	"github.com/nix-united/golang-echo-boilerplate/internal/db"
	"github.com/nix-united/golang-echo-boilerplate/tests/setup"

	"github.com/stretchr/testify/require"
)

func TestAuth(t *testing.T) {
	ctx := context.Background()

	dbConfig, teardown, err := setup.SetupMYSQL(ctx)
	require.NoError(t, err)
	defer func() {
		err := teardown(ctx)
		require.NoError(t, err)
	}()

	slog.InfoContext(ctx, "conn str", "_", dbConfig)

	gormDB, err := db.NewGormDB(config.DBConfig{
		User:     dbConfig.User,
		Password: dbConfig.Password,
		Name:     dbConfig.Name,
		Host:     dbConfig.Host,
		Port:     dbConfig.ExposedPort,
	})
	require.NoError(t, err)

	// cases := []helpers.TestCase{
	// 	{
	// 		"Auth success",
	// 		request,
	// 		requests.LoginRequest{
	// 			BasicAuth: requests.BasicAuth{
	// 				Email:    "name@test.com",
	// 				Password: "password",
	// 			},
	// 		},
	// 		handlerFunc,
	// 		[]*helpers.QueryMock{&helpers.SelectVersionMock, commonMock},
	// 		helpers.ExpectedResponse{
	// 			StatusCode: 200,
	// 			BodyPart:   "",
	// 		},
	// 	},
	// 	// {
	// 	// 	"Login attempt with incorrect password",
	// 	// 	request,
	// 	// 	requests.LoginRequest{
	// 	// 		BasicAuth: requests.BasicAuth{
	// 	// 			Email:    "name@test.com",
	// 	// 			Password: "incorrectPassword",
	// 	// 		},
	// 	// 	},
	// 	// 	handlerFunc,
	// 	// 	[]*helpers.QueryMock{&helpers.SelectVersionMock, commonMock},
	// 	// 	helpers.ExpectedResponse{
	// 	// 		StatusCode: 401,
	// 	// 		BodyPart:   "Invalid credentials",
	// 	// 	},
	// 	// },
	// 	// {
	// 	// 	"Login attempt as non-existent user",
	// 	// 	request,
	// 	// 	requests.LoginRequest{
	// 	// 		BasicAuth: requests.BasicAuth{
	// 	// 			Email:    "user.not.exists@test.com",
	// 	// 			Password: "password",
	// 	// 		},
	// 	// 	},
	// 	// 	handlerFunc,
	// 	// 	[]*helpers.QueryMock{&helpers.SelectVersionMock, commonMock},
	// 	// 	helpers.ExpectedResponse{
	// 	// 		StatusCode: 401,
	// 	// 		BodyPart:   "Invalid credentials",
	// 	// 	},
	// 	// },
	// }

	// for _, test := range cases {
	// 	t.Run(test.TestName, func(t *testing.T) {
	// 		helpers.PrepareDatabaseQueryMocks(test, sqlMock)
	// 		db := helpers.InitGorm(dbMock)
	// 		s := helpers.NewServer(db)

	// 		c, recorder := helpers.PrepareContextFromTestCase(s, test)

	// 		if assert.NoError(t, test.HandlerFunc(s, c)) {
	// 			assert.Contains(t, recorder.Body.String(), test.Expected.BodyPart)
	// 			if assert.Equal(t, test.Expected.StatusCode, recorder.Code) {
	// 				if recorder.Code == http.StatusOK {
	// 					assertTokenResponse(t, recorder)
	// 				}
	// 			}
	// 		}
	// 	})
	// }
}
