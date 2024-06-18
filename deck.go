package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

type deck []string

func newDeck() deck {
	cards := deck{}
	cardSuits := []string{"Spades", "Diamonds", "Hearts", "Clubs"}
	cardValues := []string{"Ace", "Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine", "Ten", "Jack", "Queen", "King"}

	for _, suit := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, value+" of "+suit)
		}
	}
	return cards
}

func (cards deck) print() {
	fmt.Println(cards)
	test2 := strings.Join(cards, ",")
	fmt.Println(test2)
	for i, card := range cards {
		fmt.Println(i, card)
	}
}

func (cards deck) toString() string {
	return strings.Join([]string(cards), ",")
}

func (cards deck) save(filename string) error {
	return ioutil.WriteFile(filename, []byte(cards.toString()), 0666)
}

func (cards deck) readFile(filename string) deck {
	bs, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
	fmt.Println("bs", string(bs))
	s := strings.Split(string(bs), ",")
	return s
}
