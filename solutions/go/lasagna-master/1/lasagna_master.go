package lasagnamaster

// TODO: define the 'PreparationTime()' function
func PreparationTime(layers []string, time int) int {
    if time == 0 {
        time = 2
    }

    return len(layers) * time
}

func Quantities(layers []string) (int, float64) {
    totalNoodles := 0
    totalSauce := 0.0
    for i := 0; i < len(layers); i++ {
  		if layers[i] == "sauce" {
            totalSauce = totalSauce + 0.2
        }else if layers[i] == "noodles" { 
            totalNoodles = totalNoodles + 50
        }
	}

    return totalNoodles, totalSauce
}

func AddSecretIngredient(friendList []string, myList []string){
    secretIngredient := friendList[len(friendList)-1]
    myList[len(myList)-1] = secretIngredient
}

func ScaleRecipe(qty []float64, amount int) []float64{
    multiplier := float64(amount) / 2
	result := []float64{}

    for i := 0; i < len(qty); i++ {
        result = append(result, qty[i] * multiplier)
    }

    return result
}

// Your first steps could be to read through the tasks, and create
// these functions with their correct parameter lists and return types.
// The function body only needs to contain `panic("")`.
//
// This will make the tests compile, but they will fail.
// You can then implement the function logic one by one and see
// an increasing number of tests passing as you implement more
// functionality.
