package main

import "fmt"

func main() {
	ns := GetNutritionalScore(
		NutritionalData{
			Energy:              EnergyFromKcal(0),
			Sugars:              SugarGram(10),
			SaturatedFattyAcids: StaturatedFattyAcids(),
			Sodium:              SodiumMiligram(),
			Fruits:              FruitsPercent(),
			Fibre:               FibreGram(),
			Protein:             ProteinGram(),
		}, Food)
	fmt.Println("Supplyment That I Eat's : %d\n", ns.Value)
}
