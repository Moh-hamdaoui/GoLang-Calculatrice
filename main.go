package main

import (
	"errors"
	"math"
	"fmt"
	"strconv"
	
	
)

func addition(a, b float64) float64 {
	return a + b
}

func soustraction(a, b float64) float64 {
	return a - b
}

func multiplication(a, b float64) float64 {
	return a * b
}

func division(a, b float64) (float64, error) {
	if b != 0 {
		return a / b, nil
	} else {
		return 0, errors.New("division par zero")
	}
}

func puissance(a, b float64) float64 {
	return math.Pow(a, b)
}

func racineCarree(a float64) (float64, error) {
	if a < 0 {
		return 0, errors.New("Le nombre est negative")
	}
	return math.Sqrt(a), nil
}

func factoriel(n int) (int, error) {
	if n < 0 {
		return 0, errors.New("Nombre negative")
	}
	if n == 0 {
		return 1, nil
	}
	result := 1
	for i := 1; i <= n; i++ {
		result *= i
	}
	return result, nil
}

func readFloat(prompt string) (float64, error) {
	var input string
	fmt.Print(prompt)
	fmt.Scanln(&input)
	value, err := strconv.ParseFloat(input, 64)
	if err != nil {
		return 0, fmt.Errorf("Veuillez entrer un nombre flottant valide")
	}
	return value, nil
}

func readInt(prompt string) (int, error) {
	var input string
	fmt.Print(prompt)
	fmt.Scanln(&input)
	value, err := strconv.Atoi(input)
	if err != nil {
		return 0, fmt.Errorf("Veuillez entrer un entier valide")
	}
	return value, nil	
}

func main() {
	for {
		fmt.Println("\n Menu Principal:")
		fmt.Println("1. Addition")
		fmt.Println("2. Soustraction")
		fmt.Println("3. Multiplication")
		fmt.Println("4. Division")
		fmt.Println("5. Puissance")
		fmt.Println("6. Racine Carrée")
		fmt.Println("7. Factoriel")
		fmt.Println("8. Quitter")
		fmt.Print("Choisissez une option : ")

		var choice int
		fmt.Scanln(&choice)

		switch choice {
		case 1:
			a, err := readFloat("Entrez le premier nombre : ")
			if err != nil {
				fmt.Println("Erreur :", err)
				break
			}
			b, err := readFloat("Entrez le deuxième nombre : ")
			if err != nil {
				fmt.Println("Erreur :", err)
				break
			}
			fmt.Println("Résultat :", addition(a, b))
		case 2:
			a, err := readFloat("Entrez le premier nombre : ")
			if err != nil {
				fmt.Println("Erreur :", err)
				break
			}
			b, err := readFloat("Entrez le deuxième nombre : ")
			if err != nil {
				fmt.Println("Erreur :", err)
				break
			}
			fmt.Println("Résultat :", soustraction(a, b))
		case 3:
			a, err := readFloat("Entrez le premier nombre : ")
			if err != nil {
				fmt.Println("Erreur :", err)
				break
			}
			b, err := readFloat("Entrez le deuxième nombre : ")
			if err != nil {
				fmt.Println("Erreur :", err)
				break
			}
			fmt.Println("Résultat :", multiplication(a, b))
		case 4:
			a, err := readFloat("Entrez le premier nombre : ")
			if err != nil {
				fmt.Println("Erreur :", err)
				break
			}
			b, err := readFloat("Entrez le deuxième nombre : ")
			if err != nil {
				fmt.Println("Erreur :", err)
				break
			}
			result, err := division(a, b)
			if err != nil {
				fmt.Println("Erreur :", err)
			} else {
				fmt.Println("Résultat :", result)
			}
		case 5:
			a, err := readFloat("Entrez la base : ")
			if err != nil {
				fmt.Println("Erreur :", err)
				break
			}
			b, err := readFloat("Entrez l'exposant : ")
			if err != nil {
				fmt.Println("Erreur :", err)
				break
			}
			fmt.Println("Résultat :", puissance(a, b))
		case 6:
			a, err := readFloat("Entrez le nombre : ")
			if err != nil {
				fmt.Println("Erreur :", err)
				break
			}
			result, err := racineCarree(a)
			if err != nil {
				fmt.Println("Erreur :", err)
			} else {
				fmt.Println("Résultat :", result)
			}
		case 7:
			n, err := readInt("Entrez un entier : ")
			if err != nil {
				fmt.Println("Erreur :", err)
				break
			}
			result, err := factoriel(n)
			if err != nil {
				fmt.Println("Erreur :", err)
			} else {
				fmt.Println("Résultat :", result)
			}
		case 8:
			fmt.Println("Au revoir!")
			return
		default:
			fmt.Println("Option invalide. Veuillez réessayer.")
		}
	}
}