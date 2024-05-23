package storage

import (
	"context"

	"github.com/diplom-pam/edu/internal/domain/models"
)

func (s *Storage) CreateTest(ctx context.Context, test models.Test) (int, error) {
	row := s.pg.QueryRow(ctx, `
		INSERT INTO tests (questions, title)
		VALUES ($1, $2, $3)
		RETURNING id
	`, test.Questions, test.Title)

	var id int
	err := row.Scan(&id)
	if err != nil {
		return 0, err
	}

	return id, nil
}
