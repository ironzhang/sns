package snsutil

import "fmt"

type ClusterMetadataName struct {
	ClusterName string
	PortName    string
	AppName     string
}

func NewClusterMetadataName(clusterName, portName, appName string) ClusterMetadataName {
	if clusterName == "" {
		clusterName = "default"
	}
	return ClusterMetadataName{
		ClusterName: clusterName,
		PortName:    portName,
		AppName:     appName,
	}
}

func (p *ClusterMetadataName) String() string {
	return fmt.Sprintf("%s.%s.%s", p.ClusterName, p.PortName, p.AppName)
}

func (p *ClusterMetadataName) Domain() string {
	return fmt.Sprintf("%s.%s", p.PortName, p.AppName)
}
