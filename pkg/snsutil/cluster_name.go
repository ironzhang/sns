package snsutil

import "fmt"

type ClusterMetadataName struct {
	PortName    string
	ServiceName string
	ClusterName string
}

func NewClusterMetadataName(appName, portName string) ClusterMetadataName {
	cluster, service := ParseAppName(appName)
	return ClusterMetadataName{
		PortName:    portName,
		ServiceName: service,
		ClusterName: cluster,
	}
}

func (p *ClusterMetadataName) String() string {
	return fmt.Sprintf("%s.sns.%s.%s", p.ClusterName, p.PortName, p.ServiceName)
}

func (p *ClusterMetadataName) Domain() string {
	return fmt.Sprintf("sns.%s.%s", p.PortName, p.ServiceName)
}
