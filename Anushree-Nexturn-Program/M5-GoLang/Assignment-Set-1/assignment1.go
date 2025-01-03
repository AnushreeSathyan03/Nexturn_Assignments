package main

import (
	"errors"
	"fmt"
	"strings"
)

// Staff struct to hold staff member data
type Staff struct {
	ID          int
	FullName    string
	YearsOfWork int
	Division    string
}

// Constants for divisions
const (
	Finance    = "Finance"
	Engineering = "Engineering"
	Marketing   = "Marketing"
)

// In-memory staff database
var staffList []Staff

// AddStaffMember adds a new staff member to the database
func AddStaffMember(id int, fullName string, yearsOfWork int, division string) error {
	// Validate ID uniqueness
	for _, staff := range staffList {
		if staff.ID == id {
			return errors.New("staff member with this ID already exists")
		}
	}

	// Validate Years of Work
	if yearsOfWork <= 0 {
		return errors.New("years of work must be greater than 0")
	}

	// Add staff member to the database
	newStaff := Staff{
		ID:          id,
		FullName:    fullName,
		YearsOfWork: yearsOfWork,
		Division:    division,
	}
	staffList = append(staffList, newStaff)
	return nil
}

// SearchStaffMember searches for a staff member by ID or name
func SearchStaffMember(searchTerm string) (*Staff, error) {
	for _, staff := range staffList {
		if fmt.Sprintf("%d", staff.ID) == searchTerm || strings.EqualFold(staff.FullName, searchTerm) {
			return &staff, nil
		}
	}
	return nil, errors.New("staff member not found")
}

// ListStaffByDivision lists staff members in a specific division
func ListStaffByDivision(division string) ([]Staff, error) {
	var filteredStaff []Staff
	for _, staff := range staffList {
		if strings.EqualFold(staff.Division, division) {
			filteredStaff = append(filteredStaff, staff)
		}
	}
	if len(filteredStaff) == 0 {
		return nil, errors.New("no staff members found in this division")
	}
	return filteredStaff, nil
}

// CountStaffMembers counts the staff members in a specific division
func CountStaffMembers(division string) int {
	count := 0
	for _, staff := range staffList {
		if strings.EqualFold(staff.Division, division) {
			count++
		}
	}
	return count
}

func main() {
	// Adding sample staff members
	_ = AddStaffMember(101, "John Doe", 5, Engineering)
	_ = AddStaffMember(102, "Jane Smith", 8, Marketing)
	_ = AddStaffMember(103, "Emily Brown", 3, Finance)
	_ = AddStaffMember(104, "Michael Green", 10, Engineering)

	fmt.Println("=== Staff Management System ===")

	// 1. Add Staff Member
	fmt.Println("\nAdding Staff Member...")
	err := AddStaffMember(105, "Sophia Davis", 2, Marketing) // Valid input
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Successfully Added!")
	}

	// 2. Search Staff Member
	fmt.Println("\nSearching for Staff Member by ID...")
	staff, err := SearchStaffMember("103")
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Printf("Staff Member Found: %+v\n", *staff)
	}

	// 3. List Staff by Division
	fmt.Println("\nListing Staff in Engineering Division...")
	engineeringStaff, err := ListStaffByDivision(Engineering)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		for _, s := range engineeringStaff {
			fmt.Printf("Staff Member: %+v\n", s)
		}
	}

	// 4. Count Staff Members
	fmt.Println("\nCounting Staff Members in Marketing Division...")
	marketingCount := CountStaffMembers(Marketing)
	fmt.Printf("Number of Staff Members in Marketing: %d\n", marketingCount)

	// Bonus: Handling invalid operations
	fmt.Println("\nSearching for a non-existent Staff Member...")
	_, err = SearchStaffMember("999")
	if err != nil {
		fmt.Println("Error:", err)
	}
}