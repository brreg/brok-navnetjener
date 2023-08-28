package utils

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/sirupsen/logrus"
)

// ExecuteGraphQLQuery executes a GraphQL query with variables and unmarshals the result into the given data structure.
func ExecuteGraphQLQuery(query string, variables map[string]interface{}, result interface{}) error {
	graphqlURL, exists := os.LookupEnv("GRAPHQL_URL")
	if !exists {
		logrus.Error("GRAPHQL_URL environment variable not set")
		return fmt.Errorf("GRAPHQL_URL environment variable not set")
	}

	logrus.Debug("Executing GraphQL query against:", graphqlURL)

	// Create a new GraphQL query with variables
	queryStruct := struct {
		Query     string                 `json:"query"`
		Variables map[string]interface{} `json:"variables"`
	}{
		Query:     query,
		Variables: variables,
	}

	logrus.Debug("Query:", queryStruct.Query)
	logrus.Debug("Variables:", queryStruct.Variables)

	queryBytes, err := json.Marshal(queryStruct)
	if err != nil {
		logrus.Error("Error marshaling query:", err)
		return fmt.Errorf("Error marshaling query: %s", err)
	}

	// Create a new HTTP request
	req, err := http.NewRequest("POST", graphqlURL, bytes.NewBuffer(queryBytes))
	if err != nil {
		logrus.Error("Error creating request:", err)
		return fmt.Errorf("Error creating request: %s", err)
	}
	req.Header.Set("Content-Type", "application/json")

	// Execute the HTTP request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		logrus.Error("Error executing request:", err)
		return fmt.Errorf("Error executing request: %s", err)
	}
	defer resp.Body.Close()

	// Parse the GraphQL response
	respBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		logrus.Error("Error reading response body:", err)
		return fmt.Errorf("Error reading response body: %s", err)
	}

	err = json.Unmarshal(respBytes, result)
	if err != nil {
		logrus.Error("Error unmarshaling response:", err)
		return fmt.Errorf("Error unmarshaling response: %s", err)
	}

	logrus.Debug("Successfully executed GraphQL query")
	logrus.Debug("Response:", string(respBytes))

	return nil
}
