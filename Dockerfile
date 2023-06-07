FROM                    tsuru/scratch:latest
WORKDIR                 app
COPY .                  /app
ENTRYPOINT              ["./build/server"]
