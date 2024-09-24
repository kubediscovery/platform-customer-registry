package service

import (
	"errors"
	"log"

	"github.com/kubediscovery/platform-customer-registry/pkg/cache"
	"github.com/kubediscovery/platform-customer-registry/pkg/otelpkg"

	"github.com/kubediscovery/platform-customer-registry/internal/core/entity"
)

type RegistryService struct {
	Repository entity.CustomerRegistryInterface
	Cache      cache.CacheInterface
	Tracer     *otelpkg.OtelPkgInstrument
}

// NewRegistryService creates a new RegistryService.
func NewRegistryService(repo entity.CustomerRegistryInterface, cc *cache.CacheInterface, otl *otelpkg.OtelPkgInstrument) (entity.CustomerRegistryInterface, error ){
	return &RegistryService{
		Repository: repo,
		Cache:      *cc,
		Tracer:     otl,
	}, nil
}

func (rs *RegistryService) Create(ctx context.Context , cr *entity.CustomerRegistry) (*entity.CustomerRegistryResponse, error) {

	if err := cr.Validate(); err != nil {
		return nil, err
	}

	result, err := rs.Search(cr)
	if err != nil {
		return nil, err
	}

	if len(result) > 0 {
		if result[0].ProjectName == cr.ProjectName &&
			result[0].Repository == cr.Repository &&
			result[0].UserName == cr.UserName &&
			result[0].UserEmail == cr.UserEmail && 
			result[0].IsValid{
			return nil, errors.New("the customer has one lab running")
		}
	}

	crr, err := entity.RegistryNewCustomer(cr)
	if err != nil {
		return nil, err
	}

	// crr, err = rs.Repository.Create(crr)
	// if err != nil {
	// 	return nil, err
	// }

	return crr, nil
}

func (rs *RegistryService) List() ([]entity.CustomerRegistryResponse, error) {
	return nil, nil
}
func (rs *RegistryService) Get(cr *entity.CustomerRegistryResponse) (*entity.CustomerRegistryResponse, error) {
	return nil, nil
}

func (rs *RegistryService) Search(cr *entity.CustomerRegistry) ([]entity.CustomerRegistryResponse, error) {
	return nil, nil
}

func (rs *RegistryService) IsValid(cr *entity.CustomerRegistry) (*entity.CustomerRegistryResponse, error) {
	return nil, nil
}
