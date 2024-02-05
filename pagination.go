package golangmodule

import (
	"errors"
	"fmt"
	"strconv"
	"strings"

	"gorm.io/gorm"
)

// Pagination adalah struktur yang merepresentasikan parameter pagination.
type Pagination struct {
	PerPage int
	Page    int
	Offset  int
}

// Paginator adalah struktur yang merepresentasikan detail pagination.
type Paginator struct {
	Page        uint   `json:"page"`
	PerPage     uint   `json:"per_page"`
	CurrentPage string `json:"current_page"`
	TotalPage   uint   `json:"total_page"`
	TotalData   uint   `json:"total_data"`
	FilterData  uint   `json:"filter_data"`
}

// BuildPagination membangun struktur Pagination dari parameter query yang diberikan.
func BuildPagination(perPageParam, pageParam string) (Pagination, error) {
	perPage, err := strconv.Atoi(perPageParam)
	if err != nil || perPage <= 0 {
		return Pagination{}, errors.New("Parameter per_page tidak valid")
	}

	page, err := strconv.Atoi(pageParam)
	if err != nil || page <= 0 {
		return Pagination{}, errors.New("Parameter page tidak valid")
	}

	offset := (page - 1) * perPage
	return Pagination{
		PerPage: perPage,
		Page:    page,
		Offset:  offset,
	}, nil
}

// ApplySort melakukan pengurutan dinamis berdasarkan field sort yang valid.
func ApplySort(query *gorm.DB, sortBy, sort string, validSortFields []string) *gorm.DB {
	defaultSort := "id ASC"
	if sort == "desc" {
		defaultSort = "id DESC"
	}

	if contains(validSortFields, sortBy) {
		query = query.Order(fmt.Sprintf("%s %s", sortBy, sort))
	} else {
		query = query.Order(defaultSort)
	}

	return query
}

// ApplyFilters menerapkan filter pada query berdasarkan map filter yang diberikan.
func ApplyFilters(query *gorm.DB, filters map[string]string) *gorm.DB {
	for field, value := range filters {
		if value != "" {
			query = query.Where(fmt.Sprintf("%s = ?", field), value)
		}
	}
	return query
}

// ApplyDynamicSearch melakukan pencarian dinamis pada kolom-kolom tertentu dengan satu string pencarian.
func ApplyDynamicSearch(query *gorm.DB, search string, searchFields ...string) *gorm.DB {
	if search == "" {
		return query
	}

	var orConditions []string
	var searchArgs []interface{}

	searchTerms := strings.Split(search, " ")

	for _, column := range searchFields {
		var columnConditions []string
		for _, term := range searchTerms {
			columnConditions = append(columnConditions, fmt.Sprintf("LOWER(%s) ~* ?", column))
			searchArgs = append(searchArgs, ".*"+term+".*")
		}

		// Gabungkan kondisi pencarian untuk setiap kolom dengan operator AND
		andCondition := "(" + strings.Join(columnConditions, " AND ") + ")"
		orConditions = append(orConditions, andCondition)
	}

	// Gabungkan kondisi pencarian untuk setiap kolom dengan operator OR
	searchQuery := strings.Join(orConditions, " OR ")

	return query.Where(searchQuery, searchArgs...)
}

// CountTotalData menghitung total jumlah rekord dalam tabel.
func CountTotalData(query *gorm.DB, model interface{}) (int64, error) {
	var totalData int64
	err := query.Model(model).Count(&totalData).Error
	return totalData, err
}

// CountFilteredData menghitung total jumlah rekord setelah menerapkan filter.
func CountFilteredData(query *gorm.DB) (int64, error) {
	var filteredData int64
	err := query.Count(&filteredData).Error
	return filteredData, err
}

// CountCurrentPageRange menghitung rentang halaman saat ini.
func CountCurrentPageRange(offset, length int) string {
	currentPageStart := offset + 1
	currentPageEnd := currentPageStart + length - 1
	return fmt.Sprintf("%d - %d", currentPageStart, currentPageEnd)
}

// contains memeriksa apakah sebuah string ada dalam sebuah slice dari string.
func contains(arr []string, item string) bool {
	for _, a := range arr {
		if a == item {
			return true
		}
	}
	return false
}
