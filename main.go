package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"

	"github.com/go-pg/pg"
	"github.com/go-pg/pg/orm"
	"readfile.com/practice/databasesetup"
)

type Details struct {
	Name  string
	Age   int
	Phone int64
}

//it creates table in database
func createtable(db *pg.DB) {
	m := []interface{}{
		(*Details)(nil),
	}
	for _, m := range m {
		err := db.Model(m).CreateTable(&orm.CreateTableOptions{
			IfNotExists: true,
			Varchar:     50,
		})

		if err != nil {
			fmt.Println(err)
		}
	}
}

//it inserts the data to the table
func adddatatotable(db *pg.DB, Detail *[]Details) {

	_, err := db.Model(Detail).Insert() //Inserts the details to the db
	if err != nil {
		fmt.Println(err)
	}
}

//reads the specified file and returns a struc details
func readfile() []Details {
	file, err := os.Open("sample.csv") //opens the file
	defer file.Close()
	if err != nil {
		fmt.Println("Failed to open file")

	}
	fmt.Println("File opened succesfully")

	csvfile, err := csv.NewReader(file).ReadAll() //reads the data from the files
	if err != nil {
		fmt.Println("Not able to read value")

	}
	var Detail []Details
	var det Details
	temp := csvfile[1:]
	for _, details := range temp {

		intage, _ := strconv.Atoi(details[1])   //convert string to int
		intphone, _ := strconv.Atoi(details[2]) //
		det = Details{
			Name:  details[0],
			Age:   intage,
			Phone: int64(intphone),
		}
		Detail = append(Detail, det) //adds the value from det to Detail

	}
	return Detail
}

//Selects all the details from the database
func getdatafromtable(db *pg.DB) []Details {
	var det []Details
	err := db.Model(&det).Select()
	if err != nil {
		fmt.Println(err)
	}

	return det
}

//Writes the data to the file
func writefile(detail *[]Details) {
	file, err := os.Create("detail.csv")
	defer file.Close()
	if err != nil {
		fmt.Println("file not created")
	}
	fmt.Println("File created")
	writefile := csv.NewWriter(file)
	defer writefile.Flush()
	for _, det := range *detail {
		row := []string{det.Name, strconv.Itoa(det.Age), strconv.Itoa(int(det.Phone))}
		_ = writefile.Write(row)
	}

}

func main() {
	Detail := readfile()
	db := databasesetup.Setup() //calling connection to db
	defer db.Close()
	if db == nil {
		fmt.Println("Connection Failed")
	}
	createtable(db)
	adddatatotable(db, &Detail)
	detail := getdatafromtable(db)
	writefile(&detail)

}
