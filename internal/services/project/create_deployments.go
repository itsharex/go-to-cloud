package project

import (
	"go-to-cloud/internal/models/deploy"
	"go-to-cloud/internal/repositories"
)

func CreateDeployments(_ uint64, d *deploy.Deployment) error {
	repo := d.Deployment
	return repositories.CreateDeployment(&repo)
}
