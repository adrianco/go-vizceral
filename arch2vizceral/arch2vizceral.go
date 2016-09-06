// utility to read a Spigo arch file and write out a NetflixOSS vizceral file
package main

import (
	"flag"
	//vizceral "github.com/adrianco/go-vizceral"
	"github.com/adrianco/go-vizceral/vizceralSpigo"
)

func main() {
	var arch string
	flag.StringVar(&arch, "arch", "", "Spigo architecture name found in json_arch/<arch>_arch.json")
	flag.Parse()
	if arch != "" {
		vizceralSpigo.ConvertA2V(arch)
	} else {
		flag.PrintDefaults()
	}
}
