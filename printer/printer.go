package printer

import (
	"crossword/grid"
	"log/slog"
)

type Printer struct {
	count int
	Each  int
}

func New(each int) *Printer {
	return &Printer{Each: each}
}

func (p *Printer) Print(logger *slog.Logger, msg string, g grid.Grid) {
	p.count++
	if p.count < p.Each {
		return
	}
	p.count = 0

	g.Print()
	logger.Info(msg, "completion", g.CompletionState())
}
