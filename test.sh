source ./.env
go test ./features/... -coverprofile=cover.out && go tool cover -html=cover.out
go tool cover -func cover.out
