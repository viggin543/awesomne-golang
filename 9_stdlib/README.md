# $GOROOT/pkg
contains the precompiled code of all stdlib,  ( files with an *.a extension)

#### on osx
`ls -l $GOROOT/pkg/darwin_arm64`

This speeds up golang compilation, since on every re-compile it only links with stdlib and does not need to compile it.