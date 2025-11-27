package db

import (
	"database/sql"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"sort"
	"strings"
)

const defaultMigrationsDir = "database/migrations"

// Migrate применяет все *.up.sql миграции из папки database/migrations.
// Миграции выполняются в алфавитном порядке имён файлов.
func Migrate(db *sql.DB) error {
	return migrateFromDir(db, defaultMigrationsDir)
}

func migrateFromDir(db *sql.DB, dir string) error {
	var files []string

	err := filepath.WalkDir(dir, func(path string, d fs.DirEntry, walkErr error) error {
		if walkErr != nil {
			return walkErr
		}

		if d.IsDir() {
			return nil
		}

		// Берём только *.up.sql
		if !strings.HasSuffix(path, ".up.sql") {
			return nil
		}

		files = append(files, path)
		return nil
	})
	if err != nil {
		return fmt.Errorf("scan migrations dir: %w", err)
	}

	if len(files) == 0 {
		return nil
	}

	sort.Strings(files)

	for _, filePath := range files {
		sqlBytes, readErr := os.ReadFile(filePath)
		if readErr != nil {
			return fmt.Errorf("read migration %s: %w", filePath, readErr)
		}

		sqlText := strings.TrimSpace(string(sqlBytes))
		if sqlText == "" {
			continue
		}

		fmt.Printf("Executing migration: %s\n", filePath)
		if _, execErr := db.Exec(sqlText); execErr != nil {
			return fmt.Errorf("execute migration %s: %w", filePath, execErr)
		}
		fmt.Printf("Migration %s completed successfully\n", filePath)
	}

	return nil
}


