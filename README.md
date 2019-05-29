# CLI to interact with CUPS


### Usage
```
CLI to interact with CUPS

Usage:
  go-cups [command]
  
  Available Commands:
    help        Help about any command
    options     Options from the printer's attribute list
    print       Print a file
```

#### Examples

- Print
```$ ./cups-cli print -f {path to file}```

- Options
```
$ ./cups-cli options all
Current options for  ZJ-58
  printer-type = 184324
  auth-info-required = 
  printer-is-accepting-jobs = true
  printer-is-shared = true
  printer-make-and-model = Zijiang ZJ-58
  printer-state = 3
  printer-state-change-time = 1559150386
  printer-info = Thermal 58 POS Printer
  printer-location = Kitchen
  printer-state-reasons = none
```
```
$ ./cups-cli options printer-type
Current option:  printer-type 184324
```
