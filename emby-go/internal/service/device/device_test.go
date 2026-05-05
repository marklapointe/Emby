package device

import (
	"testing"
)

func TestNewManager(t *testing.T) {
	m := NewManager(nil)
	if m == nil {
		t.Fatal("NewManager returned nil")
	}
	if m.devices == nil {
		t.Error("devices map not initialized")
	}
}

func TestRegisterDevice(t *testing.T) {
	m := NewManager(nil)

	err := m.RegisterDevice(&Device{ID: "device-1", Name: "Test Device"})
	if err != nil {
		t.Errorf("RegisterDevice returned error: %v", err)
	}

	devices := m.GetDevices()
	if len(devices) != 1 {
		t.Errorf("expected 1 device, got %d", len(devices))
	}
}

func TestRegisterDevice_Duplicate(t *testing.T) {
	m := NewManager(nil)

	m.RegisterDevice(&Device{ID: "device-1", Name: "Test Device"})
	err := m.RegisterDevice(&Device{ID: "device-1", Name: "Test Device"})
	if err == nil {
		t.Error("expected error for duplicate device")
	}
}

func TestGetDevice(t *testing.T) {
	m := NewManager(nil)

	m.RegisterDevice(&Device{ID: "device-1", Name: "Test Device"})
	got, ok := m.GetDevice("device-1")
	if !ok {
		t.Error("GetDevice returned false for existing device")
	}
	if got.Name != "Test Device" {
		t.Errorf("expected name 'Test Device', got '%s'", got.Name)
	}
}

func TestGetDevice_NotFound(t *testing.T) {
	m := NewManager(nil)
	_, ok := m.GetDevice("non-existent")
	if ok {
		t.Error("GetDevice returned true for non-existent device")
	}
}

func TestUpdateDevice(t *testing.T) {
	m := NewManager(nil)

	m.RegisterDevice(&Device{ID: "device-1", Name: "Old Name"})
	err := m.UpdateDevice("device-1", "New Name", "New Product")
	if err != nil {
		t.Errorf("UpdateDevice returned error: %v", err)
	}

	got, _ := m.GetDevice("device-1")
	if got.Name != "New Name" {
		t.Errorf("expected name 'New Name', got '%s'", got.Name)
	}
}

func TestGetDevices(t *testing.T) {
	m := NewManager(nil)

	m.RegisterDevice(&Device{ID: "d1", Name: "Device 1"})
	m.RegisterDevice(&Device{ID: "d2", Name: "Device 2"})

	devices := m.GetDevices()
	if len(devices) != 2 {
		t.Errorf("expected 2 devices, got %d", len(devices))
	}
}

func TestRemoveDevice(t *testing.T) {
	m := NewManager(nil)

	m.RegisterDevice(&Device{ID: "device-1", Name: "Test Device"})
	err := m.RemoveDevice("device-1")
	if err != nil {
		t.Errorf("RemoveDevice returned error: %v", err)
	}

	devices := m.GetDevices()
	if len(devices) != 0 {
		t.Errorf("expected 0 devices, got %d", len(devices))
	}
}

func TestGetDeviceProfile(t *testing.T) {
	m := NewManager(nil)

	profile := &DeviceProfile{ID: "profile-1", Name: "Test Profile"}
	m.RegisterDevice(&Device{ID: "device-1", Name: "Test", DeviceProfile: profile})

	got, ok := m.GetDeviceProfile("device-1")
	if !ok {
		t.Error("GetDeviceProfile returned false")
	}
	if got.ID != "profile-1" {
		t.Errorf("expected profile ID 'profile-1', got '%s'", got.ID)
	}
}

func TestGetActiveDeviceCount(t *testing.T) {
	m := NewManager(nil)

	m.RegisterDevice(&Device{ID: "d1", Name: "Device 1"})
	m.RegisterDevice(&Device{ID: "d2", Name: "Device 2"})

	count := m.GetActiveDeviceCount()
	if count != 2 {
		t.Errorf("expected 2 active devices, got %d", count)
	}
}