// Copyright (c) Citrix, Inc.

package common

import (
	"testing"

	elmsoap "github.com/citrix/packer-plugin-citrix/elm-client"
)

func TestFormatELMError(t *testing.T) {
	cases := []struct {
		name string
		in   *elmsoap.ApplicationError
		want string
	}{
		{"nil", nil, "no error details available"},
		{"empty message", &elmsoap.ApplicationError{}, "no error details available"},
		{"with message", &elmsoap.ApplicationError{Message: "boom"}, "boom"},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			if got := FormatELMError(tc.in); got != tc.want {
				t.Errorf("FormatELMError = %q, want %q", got, tc.want)
			}
		})
	}
}

func TestGetCreateLayerResultError(t *testing.T) {
	wantErr := &elmsoap.ApplicationError{Message: "create failed"}
	cases := []struct {
		name string
		in   *elmsoap.CreateLayerResult
		want *elmsoap.ApplicationError
	}{
		{"nil result", nil, nil},
		{"nil WebResultBase", &elmsoap.CreateLayerResult{}, nil},
		{"nil ResultBase", &elmsoap.CreateLayerResult{WebResultBase: &elmsoap.WebResultBase{}}, nil},
		{"nil Error in ResultBase",
			&elmsoap.CreateLayerResult{WebResultBase: &elmsoap.WebResultBase{ResultBase: &elmsoap.ResultBase{}}}, nil},
		{"populated error",
			&elmsoap.CreateLayerResult{WebResultBase: &elmsoap.WebResultBase{ResultBase: &elmsoap.ResultBase{Error: wantErr}}}, wantErr},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			got := GetCreateLayerResultError(tc.in)
			if got != tc.want {
				t.Errorf("GetCreateLayerResultError = %v, want %v", got, tc.want)
			}
		})
	}
}

func TestGetRevisionResultError(t *testing.T) {
	wantErr := &elmsoap.ApplicationError{Message: "revision failed"}
	cases := []struct {
		name string
		in   *elmsoap.CreateRevisionResult
		want *elmsoap.ApplicationError
	}{
		{"nil result", nil, nil},
		{"nil WebResultBase", &elmsoap.CreateRevisionResult{}, nil},
		{"nil ResultBase", &elmsoap.CreateRevisionResult{WebResultBase: &elmsoap.WebResultBase{}}, nil},
		{"nil Error in ResultBase",
			&elmsoap.CreateRevisionResult{WebResultBase: &elmsoap.WebResultBase{ResultBase: &elmsoap.ResultBase{}}}, nil},
		{"populated error",
			&elmsoap.CreateRevisionResult{WebResultBase: &elmsoap.WebResultBase{ResultBase: &elmsoap.ResultBase{Error: wantErr}}}, wantErr},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			got := GetRevisionResultError(tc.in)
			if got != tc.want {
				t.Errorf("GetRevisionResultError = %v, want %v", got, tc.want)
			}
		})
	}
}
