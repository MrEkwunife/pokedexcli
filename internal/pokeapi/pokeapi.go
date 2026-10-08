package pokeapi

import (
	"net/http"
	"time"

	"github.com/MrEkwunife/pokedexcli/internal/pokecache"
)

const baseUrl = "https://pokeapi.co/api/v2"

type Client struct {
	cache      pokecache.Cache
	httpClient http.Client
}

func NewClient(interval time.Duration) Client {
	return Client{
		httpClient: http.Client{
			Timeout: time.Minute,
		},

		cache: pokecache.NewCache(interval),
	}
}
