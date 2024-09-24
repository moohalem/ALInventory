package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"os/exec"
	"strconv"
)

// Ingredients Model
type Ingredients struct {
	Name   string
	Weight int
}

func NewIngredients(name string, amount int) *Ingredients {
	return &Ingredients{
		Name:   name,
		Weight: amount,
	}
}

func (i *Ingredients) display() {
	fmt.Printf("%s: %d\n", i.Name, i.Weight)
}

func (i *Ingredients) UpdateWeight(newWeight int) {
	i.Weight = newWeight
}

func main() {
	var ingredients []*Ingredients

	// Load existing ingredients from file
	err := loadIngredients(&ingredients)
	if err != nil {
		initIngredients(ingredients)
		_ = loadIngredients(&ingredients)
	}

	for {
		clearScreen()
		fmt.Println("ALInventory")
		fmt.Println("MENU")
		fmt.Println("1. Inventory Check")
		fmt.Println("2. Inventory Update")
		fmt.Println("3. Save Inventory")
		fmt.Println("4. Exit")

		choice, err := GetInputInt("Enter your choice: ")
		if err != nil {
			fmt.Println(err)
			continue
		}

		switch choice {
		case 1:
			checkInventory(ingredients)
		case 2:
			updateInventory(ingredients)
		case 3:
			saveIngredients(ingredients)
		case 4:
			return
		default:
			fmt.Println("Invalid choice. Please try again.")
		}
	}
}

func checkInventory(ingredients []*Ingredients) {
	clearScreen()
	fmt.Println("INVENTORY CHECK")
	for i, ingredient := range ingredients {
		fmt.Printf("%d. %s\n", i+1, ingredient.Name)
	}
	fmt.Println("7. Display all")
	fmt.Println("8. Back")

	choice, err := GetInputInt("Enter your choice: ")
	if err != nil {
		fmt.Println(err)
		return
	}

	if choice >= 1 && choice <= len(ingredients) {
		clearScreen()
		ingredients[choice-1].display()
		fmt.Printf("\n\nPress any key to continue ...")
		_, _ = fmt.Scanln()
		return
	} else if choice == 7 {
		clearScreen()
		for _, ingredient := range ingredients {
			ingredient.display()
		}
		fmt.Printf("\n\nPress any key to continue ...")
		_, _ = fmt.Scanln()
	} else if choice == 8 {
		return
	} else {
		fmt.Println("Invalid choice. Please try again.")
	}
}

func updateInventory(ingredients []*Ingredients) {
	clearScreen()
	for i, ingredient := range ingredients {
		fmt.Printf("%d. %s: %d Grams\n", i+1, ingredient.Name, ingredient.Weight)
	}

	choice, err := GetInputInt("Enter the ingredient number to update weight: ")
	if err != nil || choice < 1 || choice > len(ingredients) {
		fmt.Println("Invalid choice.")
		return
	}

	newWeight, err := GetInputInt("Enter new weight: ")
	if err != nil {
		fmt.Println(err)
		return
	}

	ingredients[choice-1].UpdateWeight(newWeight)
	fmt.Printf("Updated %s weight to %d.\n\n", ingredients[choice-1].Name, newWeight)
	fmt.Println("Press any key to continue ...")
	_, _ = fmt.Scanln()
	updateInventory(ingredients)
}

func saveIngredients(ingredients []*Ingredients) {
	file, err := os.Create("ingredients.json")
	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			return
		}
	}(file)

	encoder := json.NewEncoder(file)
	err = encoder.Encode(ingredients)
	if err != nil {
		fmt.Println("Error saving ingredients:", err)
	}
	fmt.Println("Inventory saved successfully.")
}

func loadIngredients(ingredients *[]*Ingredients) error {
	file, err := os.Open("ingredients.json")
	if err != nil {
		if os.IsNotExist(err) {
			fmt.Println("No existing inventory found, starting fresh.")
			return err
		}
		fmt.Println("Error opening file:", err)
		return err
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			return
		}
	}(file)

	decoder := json.NewDecoder(file)
	err = decoder.Decode(ingredients)
	if err != nil {
		fmt.Println("Error loading ingredients:", err)
	}
	return nil
}

func GetInputInt(prompt string) (int, error) {
	var x string
	fmt.Print(prompt)
	_, err := fmt.Scanln(&x)
	if err != nil {
		return 0, err
	} else if x == "" {
		return 0, errors.New("value must not be empty")
	}
	value, err := strconv.Atoi(x)
	return value, nil
}

func initIngredients(ingredients []*Ingredients) {
	ingredients = append(ingredients, NewIngredients("Transglutaminase", 0))
	ingredients = append(ingredients, NewIngredients("Wendaphos", 0))
	ingredients = append(ingredients, NewIngredients("M100FT", 0))
	ingredients = append(ingredients, NewIngredients("ISP", 0))
	ingredients = append(ingredients, NewIngredients("Karagenan", 0))
	ingredients = append(ingredients, NewIngredients("M100F", 0))
	saveIngredients(ingredients)
}

func clearScreen() {
	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	err := cmd.Run()
	if err != nil {
		fmt.Println(err)
	}
}
