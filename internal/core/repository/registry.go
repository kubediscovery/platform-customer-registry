package	repository

type RegistryRepository struct {
	Repository entity.CustomerRegistryInterface
	Tracer     *otelpkg.OtelPkgInstrument
}

// NewRegistryService creates a new RegistryService.
func NewRegistryRepository(repo entity.CustomerRegistryInterface, cc *cache.CacheInterface, otl *otelpkg.OtelPkgInstrument) (entity.CustomerRegistryInterface, error ){
	return &RegistryService{
		Repository: repo,
		Cache:      *cc,
		Tracer:     otl,
	}, nil
}