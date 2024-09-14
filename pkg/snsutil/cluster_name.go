package snsutil

import (
	"errors"
	"fmt"
	"strings"
)

type ClusterMetaName struct {
	Zone        string
	Lane        string
	Kind        string
	PortName    string
	Application string
}

func (p *ClusterMetaName) String() string {
	return fmt.Sprintf("%s.%s.%s.%s.%s", p.Zone, p.Lane, p.Kind, p.PortName, p.Application)
}

func (p *ClusterMetaName) DomainName() string {
	return fmt.Sprintf("%s.%s", p.PortName, p.Application)
}

func (p *ClusterMetaName) ClusterName() string {
	return fmt.Sprintf("%s.%s.%s", p.Zone, p.Lane, p.Kind)
}

func ParseClusterName(name string) (zone, lane, kind string, err error) {
	slice := strings.Split(name, ".")
	if len(slice) != 3 {
		return "", "", "", fmt.Errorf("%s is an invalid cluster name", name)
	}
	return slice[0], slice[1], slice[2], nil
}

func BuildClusterMetaName(clusterName, portName, appName string) (ClusterMetaName, error) {
	if clusterName == "" {
		return ClusterMetaName{}, errors.New("cluster name is invalid")
	}
	if portName == "" {
		return ClusterMetaName{}, errors.New("port name is invalid")
	}
	if appName == "" {
		return ClusterMetaName{}, errors.New("app name is invalid")
	}

	zone, lane, kind, err := ParseClusterName(clusterName)
	if err != nil {
		return ClusterMetaName{}, err
	}
	return ClusterMetaName{
		Zone:        zone,
		Lane:        lane,
		Kind:        kind,
		PortName:    portName,
		Application: appName,
	}, nil
}
