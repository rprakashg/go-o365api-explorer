FROM golang:latest

# Add pkg files to container workspace
ADD . /go/src/github.com/rprakashg/go-o365api-explorer

# install dependencies
RUN go get github.com/astaxie/beego/session

RUN go get github.com/codegangsta/negroni

RUN go get github.com/gorilla/mux

#commenting out since we want our version of the package which contain patch to work with Azure AD
#RUN go get golang.org/x/oauth2

#add the oauth2 and dependant packages
ADD ./oauth2/net /go/src/golang.org/x/net

ADD ./oauth2/oauth2 /go/src/golang.org/x/oauth2

#build golang oauth2 package
RUN go build golang.org/x/oauth2

#install golang oauth2 package
RUN go install golang.org/x/oauth2

#go rest client library for talking to RESTFUL services
RUN go get github.com/jmcvetta/napping

#build the app
RUN go build github.com/rprakashg/go-o365api-explorer

# Install the app
RUN go install github.com/rprakashg/go-o365api-explorer

#run the app when the container starts
ENTRYPOINT /go/bin/go-o365api-explorer

#expose port 10443 as the server is listening on port 10443
EXPOSE 10443