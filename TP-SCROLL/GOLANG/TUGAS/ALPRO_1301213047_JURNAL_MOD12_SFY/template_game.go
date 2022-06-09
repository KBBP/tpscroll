package main

import (
	. "fmt"
	"math/rand"
	"time"
)

type domino struct {
	left, right int
}

const nMax = 28

type arrDomino [nMax]domino
type arrInt [nMax]int

func main() {
	var (
		playerScore, dealerScore, decision, nBoneyard, pSkor, dSkor, total int
		boneyard, playerTiles                                              arrDomino
	)

	playerScore = 0
	dealerScore = 0

	Println("Welcome to the Algorithm and Programming Practice")
	rand.Seed(time.Now().UnixNano())

	createTile(&boneyard)

	nBoneyard = nMax

	total = 0

	Println("Boneyard tiles:", nBoneyard)

	for decision == 0 && nBoneyard >= 4 {

		Printf("Your score is %d/%d \n\n", playerScore, total)

		divTile(&boneyard, &playerTiles, &nBoneyard)

		for decision == 1 || decision == 2 {
			printTile(playerTiles, "player")
			getDecision(&decision)
		}

		if decision == 0 {
			if pSkor == 1 {
				Println("You won")
				playerScore++
			} else if dSkor == 1 {
				Println("You lose")
				dealerScore++
			}
			total++
		}
	}

	Printf("Your last score is %d/%d \n", playerScore, total)
	Println("Thank you for playing with us.")
	Println("Your winning rate is", playerScore*100/total, "%")
}

func getDecision(decision *int) int {
	var dec int
	Scan(&dec)
	for dec != 0 || dec != 1 || dec != 2 || dec != 9 {
		Scan(&dec)
	}

	return dec
}

func createTile(T *arrDomino) {
	var kiri, kanan, i int
	i = 0
	for kiri = 0; kiri <= 6; kiri++ {
		for kanan = 0; kanan <= kiri; kanan++ {
			T[i].left = kiri
			T[i].right = kanan
			i++
		}
	}
}

func searchInt(T arrInt, n, x int) int {
	var found, k int
	found = -1
	k = 0
	for found == -1 && k < n {
		if found == T[k] && found == x {
			found = k
		}
		k++
	}
	return found
}

func randomTile(T *arrDomino, n int) {
	Println("Dealing ...")
	//var temp arrInt
	//var temp2 arrDomino = *T
}

func divTile(T, p *arrDomino, n *int) {

}

func replaceTile(T, p *arrDomino, n *int, posisi int) {

}

func haveBalak(p domino) bool {
	return false
}

func tilePoint(p domino) int {
	return 0
}

func haveTwoBalak(p arrDomino) bool {
	return false
}

func win(p1, p2 arrDomino, p1Skor, p2Skor *int) {
	//var poin1, poin2 int

	*p1Skor = 0
	*p2Skor = 0
}

func printTile(p arrDomino, s string) {
	var text string
	if s == "player" {
		text = "Your tiles"
	} else { // s == "dealer"
		text = "Dealer tiles"
	}

	Printf("%s: (%d,%d) (%d,%d)\n", text, p[0].left, p[0].right, p[1].left, p[1].right)
}
