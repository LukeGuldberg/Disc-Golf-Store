package models

type Disc struct {
	DiscId int
	Name string
	Type string // putter, mid-range, fairway driver, distance driver
	Manufacturer string
	Speed int
	Glide int
	Turn int
	Fade int
}