package main

import (
	"fmt"
	"os"

	"github.com/mehanizm/airtable"
)

func fetchRecords(filename string) error {
	client := airtable.NewClient(os.Getenv("AIRTABLE_TOKEN"))
	table := client.GetTable("appchP0LatuPh21WF", "Table1")
	if table == nil {
		return fmt.Errorf("failed to get table")
	}
	records, err := table.GetRecords().Do()
	if err != nil {
		return fmt.Errorf("failed to fetch records: %w", err)
	}
	out, err := os.Create(filename)
	if err != nil {
		return fmt.Errorf("failed to create file '%s': %w", filename, err)
	}
	defer out.Close()
	for _, record := range records.Records {
		fmt.Fprintf(out, "%+#v\n\n\n", record.Fields)
	}
	fmt.Printf("Fetched %d records to %s\n", len(records.Records), filename)
	return nil
}
