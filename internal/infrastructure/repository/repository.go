package repository

import (
	"clean-architecture/internal/domain/repository"
	"context"
	"errors"
	"fmt"

	"github.com/go-sql-driver/mysql"
	"gorm.io/gorm"
)

var (
	ErrInvalidArgument    = errors.New("database: invalid argument")
	ErrNotFound           = errors.New("database: not found")
	ErrNotAcceptable      = errors.New("database: not acceptable")
	ErrAlreadyExists      = errors.New("database: already exists")
	ErrFailedPrecondition = errors.New("database: failed precondition")
	ErrNotImplemented     = errors.New("database: not implemented")
	ErrInternal           = errors.New("database: internal error")
	ErrCanceled           = errors.New("database: canceled")
	ErrDeadlineExceeded   = errors.New("database: deadline exceeded")
	ErrUnknown            = errors.New("database: unknown")
)

type Repository struct {
	User repository.UserRepository
}

func NewRepository(db *gorm.DB) *repository.Repository {
	return &repository.Repository{
		User: NewUserRepository(db),
	}
}

// Transaction runs the given function within a database transaction and
// provides a transactional repository instance to the callback.
func Transaction(ctx context.Context, db *gorm.DB, fn func(r *repository.Repository) error) error {
	if db == nil {
		return fmt.Errorf("%w: %s", ErrInternal, "nil db client")
	}
	if fn == nil {
		return fmt.Errorf("%w: %s", ErrInvalidArgument, "nil transaction function")
	}

	err := db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		return fn(NewRepository(tx))
	})
	return dbError(err)
}

func dbError(err error) error {
	if err == nil {
		return nil
	}

	switch err := err.(type) {
	case *mysql.MySQLError:
		if err.Number == 1062 {
			return fmt.Errorf("%w: %s", ErrAlreadyExists, err)
		}
		return fmt.Errorf("%w: %s", ErrInternal, err)
	}

	switch {
	case errors.Is(err, ErrInvalidArgument),
		errors.Is(err, ErrNotFound),
		errors.Is(err, ErrAlreadyExists),
		errors.Is(err, ErrFailedPrecondition),
		errors.Is(err, ErrNotImplemented),
		errors.Is(err, ErrInternal),
		errors.Is(err, ErrCanceled),
		errors.Is(err, ErrDeadlineExceeded),
		errors.Is(err, ErrUnknown):
		return err
	case errors.Is(err, context.Canceled):
		return fmt.Errorf("%w: %s", ErrCanceled, err)
	case errors.Is(err, context.DeadlineExceeded):
		return fmt.Errorf("%w: %s", ErrDeadlineExceeded, err)
	case errors.Is(err, gorm.ErrEmptySlice),
		errors.Is(err, gorm.ErrInvalidData),
		errors.Is(err, gorm.ErrInvalidField),
		errors.Is(err, gorm.ErrInvalidTransaction),
		errors.Is(err, gorm.ErrInvalidValue),
		errors.Is(err, gorm.ErrInvalidValueOfLength),
		errors.Is(err, gorm.ErrMissingWhereClause),
		errors.Is(err, gorm.ErrModelValueRequired),
		errors.Is(err, gorm.ErrPrimaryKeyRequired):
		return fmt.Errorf("%w: %s", ErrInvalidArgument, err)
	case errors.Is(err, gorm.ErrRecordNotFound):
		return fmt.Errorf("%w: %s", ErrNotFound, err)
	case errors.Is(err, gorm.ErrNotImplemented):
		return fmt.Errorf("%w: %s", ErrNotImplemented, err)
	case errors.Is(err, gorm.ErrDryRunModeUnsupported),
		errors.Is(err, gorm.ErrInvalidDB),
		errors.Is(err, gorm.ErrRegistered),
		errors.Is(err, gorm.ErrUnsupportedDriver),
		errors.Is(err, gorm.ErrUnsupportedRelation):
		return fmt.Errorf("%w: %s", ErrInternal, err)
	default:
		return fmt.Errorf("%w: %s", ErrUnknown, err)
	}
}
