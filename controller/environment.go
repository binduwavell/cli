package controller

import (
	"context"
	"github.com/railwayapp/cli/entity"
	CLIErrors "github.com/railwayapp/cli/errors"
)

// GetEnvironment returns the currently active environment for the Railway project
func (c *Controller) GetEnvironment(ctx context.Context) (*entity.Environment, error) {
	projectCfg, err := c.cfg.GetProjectConfigs()
	if err != nil {
		return nil, err
	}

	project, err := c.GetProject(ctx, projectCfg.Project)
	if err != nil {
		return nil, err
	}

	for _, environment := range project.Environments {
		if environment.Id == projectCfg.Environment {
			return environment, nil
		}
	}
	return nil, CLIErrors.EnvironmentNotFound
}
