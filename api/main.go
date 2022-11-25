package main

// TODO check all my fking math!!
// tests

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"
	// "time"
)

type Player struct {
	AtBats           int
	Hits             int
	PlateAppearances int
	Walks            int
	Doubles          int
	Triples          int
	HomeRuns         int
	Position         Position
}

func (p Player) Singles() int {
	return p.Hits - (p.Doubles + p.Triples + p.HomeRuns)
}

func (p Player) Avg() int {
	f := toInt(float64(p.Hits) / float64(p.AtBats))
	return f
}

func (p Player) Obp() int {
	f := toInt(float64(p.Walks) / float64(p.PlateAppearances))
	return f
}

type AtBat struct {
	IsOnBase bool
	IsHit    bool
	Hit      Hit
	Out      Out
}

type Position int64

const (
	C Position = iota
	FB
	SB
	SS
	TB
	RF
	CF
	LF
	DH
)

type Hit int64

const (
	HomeRun Hit = iota
	Single
	Double
	Triple
)

func (h Hit) String() string {
	switch h {
	case Single:
		return "Single"
	case Double:
		return "Double"
	case Triple:
		return "Triple"
	default:
		return "HomeRun"
	}
}

type Out int64

const (
	Ground Out = iota
	Fly
	Pop
	Strike
)

func (o Out) String() string {
	switch o {
	case Strike:
		return "Strike"
	case Fly:
		return "Fly"
	case Pop:
		return "Pop"
	default:
		return "Ground"
	}
}

func main() {
	i := 0

	plateAppearanceCount := 10000

	onBase := 0
	hit := 0
	out := 0
	walk := 0

	aaron := Player{
		Hits:             177,
		AtBats:           570,
		PlateAppearances: 696,
		Walks:            111,
		Doubles:          28,
		Triples:          0,
		HomeRuns:         62,
		Position:         RF,
	}

	for i < plateAppearanceCount {
		x := RandomAtBat(aaron)
		switch x {
		case "Hit":
			hit++
			onBase++
		case "Out":
			out++
		default:
			walk++
			onBase++
		}
		i++
	}
	fmt.Println("Walks " + strconv.Itoa(walk))
	fmt.Println("Hits " + strconv.Itoa(hit))

	fmt.Println("Outs " + strconv.Itoa(out))
	fmt.Println("On Base " + strconv.Itoa(onBase))
	fmt.Println("")

	obp := toInt(float64(onBase*10) / float64(plateAppearanceCount))
	fmt.Println("OBP " + strconv.Itoa(obp))

	avg := toInt(float64(hit*10) / float64(plateAppearanceCount))
	fmt.Println("AVG " + strconv.Itoa(avg))

	wavg := toInt(float64(walk*10) / float64(plateAppearanceCount))
	fmt.Println("WAVG " + strconv.Itoa(wavg))
}

func RandomAtBat(player Player) string {
	rand.Seed(time.Now().UnixNano())

	result := CalculateAtBat(player)

	if result.IsOnBase && result.IsHit {
		return "Hit"
	}

	if result.IsOnBase {
		return "Walk"
	}

	return "Out"
}

func CalculateAtBat(player Player) AtBat {
	max := 1000
	random := rand.Intn(max)

	if random <= player.Obp() {
		walkPercent := toInt(float32(player.Walks*1000) / float32(player.PlateAppearances))

		r := rand.Intn(100)

		if r <= walkPercent {
			return AtBat{
				IsOnBase: true,
				IsHit:    false,
			}
		}

		return AtBat{
			IsOnBase: true,
			IsHit:    true,
			Hit:      generateRandomHit(),
		}
	}

	return AtBat{
		IsOnBase: false,
		Out:      generateRandomOut(),
	}
}

func generateRandomHit() Hit {
	max := 4

	// Intn is 0 indexed ie. max of 4 = [0, 1, 2, 3]
	r := rand.Intn(max)

	switch r {
	case 1:
		return Single
	case 2:
		return Double
	case 3:
		return Triple
	case 0:
		return HomeRun
	}
	return HomeRun
}

func generateRandomOut() Out {
	max := 3
	r := rand.Intn(max)

	switch r {
	case 1:
		return Fly
	case 2:
		return Ground
	case 3:
		return Strike
	default:
		return Pop
	}
}

func toInt(f float64) int {
	return (int)(f * 100)
}

// import (
// 	// "github.com/gin-gonic/gin"
// 	// "net/http"
// )

// func main() {
// 	// router := gin.Default()
// 	// router.GET("/albums", GetAlbums)
// 	// router.Run("localhost:8080")
// }
//
// // func GetAlbums(c *gin.Context) {
// // 	println("Getting Albums")
// // 	c.JSON(http.StatusOK, albums)
// // }
//
// // func PostAlbum(c *gin.Context) {
// // albums := append(albums)
// //
// // }
