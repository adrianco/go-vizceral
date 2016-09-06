// vizceral generates output in the NetflixOSS vizceral format
// https://github.com/Netflix/vizceral/blob/master/DATAFORMATS.md
package vizceral

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
)

// Metadata
type VizceralMetadata struct {
	Streaming int `json:"streaming"`
}

// Notice appears in the sidebar
type VizceralNotice struct {
	Title    string `json:"title,omitempty"`
	Subtitle string `json:"subtitle,omitempty"`
	Link     string `json:"link,omitempty"`
	Severity int    `json:"severity,omitempty"`
}

// Levels of trafic in each state
type VizceralLevels struct {
	Danger  float32 `json:"danger,omitempty"`
	Warning float32 `json:"warning,omitempty"`
	Normal  float32 `json:"normal,omitempty"`
}

// One Connection
type VizceralConnection struct {
	Source   string           `json:"source,omitempty"`
	Target   string           `json:"target,omitempty"`
	Metadata VizceralMetadata `json:"metadata,omitempty"`
	Metrics  VizceralLevels   `json:"metrics,omitempty"`
	Status   VizceralLevels   `json:"status,omitempty"`
	Notices  []VizceralNotice `json:"node,omitempty"`
	Class    string           `json:"class,omitempty"`
}

// One node (region/service hierarchy)
type VizceralNode struct {
	Renderer    string               `json:"renderer,omitempty"` // 'region' or omit for service
	Name        string               `json:"name,omitempty"`
	MaxVolume   float64              `json:"maxVolume,omitempty"` // relative base for levels animation
	Updated     int64                `json:"updated,omitempty"`   // Unix timestamp. Only checked on the top-level list of nodes. Last time the data was updated
	Nodes       []VizceralNode       `json:"nodes,omitempty"`
	Connections []VizceralConnection `json:"connections,omitempty"`
	Notices     []VizceralNotice     `json:"node,omitempty"`
	Class       string               `json:"class,omitempty"` // 'normal', 'warning', or 'danger'
	Metadata    VizceralMetadata     `json:"metadata,omitempty"`
}

// Global level of graph file format
type VizceralGraph struct {
	Renderer    string               `json:"renderer"` // 'global'
	Name        string               `json:"name"`
	MaxVolume   float64              `json:"maxVolume,omitempty"` // relative base for levels animation
	Nodes       []VizceralNode       `json:"nodes,omitempty"`
	Connections []VizceralConnection `json:"connections,omitempty"`
}

// print a Vizceral graph as json
func Write(v *VizceralGraph) {
	vJson, _ := json.Marshal(*v)
	fmt.Println(string(vJson))
}

// Read a Vizceral format file into a graph
func ReadFile(fn string) *VizceralGraph {
	data, err := ioutil.ReadFile(fn)
	if err != nil {
		log.Fatal(err)
	}
	v := new(VizceralGraph)
	err = json.Unmarshal(data, v)
	if err != nil {
		log.Fatal(err)
	}
	return v
}
