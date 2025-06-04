package config

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"go.fork.vn/di"
	di_mocks "go.fork.vn/di/mocks"
)

func TestNewServiceProvider(t *testing.T) {
	provider := NewServiceProvider()
	assert.NotNil(t, provider, "NewServiceProvider should return a non-nil provider")

	// Verify that the provider object is properly initialized
	assert.IsType(t, &ServiceProvider{}, provider, "Provider should be of type *ServiceProvider")
}

func TestServiceProvider_Register(t *testing.T) {
	provider := NewServiceProvider()

	// Test with valid app
	diContainer := di.New()
	mockApp := new(di_mocks.MockApplication)
	mockApp.On("Container").Return(diContainer)

	provider.Register(mockApp)

	// Check that config was registered in the container
	cfg, err := diContainer.Make("config")
	assert.NoError(t, err)
	assert.NotNil(t, cfg, "Config should be registered in container")

	// Verify that the registered object is a Manager
	_, ok := cfg.(Manager)
	assert.True(t, ok, "Registered object should implement Manager interface")

	// Verify all expectations
	mockApp.AssertExpectations(t)
}

func TestServiceProvider_Boot(t *testing.T) {
	provider := NewServiceProvider()

	// Test with valid app
	mockApp := new(di_mocks.MockApplication)
	assert.NotPanics(t, func() {
		provider.Boot(mockApp)
	}, "Boot should not panic with valid app")

	// Test with nil app - should panic
	assert.Panics(t, func() {
		provider.Boot(nil)
	}, "Boot should panic with nil app")
}

func TestServiceProvider_Requires(t *testing.T) {
	provider := NewServiceProvider()

	// Verify that Requires returns an empty slice
	requires := provider.Requires()
	assert.NotNil(t, requires, "Requires should return a non-nil slice")
	assert.Empty(t, requires, "Config provider shouldn't have any dependencies")
}

func TestServiceProvider_Providers(t *testing.T) {
	provider := NewServiceProvider()

	// Verify that Providers returns a slice with "config"
	providers := provider.Providers()
	assert.NotNil(t, providers, "Providers should return a non-nil slice")
	assert.Len(t, providers, 1, "Config provider should register exactly one service")
	assert.Equal(t, "config", providers[0], "Config provider should register a 'config' service")
}
