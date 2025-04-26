package config

import "os"

// no need here for this file i just make for understanding purpose

func GetConfig(key string) string {
	return os.Getenv(key)
}

/*
The config file is used to manage environment-specific settings, like API keys, database credentials, or other configurations that you might need for your application.

For example, instead of hardcoding values like database credentials or API keys directly in your code, you store them in environment variables and use the config package to fetch those values.

For example

You might have an environment variable for the database URL

DATABASE_URL="postgres://user:password@localhost/dbname"

//And in your config.go, you'd fetch it like this ?


package config

import "os"

func GetConfig(key string) string {
    return os.Getenv(key) // This fetches the value of the environment variable
}
In your code you can use like this ?

dbURL := config.GetConfig("DATABASE_URL")

*/
