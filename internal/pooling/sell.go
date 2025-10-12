package pooling

import (
	"time"

	"github.com/Tulkdan/central-limit-order-book/internal/service"
)

type Pooling struct {
	sellService *service.SellService
}

func NewPooling(sellService *service.SellService) *Pooling {
	return &Pooling{sellService: sellService}
}

func (p *Pooling) StartPooling() {
	time.Sleep(5 * time.Second)

	p.sellService.MakeSales()

	p.StartPooling()
}
