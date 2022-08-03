# JEC-go
Jabacat's Easy Config

[JEC-py](https://github.com/JakeRoggenbuck/JEC-py) | JEC-rs | [JEC-go](https://github.com/JakeRoggenbuck/JEC-go) | JEC-c | JEC-c++ | JEC-zig | JEC-ts

## API
```go
ConfigFile
  - Exists
  - Remove
  - Create
  
ConfigDir
  - Exists
  - Remove
  - Create
```

## Usage
```go
conf := ConfigFile{"./test.conf"}

if !conf.Exists() {
	conf.Create()
}

dir := ConfigDir{"./config/"}

if !dir.Exists() {
	dir.Create()
}

conf.Remove()
dir.Remove()

conf = ConfigFile{""}.FromHome("./test.conf")
strings.Contains(conf.path, "home") # true at /home/user/test.conf
```
