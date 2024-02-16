// utils/general.go
package utils

import (
	"fmt"
	"time"

	"github.com/tebeka/selenium"
)

func Sleep(seconds int) {
	// Wait for the specified number of seconds
	time.Sleep(time.Duration(seconds) * time.Second)
}

// WaitForClickable waits for an element to be clickable within a specified timeout duration.
func WaitForClickable(driver selenium.WebDriver, by, value string, timeout time.Duration) (selenium.WebElement, error) {
	// Start time for timeout
	startTime := time.Now()

	// Keep trying to find the element until it's clickable or timeout occurs
	for {
		// Check if the element is present
		elem, err := driver.FindElement(by, value)
		if err != nil {
			if time.Since(startTime) >= timeout {
				return nil, fmt.Errorf("element not found within %v: %v", timeout, err)
			}
			// Sleep for a short duration before trying again
			time.Sleep(500 * time.Millisecond)
			continue
		}

		// Check if the element is clickable
		isClickable, err := IsElementClickable(elem)
		if err != nil {
			return nil, fmt.Errorf("error checking if element is clickable: %v", err)
		}

		// If the element is clickable, return it
		if isClickable {
			return elem, nil
		}

		// If the timeout has been reached, return an error
		if time.Since(startTime) >= timeout {
			return nil, fmt.Errorf("element not clickable within %v", timeout)
		}

		// Sleep for a short duration before trying again
		time.Sleep(500 * time.Millisecond)
	}
}

// IsElementClickable checks if an element is clickable.
// This function can be customized based on specific criteria for clickable elements.
func IsElementClickable(elem selenium.WebElement) (bool, error) {
	// Check if the element is enabled and visible
	isEnabled, err := elem.IsEnabled()
	if err != nil {
		return false, err
	}
	isVisible, err := elem.IsDisplayed()
	if err != nil {
		return false, err
	}
	return isEnabled && isVisible, nil
}
