# Dick's Random Link Generator

Simple go service to return a random link for a static qr code

## Build
`go build`

## Running
Should be run as a service under systemd. I have yet to create the service file yet though.
Be sure to have a randomlinks.txt file or the service won't run.
`./randomlinks`
