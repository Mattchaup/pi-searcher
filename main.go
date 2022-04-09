package main

import (
	"fmt"
	"bufio"
	"os"
	"strconv"
	"time"
)

func main(){
	var pi string

	doss, err := os.Open("pi-million.txt")
	if err != nil{
		panic(err)
	}
	scanner := bufio.NewScanner(doss)

	search := inputNumber()

	debut := time.Now()
	
	for scanner.Scan(){
		line := scanner.Text()
		pi += line
		//isWeird(line)
		//fmt.Println(line)
	}
	finScan := time.Now()

	//fmt.Println(firstNoneInt(pi)) //Réponse : 14523 17125 22801

	index := findNbr(search,pi[2:])
	fin := time.Now()

	tempsTotal := fin.Sub(debut)
	tempsChargement := finScan.Sub(debut)
	tempsCalcul := tempsTotal- tempsChargement

	if index == 0{
		fmt.Println("Ce nombre ne se trouve pas dans les milionnième décimale de pi")
	}else{
		fmt.Println("Votre nombre se trouve à la",index,"ième place")
		fmt.Println("Temps de chargement :",tempsChargement)
		fmt.Println("temps de calcul     :",tempsCalcul)
		fmt.Println("Temps total         :",tempsTotal)
	}
}

func firstNoneInt(pi string)string{
	nbr := []string{}
	for i := 1 ; true ; i++{
		search := strconv.Itoa(i)
		index := findNbr(search,pi[2:])
		if index == 0{
			nbr = append(nbr,search)
			//return ("Le nombre "+search+" est le premier entier qui n'est pas dans le premières millionième décimales de pi")
		}
		
		if i % 1000 == 0{
			fmt.Println(i)
		}

		if len(nbr) >= 3{
			return ("Voici 3 nombres impossibles à trouver : "+nbr[0]+" "+nbr[1]+" "+nbr[2])
		}
		
	}
	return ("fin")
}

func inputNumber()string{
	input := bufio.NewScanner(os.Stdin)
	fmt.Printf("Taper un nombre qui peut se trouver dans pi : ")
	input.Scan()
	return input.Text()
}

//Permet de savoir si un caractère est non désiré
func isWeird(line string) {
	for _,ltr := range line{
		letter := string(ltr)
		if inInt(letter) == false{
			fmt.Println(letter)
		}
	}
}

//Permet de savoir si c'est un entier
func inInt(letter string) bool{
	if letter != "0" && letter != "1" && letter != "2" && letter != "3" && letter != "4" && letter != "5" && letter != "6" && letter != "7" && letter != "8" && letter != "9"{
		return false
	} else{
		return true
	}
}

//Found if a number is in pi number
func findNbr(nbr string,pi string) int{
	var long int = len(nbr)
	var piLong int = len(pi)

	for index := range pi{
		if nbr[0] == pi[index]{
			if long > 1 && index + long <= piLong{
				if isSameNumber(pi[index+1:index+long],nbr[1:],long){
					return index + 1
				}
			}else if long == 1{
				return index + 1
			}
		}
	}
	return 0 
}

//vérifie si la chaine de caractère est identiques
func isSameNumber(piNbr string, inputNbr string,long int) bool{
	for i := 0 ; i < long - 1 ; i++{
		if piNbr[i] != inputNbr[i]{
			return false
		}
	}
	return true
}