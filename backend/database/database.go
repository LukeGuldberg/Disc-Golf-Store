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
		var disc models.Disc
		err := row.Scan(&disc.DiscId, &disc.Name, &disc.Type, &disc.Manufacturer, &disc.Speed, &disc.Glide, &disc.Turn, &disc.Fade, &disc.ImageFileName, &disc.Description, &disc.Price)
		if err != nil {
			log.Fatal(err)
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

func GetDiscsByManufacturer(db *sql.DB, manufacturer string) ([]models.Disc, error) {
	var discs []models.Disc

	rows, err := db.Query("SELECT * FROM Discs WHERE manufacturer = ?", manufacturer)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	for rows.Next() {
		var disc models.Disc
		err := rows.Scan(&disc.DiscId, &disc.Name, &disc.Type, &disc.Manufacturer, &disc.Speed, &disc.Glide, &disc.Turn, &disc.Fade, &disc.ImageFileName, &disc.Description, &disc.Price)
		if err != nil {
			log.Fatal(err)
		}
		discs = append(discs, disc)
	}

	return discs, nil
}

func GetDiscsByType(db *sql.DB, discType string) ([]models.Disc, error) {
	var discs []models.Disc

	rows, err := db.Query("SELECT * FROM Discs WHERE type = ?", discType)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	for rows.Next() {
		var disc models.Disc
		err := rows.Scan(&disc.DiscId, &disc.Name, &disc.Type, &disc.Manufacturer, &disc.Speed, &disc.Glide, &disc.Turn, &disc.Fade, &disc.ImageFileName, &disc.Description, &disc.Price)
		if err != nil {
			log.Fatal(err)
		}
		discs = append(discs, disc)
	}

	return discs, nil
}
