package mock_test

import (
	"github.com/ivan-zherdev/benthos/v4/internal/component/cache"
	"github.com/ivan-zherdev/benthos/v4/internal/manager/mock"
)

var _ cache.V1 = &mock.Cache{}
