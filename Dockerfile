FROM tsuru/scratch:latest
WORKDIR app
COPY ./build/server 	/app
ENTRYPOINT ["/bin/bash", "-c", "./server"]
