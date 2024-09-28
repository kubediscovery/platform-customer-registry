package repository

import (
	"context"
	"database/sql"
	"errors"
	"time"

	"github.com/kubediscovery/platform-customer-registry/internal/core/entity"
	"github.com/kubediscovery/platform-customer-registry/pkg/otelpkg"
	"github.com/kubediscovery/platform-customer-registry/scripts/db/dbsql"
)

type LabDestroyRepository struct {
	Db     *sql.DB
	Tracer *otelpkg.OtelPkgInstrument
}

func NewLabDestroyRepository(db *sql.DB, otel *otelpkg.OtelPkgInstrument) (entity.LabDestroyInterface, error) {

	if db == nil {
		return nil, errors.New("db is required")
	}

	ld := &LabDestroyRepository{
		Db:     db,
		Tracer: otel,
	}

	return ld, nil
}

func (ldr *LabDestroyRepository) Create(ctx context.Context, ld *entity.LabDestroyResponse) (*entity.LabDestroyResponse, error) {
	query := dbsql.New(ldr.Db)
	result, err := query.CreatePlatformLabDestroy(ctx, dbsql.CreatePlatformLabDestroyParams{
		ID:            ld.ID,
		ProjectName:   ld.LabDestroy.CustomerRegistry.ProjectName,
		Repository:    ld.LabDestroy.CustomerRegistry.Repository,
		Username:      ld.LabDestroy.CustomerRegistry.UserName,
		Email:         ld.LabDestroy.CustomerRegistry.UserEmail,
		TargetDestroy: ld.LabDestroy.TargetDestroy,
		CreatedAt:     ld.LabDestroy.CreatedAt,
		Available:     *ld.Avaliable,
	})
	if err != nil {
		return nil, err
	}

	return &entity.LabDestroyResponse{
		ID: result.ID,
		LabDestroy: entity.LabDestroy{
			CustomerRegistry: entity.CustomerRegistry{
				ProjectName: result.ProjectName,
				Repository:  result.Repository,
				UserName:    result.Username,
				UserEmail:   result.Email,
			},
			TargetDestroy: result.TargetDestroy,
			CreatedAt:     result.CreatedAt,
		},
		Avaliable: &result.Available,
	}, nil
}
func (ldr *LabDestroyRepository) GetAll(ctx context.Context) ([]entity.LabDestroyResponse, error) {
	query := dbsql.New(ldr.Db)
	result, err := query.GetAllPlatformLabDestroy(ctx)
	if err != nil {
		return nil, err
	}

	var lds []entity.LabDestroyResponse
	for _, r := range result {
		lds = append(lds, entity.LabDestroyResponse{
			ID:        r.ID,
			Avaliable: &r.Available,
			Error:     r.ErrorMessage.String,
			UpdatedAt: r.UpdatedAt.String,
			LabDestroy: entity.LabDestroy{
				CustomerRegistry: entity.CustomerRegistry{
					ProjectName: r.ProjectName,
					Repository:  r.Repository,
					UserName:    r.Username,
					UserEmail:   r.Email,
				},
				TargetDestroy: r.TargetDestroy,
				CreatedAt:     r.CreatedAt,
			},
		})
	}
	return lds, nil
}

func (ldr *LabDestroyRepository) GetByID(ctx context.Context, ld *string) (*entity.LabDestroyResponse, error) {
	query := dbsql.New(ldr.Db)
	result, err := query.GetByIDPlatformLabDestroy(ctx, *ld)
	if err != nil {
		return nil, err
	}

	return &entity.LabDestroyResponse{
		ID:        result.ID,
		Avaliable: &result.Available,
		Error:     result.ErrorMessage.String,
		UpdatedAt: result.UpdatedAt.String,
		LabDestroy: entity.LabDestroy{
			CustomerRegistry: entity.CustomerRegistry{
				ProjectName: result.ProjectName,
				Repository:  result.Repository,
				UserName:    result.Username,
				UserEmail:   result.Email,
			},
			TargetDestroy: result.TargetDestroy,
			CreatedAt:     result.CreatedAt,
		},
	}, nil
}
func (ldr *LabDestroyRepository) GetByFilter(ctx context.Context, ld *entity.LabDestroyResponse) ([]entity.LabDestroyResponse, error) {

	if ld == nil {
		return nil, errors.New("filter cannot be empty")
	}

	available := false
	if ld.Avaliable != nil {
		available = *ld.Avaliable
	}

	query := dbsql.New(ldr.Db)
	result, err := query.GetByFilterPlatformLabDestroy(ctx, dbsql.GetByFilterPlatformLabDestroyParams{
		Email:         ld.CustomerRegistry.UserEmail,
		IsEmail:       ld.CustomerRegistry.UserEmail != "",
		ProjectName:   ld.CustomerRegistry.ProjectName,
		IsProjectName: ld.CustomerRegistry.ProjectName != "",
		IsRepository:  ld.CustomerRegistry.Repository != "",
		Repository:    ld.CustomerRegistry.Repository,
		IsUsername:    ld.CustomerRegistry.UserName != "",
		Username:      ld.CustomerRegistry.UserName,
		IsAvailable:   ld.Avaliable != nil,
		Available:     available,
	})

	if err != nil {
		return nil, err
	}

	var results []entity.LabDestroyResponse
	for _, r := range result {
		results = append(results, entity.LabDestroyResponse{
			ID:        r.ID,
			Avaliable: &r.Available,
			Error:     r.ErrorMessage.String,
			UpdatedAt: r.UpdatedAt.String,
			LabDestroy: entity.LabDestroy{
				TargetDestroy: r.TargetDestroy,
				CreatedAt:     r.CreatedAt,
				CustomerRegistry: entity.CustomerRegistry{
					ProjectName: r.ProjectName,
					Repository:  r.Repository,
					UserName:    r.Username,
					UserEmail:   r.Email,
				},
			},
		})
	}

	return results, nil
}
func (ldr *LabDestroyRepository) PATCH(ctx context.Context, id *string, err_msg *string, available *bool) (*entity.LabDestroyResponse, error) {
	query := dbsql.New(ldr.Db)

	result, err := query.PatchlatformLabDestroy(ctx, dbsql.PatchlatformLabDestroyParams{
		ID:           *id,
		ErrorMessage: *err_msg,
		Available:    *available,
		UpdatedAt:    sql.NullString{String: time.Now().Format(time.RFC3339), Valid: true},
	})
	if err != nil {
		return nil, err
	}

	return &entity.LabDestroyResponse{
		ID:        result.ID,
		Avaliable: &result.Available,
		Error:     result.ErrorMessage.String,
		UpdatedAt: result.UpdatedAt.String,
		LabDestroy: entity.LabDestroy{
			TargetDestroy: result.TargetDestroy,
			CreatedAt:     result.CreatedAt,
			CustomerRegistry: entity.CustomerRegistry{
				ProjectName: result.ProjectName,
				Repository:  result.Repository,
				UserName:    result.Username,
				UserEmail:   result.Email,
			},
		},
	}, nil
}
