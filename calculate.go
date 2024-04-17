package golangmodule

import (
	"fmt"
	"reflect"
	"strconv"
)

// CalculatePercentage menghitung presentasi dari dua nilai nilai int atau flaot64.
func CalculatePercentageInt(currentValue, totalValue interface{}) (float64, error) {
	current, ok := convertToFloat64(currentValue)
	if !ok {
		return 0.0, fmt.Errorf("nilai current tidak dapat dikonversi ke float64")
	}

	total, ok := convertToFloat64(totalValue)
	if !ok {
		return 0.0, fmt.Errorf("nilai total tidak dapat dikonversi ke float64")
	}

	if total == 0 {
		return 0.0, fmt.Errorf("total nilai tidak boleh nol")
	}

	percentage := (current / total) * 100
	return percentage, nil
}

// convertToFloat64 mengkonversi nilai ke float64.
func convertToFloat64(value interface{}) (float64, bool) {
	switch v := value.(type) {
	case float64:
		return v, true
	case int:
		return float64(v), true
	case uint:
		return float64(v), true
	default:
		fmt.Printf("Tipe data %v tidak didukung.\n", reflect.TypeOf(value))
		return 0.0, false
	}
}

//cara menggunakan code
// func main() {
// 	// Contoh penggunaan fungsi CalculatePercentage
// 	currentValue := 35050000
// 	totalValue := 50000000

// 	percentage, err := calculate.CalculatePercentageInt(currentValue, totalValue)
// 	if err != nil {
// 		fmt.Println("Error:", err)
// 		return
// 	}

// 	fmt.Printf("Presentasi: %.2f%%\n", percentage)
// }

// CalculatePercentageString menghitung presentasi dari dua nilai dalam bentuk string.
func CalculatePercentageString(currentValue, totalValue string) (string, error) {
	current, err := strconv.ParseFloat(currentValue, 64)
	if err != nil {
		return "", fmt.Errorf("nilai current tidak dapat dikonversi ke float64")
	}

	total, err := strconv.ParseFloat(totalValue, 64)
	if err != nil {
		return "", fmt.Errorf("nilai total tidak dapat dikonversi ke float64")
	}

	if total == 0 {
		return "", fmt.Errorf("total nilai tidak boleh nol")
	}

	percentage := (current / total) * 100
	return fmt.Sprintf("%.2f", percentage), nil
}

// convertToString mengkonversi nilai ke string.
func convertToString(value interface{}) (string, bool) {
	switch v := value.(type) {
	case float64:
		return fmt.Sprintf("%.2f", v), true
	case int:
		return strconv.Itoa(v), true
	case uint:
		return strconv.FormatUint(uint64(v), 10), true
	case string:
		return v, true
	default:
		fmt.Printf("Tipe data %v tidak didukung.\n", reflect.TypeOf(value))
		return "", false
	}
}

//cara menggunakan code
// func main() {
// 	// Contoh penggunaan fungsi CalculatePercentageString
// 	currentValue := "35050000"
// 	totalValue := "50000000"
// 	percentage, err := calculate.CalculatePercentageString(currentValue, totalValue)
// 	if err != nil {
// 		fmt.Println("Error:", err)
// 		return
// 	}
// 	fmt.Printf("Presentasi: %s%%\n", percentage)
// }
