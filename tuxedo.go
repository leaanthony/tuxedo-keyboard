package tuxedokeyboard

import (
	"fmt"
	"io/ioutil"
)

// Mode represents the keyboard lighting mode
type Mode uint8

const (
	Custom Mode = iota
	Breathe
	Cycle
	Dance
	Flash
	Random
	Tempo
	Wave
)

// Writes the given value to the sysfs key
func writeSysfsValue(key string, value string) error {
	var data = []byte(value)
	return ioutil.WriteFile("/sys/devices/platform/tuxedo_keyboard/"+key, data, 0644)
}

// SetBrightness will set the brightness of the keyboard backlight
// Accepts values 0-255
func SetBrightness(value uint8) error {
	stringValue := fmt.Sprintf("%d", value)
	return writeSysfsValue("brightness", stringValue)
}

// SetKeyboardColourLeft will set the colour of the left side of the keyboard backlight
// Accepts RGB values 0-255
func SetKeyboardColourLeft(red, green, blue uint8) error {
	hex := fmt.Sprintf("0x%02x%02x%02x", red, green, blue)
	return writeSysfsValue("color_left", hex)
}

// SetKeyboardColourRight will set the colour of the right side of the keyboard backlight
// Accepts RGB values 0-255
func SetKeyboardColourRight(red, green, blue uint8) error {
	hex := fmt.Sprintf("0x%02x%02x%02x", red, green, blue)
	return writeSysfsValue("color_right", hex)
}

// SetKeyboardColourCenter will set the colour of the center of the keyboard backlight
// Accepts RGB values 0-255
func SetKeyboardColourCenter(red, green, blue uint8) error {
	hex := fmt.Sprintf("0x%02x%02x%02x", red, green, blue)
	return writeSysfsValue("color_center", hex)
}

// EnableLights turns the keyboard backlight on
func EnableLights() error {
	return writeSysfsValue("state", "1")
}

// DisableLights turns the keyboard backlight off
func DisableLights() error {
	return writeSysfsValue("state", "0")
}

// SetMode sets the keyboard backlight mode
func SetMode(mode Mode) error {
	modeString := fmt.Sprintf("%d", mode)
	return writeSysfsValue("mode", modeString)
}
