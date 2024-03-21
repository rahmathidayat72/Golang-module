package golangmodule

import (
	"crypto/rand"
	"errors"
	"fmt"
	"mime/multipart"
	"net/http"
	"net/url"
	"regexp"
	"strconv"
	"strings"
	"time"
	"unicode"
)

// ....validasi untuk inputan yang wajib diisi....
func ValidasiRequired(inputs ...string) error {
	for _, input := range inputs {
		if strings.TrimSpace(input) == "" {
			return errors.New("error, tidak boleh ada kolom yang kosong")
		}
	}
	return nil
}

// code untuk testing
// func main() {
// 	// Contoh penggunaan
// 	fmt.Print("Masukkan data: ")
// 	var input string
// 	fmt.Scanln(&input)
// 	err := validasi.ValidasiRequired(input)
// 	if err != nil {
// 		fmt.Println("Error:", err)
// 	} else {
// 		fmt.Println("Input valid, lanjutkan pengolahan data.")
// 	}
// }

//....****....

// ....validasi untuk inputan yang berformat email....

func ValidasiEmail(email string) error {
	// Gunakan regex untuk memeriksa apakah email memiliki format yang valid
	emailRegex := `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`
	match, err := regexp.MatchString(emailRegex, email)
	if err != nil {
		return err
	}
	if !match {
		return errors.New("format email tidak valid")
	}
	return nil
}

// code untuk testing
// func main() {
// 	// Contoh penggunaan
// 	fmt.Print("Masukkan alamat email: ")
// 	var email string
// 	fmt.Scanln(&email)
// 	err := validasi.ValidasiEmail(email)
// 	if err != nil {
// 		fmt.Println("Error:", err)
// 	} else {
// 		fmt.Println("Email valid, lanjutkan pengolahan data.")
// 	}
// }

//....****....

// ....validasi untuk inputan no telepon....
func ValidasiPhoneNumber(nomorTelepon string) error {
	// Gunakan regex untuk memeriksa apakah nomor telepon memiliki format yang valid
	// Pada contoh ini, diasumsikan nomor telepon menggunakan format internasional tanpa tanda '+' di depan
	teleponRegex := `^\d{1,4}\d{1,14}$`
	match, err := regexp.MatchString(teleponRegex, nomorTelepon)
	if err != nil {
		return err
	}
	if !match {
		return errors.New("format nomor telepon tidak valid")
	}
	return nil
}

//code untuk testing
// func main() {
// 	// Contoh penggunaan
// 	fmt.Print("Masukkan nomor telepon: ")
// 	var nomorTelepon string
// 	fmt.Scanln(&nomorTelepon)
// 	err := validasi.ValidasiPhoneNumber(nomorTelepon)
// 	if err != nil {
// 		fmt.Println("Error:", err)
// 	} else {
// 		fmt.Println("Nomor telepon valid, lanjutkan pengolahan data.")
// 	}
// }

//....****....

// ....input kombinasi password....
func InputCombinationPassword(input string) error {
	if len(input) < 8 {
		return errors.New("input harus memiliki setidaknya 8 karakter")
	}
	var (
		hasUppercase   bool
		hasLowercase   bool
		hasNumber      bool
		hasSpecialChar bool
	)
	for _, char := range input {
		switch {
		case unicode.IsUpper(char):
			hasUppercase = true
		case unicode.IsLower(char):
			hasLowercase = true
		case unicode.IsDigit(char):
			hasNumber = true
		case unicode.IsPunct(char) || unicode.IsSymbol(char):
			hasSpecialChar = true
		}
		// Jika telah memenuhi semua kriteria, keluar dari loop
		if hasUppercase && hasLowercase && hasNumber && hasSpecialChar {
			break
		}
	}
	// Periksa apakah setiap kriteria terpenuhi
	if !hasUppercase || !hasLowercase || !hasNumber || !hasSpecialChar {
		return errors.New("input tidak memenuhi kriteria kombinasi")
	}

	return nil
}

