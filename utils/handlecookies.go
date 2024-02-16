// utils/handlecookies.go

package utils

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/tebeka/selenium"
)

// AcceptCookieBanner finds and clicks the "Accept" button on the cookie consent banner.
func AcceptCookieBanner(driver selenium.WebDriver) error {
	// Execute JavaScript to find the button by its text content.
	script := `
        var buttons = document.querySelectorAll("button");
        for (var i = 0; i < buttons.length; i++) {
            if (buttons[i].textContent.trim() === "Cookies Accepteren") {
                buttons[i].click();
                break;
            }
        }
    `
	if _, err := driver.ExecuteScript(script, nil); err != nil {
		return fmt.Errorf("failed to execute JavaScript to click accept button: %v", err)
	}

	return nil
}

// SaveCookiesToFile saves cookies to a JSON file
func SaveCookiesToFile(cookies []selenium.Cookie, filePath string) error {
	file, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ")
	return encoder.Encode(cookies)
}

// ReloadCookiesFromJSON reloads cookies from a JSON file into WebDriver
func ReloadCookiesFromJSON(wd selenium.WebDriver, filePath string) error {
	// Read cookies from file
	file, err := os.Open(filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	var cookies []*selenium.Cookie
	decoder := json.NewDecoder(file)
	err = decoder.Decode(&cookies)
	if err != nil {
		return err
	}

	// Add cookies to WebDriver
	for _, cookie := range cookies {
		err := wd.AddCookie(cookie)
		if err != nil {
			return err
		}
	}

	// Refresh the page to apply the cookies
	err = wd.Refresh()
	if err != nil {
		return err
	}

	return nil
}
