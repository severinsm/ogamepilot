// utils/authentication.go

package utils

import (
	"fmt"
	"time"

	"github.com/tebeka/selenium"
)

// ClickLoginTab clicks on the login tab.
func ClickLoginTab(driver selenium.WebDriver, timeout time.Duration) error {
	// Wait for the login tab to be clickable
	loginTab, err := WaitForClickable(driver, selenium.ByCSSSelector, "#loginRegisterTabs > ul > li:nth-child(1)", timeout)
	if err != nil {
		return fmt.Errorf("failed to find login tab: %v", err)
	}

	// Click on the login tab
	err = loginTab.Click()
	if err != nil {
		return fmt.Errorf("failed to click login tab: %v", err)
	}

	return nil
}

// EnterCredentials enters the email and password in the respective input fields.
func EnterCredentials(driver selenium.WebDriver, email, password string, timeout time.Duration) error {
	// Wait for the email input field to be clickable
	emailInput, err := WaitForClickable(driver, selenium.ByCSSSelector, "input[type='email']", timeout)
	if err != nil {
		return fmt.Errorf("failed to find email input field: %v", err)
	}

	// Enter the email
	err = emailInput.SendKeys(email)
	if err != nil {
		return fmt.Errorf("failed to enter email: %v", err)
	}

	// Wait for the password input field to be clickable
	passwordInput, err := WaitForClickable(driver, selenium.ByCSSSelector, "input[type='password']", timeout)
	if err != nil {
		return fmt.Errorf("failed to find password input field: %v", err)
	}

	// Enter the password
	err = passwordInput.SendKeys(password)
	if err != nil {
		return fmt.Errorf("failed to enter password: %v", err)
	}

	// Press the Enter key after entering the password
	err = passwordInput.SendKeys(selenium.EnterKey)
	if err != nil {
		return fmt.Errorf("failed to press Enter key: %v", err)
	}

	return nil
}
