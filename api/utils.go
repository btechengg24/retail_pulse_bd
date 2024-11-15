package api

import (
	"fmt"
	"net/http"
	"io/ioutil"
	"time"
	"math/rand"
)

// DownloadImage downloads the image from the URL and returns the content.
func DownloadImage(url string) ([]byte, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("failed to download image: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("failed to fetch image, status code: %d", resp.StatusCode)
	}

	return ioutil.ReadAll(resp.Body)
}

// CalculatePerimeter calculates the perimeter of an image.
func CalculatePerimeter(image []byte) int {
	// For this mock function, the perimeter is a random number.
	// In a real scenario, you'd use an image processing library to get image dimensions.
	return 2 * (len(image) + len(image)) // mock perimeter calculation
}

// Sleep imitates GPU processing time.
func Sleep() {
	// Random sleep time between 0.1 and 0.4 seconds
	time.Sleep(time.Duration(rand.Intn(300)+100) * time.Millisecond)
}
