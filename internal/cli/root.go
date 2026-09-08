package cli

import (
	"context"
	"fmt"
	repo "jarvis/internal/repository"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/spf13/cobra"
)

var (
	daemonStateRepo repo.DaemonStateRepo
	dbPool          *pgxpool.Pool
)

var rootCmd = &cobra.Command{
	Use:   "jarvis",
	Short: "The AI voice assistant",
	Long:  "The AI voice assistant that runs local on youy PC",
	PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
		postgresHost := os.Getenv("POSTGRES_HOST")
		if postgresHost == "" {
			postgresHost = "localhost"
		}

		db := fmt.Sprintf("postgres://%s:%s@%s:5432/%s?sslmode=disable",
			os.Getenv("POSTGRES_USER"),
			os.Getenv("POSTGRES_PASSWORD"),
			postgresHost,
			os.Getenv("POSTGRES_DB"),
		)

		pool, err := repo.NewPool(context.Background(), db)
		if err != nil {
			return fmt.Errorf("cannot connect to database: %w", err)
		}

		dbPool = pool
		daemonStateRepo = repo.NewDaemonStatePG(pool)
		return nil
	},
	PersistentPostRunE: func(cmd *cobra.Command, args []string) error {
		if dbPool != nil {
			dbPool.Close()
		}
		return nil
	},
}

func Execute() error {
	return rootCmd.Execute()
}
