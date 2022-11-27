package main

import (
	"encoding/xml"
	"fmt"
	"io"
	"os"
	"sort"
	"strings"
)

type Entry struct {
	Type       string `xml:"type,attr"`
	Title      string `xml:"title"`
	Artist     string `xml:"artist"`
	Album      string `xml:"album"`
	Duration   uint64 `xml:"duration"`
	Location   string `xml:"location"`
	Mtime      uint64 `xml:"mtime"`
	FirstSeen  uint64 `xml:"first-seen"`
	LastSeen   uint64 `xml:"last-seen"`
	PlayCount  uint64 `xml:"play-count"`
	LastPlayed uint64 `xml:"last-played"`
}

type DB struct {
	XMLName xml.Name `xml:"rhythmdb"`
	Version string   `xml:"version,attr"`
	Entries []Entry  `xml:"entry"`
}

type By func(e1, e2 *Entry) bool

type entrySorter struct {
	entries []Entry
	by      func(e1, e2 *Entry) bool
}

func (es *entrySorter) Len() int {
	return len(es.entries)
}

func (es *entrySorter) Swap(i, j int) {
	es.entries[i], es.entries[j] = es.entries[j], es.entries[i]
}

func (es *entrySorter) Less(i, j int) bool {
	return es.by(&es.entries[i], &es.entries[j])
}

func (by By) Sort(entries []Entry) {
	es := &entrySorter{
		entries: entries,
		by:      by,
	}

	sort.Sort(es)
}

func main() {
	data, err := io.ReadAll(os.Stdin)
	if err != nil {
		panic(err)
	}

	db := DB{}
	err = xml.Unmarshal(data, &db)
	if err != nil {
		panic(err)
	}

	playCount := func(e1, e2 *Entry) bool {
		return e1.PlayCount < e2.PlayCount
	}

	playCountDecreasing := func(e1, e2 *Entry) bool {
		return playCount(e2, e1)
	}

	By(playCountDecreasing).Sort(db.Entries)

	for _, e := range db.Entries {
		if !strings.EqualFold(e.Type, "song") {
			continue
		}
		fmt.Printf("%s\n", e.Location)
	}

}
