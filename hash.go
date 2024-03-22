package golangmodule

import (
	"crypto/sha256"
	"crypto/sha512"
	"encoding/hex"

	"golang.org/x/crypto/bcrypt"
)

// ....Fungsi untuk menghasilkan hash SHA-256 dari data....
func HashSHA256(data []byte) (string, error) {
	hasher := sha256.New()
	_, err := hasher.Write(data)
	if err != nil {
		return "", err
	}
	hashedData := hex.EncodeToString(hasher.Sum(nil))
	return hashedData, nil
}

// HashSHA384 menghasilkan hash SHA-384 dari data
func HashSHA384(data []byte) (string, error) {
	hasher := sha512.New384()
	_, err := hasher.Write(data)
	if err != nil {
		return "", err
	}
	hashedData := hex.EncodeToString(hasher.Sum(nil))
	return hashedData, nil
}

// HashSHA512 menghasilkan hash SHA-512 dari data
func HashSHA512(data []byte) (string, error) {
	hasher := sha512.New()
	_, err := hasher.Write(data)
	if err != nil {
		return "", err
	}
	hashedData := hex.EncodeToString(hasher.Sum(nil))
	return hashedData, nil
}

//code untuk testing
// func main() {
// 	// Contoh penggunaan
// 	dataToHash := []byte("Hello, World!")
// 	hashedResult256, err := hashsing.HashSHA256(dataToHash)
// 	hashedResult384, err := hashsing.HashSHA384(dataToHash) // gunakan salah satu sesuai format yang diinginkan
// 	hashedResult512, err := hashsing.HashSHA512(dataToHash)
// 	if err != nil {
// 		fmt.Println("Error:", err)
// 		return
// 	}
// 	fmt.Printf("Data: %s\n", string(dataToHash))
// 	fmt.Printf("Hash SHA-256: %s\n", hashedResult256)
// 	fmt.Printf("Hash SHA-384: %s\n", hashedResult384)
// 	fmt.Printf("Hash SHA-512: %s\n", hashedResult512)
// }

func CheckPassword(password, hash string) bool {
	passwordBytes := []byte(password)
	hashBytes := []byte(hash)

	err := bcrypt.CompareHashAndPassword(hashBytes, passwordBytes)
	return err == nil
}
