package world

import (
	"bytes"
	"github.com/olsdavis/goelan/material"
)

type Block struct {
	location   SimpleLocation
	material   material.Material
	BlockState byte
}

func (b *Block) GetLocation() SimpleLocation {
	return b.location
}

func (b *Block) GetMaterial() material.Material {
	return b.material
}

func (b *Block) writeBlock(buffer *bytes.Buffer) {

}