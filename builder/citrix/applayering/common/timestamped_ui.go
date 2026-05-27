// Copyright (c) Citrix, Inc.

package common

import (
	"fmt"
	"time"

	packersdk "github.com/hashicorp/packer-plugin-sdk/packer"
)

// TimestampedUi wraps a packer.Ui and prepends a millisecond-precision
// timestamp to every Say and Sayf call.
type TimestampedUi struct {
	packersdk.Ui
}

func (t *TimestampedUi) Say(msg string) {
	t.Ui.Say(time.Now().Format("2006-01-02 15:04:05.000") + " " + msg)
}

func (t *TimestampedUi) Sayf(format string, args ...any) {
	t.Ui.Say(time.Now().Format("2006-01-02 15:04:05.000") + " " + fmt.Sprintf(format, args...))
}
