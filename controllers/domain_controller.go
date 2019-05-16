package controllers

import (
	"context"
	"database/sql"

	"github.com/jeanbenitez/servercheck/interfaces"
	"github.com/jeanbenitez/servercheck/models"
)

// NewControllerDomain returns domain interface implementation
func NewControllerDomain(Conn *sql.DB) interfaces.IDomainController {
	return &controllerDomain{
		Conn: Conn,
	}
}

type controllerDomain struct {
	Conn *sql.DB
}

func (m *controllerDomain) fetch(ctx context.Context, query string, args ...interface{}) ([]*models.Domain, error) {
	rows, err := m.Conn.QueryContext(ctx, query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	payload := make([]*models.Domain, 0)
	for rows.Next() {
		data := new(models.Domain)

		err := rows.Scan(
			&data.Domain,
			&data.ServersChanged,
			&data.SslGrade,
			&data.PreviousSslGrade,
			&data.Logo,
			&data.Title,
			&data.IsDown,
		)
		if err != nil {
			return nil, err
		}
		payload = append(payload, data)
	}
	return payload, nil
}

func (m *controllerDomain) Fetch(ctx context.Context, num int64) ([]*models.Domain, error) {
	query := "select domain, servers_changed, ssl_grade, previous_ssl_grade, logo, title, is_down from domains limit $1"

	return m.fetch(ctx, query, num)
}

func (m *controllerDomain) GetByDomain(ctx context.Context, domain string) (*models.Domain, error) {
	query := "select domain, servers_changed, ssl_grade, previous_ssl_grade, logo, title, is_down from domains where domain = $1"

	rows, err := m.fetch(ctx, query, domain)
	if err != nil {
		return nil, err
	}

	payload := &models.Domain{}
	if len(rows) > 0 {
		payload = rows[0]
	} else {
		return nil, models.ErrNotFound
	}

	return payload, nil
}

func (m *controllerDomain) Create(ctx context.Context, d *models.Domain) (bool, error) {
	query := "insert into domains (domain, servers_changed, ssl_grade, previous_ssl_grade, logo, title, is_down) VALUES ($1, $2, $3, $4, $5, $6, $7)"

	stmt, err := m.Conn.PrepareContext(ctx, query)
	if err != nil {
		return false, err
	}

	_, err2 := stmt.ExecContext(
		ctx,
		d.Domain,
		d.ServersChanged,
		d.SslGrade,
		d.PreviousSslGrade,
		d.Logo,
		d.Title,
		d.IsDown,
	)
	defer stmt.Close()

	if err2 != nil {
		return false, err2
	}

	return true, nil
}

func (m *controllerDomain) Update(ctx context.Context, d *models.Domain) (*models.Domain, error) {
	query := "update domains set servers_changed=$1, ssl_grade=$2, previous_ssl_grade=$3, logo=$4, title=$5, is_down=$6 where domain=$7"

	stmt, err := m.Conn.PrepareContext(ctx, query)
	if err != nil {
		return nil, err
	}
	_, err = stmt.ExecContext(
		ctx,
		d.ServersChanged,
		d.SslGrade,
		d.PreviousSslGrade,
		d.Logo,
		d.Title,
		d.IsDown,
		d.Domain,
	)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	return d, nil
}

func (m *controllerDomain) Delete(ctx context.Context, domain string) (bool, error) {
	query := "delete from domains where domain=$1"

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
