# blog-app

set your mongodb database uri in .env as DB_STRING

/users get request will return the list of users in the database

/create-user post request will add a user in the following format
{
    "username": string,
    "age": int
}
