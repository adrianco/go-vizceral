// utility to read a NetflixOSS vizceral file and write out an arch_json
package main

import (
	"flag"
	vizceral "github.com/adrianco/go-vizceral"
	"github.com/adrianco/go-vizceral/vizceralSpigo"
)

func main() {
	var fn string
	flag.StringVar(&fn, "file", "", "NetflixOSS vizceral format json file")
	flag.Parse()
	if fn != "" {
		v := vizceral.ReadFile(fn)
		if v != nil {
			vizceralSpigo.ConvertV2A(v)
		}
	} else {
		flag.PrintDefaults()
	}
}
