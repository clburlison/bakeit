package client

import (
	"fmt"
	"os"
	"time"

	"github.com/cavaliercoder/grab"
)

// Download a web resource; return errors
func Download(downloadURL string, savePath string) (msg string, err error) {
	// create client
	client := grab.NewClient()
	// req, _ := grab.NewRequest(".", "https://packages.chef.io/files/stable/chef/13.6.4/mac_os_x/10.13/chef-13.6.4-1.dmg")
	// req, _ := grab.NewRequest("./munki.pkg", "https://github.com/munki/munki/releases/download/v3.1.1/munkitools-3.1.1.3447.pkg")
	req, _ := grab.NewRequest(savePath, downloadURL)

	// start download
	fmt.Printf("Downloading %v...\n", req.URL())
	resp := client.Do(req)
	fmt.Printf("  %v\n", resp.HTTPResponse.Status)

	// start UI loop
	t := time.NewTicker(500 * time.Millisecond)
	defer t.Stop()

Loop:
	for {
		select {
		case <-t.C:
			fmt.Printf("  transferred %v / %v bytes (%.2f%%)\n",
				resp.BytesComplete(),
				resp.Size,
				100*resp.Progress())

		case <-resp.Done:
			// download is complete
			break Loop
		}
	}

	// check for errors
	if err := resp.Err(); err != nil {
		fmt.Fprintf(os.Stderr, "Download failed: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("Download saved to ./%v \n", resp.Filename)
	return resp.Filename, nil
}
