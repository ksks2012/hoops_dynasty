package initial

import (
	"github.com/hoops_dynasty/ecs/entities"
	"github.com/hoops_dynasty/pkg/generator"
)

func InitPlayers(count int) []*entities.Player {
	players := make([]*entities.Player, count)
	for i := 0; i < count; i++ {
		players[i] = generator.GeneratePlayer(i)
	}
	return players

}
