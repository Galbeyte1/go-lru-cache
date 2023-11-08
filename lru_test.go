package lru

import (
	"testing"
)

func TestCache(t *testing.T) {
	cache := newLRU(2)

	cache.Set("uk", "london")
	u := cache.Get("uk")

	if u != "london" {
		t.Fatalf("expected `London`, got `%s`", u)
	}

	cache.Set("uk", "manchester")
	u = cache.Get("uk")

	if u != "manchester" {
		t.Fatalf("expected `Manchester`, got `%s`", u)
	}

}

func TestLRU(t *testing.T) {
	cache := newLRU(3)

	cache.Set("uk", "london")
	cache.Set("france", "paris")
	cache.Set("germany", "berlin")

	// at this point cache is full
	// order:  germany - frace - uk

	// should remove uk
	cache.Set("belgium", "brussels")

	u := cache.Get("uk")

	if u != "" {
		t.Fatalf("uk should no longer be present, got `%s`", u)
	}

	// order: belgium - germany - france

	f := cache.Get("france")
	if f != "paris" {
		t.Fatalf("expected `paris`, got `%s`", f)
	}

	// order: france - belgium - germany

	// remove germany
	cache.Set("netherlands", "amsterdam")

	g := cache.Get("germany")
	if g != "" {
		t.Fatalf("germany should no longer be present, got `%s`", g)
	}

	// order: netherlands - france - belgium

	n := cache.Get("netherlands")
	if n != "amsterdam" {
		t.Fatalf("expected `amsterdam`, got `%s`", n)
	}

}