FROM scratch
EXPOSE 8080
VOLUME ["/etc/podcaster", "/var/podcaster"]
ADD bin/Podcaster_linux /usr/bin/podcaster
CMD ["/usr/bin/podcaster"]