package main

import (
	"encoding/xml"
	"fmt"
	"io"
	"os"
	"time"

	"Z_Interview/vast"
)

func loadFixture(path string) (*vast.VAST, string, string, error) {
	xmlFile, err := os.Open(path)
	if err != nil {
		return nil, "", "", err
	}
	defer xmlFile.Close()
	b, _ := io.ReadAll(xmlFile)

	var v vast.VAST
	err = xml.Unmarshal(b, &v)
	if err != nil {
		return nil, "", "", err
	}
	res, err := xml.MarshalIndent(v, "", "\t") // can use two spaces, four spaces or tab.
	if err != nil {
		return nil, "", "", err

	}
	return &v, string(b), string(res), err
}
func main() {
	v, _, _, err := loadFixture("testdata/vast_inline_linear.xml")
	//fmt.Println("b: ", b)
	//fmt.Println("res: ", res)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("v: ", v)

	live, _, res, err := loadFixture("testdata/liverail-vast2-linear-companion.xml")
	fmt.Println("live: ", live)
	_ = res
	if err != nil {
		fmt.Println(err)
	}

	ad := live.Ads[0]
	adTitle := ad.InLine.AdTitle
	fmt.Println(adTitle)
	fmt.Println(adTitle.CDATA)

	timeStamp := "00:00:01.666"
	var du vast.Duration
	du.UnmarshalText([]byte(timeStamp))

	fmt.Println(time.Duration(du).Seconds())

	d := time.Duration(12*time.Hour) + time.Duration(12*time.Second)
	durText, _ := vast.Duration(d).MarshalText()
	fmt.Println(string(durText))
}
