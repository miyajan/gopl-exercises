package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Metre float64
type Mile float64

func main() {
	var inputs []string
	if len(os.Args) > 1 {
		inputs = os.Args[1:]
	} else {
		scanner := bufio.NewScanner(os.Stdin)
		for scanner.Scan() {
			txt := scanner.Text()
			inputs = append(inputs, strings.Fields(txt)...)
		}
	}

	for _, arg := range inputs {
		t, err := strconv.ParseFloat(arg, 64)
		if err != nil {
			fmt.Fprintf(os.Stderr, "mm: %v\n", err)
			os.Exit(1)
		}
		metre := Metre(t)
		mile := Mile(t)
		fmt.Printf("%s = %s, %s = %s\n", metre, metreToMile(metre), mile, mileToMetre(mile))
	}
}

func (m Metre) String() string { return fmt.Sprintf("%gm", m) }
func (m Mile) String() string  { return fmt.Sprintf("%gmile", m) }

func metreToMile(meter Metre) Mile { return Mile(meter / 1609.344) }

func mileToMetre(mile Mile) Metre { return Metre(mile * 1609.344) }
