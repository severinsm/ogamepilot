package controllers

import (
	"fmt"
	"ogame-bot/utils"
	"time"

	"github.com/tebeka/selenium"
)

func Authenticate(driver selenium.WebDriver, email, password string, timeout time.Duration) error {
	utils.Sleep(2)
	// Accept the cookie consent banner
	if err := utils.AcceptCookieBanner(driver); err != nil {
		fmt.Println("Error accepting cookie consent banner:", err)
		return err
	}

	// Click on the login tab
	if err := utils.ClickLoginTab(driver, timeout); err != nil {
		fmt.Println("Error clicking login tab:", err)
		return err
	}

	// Enter credentials using the utils.EnterCredentials function
	if err := utils.EnterCredentials(driver, email, password, timeout); err != nil {
		fmt.Println("Error entering credentials:", err)
		return err
	}
	utils.Sleep(8)
	return nil
}
