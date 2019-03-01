FROM scratch
ADD ./gocks /
ADD ./settings.json /
CMD ["/gocks"]