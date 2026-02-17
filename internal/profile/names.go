package profile

import (
	"fmt"
	"math/rand"
)

var adjectives = []string{
	"dizzy", "cosmic", "fluffy", "sneaky", "turbo",
	"brave", "clever", "witty", "jolly", "nimble",
	"swift", "bold", "calm", "eager", "fancy",
	"gentle", "happy", "keen", "lively", "mighty",
	"noble", "proud", "quiet", "rapid", "sharp",
	"tough", "vivid", "warm", "zesty", "daring",
	"epic", "fierce", "grand", "humble", "iconic",
	"jazzy", "kind", "lucid", "merry", "neat",
}

var nouns = []string{
	"panda", "taco", "falcon", "phoenix", "otter",
	"narwhal", "badger", "dragon", "koala", "tiger",
	"wolf", "eagle", "fox", "hawk", "lion",
	"bear", "crow", "deer", "elk", "frog",
	"gecko", "heron", "ibis", "jay", "kite",
	"lynx", "moose", "newt", "owl", "pike",
	"quail", "raven", "seal", "toad", "viper",
	"whale", "yak", "zebra", "bison", "crane",
}

func GenerateRandomName() string {
	adj := adjectives[rand.Intn(len(adjectives))]
	noun := nouns[rand.Intn(len(nouns))]

	return fmt.Sprintf("%s-%s", adj, noun)
}
