package entity

import (
	"context"
	"errors"
	"time"

	"github.com/kubediscovery/platform-customer-registry/pkg/kd_utils"
	"github.com/kubediscovery/platform-customer-registry/pkg/kd_uuid"
)

type LabDestroyInterface interface {
	Create(ctx context.Context, ld *LabDestroyResponse) (*LabDestroyResponse, error)
	GetAll(ctx context.Context) ([]LabDestroyResponse, error)
	GetByID(ctx context.Context, ld *string) (*LabDestroyResponse, error)
	GetByFilter(ctx context.Context, ld *LabDestroyResponse) ([]LabDestroyResponse, error)
	PATCH(ctx context.Context, id, error_message *string, available *bool) (*LabDestroyResponse, error)
}

type LabDestroy struct {
	CustomerRegistry
	TargetDestroy string `json:"target_destroy,omitempty"`
	CreatedAt     string `json:"created_at,omitempty"`
}

type LabDestroyResponse struct {
	LabDestroy
	ID        string `json:"id" binding:"required"`
	UpdatedAt string `json:"updated_at,omitempty"`
	Avaliable *bool  `json:"avaliable" binding:"required"`
	Error     string `json:"error,omitempty"`
}

func NewLabDestroy(l *LabDestroy) (*LabDestroyResponse, error) {

	if err := l.Validate(); err != nil {
		return nil, err
	}

	available := true
	return &LabDestroyResponse{
		LabDestroy: LabDestroy{
			CustomerRegistry: l.CustomerRegistry,
			TargetDestroy:    l.TargetDestroy,
			CreatedAt:        time.Now().Format(time.RFC3339),
		},
		ID:        kd_uuid.NewUuidString(),
		Avaliable: &available,
	}, nil
}

func (ld *LabDestroy) Validate() error {

	if ld.TargetDestroy == "" {
		return errors.New("target destroy is required")
	}

	if ld.CustomerRegistry == (CustomerRegistry{}) {
		return errors.New("customer registry is required")
	}

	if err := ld.CustomerRegistry.Validate(); err != nil {
		return err
	}

	layout := time.RFC3339
	_, err := time.Parse(layout, ld.TargetDestroy)
	if err != nil {
		return err
	}

	return nil
}

func (ld *LabDestroyResponse) Validate() error {

	if ld.LabDestroy == (LabDestroy{}) {
		return errors.New("lab destroy is required")
	}

	if ld.LabDestroy.Validate() != nil {
		return ld.LabDestroy.Validate()
	}

	if !kd_uuid.IsValid(&ld.ID) {
		return errors.New("id invalid format")
	}

	// _, err := ld.IsAvaliable()
	// if err != nil {
	// 	return err
	// }

	return nil
}

func (ld *LabDestroyResponse) IsAvaliable() (bool, error) {

	target, err := kd_utils.TimeAfterThan(time.RFC3339, time.Now().Format(time.RFC3339), ld.LabDestroy.TargetDestroy)
	if err != nil {
		return false, err
	}

	ld.Avaliable = &target
	return target, nil
}
