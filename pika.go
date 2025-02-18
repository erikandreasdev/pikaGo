package main

import (
	_ "embed"
	"fmt"
	"github.com/getlantern/systray"
	"log"
	"os/exec"
	"strings"
)

//go:embed icon.png
var iconBytes []byte

type BrightnessControl struct {
	display string
}

func NewBrightnessControl() (*BrightnessControl, error) {
	display, err := getPrimaryDisplay()
	if err != nil {
		return nil, fmt.Errorf("failed to get primary display: %v", err)
	}

	return &BrightnessControl{
		display: display,
	}, nil
}

func getPrimaryDisplay() (string, error) {
	cmd := exec.Command("xrandr", "--query")
	output, err := cmd.Output()
	if err != nil {
		return "", fmt.Errorf("failed to execute xrandr: %v", err)
	}

	lines := strings.Split(string(output), "\n")
	for _, line := range lines {
		if strings.Contains(line, " connected") {
			parts := strings.Split(line, " ")
			return parts[0], nil
		}
	}

	return "", fmt.Errorf("no connected display found")
}

func (bc *BrightnessControl) setBrightness(percentage float64) error {
	cmd := exec.Command("xrandr", "--output", bc.display, "--brightness", fmt.Sprintf("%.2f", percentage/100))
	return cmd.Run()
}

func (bc *BrightnessControl) onReady() {
	systray.SetIcon(iconBytes)
	systray.SetTooltip("Brightness Control")

	m100 := systray.AddMenuItem("100%", "Full brightness")
	m75 := systray.AddMenuItem("75%", "75% brightness")
	m50 := systray.AddMenuItem("50%", "50% brightness")
	m25 := systray.AddMenuItem("25%", "25% brightness")
	systray.AddSeparator()
	mQuit := systray.AddMenuItem("Quit", "Quit the app")

	go func() {
		for {
			select {
			case <-m100.ClickedCh:
				err := bc.setBrightness(100)
				if err != nil {
					return
				}
			case <-m75.ClickedCh:
				err := bc.setBrightness(75)
				if err != nil {
					return
				}
			case <-m50.ClickedCh:
				err := bc.setBrightness(50)
				if err != nil {
					return
				}
			case <-m25.ClickedCh:
				err := bc.setBrightness(25)
				if err != nil {
					return
				}
			case <-mQuit.ClickedCh:
				systray.Quit()
				return
			}
		}
	}()
}

func (bc *BrightnessControl) onExit() {}

func main() {
	bc, err := NewBrightnessControl()
	if err != nil {
		log.Fatal(err)
	}

	systray.Run(bc.onReady, bc.onExit)
}
