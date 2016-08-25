// Package vizceralSpigo converts a vizceral graph into a spigo architecture
package vizceralSpigo

import (
	"fmt"
	vizceral "github.com/adrianco/go-vizceral"
	"github.com/adrianco/spigo/tooling/architecture"
	//"log"
	//"os"
	//"strings"
)

// ConvertV2A makes an architecture from a Vizceral file
func ConvertV2A(v *vizceral.VizceralGraph) {
	a := architecture.MakeArch(v.Name, "converted from vizceral")
	services := make(map[string][]string)

	for _, region := range v.Nodes {
		fmt.Println(region.Name)
		// only map out the first region we find
		if region.Name != "INTERNET" {
			for _, conn := range region.Connections {
				services[conn.Source] = append(services[conn.Source], conn.Target)
			}
			for _, node := range region.Nodes {
				if node.Name != "INTERNET" { // map the edge node separately
					architecture.AddContainer(a, node.Name, "", "", "", "", "karyon", 1, 1, services[node.Name])
				}
			}
			break
		}
	}
	architecture.AddContainer(a, "INTERNET", "", "", "", "", "denominator", 0, 0, services["INTERNET"])
	architecture.WriteFile(a, v.Name+"_arch")
}
