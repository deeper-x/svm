clean:
	git tag -l | xargs git tag -d
	git tag v1.0.0

unittest:
	/usr/bin/go test -v ./...
