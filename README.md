# README

## Setup

```
# Configure your app's naming 
./init.sh
```


```bash
# Run on local machine
go get github.com/kyleconroy/sqlc/cmd/sqlc
go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest
```

### Docker
https://docs.docker.com/engine/install/ubuntu/
```bash
sudo apt-get remove docker docker-engine docker.io containerd runc
sudo apt-get update
sudo apt-get install ca-certificates curl gnupg lsb-release
curl -fsSL https://download.docker.com/linux/ubuntu/gpg | sudo gpg --dearmor -o /usr/share/keyrings/docker-archive-keyring.gpg
echo "deb [arch=$(dpkg --print-architecture) signed-by=/usr/share/keyrings/docker-archive-keyring.gpg] https://download.docker.com/linux/ubuntu $(lsb_release -cs) stable" | sudo tee /etc/apt/sources.list.d/docker.list > /dev/null
sudo apt-get update
sudo apt-get install docker-ce docker-ce-cli containerd.io

# verify
sudo docker run hello-world
```

### Docker Compose
```bash
sudo curl -L "https://github.com/docker/compose/releases/download/1.29.2/docker-compose-$(uname -s)-$(uname -m)" -o /usr/local/bin/docker-compose
sudo chmod +x /usr/local/bin/docker-compose

# verify
docker-compose --version
```

## Development
```bash
go mod install
make dev  # Runs on port __PORT__ 
```

```bash
make docker-dev			# init dev env
make db-create			# create the database
make migrate-create name=[name]	# create a `name` migration, add sql code to migration files
make migrate			# execute migration
make sqlc			# creates go code
```


## Useful Tools
- https://dbdiagram.io/home


## Tips
- https://www.youtube.com/watch?v=wxkEQxvxs3w


## Notes
```
\l => show databases
\c [db name] => use database
\dt => show tables
```
