package sidebar

// Sidebar represents the sidebar component.
type Sidebar struct{}

// Config represents the configuration for the Sidebar component.
type Config struct{}

// Sidebar constructor.
func New(cfg Config) *Sidebar {
	return &Sidebar{}
}
