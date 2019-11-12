# openports

<img src="https://github.com/jessecorson/openports/raw/master/logo/logo.png" width="200">

openports is a port scanning tool

openports allows you to listen on or scan any designated TCP port, list of ports or range of ports

## Usage

The `-p` flag can be used for both the open and scan commands.

Single Port
```bash
openports open -p 8080
openports scan -p 8080
```

Port Range
```bash
openports open -p 30-60
openports scan -p 30-60
```

List of Ports
```bash
openports open -p 22,53,80,443
openports scan -p 22,53,80,443
```

All possible Ports
```bash
openports open -p all
openports scan -p all
```

### Scan Command
By default the scan command will scan localhost. You can specify an IP or Hostname for the command to use with the target flag.
```bash
openports scan -p 443 -t www.jessecorson.com
```

# Notice
Before using the tool on external websites be sure you are not violating any policies.