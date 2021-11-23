package svatky

import (
	"bufio"
	"log"
	"os"
	"strings"
)

type Svatek struct {
	name string
	date string
}

var Svatky []Svatek

func init() {
	fileName := "svatky.csv"
	f, err := os.Open(fileName)
	if err != nil {
		log.Panic(err)
	}
	defer f.Close()

	lines := bufio.NewScanner(f)

	for lines.Scan() {
		line := lines.Text()
		//log.Println(line)
		lineItems := strings.Split(line, ";")
		name := lineItems[0]
		date := lineItems[1]

		dateItems := strings.Split(date, ".")
		month := dateItems[1]
		day := dateItems[0]

//		log.Printf("%s -> %s - %s %s ", name, date, day, month)
		dateShort := day + "." + month + "."
		Svatky = append(Svatky, Svatek{name, dateShort})
	}

	//log.Printf("%#v", Svatky)

}

func GetSvatekByDate(date string) string {

	var jmenaOut string

	for _, j := range Svatky {
		if j.date == date {
			if jmenaOut != "" {
				jmenaOut = jmenaOut + ", " + j.name
			} else {
				jmenaOut = j.name
			}
		}
	}
//	log.Println(jmenaOut)

	return jmenaOut
}
