package main

import (
	"fmt"
	"ogame-bot/controllers"
	"ogame-bot/drivers"
	"ogame-bot/initializers"
	"ogame-bot/utils"
	"os"
	"time"
)

const cookiesFilePath = "config/cookies.json"
const timeout = 6 * time.Second

func init() {
	initializers.LoadEnvVariables("config/settings.env")

}

func main() {

	// Start ChromeDriver service
	service, err := drivers.StartChromeDriverService()
	if err != nil {
		fmt.Println("Error starting ChromeDriver service:", err)
		os.Exit(1)
	}
	defer service.Stop()

	// Create WebDriver instance
	driver, err := drivers.CreateWebDriver(service)
	if err != nil {
		fmt.Println("Error creating WebDriver instance:", err)
		os.Exit(1)
	}
	defer driver.Quit()

	// Navigate to a webpage
	if err := drivers.NavigateToPage(driver, "https://lobby.ogame.gameforge.com/nl_NL/"); err != nil {
		fmt.Println("Error navigating to webpage:", err)
		os.Exit(1)
	}

	// Check if the cookies file exists and is not empty
	fileInfo, err := os.Stat(cookiesFilePath)
	if os.IsNotExist(err) || fileInfo.Size() == 0 {
		fmt.Println("Cookies file is empty. Skipping cookie reloading and trying to authenticate..")
		err := controllers.Authenticate(driver, os.Getenv("username"), os.Getenv("password"), timeout)
		if err != nil {
			fmt.Println("error login")
			os.Exit(1)
		}
	} else {
		fileContent, err := os.ReadFile(cookiesFilePath)
		if err != nil {
			fmt.Println("Error reading cookies file:", err)
			os.Exit(1)
		}

		// Check if the file content is "[]"
		if string(fileContent) == "[]" {
			fmt.Println("Cookies file is empty. Skipping cookie reloading and trying to authenticate..")
			err := controllers.Authenticate(driver, os.Getenv("username"), os.Getenv("password"), timeout)
			if err != nil {
				fmt.Println("error login")
				os.Exit(1)
			}
		} else {
			// File exists and is non-empty, attempt to reload cookies
			err = utils.ReloadCookiesFromJSON(driver, cookiesFilePath)
			if err != nil {
				fmt.Println("Error reloading cookies from file:", err)
				return
			}
			fmt.Println("Cookies reloaded into WebDriver successfully.")
		}
	}

	// Click on the "Play" button
	if err := controllers.ClickPlayButton(driver, timeout); err != nil {
		fmt.Println("Error clicking play button:", err)
		os.Exit(1)
	}
	if err := controllers.ClickGalaxy(os.Getenv("galaxy"), driver, timeout); err != nil {
		fmt.Println("Error clicking play button:", err)
		os.Exit(1)
	}

	utils.Sleep(3)

	// Get cookies from WebDriver
	cookies, err := driver.GetCookies()
	if err != nil {
		fmt.Println("Error getting cookies from WebDriver:", err)
		return
	}

	// Save cookies to a JSON file using the function from utils package
	err = utils.SaveCookiesToFile(cookies, cookiesFilePath)
	if err != nil {
		fmt.Println("Error saving cookies to file:", err)
		return
	}
	fmt.Println("Cookies saved to file successfully.")
	utils.Sleep(20)
	// Close the browser window
	driver.Close()
}
