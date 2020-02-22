user	:=	$(shell whoami)

# GOBIN > GOPATH > INSTALLDIR
GOBIN	:=	$(shell echo ${GOBIN} | cut -d':' -f1)
GOPATH	:=	$(shell echo $(GOPATH) | cut -d':' -f1)
BIN		:= 	""

# check GOBIN
ifneq ($(GOBIN),)
	BIN=$(GOBIN)
else
# check GOPATH
ifneq ($(GOPATH),)
	BIN=$(GOPATH)/bin
else
# check INSTALL
#ifneq ($(INSTALL),)
#	BIN=$(INSTALL)
#endif
endif
endif

all: *.go
	@GO111MODULE=on go build -o oannes main.go

.PHONY: clean
.PHONY: install
.PHONY: uninstall

install:
ifeq ($(user),root)
#root, install for all user
	@cp ./oannes /usr/bin
	@[ -d /etc/oannes ] || mkdir /etc/oannes
	@cp -rf ./install/* /etc/oannes/
else
#!root, install for current user
	$(shell if [ -z $(BIN) ]; then read -p "Please select installdir: " REPLY; mkdir -p $${REPLY}; cp ./oannes $${REPLY}/; else mkdir -p $(BIN); cp ./oannes $(BIN); fi)
	@[ -d ~/.oannes ] || mkdir ~/.oannes
	@cp -rf ./install/* ~/.oannes/
endif
	@echo "install finished"

uninstall:
ifeq ($(user),root)
#root, install for all user
	@rm -f /usr/bin/oannes &> /dev/null
	@rm -rf /etc/oannes &>/dev/null
else
#!root, install for current user
	$(shell for i in `which -a oannes | grep -v '/usr/bin/oannes' 2>/dev/null | sort | uniq`; do read -p "Press to remove $${i} (y/n): " REPLY; if [ $${REPLY} = "y" ]; then rm -f $${i}; fi; done)
	@rm -rf ~/.oannes &>/dev/null
endif
	@echo "uninstall finished"

clean:
	@rm -f ./oannes
	@echo "clean finished"