//code untuk testing
// func main() {
// 	// Contoh penggunaan
// 	fmt.Print("Masukkan kata sandi: ")
// 	var password string
// 	fmt.Scanln(&password)
// 	err := validasi.InputCombinationPassword(password)
// 	if err != nil {
// 		fmt.Println("Error:", err)
// 	} else {
// 		fmt.Println("Kata sandi memenuhi kriteria kombinasi.")
// 	}
// }

//Catatan
// * inputan harus di awali dengan huruf Uppercase.
// * inputan wajib berisi kombinasi huruf Uppercase, lowercase, angka dan carakter.
// * setiap element dalam poin ke 2 wajib di penuhi.
// * inputan minimal terdiri dari 8 karakter.

//....****....

// .... Genred UUID....
// generateRandomBytes menghasilkan slice byte acak sepanjang n
func generateRandomBytes(n int) ([]byte, error) {
	b := make([]byte, n)
	_, err := rand.Read(b)
	if err != nil {
		return nil, err
	}
	return b, nil
}

// generateUUIDV4 menghasilkan UUID versi 4
func GenerateUUIDV4() string {
	randomBytes, err := generateRandomBytes(16)
	if err != nil {
		return ""
	}

	// Set versi 4
	randomBytes[6] = (randomBytes[6] & 0x0F) | 0x40
	// Set empat bit pertama menjadi "10"
	randomBytes[8] = (randomBytes[8] & 0x3F) | 0x80

	// Format sebagai string UUID
	uuid := fmt.Sprintf("%x-%x-%x-%x-%x", randomBytes[0:4], randomBytes[4:6], randomBytes[6:8], randomBytes[8:10], randomBytes[10:16])

	return uuid
}

//code untuk testing
// func main() {
// 	// Contoh penggunaan
// 	uuid, err := validasi.GenerateUUIDV4()
// 	if err != nil {
// 		fmt.Println("Error:", err)
// 		return
// 	}
// 	fmt.Println("UUID baru:", uuid)
// }

//....****....

// ....ValidasiFormatGambar memeriksa apakah file memiliki format gambar yang diizinkan....
func ImageFormatValidation(fileHeader *multipart.FileHeader) error {
	// Membuka file untuk mendapatkan informasi MIME
	file, err := fileHeader.Open()
	if err != nil {
		return err
	}
	defer file.Close()
	// Membaca 512 byte pertama dari file untuk mendapatkan tipe MIME
	buffer := make([]byte, 512)
	_, err = file.Read(buffer)
	if err != nil {
		return err
	}
	// Mendapatkan tipe MIME dari file
	fileType := http.DetectContentType(buffer)
	// Memeriksa apakah tipe MIME sesuai dengan format gambar yang diizinkan
	if !strings.HasPrefix(fileType, "image/jpeg") && !strings.HasPrefix(fileType, "image/png") && !strings.HasPrefix(fileType, "image/jpg") && !strings.HasPrefix(fileType, "image/gif") {
		return errors.New("format gambar tidak diizinkan")
	}

	return nil
}

//code untuk testing
// func main() {
// 	// Contoh penggunaan
// 	// Anda dapat mengganti pathGambar dengan path file gambar yang ingin Anda uji
// 	fileHeader := &multipart.FileHeader{Filename: "example.jpg"}
// 	err := validasi.ImageFormatValidation(fileHeader)
// 	if err != nil {
// 		fmt.Println("Error:", err)
// 	} else {
// 		fmt.Println("Format gambar diizinkan, lanjutkan pemrosesan file.")
// 	}
// }

//....****....

// ....Validasi Format File Office....
func ValidasiFormatFileOffice(namaFile string) error {
	ekstensiDokumen := []string{".doc", ".docx", ".xls", ".xlsx", ".ppt", ".pptx", ".pdf"}
	for _, ekstensi := range ekstensiDokumen {
		if strings.HasSuffix(namaFile, ekstensi) {
			return nil
		}
	}
	return errors.New("format file tidak valid")
}

