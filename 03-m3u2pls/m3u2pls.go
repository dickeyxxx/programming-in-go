// Copyright © 2011-12 Qtrac Ltd.
//
// This program or package and any associated files are licensed under the
// Apache License, Version 2.0 (the "License"); you may not use these files
// except in compliance with the License. You can get a copy of the License
// at: http://www.apache.org/licenses/LICENSE-2.0.
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"
)

type Song struct {
	Title    string
	Filename string
	Seconds  int
}

func main() {
	if len(os.Args) == 1 || !strings.HasSuffix(os.Args[1], ".m3u") {
		fmt.Printf("usage: %s <file.m3u>\n", filepath.Base(os.Args[0]))
		os.Exit(1)
	}

	if rawBytes, err := ioutil.ReadFile(os.Args[1]); err != nil {
		log.Fatal(err)
	} else {
		songs := readM3uPlaylist(string(rawBytes))
		writePlsPlaylist(songs, os.Stdout)
	}
}

func readM3uPlaylist(data string) (songs []Song) {
	var song Song
	for _, line := range strings.Split(data, "\n") {
		line = strings.TrimSpace(line)
		if line == "" || strings.HasPrefix(line, "#EXTM3U") {
			continue
		}
		if strings.HasPrefix(line, "#EXTINF:") {
			song.Title, song.Seconds = parseExtinfLine(line)
		} else {
			song.Filename = strings.Map(mapPlatformDirSeparator, line)
		}
		if song.Filename != "" && song.Title != "" && song.Seconds != 0 {
			songs = append(songs, song)
			song = Song{}
		}
	}
	return songs
}

func splitLine(line string) []string {
	return strings.Split(line, "=")
}

func readPlsPlaylist(data string) (songs []Song) {
	for _, line := range strings.Split(data, "\n") {
		if strings.HasPrefix(line, "NumberOfEntries") {
			numSongs, _ := strconv.Atoi(splitLine(line)[1])
			songs = make([]Song, numSongs)
		}
	}

	regex := regexp.MustCompile(`([a-zA-Z]+)(\d+)=(.*)`)
	var i int
	for _, line := range strings.Split(data, "\n") {
		parsedLine := regex.FindStringSubmatch(line)
		if parsedLine != nil {
			if parsedLine[1] == "File" {
				i, _ = strconv.Atoi(parsedLine[2])
				songs[i-1].Filename = parsedLine[3]
			} else if parsedLine[1] == "Title" {
				i, _ = strconv.Atoi(parsedLine[2])
				songs[i-1].Title = parsedLine[3]
			} else if parsedLine[1] == "Length" {
				i, _ = strconv.Atoi(parsedLine[2])
				songs[i-1].Seconds, _ = strconv.Atoi(parsedLine[3])
			}
		}
	}
	return songs
}

func parseExtinfLine(line string) (title string, seconds int) {
	if i := strings.IndexAny(line, "-0123456789"); i > -1 {
		const separator = ","
		line = line[i:]
		if j := strings.Index(line, separator); j > -1 {
			title = line[j+len(separator):]
			var err error
			if seconds, err = strconv.Atoi(line[:j]); err != nil {
				log.Printf("failed to read the duration for '%s': %v\n",
					title, err)
				seconds = -1
			}
		}
	}
	return title, seconds
}

func mapPlatformDirSeparator(char rune) rune {
	if char == '/' || char == '\\' {
		return filepath.Separator
	}
	return char
}

func writePlsPlaylist(songs []Song, stdout io.Writer) {
	fmt.Fprintln(stdout, "[playlist]")
	for i, song := range songs {
		i++
		fmt.Fprintf(stdout, "File%d=%s\n", i, song.Filename)
		fmt.Fprintf(stdout, "Title%d=%s\n", i, song.Title)
		fmt.Fprintf(stdout, "Length%d=%d\n", i, song.Seconds)
	}
	fmt.Fprintf(stdout, "NumberOfEntries=%d\nVersion=2\n", len(songs))
}

func writeM3uPlaylist(songs []Song, stdout io.Writer) {
	fmt.Fprint(stdout, "#EXTM3U")
	for _, song := range songs {
		fmt.Fprintln(stdout)
		fmt.Fprintf(stdout, "#EXTINF:%d,%s\n", song.Seconds, song.Title)
		fmt.Fprint(stdout, song.Filename)
	}
}
