package hasura

import (
	"bytes"
	"encoding/json"
	"errors"
	"github.com/besufikad17/hasura-gql-demo/graph/utils/constants"
	"io"
	"net/http"
	"os"
)

func CreateUser(args CreateUserArgs) (response CreateUserOutput, err error) {
	hasuraResponse, err := execute(args)
	if err != nil {
		return
	}

	if len(hasuraResponse.Errors) != 0 {
		err = errors.New(hasuraResponse.Errors[0].Message)
		return
	}

	response = hasuraResponse.Data.Insert_users_one
	return
}
func execute(variables CreateUserArgs) (response GraphQLResponse, err error) {
	mapVariables := map[string]interface{}{
		"first_name":   variables.First_name,
		"last_name":    variables.Last_name,
		"email":        variables.Email,
		"phone_number": variables.Phone_number,
		"password":     variables.Password,
		"role":         variables.Role,
	}
	reqBody := GraphQLRequest{
		Query:     constants.CreateUser,
		Variables: mapVariables,
	}
	reqBytes, err := json.Marshal(reqBody)
	if err != nil {
		return
	}

	hasuraURL := os.Getenv("HASURA_URL")

	resp, err := http.Post(hasuraURL, "application/json", bytes.NewBuffer(reqBytes))
	if err != nil {
		return
	}

	respBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return
	}

	println(string(reqBytes))

	err = json.Unmarshal(respBytes, &response)
	if err != nil {
		return
	}

	return
}
