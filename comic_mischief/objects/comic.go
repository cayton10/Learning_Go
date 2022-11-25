package comic

import f "fmt"

type Comic struct {
	Publisher  string
	Writer     string
	Artist     string
	Title      string
	Year       int
	PageNumber int8
	Grade      float32
}

func PrintComicInfo(comic Comic) {
	f.Println(comic.Title, "written by", comic.Writer, "drawn by", comic.Artist)
}
