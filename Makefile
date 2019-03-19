GO := go
RM := rm
STRIP := $(CROSS_COMPILE)strip
BINARY := ../../../../bin/autoreadme

export GOPATH := $(shell pwd)/vendor/

GO_MAIN = autoreadme.go

GO_PREREQS += \
	$(shell find . -type f -name "*.go") \

GO_BATS_LIBS += \
	bats/utilities/fileroller/ \
	bats/utilities/logger/ \
	bats/utilities/password/ \
	bats/utilities/timestamp/ \
	bats/utilities/stopwatch/ \
	bats/protocols/nmea/ \
	bats/protocols/pelcod/ \

$(BINARY): $(GO_PREREQS)
	go build -o $(BINARY)
	@$(STRIP) $(BINARY)

app: $(BINARY)

clean:
	-$(RM) $(BINARY)
