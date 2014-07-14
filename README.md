# Install Google App Engine / Go SDK

    $ brew install go
    $ brew install go-app-engine-64


# Configure as Your App

## app.yaml

Modify "helloworld" to "your-app-name".

## main.go

Write handler codes under `myhandlers`, and add them in `init()` in main.go.


# Run Development Server

    $ goapp serve .
