package guest

import "matrixone/pkg/vm/mmu/host"

type Mmu struct {
	size  int64
	limit int64
	Mmu   *host.Mmu
}