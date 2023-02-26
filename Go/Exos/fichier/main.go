package main

import(
	"os"
	"bufio"
	"fmt"
)

func main(){
	var choix string 
	scanner := bufio.NewScanner(os.Stdin)

	for(true) {
		fmt.Println("Choisissez l'action que vous souhaitez :\n1-Récupérer tout le texte\n2-Ajouter du texte dans ce fichier\n3-Supprimer tout le contenu du fichier\n4-Remplacer le contenu par du texte donné ")
		scanner.Scan()
		choix = scanner.Text()
		switch choix {
		case "1":
			arr,_ := os.ReadFile("test.txt")
			fmt.Println(string(arr))
		case "2" :
			arr,_ := os.ReadFile("test.txt")
			arr_str:= string(arr)
			fmt.Println("Rentrez le texte a ajouter:")
			scanner.Scan()
			text := scanner.Text()
			arr_str += text
			os.WriteFile("test.txt",[]byte(arr_str) ,0666)
			fmt.Println(arr_str)
		case "3" :
			reset,_ := os.OpenFile("test.txt",os.O_WRONLY,0666)
			reset.Truncate(0)
			arr,_ := os.ReadFile("test.txt")
			fmt.Println(string(arr))
		case "4" :
			fmt.Println("Rentrez le texte a ajouter:")
			scanner.Scan()
			text := scanner.Text()
			os.WriteFile("test.txt",[]byte(text) ,0666)
			arr,_ := os.ReadFile("test.txt")
			fmt.Println(string(arr))

		}
		
	}

}