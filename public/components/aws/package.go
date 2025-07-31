package aws

import (
	// Bring in the internal plugin definitions.
	_ "github.com/ivan-zherdev/benthos/v4/internal/impl/aws"
	_ "github.com/ivan-zherdev/benthos/v4/internal/impl/elasticsearch/aws"
	_ "github.com/ivan-zherdev/benthos/v4/internal/impl/kafka/aws"
	_ "github.com/ivan-zherdev/benthos/v4/internal/impl/opensearch/aws"
)
