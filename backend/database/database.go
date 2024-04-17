package discgolfdb

import (
	"database/sql"
	"discgolf/models"
	"log"
)

func GetAllDiscs(db *sql.DB) []models.Disc {
	var discs []models.Disc

	row, err := db.Query("SELECT * FROM Discs")
	if err != nil {
		log.Fatal(err)
	}
	defer row.Close()

	for row.Next() {
		var discId int
		var name string
		var typ string
		var manufacturer string
		var speed int
		var glide int
		var turn int
		var fade int
		row.Scan(&discId, name, typ, manufacturer, speed, glide, turn, fade)
		disc := models.Disc{
			DiscId:       discId,
			Name:         name,
			Type:         typ,
			Manufacturer: manufacturer,
			Speed:        speed,
			Glide:        glide,
			Turn:         turn,
			Fade:         fade,
		}
		discs = append(discs, disc)
	}
	return discs
}
