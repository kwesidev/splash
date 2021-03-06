package main

//Random Wallpapers for your Desktop
import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"os/exec"
	"strings"
	"time"
)

//constants
const (
	URL      string        = "https://source.unsplash.com/random"
	FILENAME string        = "/tmp/randompic"
	DEFAULT  time.Duration = time.Second * 30
)

//getScreenDimension returns the dimension of the screen
//this only works on linux systems
func getScreenDimension() string {
	command := "xdpyinfo | grep dimension | awk '{ print $2 }'"
	dim, err := exec.Command("sh", "-c", command).Output()
	if err != nil {
		log.Fatal(err)

	}
	//remove line feed
	return strings.Trim(string(dim), "\x0a")
}

//DownloadImg download images from the internet
//A url needs to be provided to download the image
func DownloadImg(url string) {
	resp, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
		return
	}
	//buffer := make([]byte,1024)
	f, err := os.Create(FILENAME)
	if err != nil {
		log.Fatal(err)

		return
	}

	defer resp.Body.Close()
	writer := bufio.NewWriter(f)
	reader := bufio.NewReader(resp.Body)
	//copy download to file
	written, err := io.Copy(writer, reader)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("%d bytes Downloaded", written)

}

func main() {
	var SCREEN = getScreenDimension()

	interval := flag.Duration("interval", DEFAULT, "-interval ")
	flag.Parse()
	//tick every 30seconds ,return a go channel
	change := time.Tick(*interval)
	for now := range change {
		link := URL + "/" + SCREEN
		DownloadImg(link)
		fmt.Println(now, "Setting new Wallpaper")
		//change background image
		err := exec.Command("gsettings", "set", "org.gnome.desktop.background", "picture-uri", fmt.Sprintf("file://%s", FILENAME)).Run()

		//terminate if error occured
		if err != nil {

			log.Fatal("Fail to set Wallpaper")
		}

	}

}
