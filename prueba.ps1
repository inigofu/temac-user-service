go build
$env:DB_HOST="localhost:54322"
$env:DB_USER="postgres"
$env:DB_NAME="postgres"
$env:DB_PASSWORD="mysecretpassword"
$env:MICRO_SERVER_ADDRESS=":40054"
$env:MICRO_REGISTRY="consul"
$env:DB_LOG="true"
.\temac-user-service.exe