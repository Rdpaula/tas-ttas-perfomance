package main

import (
	csvEncoder "encoding/csv"
	"fmt"
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
	threads := []int{10, 100, 1000}
	executions := []int{10, 100, 1000}
	criticalLoads := []int{1, 10, 100, 1000}
	lockTypes := []string{"TAS", "TTAS"}

	for _, thread := range threads {
		for _, exec := range executions {
			for _, load := range criticalLoads {
				for _, lockType := range lockTypes {
					fmt.Println("running for", thread, exec, load, lockType)
					var lock Lock
					if lockType == "TAS" {
						lock = &TASLock{}
					} else {
						lock = &TTASLock{}
					}

					resultTime := Run(thread, exec, load, lock)
					row := []string{strconv.Itoa(thread), strconv.Itoa(exec), strconv.Itoa(load), lockType, resultTime.String()}
					csv = append(csv, row)
				}
			}
		}
	}

	defer csvWriter.Flush()
	csvWriter.WriteAll(csv)
}
