FROM scratch
COPY dockmd2redis /dockmd2redis
ENTRYPOINT ["/dockmd2redis"]