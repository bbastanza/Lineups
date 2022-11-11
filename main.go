package main

import (
	"fmt"
	"math/rand"
	"time"
	// "time"
)

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
	Single Hit = iota
	Double
	Triple
	HomeRun
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
  rand.Seed(time.Now().UnixNano())

	player := Player{
		BattingAverage:   300,
		OnBasePercentage: 400,
		Position:         RF,
	}

	outcome := CalculateOutcome(player)

  if (!outcome.IsOnBase){
    fmt.Println("You're Out! " + outcome.Out.String())
  } else {
    fmt.Println("Nice Hit! " + outcome.Hit.String())
  }

	fmt.Println(outcome.IsOnBase)
}

type Player struct {
	BattingAverage   int
	OnBasePercentage int
	Position         Position
}

type Outcome struct {
	IsOnBase bool
	Hit   Hit
	Out   Out
}

func CalculateOutcome(player Player) Outcome {
	rand := generateRandomNumber()

  fmt.Println(rand)
	if rand <= player.OnBasePercentage {
		return Outcome{
			IsOnBase: true,
			Hit:   generateRandomHit(),
		}
	}

	return Outcome{
		IsOnBase: false,
		Out:   generateRandomOut(),
	}
}

func generateRandomHit() Hit {
  max := 4

  // Intn will not reach the max
  r := rand.Intn(max)

  fmt.Println(r)
  switch r {
  case 1:
  return Single
  case 2:
  return Double
  case 3:
  return Triple
  default:
  return HomeRun
}
}

func generateRandomOut() Out {
  max := 4
  r := rand.Intn(max)
  fmt.Println(r)
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

func generateRandomNumber() int {
  max := 1000
  r := rand.Intn(max)
	return r
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
