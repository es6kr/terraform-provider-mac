package launchctl_test

import (
	"context"
	"testing"

	"github.com/es6kr/terraform-provider-mac/internal/launchctl"
	"gotest.tools/v3/assert"
)

func TestKickstart(t *testing.T) { //nolint:tparallel
	t.Parallel()

	ctx := context.Background()

	t.Run("kickstart nonexistent service", func(t *testing.T) { //nolint:paralleltest
		err := launchctl.Kickstart(ctx, "com.example.nonexistent")
		assert.ErrorContains(t, err, "kickstart failed")
	})
}
