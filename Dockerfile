FROM ubuntu

COPY dataserver /app/dataserver
COPY config.ini /app/config.ini

ENV GIN_MODE release

CMD cd /app && ./dataserver
