go test -covermode=count -coverpkg=../... -coverprofile ../coverage/cover.out -v ../... \
    && go tool cover -html ../coverage/cover.out -o ../coverage/cover.html \
    && echo "Coverage report is available at /coverage/cover.html"