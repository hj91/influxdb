package config

// MockConfigService mocks the ConfigService.
type MockConfigService struct {
	CreateConfigFn func(Config) (Config, error)
	DeleteConfigFn func(name string) (Config, error)
	UpdateConfigFn func(Config) (Config, error)
	ParseConfigsFn func() (Configs, error)
	SwitchConfigFn func(name string) (Config, error)
	ListConfigsFn  func() (Configs, error)
}

// ParseConfigs returns the parse fn.
func (s *MockConfigService) ParseConfigs() (Configs, error) {
	return s.ParseConfigsFn()
}

// CreateConfig create a config.
func (s *MockConfigService) CreateConfig(cfg Config) (Config, error) {
	return s.CreateConfigFn(cfg)
}

// DeleteConfig will delete by name.
func (s *MockConfigService) DeleteConfig(name string) (Config, error) {
	return s.DeleteConfigFn(name)
}

// UpdateConfig will update the config.
func (s *MockConfigService) UpdateConfig(up Config) (Config, error) {
	return s.UpdateConfigFn(up)
}

// SwitchConfig active the config by name.
func (s *MockConfigService) SwitchConfig(name string) (Config, error) {
	return s.SwitchConfigFn(name)
}

// ListConfigs lists all the configs.
func (s *MockConfigService) ListConfigs() (Configs, error) {
	return s.ListConfigsFn()
}
