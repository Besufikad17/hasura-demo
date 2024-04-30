package constants

var CreateUser string = "mutation CreateUser(   $first_name: String!,   $last_name: String!,   $email: String!,   $phone_number: String!,   $password: String!,   $role: role! ) { 	insert_users_one(object: {      first_name: $first_name,      last_name: $last_name,      email: $email,      phone_number: $phone_number,      password: $password,      role: $role }   ) {     id     first_name     last_name     email     phone_number     password     role   } }"
