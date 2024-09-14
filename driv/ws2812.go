package drivers


import (
   "driver/color"
    "machine"
)
type WS2812 struct {
	spi        machine.SPI
	colorOrder colororder.ColorOrder
	buffer     []uint8
	maxLEDs    int
}
