# golearn
Learning a golang to Nomina VIP Indonesia

# install-go
Scripts to install and update Go for Darwin, FreeBSD, Linux, and Windows systems. These scripts support all precompiled architectures for the supported OS.

Bash is used for the shell scripts. Powershell is used for Windows.

## Shell
Shell scripts for installing and updating Go.  These are consistent with the instructions on https://golang.org/doc/install

These scripts should be run as the user for which you want Go.  __Do not run as root or with `sudo`__.  This breaks things.

### install_go
Installs Go to `/usr/local/go`.  By default, the `$GOPATH` is set to `~/go` and the directory is created.  This can be overridden by using either the `-w` or `--workspace` flag. `$GOPATH`, `$GOPATH/bin`, and  `/usr/local/go/bin` are added to the `~/.bashrc` file.

`install_go` defaults to the current version of Go, Linux, and amd64. Any valid Go release can be specified. For operating system, `darwin`, `freebsd`, or `linux` can be specified. The supported arch depends on what pre-compiled versions are provided for each OS.

#### Flags
short|long|alternate|description  
:--|:--|:--|:--  
-a|--arch|--goarch|specify the arch ($GOARCH)  
-h|--help||print help message  
-o|--os|--goos|specify the os ($GOOS)  
-s|--supported||list supported GOOS:GOARCH combinations  
-v|--version||specify the Go version  
-w|--workspace|--gopath|specify the workspace directory ($GOPATH)  

### upgrade_go
This script upgrades the current Go installation to the current release.  It assumes that the Go is installation directory is `/usr/local/go`.

The contents of `/usr/local/go` are removed.  Then the current release of Go is downloaded and extracted to `/usr/local/`.

If a specific version of Go is desired, it can be specified. Likewise with $GOOS and $GOARCH.

Use this when Go is already installed.

#### Flags
short|long|alternate|description  
:--|:--|:--|:--  
-a|--arch|--goarch|specify the arch ($GOARCH)  
-h|--help||print help message  
-o|--os|--goos|specify the os ($GOOS)  
-s|--supported||list supported GOOS:GOARCH combinations  
-v|--version||specify the Go version  
-w|--workspace|--gopath|specify the workspace directory ($GOPATH)  

### install_go_all
This script is a copy of [Eric Lagergren's](https://github.com/EricLagergren) [InstallGo](https://gist.github.com/EricLagergren/ddea0f327d38f8c3a918) script.  It allows for the specification of the Go version to install along with the OS and arch. IT supports a wider range of options than `install_go` and `upgrade_go`

#### Help output
This is the help output of the script, which provides all of the information needed to use it:

```
Usage: ./install_go_all [OPTIONS] versions...

  Versions must be valid Go version numbers.
  (E.g., 1.5.3, 1.6, 1.4.2, etc.)

  -o, --goos        specify the operating system to download
  -a, --goarch      specify the architecture to download
  -p, --possible    print possible GOOS/GOARCH combinations
  -v, --version     print version
  -h, --help        print help message

```
