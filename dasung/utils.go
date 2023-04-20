package dasung

import (
	"bytes"
	"os/exec"
	"strings"
	"errors"
)

func FindDasungI2CDevicePaths() ([]string, error) {
	cmd := exec.Command("ddcutil", "detect", "--verbose")
	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()

	if err != nil {
		return nil, err
	}

	output := out.String()
	lines := strings.Split(output, "\n")
	isDasung := false
	devicePaths := make([]string, 0)

	for _, line := range lines {
		if isDasung {
			if strings.Contains(line, "I2C bus:") {
				path := strings.TrimSpace(strings.Split(line, ":")[1])
				devicePaths = append(devicePaths, path)
			}
		}

		if strings.Contains(line, "DSC:Paperlike") {
			isDasung = true
		} else {
			isDasung = false
		}
	}

	if len(devicePaths) == 0 {
		return nil, errors.New("No Dasung Paperlike displays found. Please make sure your device is connected and powered on.")
	}

	return devicePaths, nil
}