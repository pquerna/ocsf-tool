package main

import (
	"os"

	"github.com/valllabh/ocsf-schema-processor/ocsf"
	"github.com/valllabh/ocsf-schema-processor/ocsf/mappers/protobuff"
	"golang.org/x/exp/maps"
)

func main() {

	// Loads to provided OCSF schema version in schema.json
	ocsfSchema, _ := ocsf.LoadOCSFSchema()

	mapToProtoFile(ocsfSchema)
}

func mapToProtoFile(ocsfSchema ocsf.OCSFSchema) {

	mapper := protobuff.NewMapper(ocsfSchema)

	output := mapper.Marshal(maps.Values(ocsfSchema.Classes))

	WriteToFile("output.proto", []byte(output))
}
func WriteToFile(filePath string, data []byte) error {
	// Open the file with write permissions, create it if it doesn't exist, and truncate it
	file, err := os.OpenFile(filePath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		return err
	}
	defer file.Close()

	// Write the data to the file
	_, err = file.Write(data)
	if err != nil {
		return err
	}

	return nil
}
