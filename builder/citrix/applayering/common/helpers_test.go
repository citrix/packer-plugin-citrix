// Copyright (c) Citrix, Inc.

package common

import "io"

// recordingUi is a minimal packersdk.Ui implementation that captures Say calls
// for assertion. All other methods are no-ops.
type recordingUi struct {
	said []string
}

func (u *recordingUi) Ask(string) (string, error)          { return "", nil }
func (u *recordingUi) Askf(string, ...any) (string, error) { return "", nil }
func (u *recordingUi) Say(s string)                        { u.said = append(u.said, s) }
func (u *recordingUi) Sayf(string, ...any)                 {}
func (u *recordingUi) Message(string)                      {}
func (u *recordingUi) Error(string)                        {}
func (u *recordingUi) Errorf(string, ...any)               {}
func (u *recordingUi) Machine(string, ...string)           {}

// getter.ProgressTracker method (required by packersdk.Ui).
func (u *recordingUi) TrackProgress(src string, currentSize, totalSize int64, stream io.ReadCloser) io.ReadCloser {
	return stream
}
