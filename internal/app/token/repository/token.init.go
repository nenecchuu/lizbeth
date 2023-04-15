package repository

import (
	"github.com/nenecchuu/lizbeth-be-core/internal/integration/spotify_api"
	"go.mongodb.org/mongo-driver/mongo"
)

type Module struct {
	spotifyAuthApiCall spotify_api.SpotifyApiCallIntegration
	mongoManager       *mongo.Database
}

type Opts struct {
	SpotifyAuthApiCall spotify_api.SpotifyApiCallIntegration
	MongoManager       *mongo.Database
}

func New(o Opts) *Module {
	return &Module{
		spotifyAuthApiCall: o.SpotifyAuthApiCall,
		mongoManager:       o.MongoManager,
	}
}
