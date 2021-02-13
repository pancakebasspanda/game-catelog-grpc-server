package server

import (
	"context"
	protos "grpc-service/generated-protos"
)

type GameCatalogServer struct {
}

func New() *GameCatalogServer {
	return &GameCatalogServer{}
}

func (g *GameCatalogServer) InStock(ctx context.Context, req *protos.InStockRequest) (*protos.InStockResponse, error) {
	// TODO connect to DB
	return &protos.InStockResponse{
		Stock: true,
	}, nil
}

func (g *GameCatalogServer) Search(ctx context.Context, req *protos.SearchRequest) (*protos.SearchResponse, error) {
	// TODO connect to DB
	return &protos.SearchResponse{
		Games: []*protos.Game{
			&protos.Game{
				Title:   "Plants vs Zombies",
				Format:  protos.Format_XBOX,
				Price:   10.0,
				Genre:   protos.Genre_ACTION,
				Players: protos.Game_TWO_PLAYER,
			},
			&protos.Game{
				Title:   "The Witcher 3: The Wild Hunt",
				Format:  protos.Format_NINTENDO_SWITCH,
				Price:   35.99,
				Genre:   protos.Genre_RPG,
				Players: protos.Game_ONE_PLAYER,
			},
		},
	}, nil
}
