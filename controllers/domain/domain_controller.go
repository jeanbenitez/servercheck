package domain

import (
	"context"
	"database/sql"

	interfaces "github.com/jeanbenitez/servercheck/interfaces"
	models "github.com/jeanbenitez/servercheck/models"
)

// NewSQLDomain returns domain interface implementation
func NewSQLDomain(Conn *sql.DB) interfaces.IDomainController {
	return &mysqlDomain{
		Conn: Conn,
	}
}

type mysqlDomain struct {
	Conn *sql.DB
}

func (m *mysqlDomain) fetch(ctx context.Context, query string, args ...interface{}) ([]*models.Domain, error) {
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
			&data.IsDown,
		)
		if err != nil {
			return nil, err
		}
		payload = append(payload, data)
	}
	return payload, nil
}

func (m *mysqlDomain) Fetch(ctx context.Context, num int64) ([]*models.Domain, error) {
	query := "select id, title, content From domains limit ?"

	return m.fetch(ctx, query, num)
}

func (m *mysqlDomain) GetByDomain(ctx context.Context, domain string) (*models.Domain, error) {
	query := "select * from domains where domain=?"

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

func (m *mysqlDomain) Create(ctx context.Context, d *models.Domain) (bool, error) {
	query := "insert domains SET domain=?, servers_changed=?, ssl_grade=?, 	previous_ssl_grade=?, logo=?, is_down=?"

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
		d.IsDown,
	)
	defer stmt.Close()

	if err2 != nil {
		return false, err2
	}

	return true, nil
}

func (m *mysqlDomain) Update(ctx context.Context, d *models.Domain) (*models.Domain, error) {
	query := "Update domains set servers_changed=?, ssl_grade=?, 	previous_ssl_grade=?, logo=?, is_down=? where domain=?"

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
		d.IsDown,
		d.Domain,
	)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	return d, nil
}

func (m *mysqlDomain) Delete(ctx context.Context, domain string) (bool, error) {
	query := "delete From domains Where domain=?"

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
