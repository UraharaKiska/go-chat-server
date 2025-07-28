package service

//go:generate sh -c "rm -rf mocks && mkdir -p mock"
//go:generate ../../bin/minimock -i ChatService -o ./mock -s "_minimock.go"
