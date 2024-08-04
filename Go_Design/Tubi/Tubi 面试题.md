# Tubi 面试题

https://gist.githubusercontent.com/rob-brown/e261dde3a78f4b5e3673d65f55ff6ff5/raw/f54a619bcb4161e6256227f77babb4b18a19f6c9/access.log
Log Parse Coding Challenge

For this coding challenge, you may use whatever language, libraries, tools, references, and other resources you like.

As you work through the problem, talk out your ideas and explain what you are doing. This will help us understand your thought process and help you work out problems sooner.

Download the following log file.

https://gist.githubusercontent.com/rob-brown/e261dde3a78f4b5e3673d65f55ff6ff5/raw/f54a619bcb4161e6256227f77babb4b18a19f6c9/access.log

Your task is to write a script to do the following:

1. Parse the log file (standard AWS access log)
2. Calculate the TCP_HIT percentage per video ID
3. Sort the results by video ID (video id is an integer)
4. Print the results to the console or write to a file
5. Add tests if there is still time
A response will indicate `TCP_HIT` or `TCP_MISS` along with the status code, ex. `TCP_HIT/206`. You can ignore the status code.

The video ID is embedded in the URLs. There are a couple formats to catch.

In the following URL the video ID is `275211`.

http://interview.tubi.tv/04C0BF/v2/sources/content-owners/any-enter-films/275211/v0401185814-1389k.mp4+740005.ts

For the next URL the video ID is `2828870`.

http://interview.tubi.tv/04C0BF/ads/transcodes/008563/2828870/v0413163056-640x360-SD-592k.mp4_1.m3u8

Have fun and thank you for interviewing with Tubi.


```go
package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"sort"
	"strconv"
	"strings"
)

type TCPCount struct {
	TCPHit  int
	TCPMIss int
}

func ParseLog(r io.Reader) map[int]*TCPCount {
	res := make(map[int]*TCPCount)
	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		videoID, HasTCPHit, HasTCPMiss, err := ParseOnceLine(line)
		//fmt.Println("videoId:", videoID, "HasTCPHit:", HasTCPHit, "HasTCPMiss:", HasTCPMiss)
		if err != nil {
			fmt.Println("Parse one line failed: ", err)
			continue
		}

		if v, ok := res[videoID]; ok {
			if HasTCPHit {
				v.TCPHit += 1
			} else if HasTCPMiss {
				v.TCPMIss += 1
			}
		} else {
			tcpCount := &TCPCount{
				TCPHit:  0,
				TCPMIss: 0,
			}
			if HasTCPHit {
				tcpCount.TCPHit += 1
			} else {
				tcpCount.TCPMIss += 1
			}
			res[videoID] = tcpCount
		}
	}

	return res
}

func ParseOnceLine(str string) (int, bool, bool, error) {
	strArray := strings.Split(str, " ")
	var videoID int
	var HasTCPHit, HasTCPMiss bool
	var err error
	for _, v := range strArray {
		if strings.HasPrefix(v, "http://") {
			videoID, err = ParseURL(v)
			//fmt.Println("videoId:", videoID)
			if err != nil {
				// fmt.Println("error: ", videoID, false, false, err)
				continue
			}
		}
		if strings.HasPrefix(v, "TCP_HIT") {
			HasTCPHit = true
		} else if strings.HasPrefix(v, "TCP_MISS") {
			HasTCPMiss = true
		}
	}
	//fmt.Println("videoId:", videoID, "HasTCPHit:", HasTCPHit, "HasTCPMiss:", HasTCPMiss)

	return videoID, HasTCPHit, HasTCPMiss, nil
}

func ParseURL(str string) (int, error) {
	strArray := strings.Split(str, "/")
	var videoID int
	var err error
	if len(strArray) > 8 {
		videoID, err = strconv.Atoi(strArray[8])
		if err != nil {
			return -1, err
		}
	}
	//fmt.Println("VideoID: ", videoID)
	return videoID, nil
}

func main() {
	logFile, err := os.Open("testdata/test.log")
	if err != nil {
		fmt.Println("open log faield: ", err)
		return
	}
	defer logFile.Close()
	res := ParseLog(logFile)
	var ratio float64

	arr := []int{}
	for k, _ := range res {
		arr = append(arr, k)
	}
	sort.Slice(arr, func(i, j int) bool {
		return arr[i] < arr[j]
	})

	for _, k := range arr {
		if (res[k].TCPHit + res[k].TCPMIss) > 0 {
			ratio = float64(res[k].TCPHit) / float64(res[k].TCPHit+res[k].TCPMIss)
		}
		if k == 178315 || k == 268547 || k == 285897 || k == 391687 || k == 314152 || k == 409830 {
			fmt.Println("VideoID: ", k, "Ratio: ", ratio)
		}
	}

}
```