### svm - Semantic Version management [POC]

With `svm` you can easily manage your project's semver. Set incrementally major, minor or patch number. 

Usage:
```sh
./svm [ show | all | major | minor | patch | undo ]
```

Example:
```sh
justorius@XC-1660  (main)$ go build 
justorius@XC-1660  (main)$ ./svm patch
Setting new patch tag:
v9.7.9
justorius@XC-1660  (main)$ ./svm minor
Setting new minor tag:
v9.8.9
justorius@XC-1660  (main)$ ./svm major
Setting new major tag:
v10.8.9
justorius@XC-1660  (main)$ ./svm minor
Setting new minor tag:
v9.9.9
justorius@XC-1660  (main)$ ./svm undo
Deleted tag 'v9.9.9' (was 48422ba)

justorius@XC-1660  (main)$ ./svm show
Current version: v9.8.9

justorius@XC-1660  (main)$ ./svm all
v1.0.0
v1.0.1
v1.1.0
[...]
[...omissis...]
[...]
v9.7.7
v9.7.8
v9.7.9
v9.8.9
```

# test

```sh
$ go test -v ./...
=== RUN   TestShow
Current version: v9.8.9

--- PASS: TestShow (0.00s)
=== RUN   TestPatch
Setting new patch tag:
v9.8.10
--- PASS: TestPatch (0.00s)
=== RUN   TestMinor
Setting new minor tag:
v9.9.9
--- PASS: TestMinor (0.00s)
=== RUN   TestMajor
Setting new major tag:
v10.9.9
--- PASS: TestMajor (0.00s)
=== RUN   TestUndo
Deleted tag 'v9.9.9' (was 48422ba)

--- PASS: TestUndo (0.01s)
=== RUN   TestAll
v1.0.0
v1.0.1
v1.1.0
v10.5.3
v10.5.4
v10.6.5
v10.6.6
v10.8.9
v10.9.9
v2.0.0
v2.0.1
v2.1.1
v3.1.1
v4.1.1
v5.1.1
v6.1.1
v7.1.1
v7.2.1
v7.2.2
v8.2.2
v8.3.2
v8.3.3
v8.4.3
v8.5.3
v9.5.3
v9.5.4
v9.5.5
v9.5.6
v9.6.6
v9.7.6
v9.7.7
v9.7.8
v9.7.9
v9.8.10
v9.8.9

--- PASS: TestAll (0.00s)
PASS
ok  	github.com/deeper-x/svm	0.023s


```

