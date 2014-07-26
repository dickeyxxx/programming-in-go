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

package m3u2pls

import (
	"fmt"
	"log"
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

func ReadM3uPlaylist(data string) (songs []Song) {
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
	return
}

func splitLine(line string) []string {
	return strings.Split(line, "=")
}

func ReadPlsPlaylist(data string) (songs []Song) {
	for _, line := range strings.Split(data, "\n") {
		if strings.HasPrefix(line, "NumberOfEntries") {
			numSongs, _ := strconv.Atoi(splitLine(line)[1])
			songs = make([]Song, numSongs)
			break
		}
	}

	regex := regexp.MustCompile(`([a-zA-Z]+)(\d+)=(.*)`)
	var i int
	for _, line := range strings.Split(data, "\n") {
		if parsedLine := regex.FindStringSubmatch(line); parsedLine != nil {
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
	return
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

func WritePlsPlaylist(songs []Song) string {
	lines := []string{"[playlist]"}
	for i, song := range songs {
		i++
		lines = append(lines, fmt.Sprintf("File%d=%s", i, song.Filename))
		lines = append(lines, fmt.Sprintf("Title%d=%s", i, song.Title))
		lines = append(lines, fmt.Sprintf("Length%d=%d", i, song.Seconds))
	}
	lines = append(lines, fmt.Sprintf("NumberOfEntries=%d", len(songs)))
	lines = append(lines, "Version=2")
	return strings.Join(lines, "\n")
}

func WriteM3uPlaylist(songs []Song) string {
	lines := []string{"#EXTM3U"}
	for _, song := range songs {
		lines = append(lines, fmt.Sprintf("#EXTINF:%d,%s", song.Seconds, song.Title))
		lines = append(lines, song.Filename)
	}
	return strings.Join(lines, "\n")
}
