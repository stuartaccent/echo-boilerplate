package cmd

import (
	"context"
	"echo.go.dev/pkg/domain/auth"
	"echo.go.dev/pkg/storage/db"
	"echo.go.dev/pkg/storage/db/dbx"
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

var (
	setPWEmail, setPWPassword string
)

var cmdSetPassword = &cobra.Command{
	Use:   "setpassword",
	Short: "Set a user's password",
	Run: func(cmd *cobra.Command, args []string) {
		ctx := context.Background()

		conn, err := db.GetConn(ctx)
		if err != nil {
			fmt.Printf("Error connecting to the database: %v\n", err)
			os.Exit(1)
		}
		defer conn.Close(ctx)

		hashed, err := auth.GeneratePassword([]byte(setPWPassword))
		if err != nil {
			fmt.Printf("Error: %v\n", err)
			os.Exit(1)
		}

		queries := dbx.New(conn)
		if err = queries.SetUserPasswordByEmail(ctx, dbx.SetUserPasswordByEmailParams{
			Email:          setPWEmail,
			HashedPassword: hashed,
		}); err != nil {
			fmt.Printf("Error: %v\n", err)
			os.Exit(1)
		}

		fmt.Printf("Password set for user: %v\n", setPWEmail)
	},
}

func init() {
	cmdSetPassword.Flags().StringVarP(&setPWEmail, "email", "e", "", "The email address of the user")
	cmdSetPassword.Flags().StringVarP(&setPWPassword, "password", "p", "", "The password of the user")
	_ = cmdSetPassword.MarkFlagRequired("email")
	_ = cmdSetPassword.MarkFlagRequired("password")
}
