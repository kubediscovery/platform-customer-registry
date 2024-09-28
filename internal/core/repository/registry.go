package repository

import (
	"context"
	"database/sql"

	"github.com/kubediscovery/platform-customer-registry/internal/core/entity"
	"github.com/kubediscovery/platform-customer-registry/pkg/otelpkg"
	"github.com/kubediscovery/platform-customer-registry/scripts/db/dbsql"
)

type RegistryRepository struct {
	Db     *sql.DB
	Tracer *otelpkg.OtelPkgInstrument
}

// NewRegistryService creates a new RegistryService.
func NewRegistryRepository(db *sql.DB, otl *otelpkg.OtelPkgInstrument) (entity.CustomerRegistryInterface, error) {
	return &RegistryRepository{
		Db:     db,
		Tracer: otl,
	}, nil
}

func (rr *RegistryRepository) Create(ctx context.Context, cr *entity.CustomerRegistryResponse) (*entity.CustomerRegistryResponse, error) {

	queries := dbsql.New(rr.Db)
	result, err := queries.CreatePlatformRegistry(ctx, dbsql.CreatePlatformRegistryParams{
		ProjectName: cr.ProjectName,
		Repository:  cr.Repository,
		Username:    cr.UserName,
		Email:       cr.UserEmail,
		CreatedAt:   cr.CreateAt,
		DeletedAt:   cr.EndAt,
		ID:          cr.ID,
	})
	if err != nil {
		return nil, err
	}

	return &entity.CustomerRegistryResponse{
		ID:       result.ID,
		CreateAt: result.CreatedAt,
		EndAt:    result.DeletedAt,
		CustomerRegistry: entity.CustomerRegistry{
			ProjectName: result.ProjectName,
			Repository:  result.Repository,
			UserName:    result.Username,
			UserEmail:   result.Email,
		},
	}, nil
}

func (rr *RegistryRepository) GetAll(ctx context.Context) ([]entity.CustomerRegistryResponse, error) {

	queries := dbsql.New(rr.Db)
	result, err := queries.GetAllPlatformRegistry(ctx)
	if err != nil {
		return nil, err
	}

	var response []entity.CustomerRegistryResponse
	for _, r := range result {
		response = append(response, entity.CustomerRegistryResponse{
			ID:       r.ID,
			CreateAt: r.CreatedAt,
			EndAt:    r.DeletedAt,
			CustomerRegistry: entity.CustomerRegistry{
				ProjectName: r.ProjectName,
				Repository:  r.Repository,
				UserName:    r.Username,
				UserEmail:   r.Email,
			},
		})
	}
	return response, nil
}

func (rr *RegistryRepository) Get(cr *entity.CustomerRegistryResponse) (*entity.CustomerRegistryResponse, error) {
	return nil, nil
}

func (rr *RegistryRepository) GetByFilter(cr *entity.CustomerRegistry) ([]entity.CustomerRegistryResponse, error) {

	queries := dbsql.New(rr.Db)
	result, err := queries.GetByFilterPlatformRegistry(context.Background(), dbsql.GetByFilterPlatformRegistryParams{
		ProjectName:   cr.ProjectName,
		IsProjectName: (len(cr.ProjectName) > 0),
		Repository:    cr.Repository,
		IsRepository:  (len(cr.Repository) > 0),
		Username:      cr.UserName,
		IsUsername:    (len(cr.UserName) > 0),
		Email:         cr.UserEmail,
		IsEmail:       (len(cr.UserEmail) > 0),
	})
	if err != nil {
		return nil, err
	}

	var response []entity.CustomerRegistryResponse
	for _, r := range result {
		response = append(response, entity.CustomerRegistryResponse{
			ID:       r.ID,
			CreateAt: r.CreatedAt,
			EndAt:    r.DeletedAt,
			CustomerRegistry: entity.CustomerRegistry{
				ProjectName: r.ProjectName,
				Repository:  r.Repository,
				UserName:    r.Username,
				UserEmail:   r.Email,
			},
		})
	}

	return response, nil
}

func (rr *RegistryRepository) GetByID(ctx context.Context, id *string) (*entity.CustomerRegistryResponse, error) {

	queries := dbsql.New(rr.Db)
	result, err := queries.GetByIDPlatformRegistry(ctx, *id)
	if err != nil {
		return nil, err
	}

	return &entity.CustomerRegistryResponse{
		ID:       result.ID,
		CreateAt: result.CreatedAt,
		EndAt:    result.DeletedAt,
		CustomerRegistry: entity.CustomerRegistry{
			ProjectName: result.ProjectName,
			Repository:  result.Repository,
			UserName:    result.Username,
			UserEmail:   result.Email,
		},
	}, nil

}
