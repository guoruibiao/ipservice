# ipservice
[![ipservice](https://travis-ci.org/guoruibiao/ipservice.svg?branch=master)](https://travis-ci.org/guoruibiao/ipservice)

A simple demo for deploy application with docker swarm.

```
travis:
    go get dependencies
    bash run_test_case.sh
        go build
        application runs in background
        run testcases
        if success:
            docker build -t xxx .
            docker push xxx 
            docker swarm ...
        else:
            hook for IM with the purpost of noticing teammates.

```

docker.io/marksinoberg/ipservice:
- v1: linux
- v2: linux
- v3: macOS

