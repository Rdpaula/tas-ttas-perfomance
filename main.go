package main

import (
	"fmt"
	"os"
	"strconv"
	"sync"
	"encoding/csv"
)

func main(){
	csv := [][]interface{
		{"num_threads", "executions", "critical_load", "lock_type, time"},
	}

	threads := []int{10, 100, 1000}
	executions := []int{10, 100, 1000}
	criticalLoad := []int{1, 100, 10000}
	lockType := []string{"TAS", "TTAS"}

	for thread,indThread in threads{
		for exec,indExec in executions{
			for load,indLoad in criticalLoad{
				for lock,indLock in lockType{
					resultTime := Run(thread, exec, load, lock)
					row := []interface{thread, exec, load, lock, resultTime}
					csv = append(csv, row)
				}
			}
		}
	}
	csvWriter := csv.NewWriter("csvTable")
	defer csvWriter.flush()
	csvWriter.WriteAll(csv)
}