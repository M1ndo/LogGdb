##
# LogGdb MakeFile
#
# @file
# @version 0.1
USER=m1ndo
PACKAGE=LogGdb
update-pkg-cache:
    GOPROXY=https://proxy.golang.org GO111MODULE=on \
    go get github.com/$(USER)/$(PACKAGE)@v$(VERSION)


# end
