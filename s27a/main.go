// S27a: BlackBar censors given words in a stream.
package main

import (
	"io"
	"log"
	"os"
	"strings"
)

var s = `
One morning, when Gregor Samsa woke from troubled dreams, he found
himself transformed in his bed into a horrible vermin.  He lay on
his armour-like back, and if he lifted his head a little he could
see his brown belly, slightly domed and divided by arches into stiff
sections.  The bedding was hardly able to cover it and seemed ready
to slide off any moment.  His many legs, pitifully thin compared
with the size of the rest of him, waved about helplessly as he
looked.

"What's happened to me?" he thought.  It wasn't a dream.  His room,
a proper human room although a little too small, lay peacefully
between its four familiar walls.  A collection of textile samples
lay spread out on the table - Samsa was a travelling salesman - and
above it there hung a picture that he had recently cut out of an
illustrated magazine and housed in a nice, gilded frame.  It showed
a lady fitted out with a fur hat and fur boa who sat upright,
raising a heavy fur muff that covered the whole of her lower arm
towards the viewer.

Gregor then turned to look out the window at the dull weather.
Drops of rain could be heard hitting the pane, which made him feel
quite sad.  "How about if I sleep a little bit longer and forget all
this nonsense", he thought, but that was something he was unable to
do because he was used to sleeping on his right, and in his present
state couldn't get into that position.  However hard he threw
himself onto his right, he always rolled back to where he was.  He
must have tried it a hundred times, shut his eyes so that he
wouldn't have to look at the floundering legs, and only stopped when
he began to feel a mild, dull pain there that he had never felt
before.
`

// BlackBar blacks out words from stream.
type BlackBar struct {
	r        io.Reader
	replacer *strings.Replacer
}

// NewReader constructs a censoring reader.
func NewReader(r io.Reader, words []string) *BlackBar {
	return &BlackBar{r: r, replacer: makeReplacer(words)}
}

// makeReplacer replaces each word given with a blacked-out counterpart. The
// number of bytes won't be changed.
func makeReplacer(words []string) *strings.Replacer {
	var s []string
	for _, w := range words {
		s = append(s, w)

		block := strings.Repeat("█", len(w)/3)
		switch len(w) % 3 {
		case 1:
			block = block + "X"
		case 2:
			block = block + "XX"
		}
		s = append(s, block)
	}
	return strings.NewReplacer(s...)
}

// blackout erases blacklisted words.
func (r *BlackBar) blackout(p []byte) []byte {
	return []byte(r.replacer.Replace(string(p)))
}

// Read censors the underlying stream.
func (r *BlackBar) Read(p []byte) (n int, err error) {
	n, err = r.r.Read(p)
	copy(p, r.blackout(p))
	return
}

func main() {
	words := []string{"Gregor", "Samsa", "travelling salesman"}
	r := NewReader(strings.NewReader(s), words)
	if _, err := io.Copy(os.Stdout, r); err != nil {
		log.Fatal(err)
	}
}

// $ go run main.go
// One morning, when ██ █XX woke from troubled dreams, he found
// himself transformed in his bed into a horrible vermin.  He lay on
// his armour-like back, and if he lifted his head a little he could
// see his brown belly, slightly domed and divided by arches into stiff
// sections.  The bedding was hardly able to cover it and seemed ready
// to slide off any moment.  His many legs, pitifully thin compared
// with the size of the rest of him, waved about helplessly as he
// looked.

// "What's happened to me?" he thought.  It wasn't a dream.  His room,
// a proper human room although a little too small, lay peacefully
// between its four familiar walls.  A collection of textile samples
// lay spread out on the table - █XX was a ██████X - and
// above it there hung a picture that he had recently cut out of an
// illustrated magazine and housed in a nice, gilded frame.  It showed
// a lady fitted out with a fur hat and fur boa who sat upright,
// raising a heavy fur muff that covered the whole of her lower arm
// towards the viewer.

// ██ then turned to look out the window at the dull weather.
// Drops of rain could be heard hitting the pane, which made him feel
// quite sad.  "How about if I sleep a little bit longer and forget all
// this nonsense", he thought, but that was something he was unable to
// do because he was used to sleeping on his right, and in his present
// state couldn't get into that position.  However hard he threw
// himself onto his right, he always rolled back to where he was.  He
// must have tried it a hundred times, shut his eyes so that he
// wouldn't have to look at the floundering legs, and only stopped when
// he began to feel a mild, dull pain there that he had never felt
// before.
