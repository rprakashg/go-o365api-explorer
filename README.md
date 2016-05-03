# go-o365api-explorer
Containerized Office365 API explorer application written in golang

Building the App


Installing the App


Running


Build Container Image :
docker build -t "rprakashg/go-o365api-explorer" .

Running the Container :
docker run -itp 443:10443 "rprakashg/go-o365api-explorer"


