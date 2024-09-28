package service

import (
	"context"
	"errors"

	"github.com/kubediscovery/platform-customer-registry/internal/core/entity"
	"github.com/kubediscovery/platform-customer-registry/pkg/cache"
	"github.com/kubediscovery/platform-customer-registry/pkg/kd_utils"
	"github.com/kubediscovery/platform-customer-registry/pkg/kd_uuid"
	"github.com/kubediscovery/platform-customer-registry/pkg/otelpkg"
)

type LabDestroyService struct {
	Repository entity.LabDestroyInterface
	Cache      cache.CacheInterface
	Tracer     *otelpkg.OtelPkgInstrument
}

func NewServiceLabDestroy(repo *entity.LabDestroyInterface, cc *cache.CacheInterface, otel *otelpkg.OtelPkgInstrument) (entity.LabDestroyInterface, error) {
	if repo == nil {
		return nil, errors.New("repository is required")
	}

	if cc == nil {
		return nil, errors.New("cache is required")
	}

	return LabDestroyService{
		Repository: *repo,
		Cache:      *cc,
		Tracer:     otel,
	}, nil
}

func (lds LabDestroyService) Create(ctx context.Context, ld *entity.LabDestroyResponse) (*entity.LabDestroyResponse, error) {
	_, span := lds.Tracer.Tracer.Start(ctx, "LabDestroyService.Create")
	defer span.End()

	err := ld.Validate()
	if err != nil {
		span.RecordError(err)
		return nil, err
	}

	ld, err = lds.Repository.Create(ctx, ld)
	if err != nil {
		span.RecordError(err)
		return nil, err
	}

	return ld, nil
}
func (lds LabDestroyService) GetAll(ctx context.Context) ([]entity.LabDestroyResponse, error) {

	result, err := lds.Repository.GetAll(ctx)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (lds LabDestroyService) GetByID(ctx context.Context, ld *string) (*entity.LabDestroyResponse, error) {

	uuidIsValid := kd_uuid.IsValid(ld)
	emailIsValid := kd_utils.IsValidEmail(ld)
	if !uuidIsValid && !emailIsValid {
		return nil, errors.New("id or email are invalid")
	}

	var err error
	result := &entity.LabDestroyResponse{}
	if uuidIsValid {
		result, err = lds.Repository.GetByID(ctx, ld)
		if err != nil {
			return nil, err
		}
	}

	return result, err
}
func (lds LabDestroyService) GetByFilter(ctx context.Context, ld *entity.LabDestroyResponse) ([]entity.LabDestroyResponse, error) {

	if ld == nil {
		return nil, errors.New("filter cannot be empty")
	}

	if ld.CustomerRegistry.ProjectName == "" && ld.CustomerRegistry.UserEmail == "" && ld.CustomerRegistry.UserName == "" && ld.Avaliable == nil {
		return nil, errors.New("filter cannot be empty")
	}

	if ld.CustomerRegistry.UserEmail != "" {
		if !kd_utils.IsValidEmail(&ld.CustomerRegistry.UserEmail) {
			return nil, errors.New("email is invalid")
		}
	}

	result, err := lds.Repository.GetByFilter(ctx, ld)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (lds LabDestroyService) PATCH(ctx context.Context, id *string, err_msg *string, available *bool) (*entity.LabDestroyResponse, error) {

	if id == nil || available == nil {
		return nil, errors.New("id and available are required")
	}

	if !kd_uuid.IsValid(id) {
		return nil, errors.New("id is invalid")
	}

	result, err := lds.Repository.PATCH(ctx, id, err_msg, available)
	if err != nil {
		return nil, err
	}
	return result, nil
}
