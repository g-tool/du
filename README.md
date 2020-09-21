<div align="center">
	<h1>DU.</h1>
	<h4 align="center">
		Simple and fast concurrent downloader & uploader.
	</h4>
</div>

<p align="center">
	<a href="#installation">Installation</a> ❘
	<a href="#command-line-tool-usage">CLI Usage</a> ❘
	<a href="#module-usage">Module Usage</a> ❘
	<a href="#license">License</a>
</p>


## Comparison

Comparison in cloud server:




## Installation

#### Download and install the latest [release](https://github.com/g-tool/du/releases):
```bash
# go to tmp dir.
cd /tmp

# Download latest version.
curl -sfL https://git.io/getdu | sh

# Make the binary executable.
chmod +x /tmp/bin/du

# Move the binary to your PATH
sudo mv /tmp/bin/du /usr/bin/du
```

#### Or Go ahead compile it yourself:
```bash
go get github.com/g-tool/du/cmd/du
```



> **Note:** these packages are not maintained by g-tool

## Command Line Tool Usage

#### Simple usage:
```bash
du https://example.com/file.mp4
```

#### You can specify destination path:
```bash
du -o /path/to/save https://example.com/file.mp4
```

#### You can download multiple URLs and save them to directory:
```bash
du --dir /path/to/dir https://example.com/file.mp4 https://example.com/file2.mp4
```

#### You can download multiple URLs from a file:
```bash
du --dir /path/to/dir -f urls.txt
```

### You can pipe multiple URLs:
```bash
cat urls.txt | du --dir /path/to/dir
```

#### Docs for available flags:
```bash
du help
```


## Module Usage

You can use DU to download large files in your go code, the usage is simple as the CLI tool:

```bash
package main

import "github.com/go-tool/du"

func main() {

	d := du.New()

	err := d.Download("http://localhost/file.ext", "/path/to/save")

	if err != nil {
		// ..
	}
}

```

For more see [PkgDocs](https://pkg.go.dev/github.com/g-tool/du).

## How It Works?

DU takes advantage of the HTTP range requests support in servers [RFC 7233](https://tools.ietf.org/html/rfc7233), if the server supports partial content DU split the file into chunks, then starts downloading and merging the chunks into the destinaton file concurrently.


## License

DU is provided under the [MIT License](https://github.com/g-tool/du/blob/master/LICENSE) © Tacey Wong.
