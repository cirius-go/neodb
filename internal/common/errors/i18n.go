package errors

import (
	"cirius-go/neodb/internal/common"
	"cirius-go/neodb/internal/common/slice"
	"embed"
	"fmt"
	"net/http"
	"sync/atomic"
)

type localizerHolder struct {
	localizer common.Localizer
}

// localizer is a package-level variable for localization.
var localizer atomic.Pointer[localizerHolder]

// SetLocalizer sets the localizer for localization.
func SetLocalizer(l common.Localizer) {
	ol := localizer.Load()
	if ol != nil {
		return
	}
	localizer.CompareAndSwap(nil, &localizerHolder{
		localizer: l,
	})
}

// GetLocalizer retrieves the localizer for localization.
func GetLocalizer() common.Localizer {
	l := localizer.Load()
	if l == nil {
		panic("localizer is not set")
	}
	return l.localizer
}

// LocalizerMessageData represents the data needed for localizing a message.
type LocalizerMessageData struct {
	MessageID    string `json:"-" yaml:"-"`
	TemplateData any    `json:"-" yaml:"-"`
	PluralCount  any    `json:"-" yaml:"-"`
}

// Localize Localizes the message and returns it as a string.
func (l *LocalizerMessageData) Localize(langTag string) string {
	if l == nil {
		return ""
	}
	localizer := GetLocalizer()
	return localizer.Localize(langTag, common.LocalizeConfig{
		MessageID:    l.MessageID,
		TemplateData: l.TemplateData,
		PluralCount:  l.PluralCount,
	})
}

// NewLocalizedMsgByID creates a new instance of LocalizerMessageData.
func NewLocalizedMsgByID(messageID string) LocalizerMessageData {
	return LocalizerMessageData{
		MessageID: messageID,
	}
}

type (
	// I18nErrorDetail represents detailed information about an internationalization error.
	I18nErrorDetail struct {
		// Location of the error, e.g., field name or parameter.
		Location string `json:"location" yaml:"location"`
		// Value associated with the error, e.g., invalid value.
		Value any `json:"value,omitempty" yaml:"value,omitempty"`
		// Message is the human-readable error message.
		Message LocalizerMessageData `json:"message" yaml:"message"`
	}
	// I18nStatusError represents an internationalization error.
	I18nStatusError struct {
		Status   int                  `json:"status" yaml:"status"`
		Title    string               `json:"title" yaml:"title"`
		Detail   LocalizerMessageData `json:"detail" yaml:"detail"`
		Errors   []*I18nErrorDetail   `json:"errors,omitempty" yaml:"errors,omitempty"`
		internal error                `json:"-" yaml:"-"`
	}
)

// Error implements error.
func (i *I18nErrorDetail) Error() string {
	msg := i.Message.Localize(DefaultLanguage)
	return fmt.Sprintf("%s is having error: %s", i.Location, msg)
}

// Error implements common.I18nError.
func (i *I18nStatusError) Error() string {
	return i.Detail.Localize(DefaultLanguage)
}

// Is implements common.I18nError.
func (i *I18nStatusError) Is(target error) bool {
	if t, ok := target.(*I18nStatusError); ok {
		return i.Detail.Localize(DefaultLanguage) == t.Detail.Localize(DefaultLanguage) && i.Status == t.Status
	}
	return false
}

// Localize implements common.I18nError.
func (i *I18nStatusError) Localize(langTag string) *StatusError {
	errs := make([]*ErrorDetail, 0, len(i.Errors))
	for _, e := range i.Errors {
		errs = append(errs, &ErrorDetail{
			Location: e.Location,
			Value:    e.Value,
			Message:  e.Message.Localize(langTag),
		})
	}
	return &StatusError{
		Status:   i.Status,
		Title:    i.Title,
		Detail:   i.Detail.Localize(langTag),
		Errors:   errs,
		internal: i.internal,
	}
}

// Unwrap implements common.I18nError.
func (i *I18nStatusError) Unwrap() error {
	return i.internal
}

// WithDetail implements common.I18nError.
func (i *I18nStatusError) WithDetail(detail *I18nErrorDetail) *I18nStatusError {
	i.Errors = append(i.Errors, detail)
	return i
}

// WithInternal implements common.I18nError.
func (i *I18nStatusError) WithInternal(err error) *I18nStatusError {
	i.internal = err
	return i
}

var _ common.I18nError[*I18nStatusError, *I18nErrorDetail, *StatusError] = (*I18nStatusError)(nil)

// NewWithI18n creates a new instance of I18nStatusError.
func NewWithI18n(status int, msg LocalizerMessageData, errs ...error) *I18nStatusError {
	return &I18nStatusError{
		Status: status,
		Title:  http.StatusText(status),
		Detail: msg,
		Errors: slice.Map(errs, ConvertToI18nErrDetail),
	}
}

// ConvertToI18nErrDetail converts a generic error to an I18nErrorDetail.
func ConvertToI18nErrDetail(err error) *I18nErrorDetail {
	if err == nil {
		return nil
	}
	if af, ok := err.(*I18nErrorDetail); ok {
		return af
	}
	if ef, ok := err.(*ErrorDetail); ok {
		return &I18nErrorDetail{
			Location: ef.Location,
			Value:    ef.Value,
			Message: LocalizerMessageData{
				MessageID: ef.Message,
			},
		}
	}
	return &I18nErrorDetail{
		Location: UnknownErrorLocation,
		Value:    nil,
		Message: LocalizerMessageData{
			MessageID: err.Error(),
		},
	}
}

// some common errors with i18n support.
var (
	//go:embed locales/*.toml
	LocaleFS embed.FS

	ErrInvalidRequestBody = NewWithI18n(http.StatusBadRequest, NewLocalizedMsgByID("error.invalid_request_body"))
	ErrUpstreamService    = NewWithI18n(http.StatusBadGateway, NewLocalizedMsgByID("error.upstream_service_error"))
	ErrInternalServer     = NewWithI18n(http.StatusInternalServerError, NewLocalizedMsgByID("error.internal_server_error"))
	ErrUnauthorized       = NewWithI18n(http.StatusUnauthorized, NewLocalizedMsgByID("error.unauthorized"))
	ErrForbidden          = NewWithI18n(http.StatusForbidden, NewLocalizedMsgByID("error.forbidden"))
	ErrNotFound           = NewWithI18n(http.StatusNotFound, NewLocalizedMsgByID("error.resource_not_found"))
	ErrConflict           = NewWithI18n(http.StatusConflict, NewLocalizedMsgByID("error.resource_conflict"))
)
