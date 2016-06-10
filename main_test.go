package main

import (
	"flag"
	"os"
	"testing"
)

//test main function
func TestMain(m *testing.M) {

	flag.Parse()
	os.Exit(m.Run())

}

//benchmark download
func BenchmarkDownloadImg(b *testing.B) {
	for n := 0; n < b.N; n++ {

		DownloadImg("https://source.unsplash.com/random")
	}

}
