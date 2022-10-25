package twelve

var verse = map[string]string{
	"first":    "a Partridge in a Pear Tree.",
	"twelfth":  "twelve Drummers Drumming",
	"eleventh": "eleven Pipers Piping",
	"tenth":    "ten Lords-a-Leaping",
	"ninth":    "nine Ladies Dancing",
	"eighth":   "eight Maids-a-Milking",
	"seventh":  "seven Swans-a-Swimming",
	"sixth":    "six Geese-a-Laying",
	"fifth":    "five Gold Rings",
	"fourth":   "four Calling Birds",
	"third":    "three French Hens",
	"second":   "two Turtle Doves",
}

var wording = map[int]string{
	1:  "first",
	2:  "second",
	3:  "third",
	4:  "fourth",
	5:  "fifth",
	6:  "sixth",
	7:  "seventh",
	8:  "eighth",
	9:  "ninth",
	10: "tenth",
	11: "eleventh",
	12: "twelfth",
}

func Verse(i int) string {

	var (
		firstSentenceStr  = "On the "
		secondSentenceStr = " day of Christmas my true love gave to me, "
		str               = ""
	)

	for j := 1; j <= i; j++ {
		isAnd := func(x int) string {
			if j == 2 {
				return "and "
			}
			return ""
		}
		isDot := func(x int) string {
			if j == 1 {
				return ""
			}
			return ", "
		}
		str = verse[wording[j]] + isDot(j) + isAnd(j) + str
	}

	return firstSentenceStr + wording[i] + secondSentenceStr + str
}

func Song() string {
	var song = ""

	for i := 1; i <= len(wording); i++ {
		song = song + Verse(i) + "\n"
	}

	return song
}
