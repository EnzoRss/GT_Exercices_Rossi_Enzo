package main

import (
	"fmt"
)

func main() {
	var nbr_allumette, nbr_enlever, cmpt int
	cmpt = 0

	fmt.Println("Choisissez un nombre d'allumettes :")
	fmt.Scan(&nbr_allumette)

	for nbr_allumette > 0 {
		fmt.Println("Combien d'allumettes souhaitez-vous enlever ?")
		fmt.Scan(&nbr_enlever)
		if nbr_enlever > 3 || nbr_enlever > nbr_allumette {
			fmt.Println("Nombre saisi incorrect, veuillez recommencer")
			continue
		}
		nbr_allumette -= nbr_enlever
		if nbr_allumette < 0 {
			nbr_allumette *= nbr_allumette
			nbr_allumette -= nbr_allumette
		}
		if cmpt%2 == 0 {
			fmt.Printf("Joueur 1 enlève %d allumettes, il reste : %d\n", nbr_enlever, nbr_allumette)
		} else {
			fmt.Printf("Joueur 2 enlève %d allumettes, il reste : %d\n", nbr_enlever, nbr_allumette)
		}
		cmpt++
	}
	if cmpt%2 == 0 {
		fmt.Println("Le joueur 1 a perdu")
	} else {
		fmt.Println("Le joueur 2 a perdu")
	}
	
}