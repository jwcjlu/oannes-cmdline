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
	@GO111MODULE=on go build -o oann main.go

.PHONY: clean
.PHONY: install
.PHONY: uninstall

install:
ifeq ($(user),root)
#root, install for all user
	@cp ./oann /usr/bin
	@[ -d /etc/oann ] || mkdir /etc/oann
	@cp -rf ./install/* /etc/oann/
else
#!root, install for current user
	$(shell if [ -z $(BIN) ]; then read -p "Please select installdir: " REPLY; mkdir -p $${REPLY}; cp ./oann $${REPLY}/; else mkdir -p $(BIN); cp ./oann $(BIN); fi)
	@[ -d ~/.oann ] || mkdir ~/.oann
	@cp -rf ./install/* ~/.oann/
endif
	@echo "install finished"

uninstall:
ifeq ($(user),root)
#root, install for all user
	@rm -f /usr/bin/oann &> /dev/null
	@rm -rf /etc/oann &>/dev/null
else
#!root, install for current user
	$(shell for i in `which -a oann | grep -v '/usr/bin/oann' 2>/dev/null | sort | uniq`; do read -p "Press to remove $${i} (y/n): " REPLY; if [ $${REPLY} = "y" ]; then rm -f $${i}; fi; done)
	@rm -rf ~/.oann &>/dev/null
endif
	@echo "uninstall finished"

clean:
	@rm -f ./oann
	@echo "clean finished"