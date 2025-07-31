package mock_test

import (
	"github.com/ivan-zherdev/benthos/v4/internal/component/output"
	"github.com/ivan-zherdev/benthos/v4/internal/manager/mock"
)

var _ output.Sync = mock.OutputWriter(nil)
