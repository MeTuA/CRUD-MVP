$env:PORT="1345"
$env:DB_URL="postgres://admin:admin@localhost:5435/users?sslmode=disable"

cd cmd
go build -mod vendor
cd ..
./cmd/cmd.exe