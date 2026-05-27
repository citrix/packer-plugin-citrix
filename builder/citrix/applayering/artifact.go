// Copyright (c) Citrix, Inc.

package applayering

import "fmt"

// packersdk.Artifact implementation
type Artifact struct {
	// StateData should store data such as GeneratedData
	// to be shared with post-processors
	OperationType          string
	TargetLayerName        string
	TargetLayerVersionName string
	StateData              map[string]interface{}
}

func (*Artifact) BuilderId() string {
	return BuilderId
}

func (a *Artifact) Files() []string {
	return []string{}
}

func (*Artifact) Id() string {
	return ""
}

func (a *Artifact) String() string {
	return fmt.Sprintf("Operation Type: %s, Target Layer Name: %s, Target Layer Version: %s",
		a.OperationType, a.TargetLayerName, a.TargetLayerVersionName)
}

func (a *Artifact) State(name string) interface{} {
	return a.StateData[name]
}

func (a *Artifact) Destroy() error {
	return nil
}
