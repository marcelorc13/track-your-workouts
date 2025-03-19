package main

import (
	"fmt"
	"server/internal/models"
)

func main() {
	var exercicios []models.Exercicio

	e1 := models.Exercicio{
		Nome:   "Supino Inclinado com Halteres",
		Series: 2,
	}
	e2 := models.Exercicio{
		Nome:   "Rosca Direta Livre na Barra W",
		Series: 2,
	}
	e3 := models.Exercicio{
		Nome:   "Tríceps Testa Livre na Barra W",
		Series: 2,
	}

	exercicios = append(exercicios, e1, e2, e3)

	treino := models.Treino{
		Nome:       "Upper I",
		Exercicios: exercicios,
	}

	fmt.Println(treino.Nome)
	fmt.Println("N de Exercicios", len(treino.Exercicios))

	for _, e := range exercicios {
		fmt.Println()
		fmt.Println("Exercicio:", e.Nome)
		fmt.Println("Séries Válidas:", e.Series)
	}
}
