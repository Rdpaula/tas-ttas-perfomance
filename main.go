package main

import (
	csvEncoder "encoding/csv"
	"log"
	"os"
	"strconv"
)

func main() {
	file, err := os.Create("csv_table.csv")
	if err != nil {
		log.Fatalln("failed to open file", err)
	}
	defer file.Close()
	csvWriter := csvEncoder.NewWriter(file)
	defer csvWriter.Flush()

	csv := [][]string{
		{"num_threads", "executions", "critical_load", "lock_type", "time"},
	}
	threads := []int{1, 10, 100}
	executions := []int{1, 10, 100}
	criticalLoad := []int{1, 10, 100}
	lockType := []string{"TAS", "TTAS"}

	for _, thread := range threads {
		for _, exec := range executions {
			for _, load := range criticalLoad {
				for _, lock := range lockType {
					resultTime := Run(thread, exec, load, lock)
					row := []string{strconv.Itoa(thread), strconv.Itoa(exec), strconv.Itoa(load), lock, resultTime.String()}
					csv = append(csv, row)
				}
			}
		}
	}

	defer csvWriter.Flush()
	csvWriter.WriteAll(csv)
}
