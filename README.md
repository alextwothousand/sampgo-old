# sampgo
Go wrapper around SAMPGDK

## Does this work?
No. It's currently a work in progress, and doesn't even compile correctly. I open sourced this in the hopes that together, collaboratively, we may be able to get this to function.

## How can I compile?
`CGO_ENABLED=1 CGO_FLAGS_ALLOW=-Wno-attributes GOARCH=386 GOOS=linux go build -buildmode=c-shared`
You change the GOOS environment variable depending on your OS.

## How can I contribute?
When you compile you will encounter an issue.

```usr/bin/ld: $WORK/b001/_x002.o: in function `OnGameModeInit':
./main.go:19: undefined reference to `goOnGameModeInit'
/usr/bin/ld: $WORK/b001/_x004.o: in function `sampgdk_plugin_get_symbol':
./sampgdk.c:1070: undefined reference to `dlsym'
/usr/bin/ld: $WORK/b001/_x004.o: in function `sampgdk_plugin_get_handle':
./sampgdk.c:1076: undefined reference to `dladdr'
/usr/bin/ld: ./sampgdk.c:1077: undefined reference to `dlopen'
collect2: error: ld returned 1 exit status```

Even though these are all defined, you still get the issue. If you can tackle them, god speed to you sir (or ma'am).
