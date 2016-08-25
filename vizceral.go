// vizceral generates output in the NetflixOSS vizceral format
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

// Notice
type VizceralNotice struct {
	Title    string `json:"title,omitempty"`
	Link     string `json:"link,omitempty"`
	Severity int    `json:"severity,omitempty"`
}

// Levels
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
}

// One Service
type VizceralService struct {
	Renderer string           `json:"renderer"`
	Name     string           `json:"name"`
	Metadata VizceralMetadata `json:"metadata,omitempty"`
}

// One region
type VizceralRegion struct {
	Renderer    string               `json:"renderer,omitempty"`
	Name        string               `json:"name,omitempty"`
	MaxVolume   float32              `json:"maxVolume,omitempty"`
	Updated     int64                `json:"updated,omitempty"`
	Nodes       []VizceralService    `json:"nodes,omitempty"`
	Connections []VizceralConnection `json:"connections,omitempty"`
	Class       string               `json:"class,omitempty"`
}

// Global level of graph file format
type VizceralGraph struct {
	Renderer    string               `json:"renderer"`
	Name        string               `json:"name"`
	Nodes       []VizceralRegion     `json:"nodes,omitempty"`
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
