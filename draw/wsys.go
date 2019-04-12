package draw

import "image"

// TODO

func (d *Display) Resize(r image.Rectangle) error {
	return d.conn.Resize(r)
}
