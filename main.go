package main
//Random Wallpapers
//William Kwasidev

import ("fmt"
		"os/exec"
		"os"
		"time"
		"net/http"
		"log"
		"io"
		"bufio"
    "flag"

)

const URL string = "https://source.unsplash.com/random"
const FILENAME string = "/tmp/randompic"
const DEFAULT time.Duration = time.Second * 30
//Downloads Image and saves it to /tmp
func DownloadImg(url string){
    resp,err := http.Get(URL)
    if err != nil{

    	log.Fatal("Failed to Download ... Problem with your internet connection")
    	return
    }
   //buffer := make([]byte,1024)
   f,err := os.Create(FILENAME)
   if err != nil{

   	log.Fatal(err)

   	return
   }

   defer  resp.Body.Close()
   writer := bufio.NewWriter(f)
   reader := bufio.NewReader(resp.Body)
   //copy download to file
   written,err := io.Copy(writer,reader)

   if err != nil{

   	log.Fatal(err)
   }
   log.Printf("%d bytes Downloaded",written)

}

func main(){
 //how often it should change the Wallpaper
  interval := flag.Duration("interval",DEFAULT,"-interval how often do you want to change the Wallpaper?")
  flag.Parse()

  //tick every 20seconds ,return a go channel
	change := time.Tick(*interval);
	for now := range change {
		DownloadImg(URL)
		fmt.Println(now,"Setting new Wallpaper")
		//change background image
		err := exec.Command("gsettings","set","org.gnome.desktop.background","picture-uri",fmt.Sprintf("file://%s",FILENAME)).Run()

		//terminate if error occured
		if err != nil{

			log.Fatal("Fail to set Wallpaper")
		}

	}

}
