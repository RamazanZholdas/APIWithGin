package initialData

import (
	"time"

	"github.com/RamazanZholdas/APIWithGin/structs"
)

var (
	Songs = []structs.Song{}
)

func PutData() {
	Songs = append(Songs, structs.Song{
		Id:       1,
		Name:     "Bohemian Rhapsody",
		Duration: time.Now(),
		Genre:    "R&B",
		Artist: &structs.Artist{
			FirstName: "Freddie",
			LastName:  "Mercury",
			Label:     "Queen",
		},
	}, structs.Song{
		Id:       2,
		Name:     "Crazy Train",
		Duration: time.Now().Add(time.Hour * 3),
		Genre:    "Metal",
		Artist: &structs.Artist{
			FirstName: "Ozzy",
			LastName:  "Osbourne",
			Label:     "Black Sabbath",
		},
	})
}
