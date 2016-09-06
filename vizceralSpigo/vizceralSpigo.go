// Package vizceralSpigo converts a vizceral graph into a spigo architecture and vice versa
package vizceralSpigo

import (
	"fmt"
	vizceral "github.com/adrianco/go-vizceral"
	"github.com/adrianco/spigo/tooling/architecture"
	"log"
	//"time"
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
				if conn.Target == "" {
					log.Printf("Empty target for\n%v\n", conn)
				} else {
					services[conn.Source] = append(services[conn.Source], conn.Target)
				}
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

// ConvertA2V makes a Vizceral example file from a Spigo architecture with faked flows on the connections
func ConvertA2V(arch string) {
	a := architecture.ReadArch(arch)
	var services []string
	var deps []architecture.Connection
	architecture.ListDependencies(a, &services, &deps)
	// root of the vizceral graph
	vg := vizceral.VizceralGraph{"global", arch, 0.0, nil, nil}
	// regional nodes at the top level of the graph
	regions := []vizceral.VizceralNode{
		{"region", "INTERNET", 0.0, 0, // time.Now() change this to unix seconds
			nil, nil, nil, "normal", vizceral.VizceralMetadata{0},
		},
		{"region", "us-east-1", 20000.0, 0, // time.Now() change this to unix seconds
			nil, nil, nil, "normal", vizceral.VizceralMetadata{1},
		},
	}
	vg.Nodes = regions
	regcons := []vizceral.VizceralConnection{
		{"INTERNET", "us-east-1", vizceral.VizceralMetadata{0},
			vizceral.VizceralLevels{92.3, 0.0, 10000.0},
			vizceral.VizceralLevels{0, 0, 0}, nil, "normal",
		},
	}
	vg.Connections = regcons

	// service nodes inside a region
	nodes := []vizceral.VizceralNode{
		{"", "INTERNET", 0.0, 0, nil, nil, nil, "normal", vizceral.VizceralMetadata{1}},
	}
	for _, s := range services {
		nodes = append(nodes, vizceral.VizceralNode{
			"", s, 0.0, 0, nil, nil, nil, "normal", vizceral.VizceralMetadata{1},
		})
	}
	regions[1].Nodes = nodes
	// service connections inside a region
	var conns []vizceral.VizceralConnection
	for _, c := range deps {
		if c.Source != c.Dest {
			conns = append(conns, vizceral.VizceralConnection{c.Source, c.Dest, vizceral.VizceralMetadata{1},
				vizceral.VizceralLevels{92.3, 0.0, 5000.0},
				vizceral.VizceralLevels{0, 0, 0}, nil, "normal",
			})
		}
	}
	// last service in the spigo arch list is the one that drives incoming traffic
	conns = append(conns, vizceral.VizceralConnection{"INTERNET", services[len(services)-1], vizceral.VizceralMetadata{1},
		vizceral.VizceralLevels{92.3, 0.0, 10000.0},
		vizceral.VizceralLevels{0, 0, 0}, nil, "normal",
	})
	regions[1].Connections = conns
	vizceral.Write(&vg)
}