// code untuk testing
// func main() {
// 	// Contoh penggunaan
// 	fmt.Print("Masukkan nama file: ")
// 	var namaFile string
// 	fmt.Scanln(&namaFile)
// 	err := validasi.ValidasiFormatFileOffice(namaFile)
// 	if err != nil {
// 		fmt.Println("Error:", err)
// 	} else {
// 		fmt.Println("Format file valid untuk dokumen kantor, lanjutkan pemrosesan file.")
// 	}
// }

//....****....

// ....ValidasiURL memeriksa apakah input memiliki format URL yang valid....
func ValidasiURL(input string) error {
	u, err := url.Parse(input)
	if err != nil || u.Scheme == "" || u.Host == "" {
		return errors.New("format URL tidak valid")
	}
	return nil
}

//code untuk testing
// func main() {
// 	// Contoh penggunaan
// 	fmt.Print("Masukkan URL: ")
// 	var inputURL string
// 	fmt.Scanln(&inputURL)
// 	err := validasi.ValidasiURL(inputURL)
// 	if err != nil {
// 		fmt.Println("Error:", err)
// 	} else {
// 		fmt.Println("Format URL valid, lanjutkan pemrosesan.")
// 	}
// }

//....****....

// ....ValidasiFormatVideo....
func ValidasiFormatVideo(namaFile string) error {
	// Daftar ekstensi file video yang diizinkan
	ekstensiVideo := []string{".mp4", ".mkv", ".avi", ".mov", ".wmv"}
	// Mengonversi nama file menjadi huruf kecil agar case-insensitive
	namaFile = strings.ToLower(namaFile)
	// Memeriksa apakah nama file berakhir dengan salah satu ekstensi video yang diizinkan
	valid := false
	for _, ekstensi := range ekstensiVideo {
		if strings.HasSuffix(namaFile, ekstensi) {
			valid = true
			break
		}
	}
	if !valid {
		return errors.New("format video tidak diizinkan")
	}
	return nil
}

//code untuk testing
// func main() {
// 	// Contoh penggunaan
// 	fmt.Print("Masukkan nama file video: ")
// 	var namaFile string
// 	fmt.Scanln(&namaFile)
// 	err := validasi.ValidasiFormatVideo(namaFile)
// 	if err != nil {
// 		fmt.Println("Error:", err)
// 	} else {
// 		fmt.Println("Format video valid, lanjutkan pemrosesan file.")
// 	}
// }

//....****....

// ....ValidasiFormatAudio....
func ValidasiFormatAudio(namaFile string) error {
	// Daftar ekstensi file audio yang diizinkan
	ekstensiAudio := []string{".mp3", ".wav", ".ogg", ".flac", ".aac"}
	// Mengonversi nama file menjadi huruf kecil agar case-insensitive
	namaFile = strings.ToLower(namaFile)
	// Memeriksa apakah nama file berakhir dengan salah satu ekstensi audio yang diizinkan
	valid := false
	for _, ekstensi := range ekstensiAudio {
		if strings.HasSuffix(namaFile, ekstensi) {
			valid = true
			break
		}
	}
	if !valid {
		return errors.New("format audio tidak diizinkan")
	}
	return nil
}

//code untuk testing
// func main() {
// 	// Contoh penggunaan
// 	fmt.Print("Masukkan nama file audio: ")
// 	var namaFile string
// 	fmt.Scanln(&namaFile)
// 	err := validasi.ValidasiFormatAudio(namaFile)
// 	if err != nil {
// 		fmt.Println("Error:", err)
// 	} else {
// 		fmt.Println("Format audio valid, lanjutkan pemrosesan file.")
// 	}
// }

//....****....

