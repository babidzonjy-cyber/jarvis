package repo

import (
	"context"
	"errors"
	"jarvis/internal/domain"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type DaemonStateRepo interface {
	SetMode(ctx context.Context, mode string) error
	Heartbeat(ctx context.Context) error
	GetState(ctx context.Context) (domain.DaemonState, error)
}

type daemonStatePG struct {
	pool *pgxpool.Pool
}

func NewDaemonStatePG(pool *pgxpool.Pool) *daemonStatePG {
	return &daemonStatePG{pool: pool}
}

func (d *daemonStatePG) SetMode(ctx context.Context, mode string) error {
	query := `
		UPDATE jarvis.daemon_state SET current_mode = $1, updated_at = now() WHERE id = 1;
	`

	_, err := d.pool.Exec(
		ctx, query, mode,
	)

	return err
}

func (d *daemonStatePG) Heartbeat(ctx context.Context) error {
	query := `
		UPDATE jarvis.daemon_state SET updated_at = NOW() WHERE id = 1;
	`

	_, err := d.pool.Exec(
		ctx, query,
	)

	return err
}

func (d *daemonStatePG) GetState(ctx context.Context) (domain.DaemonState, error) {
	query := `SELECT current_mode, updated_at from jarvis.daemon_state WHERE id = 1;`

	var daemonState domain.DaemonState

	if err := d.pool.QueryRow(
		ctx,
		query,
	).Scan(
		&daemonState.Mode,
		&daemonState.UpdatedAt,
	); err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return domain.DaemonState{}, errors.New("not found")
		}
		return domain.DaemonState{}, err
	}

	return daemonState, nil
}
