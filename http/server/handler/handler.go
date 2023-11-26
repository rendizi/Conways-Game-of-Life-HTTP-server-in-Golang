package handler

import (
	"context"
	"fmt"
	"github.com/aivanov/game/internal/service"
	"net/http"
)

type Decorator func(http.Handler) http.Handler

// object to store the game state
type LifeStates struct {
	service.LifeService
}
type World struct {
	Height int
	Width  int
	Cells  [][]bool
}

func New(ctx context.Context,
	lifeService service.LifeService,
) (http.Handler, error) {
	serveMux := http.NewServeMux()

	lifeState := LifeStates{
		LifeService: lifeService,
	}

	serveMux.HandleFunc("/nextstate", lifeState.nextState)

	return serveMux, nil
}

// function to add middleware
func Decorate(next http.Handler, ds ...Decorator) http.Handler {
	decorated := next
	for d := len(ds) - 1; d >= 0; d-- {
		decorated = ds[d](decorated)
	}

	return decorated
}
func (w *World) String() string {
	brownSquare := "\xF0\x9F\x9F\xAB"
	greenSquare := "\xF0\x9F\x9F\xA9"
	result := ""

	for i := 0; i < w.Height; i++ {
		for j := 0; j < w.Width; j++ {
			if w.Cells[i][j] {
				result += greenSquare
			} else {
				result += brownSquare
			}
		}
		if i != w.Height-1 {
			result += "\n"
		}
	}

	return result
}

// getting the next game state
func (ls *LifeStates) nextState(w http.ResponseWriter, r *http.Request) {
	worldState := ls.LifeService.NewState()

	world := World{
		Height: len(worldState.Cells),
		Width:  len(worldState.Cells[0]),
		Cells:  worldState.Cells,
	}

	worldString := world.String()

	_, err := fmt.Fprint(w, worldString)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