// ....ValidasiFormatTanggal saat ini (DD-MM-YYYY)....
// untuk fungsi ValidateCurrentDateFormat tahun maksimal di isi adalah tahun saat aplikasi di jalankan (tahun saat di inputkan)
func ValidateCurrentDateFormat(input string) error {
	// Hapus tanda '-' jika ada
	input = strings.ReplaceAll(input, "-", "")

	// Format regex untuk DDMMYYYY
	formatRegex := `^\d{2}\d{2}\d{4}$`

	// Memeriksa apakah input sesuai dengan format regex
	match, err := regexp.MatchString(formatRegex, input)
	if err != nil {
		return errors.New("error dalam memeriksa format tanggal")
	}

	if !match {
		return errors.New("format tanggal tidak valid. Gunakan format DD-MM-YYYY atau DDMMYYYY")
	}

	// Pecah input menjadi tanggal, bulan, dan tahun
	tanggalStr := input[:2]
	bulanStr := input[2:4]
	tahunStr := input[4:]

	// Konversi ke tipe data numerik
	tanggal, err := strconv.Atoi(tanggalStr)
	if err != nil || tanggal < 1 || tanggal > 31 {
		return errors.New("tanggal tidak valid. Harus antara 1 dan 31")
	}

	bulan, err := strconv.Atoi(bulanStr)
	if err != nil || bulan < 1 || bulan > 12 {
		return errors.New("bulan tidak valid. Harus antara 1 dan 12")
	}

	tahun, err := strconv.Atoi(tahunStr)
	if err != nil || tahun > time.Now().Year() {
		return errors.New("tahun tidak valid. Harus setara atau sebelum tahun sekarang")
	}
	// Tambahkan tanda '-' untuk memformat ulang input
	formattedInput := fmt.Sprintf("%s-%s-%s", input[:2], input[2:4], input[4:])
	fmt.Println("Tanggal telah diformat ulang:", formattedInput)

	return nil
}

// untuk fungsi DateFormatValidation tahun dapat di isi dengan bebas asal terdiri dari 4 angka
func DateFormatValidation(input string) error {
	// Hapus tanda '-' jika ada
	input = strings.ReplaceAll(input, "-", "")

	// Format regex untuk DDMMYYYY
	formatRegex := `^\d{2}\d{2}\d{4}$`

	// Memeriksa apakah input sesuai dengan format regex
	match, err := regexp.MatchString(formatRegex, input)
	if err != nil {
		return errors.New("error dalam memeriksa format tanggal")
	}

	if !match {
		return errors.New("format tanggal tidak valid. Gunakan format DD-MM-YYYY atau DDMMYYYY")
	}

	// Pecah input menjadi tanggal dan bulan
	tanggalStr := input[:2]
	bulanStr := input[2:4]

	// Konversi ke tipe data numerik
	tanggal, err := strconv.Atoi(tanggalStr)
	if err != nil || tanggal < 1 || tanggal > 31 {
		return errors.New("tanggal tidak valid. Harus antara 1 dan 31")
	}

	bulan, err := strconv.Atoi(bulanStr)
	if err != nil || bulan < 1 || bulan > 12 {
		return errors.New("bulan tidak valid. Harus antara 1 dan 12")
	}

	// Tambahkan tanda '-' untuk memformat ulang input
	formattedInput := fmt.Sprintf("%s-%s-%s", input[:2], input[2:4], input[4:])
	fmt.Println("Tanggal telah diformat ulang:", formattedInput)

	return nil
}

//code untuk testing
// func main() {
// 	// Contoh penggunaan
// 	fmt.Print("Masukkan tanggal (DD-MM-YYYY atau DDMMYYYY): ")
// 	var tanggalInput string
// 	fmt.Scanln(&tanggalInput)

// 	err := validasi.ValidateCurrentDateFormat(tanggalInput)

// 	if err != nil {
// 		fmt.Println("Error:", err)
// 	} else {
// 		fmt.Println("Format tanggal valid, lanjutkan pemrosesan.")
// 	}
// }

//....****....
