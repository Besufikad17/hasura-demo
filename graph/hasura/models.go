package hasura

type Role string

type CreateUserArgs struct {
	First_name   string
	Last_name    string
	Email        string
	Phone_number string
	Password     string
	Role         Role
}

type CreateUserOutput struct {
	Email        *string
	First_name   *string
	Id           int
	Last_name    *string
	Password     *string
	Phone_number *string
	Role         Role
}

type ActionPayload struct {
	SessionVariables map[string]interface{} `json:"session_variables"`
	Input            CreateUserArgs         `json:"input"`
}

type GraphQLError struct {
	Message string `json:"message"`
}

type GraphQLRequest struct {
	Query     string                 `json:"query"`
	Variables map[string]interface{} `json:"variables"`
}
type GraphQLData struct {
	Insert_users_one CreateUserOutput `json:"insert_users_one"`
}
type GraphQLResponse struct {
	Data   GraphQLData    `json:"data,omitempty"`
	Errors []GraphQLError `json:"errors,omitempty"`
}
