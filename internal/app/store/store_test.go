package store_test

import (
    "os"
    "testing"
)

var (
    databaseUrl string
)

func TestMain(m *testing.M) {
    databaseUrl = os.Getenv("DATABASE_URL")
    if databaseUrl == "" {
        databaseUrl = "host=localhost dbname=restapi_test user=go_user_test password=go_user_test"
    }

    os.Exit(m.Run())
}
