package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

var missions []*Mission
var asJson = flag.Bool("json", false, "Output JSON instead of tab delimited")

func main() {
	fmt.Println("starting...")
	flag.Parse()
	var doc *goquery.Document
	var err error

	if doc, err = goquery.NewDocument("http://science.nasa.gov/missions/?group=all"); err != nil {
		log.Fatal("Failed to fetch page")
	}
	doc.Find(".missions").Find("tbody").Children().Each(func(i int, s *goquery.Selection) {
		m := unpackMission(s)
		if m.Phase == "Operating" {
			missions = append(missions, m)
		}
	})

	if *asJson == true {
		b, err := json.Marshal(missions)
		if err != nil {
			log.Fatal(err)
		}
		os.Stdout.Write(b)
	} else {
		for _, m := range missions {
			fmt.Println(m)
		}
	}
}

func unpackMission(s *goquery.Selection) *Mission {
	m := Mission{}
	tds := s.Children()
	r, err := tds.First().Html()
	if err != nil {
		log.Printf("Error parsing HTML: %+v\n", err)
	} else {
		m.Division = r
	}
	node := tds.Next().Children()
	name, err := node.Html()
	if err != nil {
		log.Println("Error getting name: ", err)
	}
	m.Name = strings.TrimSpace(name)
	href, ok := node.Attr("href")
	if !ok {
		log.Println("No href")
	}
	m.Url = href
	node = tds.Next()
	desc, err := node.Find(".desc").Children().Html()
	if err != nil {
		log.Println("Err getting desc", err)
	}
	m.Description = strings.TrimSpace(desc)
	node = tds.Next()
	date, err := node.Next().Children().Html()
	m.LaunchDate = date
	date2 := strings.Trim(node.Next().First().Text(), "1234567890")
	m.LaunchDateHuman = strings.TrimSpace(date2)
	m.Phase = strings.TrimLeft(tds.Last().Text(), "1234567890")
	return &m
}

// Mission represents a mission table row
type Mission struct {
	Division        string
	Name            string
	Url             string
	Description     string
	LaunchDate      string
	LaunchDateHuman string
	Phase           string
}

// String returns a tab delimited representation of a mission table row
func (m *Mission) String() string {
	return fmt.Sprintf("%s\t%s\t%s\t\"%s\"\t%s\t\"%s\"\t%s",
		m.Division, m.Name, m.Url, m.Description, m.LaunchDate, m.LaunchDateHuman, m.Phase)
}

/*
<!-- Canonical row extracted from table for reference. -->

<tr class="odd" style="display: table-row;">
	<td>Astrophysics</td>
	<td scope="row" class="">
		<a href="/missions/kepler/">
		    Kepler
		</a>
    	<div class="desc">
    		<div>
    			The Kepler Mission, a NASA Discovery mission, is specifically designed to survey our region of the Milky Way Galaxy to detect and characterize hundreds of Earth-size and smaller planets in or nearby the habitable zone.
    		</div>
    	</div>
    </td>
    <td>
    	<span class="hide">
    		20090306
    	</span>
    		March 06, 2009
    </td>
    <td>
	    <span class="hide">
	    	3
	    </span>
	    Operating
    </td>
</tr>
*/
