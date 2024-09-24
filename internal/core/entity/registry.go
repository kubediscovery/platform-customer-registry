package entity

import (
	"errors"
	"context"
	"github.com/kubediscovery/platform-customer-registry/pkg/kd_utils"
	"github.com/kubediscovery/platform-customer-registry/pkg/kd_uuid"
	"time"
)

type CustomerRegistryInterface interface {
	Create(ctx context.Context, cr *CustomerRegistry) (*CustomerRegistryResponse, error)
	List() ([]CustomerRegistryResponse, error)
	Get(cr *CustomerRegistryResponse) (*CustomerRegistryResponse, error)
	Search(cr *CustomerRegistry) ([]CustomerRegistryResponse, error)
	IsValid(cr *CustomerRegistry) (*CustomerRegistryResponse, error)
}

type CustomerRegistry struct {
	ProjectName string `json:'project_name' binding:'required'`
	Repository  string `json:'repository' binding:'required'`
	UserName    string `json:'username' binding:'required'`
	UserEmail   string `json:'email' binding:'required'`
}

type CustomerRegistryResponse struct {
	ID          string `json:'id' binding:'required'`
	ProjectName string `json:'project_name' binding:'required'`
	Repository  string `json:'repository' binding:'required'`
	UserName    string `json:'username' binding:'required'`
	UserEmail   string `json:'email' binding:'required'`
	CreateAt    string `json:'create_at,omitempety'`
	EndAt       string `json:'end_at,omitempety'`
}

func (cr *CustomerRegistry) Validate() error {

	if cr.ProjectName == "" {
		return errors.New("project_name is required")
	}

	if cr.Repository == "" {
		return errors.New("repository is required")
	}

	if cr.UserName == "" {
		return errors.New("username is required")
	}

	if cr.UserEmail == "" {
		return errors.New("email is required")
	}

	if !kd_utils.IsValidEmail(cr.UserEmail) {
		return errors.New("invalid email format")
	}

	return nil
}

func (cr *CustomerRegistryResponse) checkCreatAndEnd() error {

	if cr.CreateAt == "" {
		cr.CreateAt = time.Now().Format(time.RFC3339)
	}

	if cr.EndAt == "" {
		// Parse the time string to time.Time
		createAtTime, err := time.Parse(time.RFC3339, cr.CreateAt)
		if err != nil {
			return errors.New("invalid CreateAt format")
		}

		// Add one hour to the parsed time
		endAtTime := createAtTime.Add(time.Hour)

		// Format the new time back to string
		cr.EndAt = endAtTime.Format(time.RFC3339)
	}

	return nil
}

func RegistryNewCustomer(cr CustomerRegistry) (*CustomerRegistryResponse, error) {

	err := cr.Validate()
	if err != nil {
		return nil, err
	}

	crr := &CustomerRegistryResponse{
		ID:          kd_uuid.NewUuidString(),
		ProjectName: cr.ProjectName,
		Repository:  cr.Repository,
		UserName:    cr.UserName,
		UserEmail:   cr.UserEmail,
	}

	err = crr.checkCreatAndEnd()
	if err != nil {
		return nil, err
	}

	return crr, nil
}
