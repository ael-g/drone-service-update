FROM alpine

ADD service-update /
ENTRYPOINT ["/service-update"]
