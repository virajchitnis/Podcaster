FROM scratch
EXPOSE 8080
VOLUME ["/etc/podcaster"]
ADD bin/Podcaster_linux /usr/bin/podcaster
CMD ["/usr/bin/podcaster"]