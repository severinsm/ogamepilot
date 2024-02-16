package controllers

import (
	"fmt"
	"ogame-bot/utils"
	"time"

	"github.com/tebeka/selenium"
)

// ClickPlayButton clicks on the "Play" button.
func ClickPlayButton(driver selenium.WebDriver, timeout time.Duration) error {
	// Wait for the "Play" button to become clickable within the specified timeout
	playButton, err := utils.WaitForClickable(driver, selenium.ByCSSSelector, "#joinGame > a > button", timeout)
	if err != nil {
		return fmt.Errorf("failed to wait for play button: %v", err)
	}

	// Click on the "Play" button
	err = playButton.Click()
	if err != nil {
		return fmt.Errorf("failed to click play button: %v", err)
	}

	return nil
}

func ClickGalaxy(Galaxy string, driver selenium.WebDriver, timeout time.Duration) error {
	// Execute JavaScript to find the div by its text content.
	script := fmt.Sprintf(`
        var divs = document.querySelectorAll("div");
        for (var i = 0; i < divs.length; i++) {
            if (divs[i].textContent.trim() === "%s") {
                divs[i].click();
                break;
            }
        }
    `, Galaxy)
	if _, err := driver.ExecuteScript(script, nil); err != nil {
		return fmt.Errorf("failed to execute JavaScript to click div with text '%s': %v", Galaxy, err)
	}
	// Find the "Spelen" button

	// Wait for the password input field to be clickable
	spelenButton, err := utils.WaitForClickable(driver, selenium.ByCSSSelector, "button.btn.btn-primary", timeout)
	if err != nil {
		return fmt.Errorf("failed to find 'Spelen' button: %v", err)
	}

	// Click on the "Spelen" button
	err = spelenButton.Click()
	if err != nil {
		return fmt.Errorf("failed to click 'Spelen' button: %v", err)
	}
	return nil
}
