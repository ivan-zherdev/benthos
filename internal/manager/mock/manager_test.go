package mock_test

import (
	"github.com/ivan-zherdev/benthos/v4/internal/bundle"
	"github.com/ivan-zherdev/benthos/v4/internal/manager/mock"
)

var _ bundle.NewManagement = &mock.Manager{}
