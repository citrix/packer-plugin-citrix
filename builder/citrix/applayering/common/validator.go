// Copyright (c) Citrix, Inc.

package common

import (
	"errors"
	"fmt"
	"regexp"
	"strings"
	"time"

	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	packersdk "github.com/hashicorp/packer-plugin-sdk/packer"
)

// Validate validates a struct using go-playground/validator struct tags.
// Supported tags: required, required_without_all.
// Returns a combined error listing all validation failures, or nil if valid.
func Validate(data interface{}) error {
	re := regexp.MustCompile(`\s+`)
	validate := validator.New(validator.WithRequiredStructEnabled())

	enLocale := en.New()
	translator := ut.New(enLocale, enLocale)
	trans, _ := translator.GetTranslator("en")

	_ = validate.RegisterTranslation("required_without_all", trans,
		func(ut ut.Translator) error {
			return ut.Add("required_without_all",
				"{0}: At least one of the following fields must be present: {1}", true)
		},
		func(ut ut.Translator, fe validator.FieldError) string {
			t, _ := ut.T("required_without_all",
				strings.ToLower(fe.Field()),
				strings.ToLower(fe.Field())+", "+strings.ToLower(re.ReplaceAllString(fe.Param(), ", ")))
			return t
		},
	)

	err := validate.Struct(data)
	if err == nil {
		return nil
	}

	var combined error
	if errs, ok := err.(validator.ValidationErrors); ok {
		for _, fieldErr := range errs {
			combined = errors.Join(combined, errors.New(fieldErr.Translate(trans)))
		}
		return combined
	}
	return err
}

// UiSay prints a timestamped message to the Packer console (==>).
func UiSay(ui packersdk.Ui, msg string) {
	ui.Say(time.Now().Format("2006-01-02 15:04:05.000") + " " + msg)
}

// UiSayf prints a timestamped formatted message to the Packer console (==>).
func UiSayf(ui packersdk.Ui, format string, args ...interface{}) {
	ui.Say(time.Now().Format("2006-01-02 15:04:05.000") + " " + fmt.Sprintf(format, args...))
}
