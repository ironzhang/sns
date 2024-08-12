package engine

import (
	"sort"

	"github.com/ironzhang/superlib/superutil/supermodel"
)

func compare(d, x, y string) int {
	if x == y {
		return 0
	}
	if x == d {
		return -1
	}
	if y == d {
		return 1
	}
	if x < y {
		return -1
	}
	return 1
}

type clusterSorter struct {
	DefaultZone string
	DefaultLane string
	DefaultKind string
}

func (p *clusterSorter) Sort(clusters []supermodel.Cluster) {
	sort.Slice(clusters, func(i, j int) bool {
		var r int

		r = compare(p.DefaultZone, clusters[i].Labels[supermodel.ZoneKey], clusters[j].Labels[supermodel.ZoneKey])
		if r < 0 {
			return true
		} else if r > 0 {
			return false
		}

		r = compare(p.DefaultLane, clusters[i].Labels[supermodel.LaneKey], clusters[j].Labels[supermodel.LaneKey])
		if r < 0 {
			return true
		} else if r > 0 {
			return false
		}

		r = compare(p.DefaultKind, clusters[i].Labels[supermodel.KindKey], clusters[j].Labels[supermodel.KindKey])
		if r < 0 {
			return true
		}
		return false
	})
}
