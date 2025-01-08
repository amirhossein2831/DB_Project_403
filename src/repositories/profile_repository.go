package repositories

import (
	"DB_Project/src/database/connection/pgx"
	"context"
	"fmt"
)

type ProfileRepository struct {
}

func NewProfileRepository() *ProfileRepository {
	return &ProfileRepository{}
}

func (repository *ProfileRepository) UpdateField(name, id string, value interface{}) error {
	_, err := pgx.GetInstance().Exec(context.Background(), fmt.Sprintf("UPDATE profile SET %s = $1 WHERE id = $2", name), value, id)
	return err
}

func (repository *ProfileRepository) Delete(id int) error {
	_, err := pgx.GetInstance().Exec(context.Background(), "DELETE FROM profile WHERE id = $1", id)
	return err
}
