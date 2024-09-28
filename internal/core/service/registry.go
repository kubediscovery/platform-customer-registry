package service

import (
	"context"
	"errors"
	"sort"
	"sync"
	"time"

	"github.com/kubediscovery/platform-customer-registry/pkg/cache"
	"github.com/kubediscovery/platform-customer-registry/pkg/kd_utils"
	"github.com/kubediscovery/platform-customer-registry/pkg/kd_uuid"
	"github.com/kubediscovery/platform-customer-registry/pkg/otelpkg"

	"github.com/kubediscovery/platform-customer-registry/internal/core/entity"
)

type RegistryServiceInterface interface {
	entity.CustomerRegistryInterface
}

type RegistryService struct {
	Repository entity.CustomerRegistryInterface
	Cache      cache.CacheInterface
	Tracer     *otelpkg.OtelPkgInstrument
}

// NewRegistryService creates a new RegistryService.
func NewRegistryService(repo entity.CustomerRegistryInterface, cc *cache.CacheInterface, otl *otelpkg.OtelPkgInstrument) (entity.CustomerRegistryInterface, error) {
	return &RegistryService{
		Repository: repo,
		Cache:      *cc,
		Tracer:     otl,
	}, nil
}

func (rs *RegistryService) Create(ctx context.Context, cr *entity.CustomerRegistryResponse) (*entity.CustomerRegistryResponse, error) {

	if err := cr.Validate(); err != nil {
		return nil, err
	}

	one, err := rs.GetByID(ctx, &cr.CustomerRegistry.UserEmail)
	if err != nil {
		return nil, err
	}

	if one.Avaliable {
		return nil, errors.New("the customer has one lab running")
	}

	crr, err := entity.RegistryNewCustomer(cr.CustomerRegistry)
	if err != nil {
		return nil, err
	}

	_, err = rs.Repository.Create(ctx, crr)
	if err != nil {
		return nil, err
	}

	// if err := rs.Cache.Set(crr.ID, crr); err != nil {
	// }
	return crr, nil
}

func (rs *RegistryService) GetAll(ctx context.Context) ([]entity.CustomerRegistryResponse, error) {

	result, err := rs.Repository.GetAll(ctx)
	if err != nil {
		return nil, err
	}

	var wg sync.WaitGroup
	for i := range result {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			result[i].Avaliable = rs.isAvaliable(&result[i])
		}(i)
	}
	wg.Wait()

	return result, nil
}
func (rs *RegistryService) Get(cr *entity.CustomerRegistryResponse) (*entity.CustomerRegistryResponse, error) {
	return nil, nil
}

func (rs *RegistryService) GetByFilter(cr *entity.CustomerRegistry) ([]entity.CustomerRegistryResponse, error) {

	if cr == nil {
		return nil, errors.New("filter is required")
	}

	return rs.Repository.GetByFilter(cr)
}

func (rs *RegistryService) GetByID(ctx context.Context, id *string) (*entity.CustomerRegistryResponse, error) {

	uuidIsValid := kd_uuid.IsValid(id)

	emailIsValid := kd_utils.IsValidEmail(id)
	if !uuidIsValid && !emailIsValid {
		return nil, errors.New("invalid id or email")
	}

	var err error
	result := &entity.CustomerRegistryResponse{}

	if uuidIsValid {
		result, err = rs.Repository.GetByID(ctx, id)
		if err != nil {
			return nil, err
		}
	}

	if emailIsValid {
		resultList, err := rs.GetByFilter(&entity.CustomerRegistry{UserEmail: *id})
		if err != nil {
			return nil, err
		}

		// Sort the resultList by EndAt field
		sort.Slice(resultList, func(i, j int) bool {
			layout := time.RFC3339
			timeI, errI := time.Parse(layout, resultList[i].EndAt)
			timeJ, errJ := time.Parse(layout, resultList[j].EndAt)
			if errI != nil || errJ != nil {
				return false
			}
			return timeI.Before(timeJ)
		})

		if len(resultList) > 0 {
			result = &resultList[len(resultList)-1]
		}
	}

	if result == nil {
		return result, nil
	}

	result.Avaliable = rs.isAvaliable(result)
	return result, nil
}

func (rs *RegistryService) isAvaliable(cr *entity.CustomerRegistryResponse) bool {

	target, _ := kd_utils.TimeAfterThan(time.RFC3339, time.Now().Format(time.RFC3339), cr.EndAt)

	return target
}
