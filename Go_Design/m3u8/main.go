package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

type Playlist struct {
	Version     int
	MediaSeq    int
	TSDurations []float64
	TSList      []string
}

func main() {
	m3u8File, err := os.Open("testdata/manifest.m3u8")
	if err != nil {
		fmt.Println("Error fetching playlist:", err)
		return
	}
	defer m3u8File.Close()

	playlist, err := parseM3U8(m3u8File)
	if err != nil {
		fmt.Println("Error parsing playlist:", err)
		return
	}

	fmt.Printf("Playlist Version: %d\n", playlist.Version)
	fmt.Printf("Media Sequence: %d\n", playlist.MediaSeq)
	for i, segment := range playlist.TSList {
		fmt.Println(playlist.TSDurations[i])
		fmt.Println(segment)
	}
}

func parseM3U8(r io.Reader) (*Playlist, error) {
	scanner := bufio.NewScanner(r)
	var version int
	var mediaSeq int
	var tsDurations []float64
	var tsList []string

	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "" || strings.HasPrefix(line, "#EXTM3U") {
			continue
		} else if strings.HasPrefix(line, "#EXT-X-VERSION:") {
			version, _ = strconv.Atoi(strings.TrimPrefix(line, "#EXT-X-VERSION:"))
		} else if strings.HasPrefix(line, "#EXT-X-MEDIA-SEQUENCE:") {
			mediaSeq, _ = strconv.Atoi(strings.TrimPrefix(line, "#EXT-X-MEDIA-SEQUENCE:"))
		} else if strings.HasPrefix(line, "#EXTINF:") {
			line = strings.TrimPrefix(line, "#EXTINF:")
			line = strings.TrimSuffix(line, ",")
			segmentDuration, err := strconv.ParseFloat(line, 32)
			if err != nil {
				fmt.Println(err)
			}
			tsDurations = append(tsDurations, float64(segmentDuration))
		} else if !strings.HasPrefix(line, "#") {
			tsList = append(tsList, line)
		}
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return &Playlist{
		Version:     version,
		MediaSeq:    mediaSeq,
		TSDurations: tsDurations,
		TSList:      tsList,
	}, nil
}
