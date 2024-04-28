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
		var imageFileName string
		var description string
		var price float64
		row.Scan(&discId, &name, &typ, &manufacturer, &speed, &glide, &turn, &fade, &imageFileName, &description, &price)
		disc := models.Disc{
			DiscId:        discId,
			Name:          name,
			Type:          typ,
			Manufacturer:  manufacturer,
			Speed:         speed,
			Glide:         glide,
			Turn:          turn,
			Fade:          fade,
			ImageFileName: imageFileName,
			Description:   description,
			Price:         price,
		}
		discs = append(discs, disc)
	}
	return discs
}

func GetDiscByName(db *sql.DB, name string) (models.Disc, error) {
	var disc models.Disc

	row := db.QueryRow("SELECT * FROM Discs WHERE name = ?", name)
	err := row.Scan(&disc.DiscId, &disc.Name, &disc.Type, &disc.Manufacturer, &disc.Speed, &disc.Glide, &disc.Turn, &disc.Fade, &disc.ImageFileName, &disc.Description, &disc.Price)
	if err != nil {
		if err == sql.ErrNoRows {
			return disc, err
		}
		log.Fatal(err)
	}

	return disc, nil
}
