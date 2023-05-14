MAKEFLAGS += --silent

TARGET = notify-mail
PLATFORM := $(shell uname -m)

ARCH :=
	ifeq ($(PLATFORM),x86_64)
		ARCH = amd64
	endif
	ifeq ($(PLATFORM),aarch64)
		ARCH = arm64
	endif
	ifeq ($(PLATFORM),armv7l)
		ARCH = armhf
	endif
GOARCH :=
	ifeq ($(ARCH),amd64)
		GOARCH = amd64
	endif
	ifeq ($(ARCH),i386)
		GOARCH = 386
	endif
	ifeq ($(ARCH),arm64)
		GOARCH = arm64
	endif
	ifeq ($(ARCH),armhf)
		GOARCH = arm
	endif

ifeq ($(GOARCH),)
  $(error Invalid ARCH: $(ARCH))
endif

ifdef DESTDIR
# dh_auto_install (Debian) sets this variable
  TARGET_DIR = $(DESTDIR)/usr
else
  TARGET_DIR ?= /usr/bin
endif

all: $(TARGET)

distclean: clean

clean:
	rm -f $(TARGET)

install: $(TARGET)
	install -Dm755 $(TARGET) $(TARGET_DIR)/$(TARGET)

uninstall:
	rm -f $(TARGET_DIR)/sbin/$(TARGET)

$(TARGET):
	GOOS=linux GOARCH=$(GOARCH) go build

test:
	go test ./... -race -cover

tidy:
	go mod tidy

vendor: tidy
	go mod vendor