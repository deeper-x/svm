### svm - Semantic Version management

With `svm` you can easily manage your project's semver. Set incrementally major, minor or patch number. 

Usage:
```sh
./svm [ show | all | major | minor | patch | undo | write <file>]
```

Example:
```sh
justorius@XC-1660  (main)$ go build 
justorius@XC-1660  (main)$ ./svm show
Current version: v1.4.0
justorius@XC-1660  (main)$ ./svm patch
v1.4.1
justorius@XC-1660  (main)$ ./svm minor
v1.5.0
justorius@XC-1660  (main)$ ./svm major
v2.0.0
justorius@XC-1660  (main)$ ./svm patch
v2.0.1
justorius@XC-1660  (main)$ ./svm patch
v2.0.2
justorius@XC-1660  (main)$ ./svm minor
v2.1.0
justorius@XC-1660  (main)$ ./svm undo
Deleted tag 'v2.1.0' (was 2864110)
justorius@XC-1660  (main)$ ./svm all
v1.0.0
v1.0.1
v1.1.1
v1.1.2
v1.2.2
v1.2.3
v1.3.3
v1.3.4
v1.4.0
v1.4.1
v1.5.0
v2.0.0
v2.0.1
v2.0.2

```

# test

```sh
$ make unittest
/usr/bin/go test -v ./...
?   	github.com/deeper-x/svm/settings	[no test files]
=== RUN   TestShow
--- PASS: TestShow (0.00s)
=== RUN   TestPatch
--- PASS: TestPatch (0.00s)
=== RUN   TestMinor
2023/11/16 20:58:51 redo
--- PASS: TestMinor (0.00s)
=== RUN   TestMajor
--- PASS: TestMajor (0.00s)
=== RUN   TestUndo
--- PASS: TestUndo (0.00s)
=== RUN   TestAll
--- PASS: TestAll (0.00s)
=== RUN   TestWritre
--- PASS: TestWritre (0.00s)
=== RUN   TestSetNewVer
--- PASS: TestSetNewVer (0.00s)
=== RUN   TestDelVer
--- PASS: TestDelVer (0.00s)
PASS
ok  	github.com/deeper-x/svm	0.032s
=== RUN   TestNewTag
--- PASS: TestNewTag (0.00s)
=== RUN   TestCheckNumArgs
--- PASS: TestCheckNumArgs (0.00s)
=== RUN   TestCheckNumArgsMinimum
--- PASS: TestCheckNumArgsMinimum (0.00s)
=== RUN   TestSetNewVer
--- PASS: TestSetNewVer (0.00s)
=== RUN   TestDelVer
--- PASS: TestDelVer (0.00s)
=== RUN   TestShowAll
--- PASS: TestShowAll (0.00s)
PASS
ok  	github.com/deeper-x/svm/git	(cached)
```

