// S24b: A simple language guesser.
//
// OUTPUT:
//
//     $ go run main.go
//     en en
//     de de
//     it it
package main

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"strings"
)

var examples = map[string]string{
	"it": `
La morfosintassi dell'italiano è conforme al modello delle altre lingue
italo-occidentali in generale, possedendo un ricco sistema verbale e
configurandosi come lingua SVO. I nomi non possiedono distinzioni di caso.
`,
	"en": `
Italian often was an official language of the various Italian states predating
unification, slowly replacing Latin, even when ruled by foreign powers (such as
the Spanish in the Kingdom of Naples, or the Austrians in the Kingdom of
Lombardy-Venetia), even though the masses spoke primarily vernacular languages
and dialects. Italian was also one of the many recognised languages in the
Austro-Hungarian Empire.
`,
	"de": `
Die wirkliche Vereinheitlichung, besonders der gesprochenen Sprache, erfolgte
allerdings erst aufgrund der nationalen Einigung. Als italienische
Einheitssprache setzte sich im 19. Jahrhundert im vereinigten Italien der
„florentinische“ Dialekt durch. Zu verdanken ist dies unter anderem der zweiten
Fassung des Romans I Promessi Sposi von Alessandro Manzoni.
`,
}

// TrigramMap from the internet.
// * it: http://stefantrost.de, http://www.sttmedia.com/syllablefrequency-italian
// * en: http://norvig.com/mayzner.html
// * de: http://www.mathe.tu-freiberg.de/~hebisch/cafe/kryptographie/trigramme.html
//
// Early collection (1969): http://digitalcommons.butler.edu/wordways/vol2/iss3/17/
var TrigramMap = map[string][]string{
	"it": []string{
		"ale", "all", "anc", "and", "ant", "are", "ato", "att", "che", "chi", "com", "con",
		"del", "ell", "ent", "era", "ere", "ess", "est", "ett", "gli", "ion", "lla", "men",
		"non", "nte", "nti", "nto", "olo", "one", "ono", "per", "que", "son", "sta", "ver",
	},
	"en": []string{
		"the", "ing", "and", "her", "ere", "ent", "hat", "tha", "nth", "was", "eth", "for",
		"dth", "his", "ion", "ter", "you", "ith", "ver", "all", "wit", "thi", "tio", "eve",
		"ate", "con", "nce", "ted", "ive", "sta", "cti", "ess", "not", "iti", "rat", "one",
	},
	"de": []string{
		"ein", "ich", "nde", "die", "und", "der", "che", "end", "gen", "sch", "cht", "den",
		"ine", "nge", "nun", "ung", "das", "hen", "ind", "enw", "ens", "ies", "ste", "ten",
		"ere", "lic", "ach", "ndi", "sse", "aus", "ers", "ebe", "erd", "enu", "nen", "rau",
	},
}

// TrigramGuesser guesses the language of a byte stream.
type TrigramGuesser struct {
	r           io.Reader
	trigramFreq map[string]uint64
}

// New creates a trigramGuesser.
func New(r io.Reader) *TrigramGuesser {
	return &TrigramGuesser{r: r, trigramFreq: make(map[string]uint64)}
}

// Read counts overlapping trigrams in byte stream.
func (r *TrigramGuesser) Read(p []byte) (n int, err error) {
	n, err = r.r.Read(p)
	for i := 0; i < n-3; i++ {
		r.trigramFreq[string(p[i:i+3])]++
	}
	return
}

// String prints the guessers trigram frequencies.
func (r *TrigramGuesser) String() string {
	var buf bytes.Buffer
	for k, v := range r.trigramFreq {
		io.WriteString(&buf, fmt.Sprintf("%s -> %d\n", k, v))
	}
	return buf.String()
}

// Guess makes a guess.
// Exercise: Return list ordered by "confidence".
func (r *TrigramGuesser) Guess() string {
	var guess string
	var best int
	for lang, trigrams := range TrigramMap {
		var hits int
		for _, g := range trigrams {
			if _, ok := r.trigramFreq[g]; ok {
				hits++
			}
		}
		if hits > best {
			best, guess = hits, lang
		}
	}
	return guess
}

func main() {
	for k, v := range examples {
		// Exercise: Rewrite and test it on the contents of
		// http://www.corriere.it/, https://www.nytimes.com/, http://www.sueddeutsche.de/.
		r := New(strings.NewReader(v))
		if _, err := io.Copy(ioutil.Discard, r); err != nil {
			log.Fatal(err)
		}
		fmt.Printf("%s %s\n", k, r.Guess())
	}
}
