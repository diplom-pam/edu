package storage

import (
	"context"

	"github.com/diplom-pam/edu/internal/domain/models"
)

func (s *Storage) CreateQuestion(ctx context.Context, question models.Question) (id int, err error) {
	row := s.pg.QueryRow(ctx, `
		INSERT INTO questions (question, body, content)
		VALUES ($1, $2, $3)
		RETURNING id
	`, question.Question, question.Body, question.Content)

	err = row.Scan(&id)
	if err != nil {
		return 0, err
	}

	return id, nil
}

func (s *Storage) GetQuestion(ctx context.Context, questionID int) (*models.Question, error) {
	row := s.pg.QueryRow(ctx, `
		SELECT id, question, body, content
		FROM questions
		WHERE id = $1
	`, questionID)

	q := models.Question{}
	err := row.Scan(&q.ID, &q.Question, &q.Body, &q.Content)
	if err != nil {
		return nil, err
	}

	return &q, nil
}
