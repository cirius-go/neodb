package i18n

import (
	"io/fs"
	"maps"
	"sync/atomic"
	"text/template"

	"cirius-go/neodb/pkg/common"
	"cirius-go/neodb/pkg/common/logger"

	"github.com/nicksnyder/go-i18n/v2/i18n"
	"golang.org/x/text/language"
)

type (
	// Config represents the configuration for the Localizer.
	Config struct {
		DefaultLanguage language.Tag
		Pipes           template.FuncMap
		Logger          common.Logger
	}
	// Localizer wraps the i18n Localizer.
	Localizer struct {
		bundle *i18n.Bundle
		pipes  template.FuncMap
		logger common.Logger
	}
)

// LoadMessageFileFS implements common.Localizer.
func (l *Localizer) LoadMessageFileFS(fsys fs.FS, path ...string) error {
	for _, p := range path {
		l.logger.Debug("loading fs file", "file", p)
		if _, err := l.bundle.LoadMessageFileFS(fsys, p); err != nil {
			return err
		}
	}
	return nil
}

// RegisterUnmarshalFunc implements common.Localizer.
func (l *Localizer) RegisterUnmarshalFunc(name string, fn common.LocalizeUnmarshalFunc) {
	l.bundle.RegisterUnmarshalFunc(name, fn)
}

// Localize implements common.Localizer.
func (l *Localizer) Localize(langTag string, conf common.LocalizeConfig) string {
	funcs := maps.Clone(l.pipes)
	maps.Copy(funcs, conf.Funcs)
	msg, err := i18n.NewLocalizer(l.bundle, langTag).Localize(&i18n.LocalizeConfig{
		MessageID:    conf.MessageID,
		TemplateData: conf.TemplateData,
		PluralCount:  conf.PluralCount,
		Funcs:        funcs,
	})
	if err != nil {
		l.logger.Error("i18n localize failed", "error", err)
		return conf.MessageID
	}
	return msg
}

var _ common.Localizer = (*Localizer)(nil)

var glob atomic.Pointer[Localizer]

// NewLocalize creates a new Localizer instance.
// It loads the language files from the embedded filesystem
// by the provided language codes.
func NewLocalizer(cfg Config) *Localizer {
	bundle := i18n.NewBundle(cfg.DefaultLanguage)
	lg := logger.Get()
	if cfg.Logger != nil {
		lg = cfg.Logger
	}
	return &Localizer{
		bundle: bundle,
		pipes:  cfg.Pipes,
		logger: lg.With("module", "i18n"),
	}
}

// SetLocalizer sets the global Localizer instance.
func SetLocalizer(l *Localizer) {
	ol := glob.Load()
	if ol != nil {
		return
	}
	glob.CompareAndSwap(nil, l)
}

// GetLocalizer retrieves the global Localizer instance.
func GetLocalizer() *Localizer {
	ol := glob.Load()
	if ol == nil {
		panic("localizer is not set")
	}
	return ol
}
