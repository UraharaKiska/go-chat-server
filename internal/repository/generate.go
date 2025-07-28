package repository

//go:generate sh -c "rm -rf mocks && mkdir -p mock"
//go:generate ../../bin/minimock -i ChatRepository -o ./mock -s "_minimock.go"
//go:generate ../../bin/minimock -i ChatUserRepository -o ./mock -s "_minimock.go"
//go:generate ../../bin/minimock -i ChatMessageRepository -o ./mock -s "_minimock.go"

