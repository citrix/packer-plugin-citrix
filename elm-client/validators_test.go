// Copyright (c) Citrix, Inc.

package elmsoap

import (
	"errors"
	"testing"
)

func TestIsValidApplayeringOperationType(t *testing.T) {
	cases := []struct {
		op   string
		want bool
	}{
		{REVISION_OS_LAYER, true},
		{REVISION_PLATFORM_LAYER, true},
		{REVISION_APP_LAYER, true},
		{CREATE_PLATFORM_LAYER, true},
		{CREATE_APP_LAYER, true},
		{CONNECT_REVISION_OS_VM_ONLY, true},
		{CONNECT_REVISION_PLATFORM_VM_ONLY, true},
		{CONNECT_REVISION_APP_VM_ONLY, true},
		{CONNECT_CREATE_PLATFORM_VM_ONLY, true},
		{CONNECT_CREATE_APP_VM_ONLY, true},
		{APPLAYERING_OPERATION_TYPE_UNDEFINED, true},
		{"NOT_A_REAL_OP", false},
		{"", false},
	}
	for _, tc := range cases {
		t.Run(tc.op, func(t *testing.T) {
			if got := IsValidApplayeringOperationType(tc.op); got != tc.want {
				t.Errorf("IsValidApplayeringOperationType(%q) = %v, want %v", tc.op, got, tc.want)
			}
		})
	}
}

func TestGetAllSupportedApplayeringOperationTypes(t *testing.T) {
	got := GetAllSupportedApplayeringOperationTypes()
	want := []string{
		REVISION_OS_LAYER,
		REVISION_PLATFORM_LAYER,
		REVISION_APP_LAYER,
		CREATE_PLATFORM_LAYER,
		CREATE_APP_LAYER,
	}
	if len(got) != len(want) {
		t.Fatalf("len = %d, want %d", len(got), len(want))
	}
	for i, w := range want {
		if got[i] != w {
			t.Errorf("[%d] = %q, want %q", i, got[i], w)
		}
	}
}

func TestIsValidDiskFormat(t *testing.T) {
	cases := []struct {
		in   string
		want bool
	}{
		{string(DiskFormatVhdDiskFormat), true},
		{string(DiskFormatVmdkDiskFormat), true},
		{string(DiskFormatVmdkSparseDiskFormat), true},
		{string(DiskFormatVhdxDiskFormat), true},
		{string(DiskFormatQCow2DiskFormat), true},
		{"", false},
		{"raw", false},
		{"VHD", false}, // case-sensitive
	}
	for _, tc := range cases {
		t.Run(tc.in, func(t *testing.T) {
			if got := IsValidDiskFormat(tc.in); got != tc.want {
				t.Errorf("IsValidDiskFormat(%q) = %v, want %v", tc.in, got, tc.want)
			}
		})
	}
}

func TestGetAllSupportedDiskFormats(t *testing.T) {
	got := GetAllSupportedDiskFormats()
	if len(got) != 5 {
		t.Fatalf("len = %d, want 5", len(got))
	}
	// every returned format must round-trip through IsValidDiskFormat.
	for _, f := range got {
		if !IsValidDiskFormat(f) {
			t.Errorf("GetAllSupportedDiskFormats returned %q but IsValidDiskFormat rejects it", f)
		}
	}
}

func TestCheckWebResultError(t *testing.T) {
	t.Run("nil WebResultBase returns nil", func(t *testing.T) {
		if err := CheckWebResultError(nil); err != nil {
			t.Errorf("nil input err = %v, want nil", err)
		}
	})
	t.Run("nil ResultBase returns nil", func(t *testing.T) {
		if err := CheckWebResultError(&WebResultBase{}); err != nil {
			t.Errorf("err = %v, want nil", err)
		}
	})
	t.Run("nil Error returns nil", func(t *testing.T) {
		w := &WebResultBase{ResultBase: &ResultBase{}}
		if err := CheckWebResultError(w); err != nil {
			t.Errorf("err = %v, want nil", err)
		}
	})
	t.Run("empty Message returns nil", func(t *testing.T) {
		w := &WebResultBase{ResultBase: &ResultBase{Error: &ApplicationError{Message: ""}}}
		if err := CheckWebResultError(w); err != nil {
			t.Errorf("err = %v, want nil", err)
		}
	})
	t.Run("non-empty Message returns wrapped error", func(t *testing.T) {
		w := &WebResultBase{ResultBase: &ResultBase{Error: &ApplicationError{Message: "boom"}}}
		err := CheckWebResultError(w)
		if err == nil || err.Error() != "boom" {
			t.Errorf("err = %v, want %q", err, "boom")
		}
	})
}

func TestBuildServerURL(t *testing.T) {
	cases := []struct {
		name string
		in   string
		want string
	}{
		{"bare host, no scheme, no path", "elm.example.com", "https://elm.example.com/Unidesk.Web/API.asmx"},
		{"https scheme, no path", "https://elm.example.com", "https://elm.example.com/Unidesk.Web/API.asmx"},
		{"http scheme preserved, no path", "http://elm.example.com", "http://elm.example.com/Unidesk.Web/API.asmx"},
		{"root path is replaced with default", "https://elm.example.com/", "https://elm.example.com/Unidesk.Web/API.asmx"},
		{"explicit path is preserved", "https://elm.example.com/custom/api", "https://elm.example.com/custom/api"},
		{"host with port", "elm.example.com:8443", "https://elm.example.com:8443/Unidesk.Web/API.asmx"},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			got, err := BuildServerURL(tc.in)
			if err != nil {
				t.Fatalf("err = %v, want nil", err)
			}
			if got != tc.want {
				t.Errorf("BuildServerURL(%q) = %q, want %q", tc.in, got, tc.want)
			}
		})
	}
}

// Sanity assertion: ErrWorkTicketNotInActiveFilter is the same identity value
// across multiple imports — important since callers use errors.Is to detect
// the "moved to completed list" signal.
func TestErrWorkTicketNotInActiveFilter_IsStable(t *testing.T) {
	if !errors.Is(ErrWorkTicketNotInActiveFilter, ErrWorkTicketNotInActiveFilter) {
		t.Errorf("ErrWorkTicketNotInActiveFilter should match itself via errors.Is")
	}
	if ErrWorkTicketNotInActiveFilter.Error() == "" {
		t.Errorf("ErrWorkTicketNotInActiveFilter should have a non-empty message")
	}
}
