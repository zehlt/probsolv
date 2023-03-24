package sol

import "github.com/zehlt/datt"

func findRedundantConnection(edges [][]int) []int {
	uf := datt.NewUnionFind(len(edges) + 1)

	for _, edge := range edges {
		if uf.Find(edge[0], edge[1]) {
			return edge
		}

		uf.Union(edge[0], edge[1])
	}

	return nil
}
