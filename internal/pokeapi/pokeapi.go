package pokeapi

import (
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/marc-enzmann/pokedexcli/internal/pokecache"
)

var cache = pokecache.NewCache(5 * time.Minute)

func CallPokeApi(url string) ([]byte, error) {
	val, ok := cache.Get(url)

	if ok {
		fmt.Println("Cache hit")
		return val, nil
	} else {
		res, err := http.Get(url)
		if err != nil {
			return nil, err
		}

		body, err := io.ReadAll(res.Body)
		if err != nil {
			return nil, err
		}

		if err = res.Body.Close(); err != nil {
			return nil, err
		}

		cache.Add(url, body)
		return body, nil
	}
}
