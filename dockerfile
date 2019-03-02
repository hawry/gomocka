FROM scratch
ADD ./gomocka /
ADD ./settings.json /
CMD ["/gomocka"]