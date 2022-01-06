package we

import (
	"context"

	"github.com/pkg/errors"
)

type StarterImpl struct {
	Runner Runner
}

func (s *StarterImpl) Start(ctx context.Context) error {
	if err := s.Runner.Run(ctx); err != nil {
		return errors.Errorf("command failed: %w", err)
	}

	return nil
}
