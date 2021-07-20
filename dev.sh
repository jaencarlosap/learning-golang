source config/config.sh
echo "Success load env data"
#version 1
#CompileDaemon --command="./`basename $PWD`" -color
#version 2
CompileDaemon -build="go build src/main.go" --command="./main" -color