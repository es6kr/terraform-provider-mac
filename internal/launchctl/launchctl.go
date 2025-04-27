package launchctl

import (
	"context"
	"os/exec"
	"strings"

	"github.com/cockroachdb/errors"
)

func Bootout(ctx context.Context, label string) error {
	cmd := exec.CommandContext(ctx, "launchctl", "bootout", "gui/$(id -u)", label)
	out, err := cmd.CombinedOutput()
	if err != nil {
		return errors.Wrapf(err, "bootout failed: %s", string(out))
	}
	return nil
}

func Bootstrap(ctx context.Context, plistPath string) error {
	cmd := exec.CommandContext(ctx, "launchctl", "bootstrap", "gui/$(id -u)", plistPath)
	out, err := cmd.CombinedOutput()
	if err != nil {
		return errors.Wrapf(err, "bootstrap failed: %s", string(out))
	}
	return nil
}

func IsLoaded(ctx context.Context, label string) (bool, error) {
	cmd := exec.CommandContext(ctx, "launchctl", "list")
	out, err := cmd.Output()
	if err != nil {
		return false, errors.Wrap(err, "list services failed")
	}

	return strings.Contains(string(out), label), nil
}

func Kickstart(ctx context.Context, label string) error {
	cmd := exec.CommandContext(ctx, "launchctl", "kickstart", "gui/$(id -u)/"+label)
	out, err := cmd.CombinedOutput()
	if err != nil {
		return errors.Wrapf(err, "kickstart failed: %s", string(out))
	}
	return nil
}
