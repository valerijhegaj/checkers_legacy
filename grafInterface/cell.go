package grafInterface

import (
	"fmt"

	"checkers/core"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
)

func NewCell(x, y int, interactor *Interface) *Cell {
	return &Cell{ptr: core.Coordinate{x, y}, interactor: interactor}
}

type Cell struct {
	canvas.Rectangle
	interactor *Interface
	ptr        core.Coordinate
}

func (c Cell) Tapped(*fyne.PointEvent) {
	eventor.Tapped(c.ptr, c.interactor)
}

func (c Cell) TappedSecondary(*fyne.PointEvent) {
	eventor.TappedSecondary(c.ptr)
}

var eventor event

type event struct {
	from core.Coordinate
	to   []core.Coordinate
}

func (c *event) Tapped(
	coordinate core.Coordinate,
	interactor *Interface,
) {
	if interactor.IsStartCoordinate(coordinate) {
		c.from = coordinate
		c.to = nil
		fmt.Println("from", coordinate)
		return
	}
	fmt.Println("to", coordinate)
	c.to = append(c.to, coordinate)
	interactor.Move(c.from, c.to)
	c.to = nil

}

func (c *event) TappedSecondary(coordinate core.Coordinate) {
	fmt.Println("way", coordinate)
	c.to = append(c.to, coordinate)
}
