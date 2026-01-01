package launchctl

import (
	"context"
	"os/exec"
	"strings"

	"github.com/cockroachdb/errors"
)

func Bootout(ctx context.Context, serviceTarget string) error {
	cmd := exec.CommandContext(ctx, "launchctl", "bootout", serviceTarget)
	out, err := cmd.CombinedOutput()
	if err != nil {
		return errors.Wrapf(err, "bootout failed: %s", string(out))
	}
	return nil
}

func Bootstrap(ctx context.Context, domainTarget string, plistPath string) error {
	cmd := exec.CommandContext(ctx, "launchctl", "bootstrap", domainTarget, plistPath)
	out, err := cmd.CombinedOutput()
	if err != nil {
		return errors.Wrapf(err, "bootstrap failed: %s", string(out))
	}
	return nil
}

func IsLoaded(ctx context.Context, serviceTarget string) (bool, error) {
	cmd := exec.CommandContext(ctx, "launchctl", "info", serviceTarget)
	out, err := cmd.Output()
	if err != nil {
		return false, errors.Wrap(err, "info failed")
	}

	return strings.Contains(string(out), serviceTarget), nil
}

func Kickstart(ctx context.Context, serviceTarget string) error {
	cmd := exec.CommandContext(ctx, "launchctl", "kickstart", serviceTarget)
	out, err := cmd.CombinedOutput()
	if err != nil {
		return errors.Wrapf(err, "kickstart failed: %s", string(out))
	}
	return nil
}
