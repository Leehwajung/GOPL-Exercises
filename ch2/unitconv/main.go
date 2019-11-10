//!+

package main

import (
	"fmt"
	"os"
	"strings"
	"strconv"

	"gopl.io/ch2/tempconv"
	"gopl.io/ch2/wghtconv"
	"gopl.io/ch2/lengconv"
)

func main() {
	argsLen := len(os.Args)
	useArgs := argsLen > 1
	
	var converter string
	if useArgs {
		converter = os.Args[1]
	} else {
		for {
			fmt.Print("Input unit converter: ")
			_, err := fmt.Scan(&converter)
			if err == nil {
				break
			} else {
				fmt.Fprintf(os.Stderr, "unitconv: %v\n", err)
			}
		}
	}

	var forwardConv func(float64) string
	var inverseConv func(float64) string
	switch strings.ToLower(converter) {
	case "cf":
		forwardConv = func(v float64) string {
			f := tempconv.Fahrenheit(v)
			return fmt.Sprintf("%s = %s", f, tempconv.FToC(f))
		}
		inverseConv = func(v float64) string {
			c := tempconv.Celsius(v)
			return fmt.Sprintf("%s = %s", c, tempconv.CToF(c))
		}
	case "lbskg":
		forwardConv = func(v float64) string {
			lbs := wghtconv.Pound(v)
			return fmt.Sprintf("%s = %s", lbs, wghtconv.LbsToKg(lbs))
		}
		inverseConv = func(v float64) string {
			kg := wghtconv.Kilogram(v)
			return fmt.Sprintf("%s = %s", kg, wghtconv.KgToLbs(kg))
		}
	case "ftm":
		forwardConv = func(v float64) string {
			ft := lengconv.Feet(v)
			return fmt.Sprintf("%s = %s", ft, lengconv.FtToM(ft))
		}
		inverseConv = func(v float64) string {
			m := lengconv.Meter(v)
			return fmt.Sprintf("%s = %s", m, lengconv.MToFt(m))
		}
	default:
		fmt.Fprintf(os.Stderr, "unitconv: unknown unit converter: %s\n", converter)
		os.Exit(1)
	}

	for i := 2; ; i++ {
		var vstr string
		if useArgs {
			if i < argsLen {
				vstr = os.Args[i]
			} else {
				break
			}
		} else {
			for {
				fmt.Print("Input value: ")
				_, err := fmt.Scan(&vstr)
				if err == nil {
					break
				} else {
					fmt.Fprintf(os.Stderr, "unitconv: %v\n", err)
				}
			}
			if strings.Contains(vstr, "exit") || strings.Contains(vstr, "quit") {
				break
			}
		}
		
		v, err := strconv.ParseFloat(vstr, 64)
		if err != nil {
			fmt.Fprintf(os.Stderr, "unitconv: %v\n", err)
			os.Exit(1)
		}
		fmt.Printf("%s, %s\n", forwardConv(v), inverseConv(v))
	}
}

//!-
