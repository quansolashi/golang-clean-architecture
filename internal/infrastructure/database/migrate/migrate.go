package migrate

import (
	database "clean-architecture/internal/infrastructure/database/gorm"
	"clean-architecture/internal/infrastructure/model"
	"context"
	"fmt"
)

var Tables = map[string]interface{}{
	"users": model.User{},
}

func Run(ctx context.Context, db *database.Client) error {
	for _, table := range Tables {
		if err := db.DB.WithContext(ctx).AutoMigrate(table); err != nil {
			fmt.Println("Failed to migrate table", err)
			return err
		}
	}
	return nil
}
