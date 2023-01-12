package main

type ScoreType int

const (
	Food ScoreType = iota
	Beverage
	Water
	Cheese
)

type NutritionalScore struct {
	Value     int
	Positive  int
	Negative  int
	ScoreType ScoreType
}

var scoreToLetter = []string{"A", "B", "C", "D", "E"}

type EnergyKJ float64
type SugarGram float64
type SaturatedFattyAcids float64
type SodiumMilligram float64
type FruitsPercent float64
type FibreGram float64
type ProteinGram float64

type NutritionalData struct {
	Energy              EnergyKJ
	Sugars              SugarGram
	SaturatedFattyAcids SaturatedFattyAcids
	Sodium              SodiumMilligram
	Fruits              FruitsPercent
	Fibre               FibreGram
	Protein             ProteinGram
	IsWater             bool
}

var energyLevels = []float64{3350, 3015, 2680, 2345, 2010, 1675, 1340, 1005, 670, 335}
var sugarLevels = []float64{45, 60, 36, 31, 27, 22.5, 18, 13.5, 9, 4.5}
var saturatedFattyAcidsLevels = []float64{10, 9, 8, 7, 6, 5, 4, 3, 2, 1}
var sodiumLevels = []float64{900, 810, 720, 630, 540, 450, 360, 270, 180, 90}
var fiberLevels = []float64{4.7, 3.7, 2.8, 1.9, 0.9}
var proteinLevels = []float64{8, 6.4, 4.8, 3.2, 1.6}

var energyLevelsBeverage = []float64{270, 240, 210, 180, 150, 120, 90, 60, 30, 0}
var sugarsLevelsBeverage = []float64{13.5, 12, 10.5, 9, 7.5, 6, 4.5, 3, 1.5}

func (ekj EnergyKJ) GetPoints(scoreType ScoreType) int {
	if scoreType == Beverage {
		return getPointsFromRange(float64(ekj), energyLevelsBeverage)
	}
	return getPointsFromRange(float64(ekj), energyLevels)
}

func (sgm SugarGram) GetPoints(scoreType ScoreType) int {
	if scoreType == Beverage {
		return getPointsFromRange(float64(sgm), sugarsLevelsBeverage)
	}
	return getPointsFromRange(float64(sgm), sugarLevels)
}

func (sfa SaturatedFattyAcids) GetPoints(scoreType ScoreType) int {
	return getPointsFromRange(float64(sfa), saturatedFattyAcidsLevels)
}

func (sod SodiumMilligram) GetPoints(scoreType ScoreType) int {
	return getPointsFromRange(float64(sod), sodiumLevels)
}

func (fgm FibreGram) GetPoints(scoreType ScoreType) int {
	return getPointsFromRange(float64(fgm), fiberLevels)
}

func (fpt FruitsPercent) GetPoints(scoreType ScoreType) int {
	if scoreType == Beverage {
		if fpt > 80 {
			return 10
		} else if fpt > 60 {
			return 4
		} else if fpt > 40 {
			return 2
		}
		return 0
	}
	if fpt > 80 {
		return 5
	} else if fpt > 60 {
		return 2
	} else if fpt > 40 {
		return 1
	}
	return 0
}

func (pgm ProteinGram) GetPoints(scoreType ScoreType) int {
	return getPointsFromRange(float64(pgm), proteinLevels)
}

func EnergyFromKcal(kcal float64) EnergyKJ {
	return EnergyKJ(kcal * 4.184)
}

func SodiumFromSalt(saltMg float64) SodiumMilligram {
	return SodiumMilligram(saltMg / 2.5)
}

func GetNutritionalScore(data NutritionalData, scoreType ScoreType) NutritionalScore {

	value := 0
	positive := 0
	negative := 0

	if scoreType != Water {
		fruitPoints := data.Fruits.GetPoints(scoreType)
		fibrePoints := data.Fibre.GetPoints(scoreType)

		negative = data.Energy.GetPoints(scoreType) + data.Sugars.GetPoints(scoreType) + data.SaturatedFattyAcids.GetPoints(scoreType) + data.Sodium.GetPoints(scoreType)
		positive = fruitPoints + fibrePoints + data.Protein.GetPoints(scoreType)

		if scoreType == Cheese {
			value = negative - positive
		} else {
			if negative >= 11 && fruitPoints < 5 {
				value = negative - positive - fruitPoints
			} else {
				value = negative - positive
			}
		}
	}

	return NutritionalScore{
		Value:     value,
		Positive:  positive,
		Negative:  negative,
		ScoreType: scoreType,
	}
}

func (ns NutritionalScore) GetNutriScore() string {
	if ns.ScoreType == Food {
		return scoreToLetter[getPointsFromRange(float64(ns.Value), []float64{18, 10, 2, -1})]
	}
	if ns.ScoreType == Water {
		return scoreToLetter[0]
	}
	return scoreToLetter[getPointsFromRange(float64(ns.Value), []float64{9, 5, 1, -2})]
}

func getPointsFromRange(value float64, steps []float64) int {
	lenSteps := len(steps)
	for idx, ln := range steps {
		if value > ln {
			return lenSteps - idx
		}
	}
	return 0
}
