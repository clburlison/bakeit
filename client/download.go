package client

import (
	"fmt"
	"time"

	"github.com/cavaliercoder/grab"
)

// Download a web resource; return file location and errors
func Download(downloadURL string, savePath string, progress bool) (msg string, err error) {
	// create client
	client := grab.NewClient()
	req, _ := grab.NewRequest(savePath, downloadURL)

	// start download
	if progress {
		fmt.Printf("Downloading %v...\n", req.URL())
	}
	resp := client.Do(req)
	if progress {
		fmt.Printf("  %v\n", resp.HTTPResponse.Status)
	}

	// start UI loop
	t := time.NewTicker(500 * time.Millisecond)
	defer t.Stop()

Loop:
	for {
		select {
		case <-t.C:
			if progress {
				fmt.Printf("  transferred %v / %v bytes (%.2f%%)\n",
					resp.BytesComplete(),
					resp.Size,
					100*resp.Progress())
			}
		case <-resp.Done:
			// download is complete
			break Loop
		}
	}

	// check for errors
	if err := resp.Err(); err != nil {
		return "", err
	}

	return resp.Filename, nil
}
