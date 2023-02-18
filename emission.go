// EMISSION -- do calculations for valve TESTER
//
// svm 17-FEB-2023 - 18-FEB-2023
//
// a. in valve_info and prpr(), clearly distinguish typical, testing and max values
//

package main

import (
	"fmt"
	"github.com/dsnet/golib/unitconv"
	"math"
	"os"
	"strings"
)

var valve_info = []struct {
	name      string
	bias      float64 // typical, not max
	max_pwr   float64
	max_curr  float64
	max_anode float64
}{
	{"EL84", -14, 12, 0.065, 300},
	{"EL34", -36, 25, 0.150, 800},
	{"ECC81", -2, 2.5, 0.015, 300}, // -2?
	{"ECC82", -8.5, 2.75, 0.020, 300},
	{"ECC83", -2, 1, 0.008, 300},
	{"EF86", 1.3, 1, 0.006, 300},
}

func main() {
	if len(os.Args) < 4 {
		// we expected at least $0, valve, ht= and curr=
		bail("bad arg")
	}
	desig := strings.ToUpper(os.Args[1])
	do_pwr(desig, os.Args[2:])
}

func do_pwr(tt string, args []string) {
	var ht, nb, curr float64

	for _, av := range args {
		xx := strings.Split(av, "=")
		switch xx[0] {
		case "ht":
			ht, _ = unitconv.ParsePrefix(xx[1], unitconv.AutoParse)
		case "nb":
			nb, _ = unitconv.ParsePrefix(xx[1], unitconv.AutoParse)
			nb = -(math.Abs(nb))
		case "curr":
			curr, _ = unitconv.ParsePrefix(xx[1], unitconv.AutoParse)
		}
	}
	if ht == 0 {
		bail("missing ht=")
		// } else if nb == 0 {
		// bail("missing nb=")
	} else if curr == 0 {
		bail("missing curr=")
	}

	// a. use ht voltage and cathode current to calculate power dissipation
	// b. compare ht, negative bias, curr, pwr against values in table

	var nht, ncurr, nnb, ndiss float64
	for i := range valve_info {
		if tt == valve_info[i].name {
			nnb = valve_info[i].bias
			ndiss = valve_info[i].max_pwr
			ncurr = valve_info[i].max_curr
			nht = valve_info[i].max_anode
			break
		}
	}
	if nnb == 0 {
		bail("no data for " + tt + " available")
	}

	println("valve:", tt, "(power valve)")
	println()
	println("value                     measured   nominal/max.")
	prpr("anode voltage", "V", "V", ht, nht, ht-nht)
	if nb != 0 {
		foo := nb / nnb
		if foo > 1.1 || foo < 0.9 {
			foo = 1
		} else {
			foo = 0
		}
		prpr("negative bias", "V", "V", nb, nnb, foo)
	}
	prpr("cathode current", "I", "A", curr, ncurr, curr-ncurr)
	diss := ht * curr // p = i * v
	prpr("dissipation", "P", "W", diss, ndiss, diss-ndiss)
	println()
}

func bail(err string) {
	println(err)
	os.Exit(1)
}

// test is a value that will cause "WARNING" to be printed if it is > 0
func prpr(label string, abbrev string, unit string, val float64, nom float64, test float64) {

	// should limit significant figures rather than digits after
	// decimal point, but well...
	val = math.Round(val*1000) / 1000
	vval := unitconv.FormatPrefix(val, unitconv.SI, -1)
	nom = math.Round(nom*1000) / 1000
	nnom := unitconv.FormatPrefix(nom, unitconv.SI, -1)

	var msg string
	if test > 0 {
		msg = "  WARNING"
	} else {
		msg = ""
	}

	fmt.Printf("%-20s %s = %8s%s %8s%s%s\n", label, abbrev, vval, unit, nnom, unit, msg)
}