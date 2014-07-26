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
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestReadM3uPlaylist(t *testing.T) {
	Convey("It reads all the songs in", t, func() {
		songs := ReadM3uPlaylist(M3U)
		So(songs, ShouldResemble, Songs)
	})
}

func TestReadPlsPlaylist(t *testing.T) {
	Convey("It reads all the songs in", t, func() {
		songs := ReadPlsPlaylist(Pls)
		So(songs, ShouldResemble, Songs)
	})
}

func TestWritePlsPlaylist(t *testing.T) {
	Convey("It writes the songs to the buffer", t, func() {
		pls := WritePlsPlaylist(Songs)
		So(pls, ShouldEqual, Pls)
	})
}

func TestWriteM3uPlaylist(t *testing.T) {
	Convey("It writes the songs to the buffer", t, func() {
		m3u := WriteM3uPlaylist(Songs)
		So(m3u, ShouldEqual, M3U)
	})
}

const M3U = `#EXTM3U
#EXTINF:315,David Bowie - Space Oddity
Music/David Bowie/Singles 1/01-Space Oddity.ogg
#EXTINF:-1,David Bowie - Changes
Music/David Bowie/Singles 1/02-Changes.ogg
#EXTINF:258,David Bowie - Starman
Music/David Bowie/Singles 1/03-Starman.ogg`

var Songs = []Song{
	{"David Bowie - Space Oddity",
		"Music/David Bowie/Singles 1/01-Space Oddity.ogg", 315},
	{"David Bowie - Changes",
		"Music/David Bowie/Singles 1/02-Changes.ogg", -1},
	{"David Bowie - Starman",
		"Music/David Bowie/Singles 1/03-Starman.ogg", 258},
}

const Pls = `[playlist]
File1=Music/David Bowie/Singles 1/01-Space Oddity.ogg
Title1=David Bowie - Space Oddity
Length1=315
File2=Music/David Bowie/Singles 1/02-Changes.ogg
Title2=David Bowie - Changes
Length2=-1
File3=Music/David Bowie/Singles 1/03-Starman.ogg
Title3=David Bowie - Starman
Length3=258
NumberOfEntries=3
Version=2`
