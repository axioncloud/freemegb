# TODO: Create Makefile to automate deployment
#windres		:= x86_64-w64-mingw32-windres
arch		:= $(shell go env GOARCH)
os			:= $(shell go env GOOS)
version		:= $(shell cat VERSION)

.PHONY: all
# linux_i386 linux_amd64 additional gcc targets
all: clean host #cross-tools linux_i386 linux_amd64 linux_arm windows_i386 windows_amd64

windows_amd64:
	@echo "PATH: ${PATH}"
	@export PATH='/mingw64/bin:${PATH}'
	@echo "GCC: $(shell which gcc)"
	@echo "G++: $(shell which g++)"
	@echo "PKG-CONFIG: $(shell which pkg-config)"
	windres --input=freemegb_windows.rc --output=freemegb_windows.syso
# ifeq ($(arch),arm\)
	PKG_CONFIG_PATH=/mingw64/lib/pkgconfig:/usr/lib/pkgconfig:/usr/share/pkgconfig:/lib/pkgconfig GOOS=windows GOARCH=amd64 CGO_ENABLED=1 go build -v -o bin/windows_amd64/freemegb.exe
# endif
	cp -rf ui bin
windows_i386:
	@echo "PATH: ${PATH}"
	PATH='/mingw32/bin:${PATH}'
	@echo "GCC: $(shell which gcc)"
	@echo "G++: $(shell which g++)"
	@echo "PKG-CONFIG: $(shell which pkg-config)"
	windres --input=freemegb_windows.rc --output=freemegb_windows.syso
#ifeq ($(arch),arm)
	GOOS=windows GOARCH=386 CGO_ENABLED=1 go build -v -o bin/windows_386/freemegb.exe -ldflags "-H windowsgui"
#endif
	cp -rf ui bin
linux_arm:
	GOOS=linux GOARCH=arm PKG_CONFIG_PATH=/usr/arm-linux-gnueabi/lib/pkgconfig CGO_ENABLED=1 CC=arm-linux-gnueabi-gcc-8 go build -v -o bin/linux_arm/freemegb 
	cp -rf ui bin/linux_arm
linux_i386:
	GOOS=linux GOARCH=386 go build -v -o bin/linux_i386/freemegb
	cp -rf ui bin/linux_i386
linux_amd64:
	GOOS=linux GOARCH=amd64 go build -v -o bin/linux_amd64/freemegb
	cp -rf ui bin/linux_amd64
travisci:
	mkdir -p bin/linux_amd64
	go build -v -tags gtk_$(GTK_VERSION) -o bin/linux_amd64/freemegb
	cp -rf ui bin/linux_amd64
host:
ifeq ($(os),windows)
	windres --input=freemegb_windows.rc --output=freemegb_windows.syso
	GOOS=$(os) GOARCH=$(arch) CGO_ENABLED=1 go build -v -o bin/freemegb-$(version)_$(os)-$(arch)/freemegb.exe -ldflags "-H windowsgui"
endif
ifeq ($(os),linux)
	GOOS=$(os) GOARCH=$(arch) go build -v -o bin/freemegb-$(version)_$(os)-$(arch)/freemegb
endif
	cp -rf ui bin/freemegb-$(version)_$(os)-$(arch)
	cp -rf shaders bin/freemegb-$(version)_$(os)-$(arch)
host_deb:
	mkdir -p build/freemegb-$(version)_$(arch)
	cp -rf installer_files/$(os)_$(arch)/* build/freemegb-$(version)_$(arch)
	mkdir -p build/freemegb-$(version)_$(arch)/usr/bin/
	cp bin/freemegb-$(version)_$(arch)/freemegb build/freemegb-$(version)_$(arch)/usr/bin/
	mkdir -p build/freemegb-$(version)_$(arch)/usr/share/freemegb
	cp -rf ui build/freemegb-$(version)_$(arch)/usr/share/freemegb
	cp -rf shaders build/freemegb-$(version)_$(arch)/usr/share/freemegb
	sudo chown -R root:root build/freemegb-$(version)_$(arch)/
	mkdir dist
	dpkg -b build/freemegb-$(version)_$(arch)
	mv build/freemegb-$(version)_$(arch).deb dist
clean:
	rm -rf bin
	rm -rf build
	rm -rf dist
	rm -f *.syso
