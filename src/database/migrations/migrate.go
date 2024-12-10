package migrations

import (
	"DB_Project/src/database"
	"DB_Project/src/utils"
)

func MigrateUp() error {
	// Read all SQL files in the "migrations" directory
	files, err := utils.ReadSQLFiles("src/database/migrations/up")
	if err != nil {
		return err
	}

	// Execute each migration file
	for i := 0; i < len(files); i++ {
		_, err = database.RunFileQuery(files[i])
		if err != nil {
			return err
		}
	}
	return nil
}

func MigrateDown() error {
	// Read all SQL files in the "migrations" directory
	files, err := utils.ReadSQLFiles("src/database/migrations/down")
	if err != nil {
		return err
	}

	// Execute each migration file
	for i := len(files) - 1; i >= 0; i-- {
		_, err = database.RunFileQuery(files[i])
		if err != nil {
			return err
		}
	}
	return nil
}
