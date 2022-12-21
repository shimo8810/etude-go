package animals

import (
	"testing"
)

func TestFeedElephant(t *testing.T) {
	expect := "Grass"
	actual := FeedElephant()

	if expect != actual {
		t.Errorf("%s != %s", expect, actual)
	}
}

func TestFeedMonkey(t *testing.T) {
	expect := "Banana"
	actual := FeedMonkey()

	if expect != actual {
		t.Errorf("%s != %s", expect, actual)
	}
}

func TestFeedRabbit(t *testing.T) {
	expect := "Carrot"
	actual := FeedRabbit()

	if expect != actual {
		t.Errorf("%s != %s", expect, actual)
	}
}
