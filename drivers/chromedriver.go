// drivers/chromedriver.go
package drivers

import (
	"github.com/tebeka/selenium"
	"github.com/tebeka/selenium/chrome"
)

func StartChromeDriverService() (*selenium.Service, error) {
	// Start ChromeDriver service
	return selenium.NewChromeDriverService("chromedriver", 4444)
}

func CreateWebDriver(service *selenium.Service) (selenium.WebDriver, error) {
	// Configure Chrome options
	chromeCaps := chrome.Capabilities{
		Args: []string{
			"window-size=1920x1080",
			"--no-sandbox",
			"--disable-dev-shm-usage",
			"--disable-gpu",
			// "--headless",  // comment out this line to see the browser
		},
	}

	// Create WebDriver capabilities
	caps := selenium.Capabilities{}
	caps.AddChrome(chromeCaps)

	// Create WebDriver instance
	return selenium.NewRemote(caps, "")
}

func NavigateToPage(driver selenium.WebDriver, url string) error {
	// Navigate to a webpage
	return driver.Get(url)
}
