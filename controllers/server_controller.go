package controllers

import (
	"context"
	"database/sql"

	"github.com/jeanbenitez/servercheck/interfaces"
	"github.com/jeanbenitez/servercheck/models"
)

// NewControllerServer returns server interface implementation
func NewControllerServer(Conn *sql.DB) interfaces.IServerController {
	return &controllerServer{
		Conn: Conn,
	}
}

type controllerServer struct {
	Conn *sql.DB
}

func (m *controllerServer) fetch(ctx context.Context, query string, args ...interface{}) ([]*models.Server, error) {
	rows, err := m.Conn.QueryContext(ctx, query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	payload := make([]*models.Server, 0)
	for rows.Next() {
		data := new(models.Server)

		err := rows.Scan(
			&data.Address,
			&data.SslGrade,
			&data.Country,
			&data.Owner,
		)
		if err != nil {
			return nil, err
		}
		payload = append(payload, data)
	}
	return payload, nil
}

func (m *controllerServer) FetchByDomain(ctx context.Context, domain string) ([]*models.Server, error) {
	query := "select address, ssl_grade, country, owner from servers where domain = $1"
	return m.fetch(ctx, query, domain)
}

func (m *controllerServer) Create(ctx context.Context, d *models.Server) (bool, error) {
	query := "insert servers SET address=?, ssl_grade=?, country=?, owner=?"
	stmt, err := m.Conn.PrepareContext(ctx, query)
	if err != nil {
		return false, err
	}

	_, err2 := stmt.ExecContext(
		ctx,
		d.Address,
		d.SslGrade,
		d.Country,
		d.Owner,
	)
	defer stmt.Close()

	if err2 != nil {
		return false, err2
	}

	return true, nil
}

func (m *controllerServer) Delete(ctx context.Context, domain string) (bool, error) {
	query := "delete from servers where domain=?"

	stmt, err := m.Conn.PrepareContext(ctx, query)
	if err != nil {
		return false, err
	}
	_, err = stmt.ExecContext(ctx, domain)
	if err != nil {
		return false, err
	}
	return true, nil
}
