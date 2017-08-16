unzipOpenXML README
===================

unzipOpenXML is a go commandline tool that unpacks [Office Open XML](https://en.wikipedia.org/wiki/Office_Open_XML) files.

Install
-------

```go get github.com/scusi/unzipOpenXML```

```
cd $GOPATH/src/github.com/scusi/unzipOpenXML
go install
```

Usage
-----

```
user@host:~/someDir/foo$ unzipOpenXML some.xlsx
```

Above command would unzip some.xlsx into a local directory called ```unzipped_some.xlsx```.
You can specify the directory to unpack to via the '-o' flag. See next example.

```
user@host:~/someDir/foo$ unzipOpenXML -o=/tmp/unzippedOpenXMLFiles/some some.xlsx
```

NOTE: output directories that do not exist will be created.

There is also a ```debug``` flag, which makes unzipOpenXML spills log messages for every step it does.

```
user@host:~/someDir/foo$ unzipOpenXML -debug=true some.xlsx
```

Disclaimer
----------

This software is provided as is without any liability.

By intention there is no license associated with this code. 

_I do care about my peers using my software. I don't give a damn about whether the lawyers and mega-corporations they work for use it. So, if you are like me and you don't care about all the intellectual property antics, here's my project, feel free to use it. If you are the kind of moron who wants to have their legal ass covered, go screw yourself._

See [Software Licenses and Failed States](http://250bpm.com/blog:82) for more thoughts about this.
