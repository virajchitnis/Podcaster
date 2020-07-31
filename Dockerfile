FROM scratch
EXPOSE 8080
VOLUME ["/etc/podcaster", "/var/podcaster"]
ADD bin/Podcaster_linux_amd64 /usr/bin/podcaster
CMD ["/usr/bin/podcaster"]