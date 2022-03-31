FROM ubuntu

COPY dataserver /app/dataserver
COPY config.ini /app/config.ini
COPY ssl.key /app/ssl.key
COPY ssl.pem /app/ssl.pem

ENV GIN_MODE release

CMD cd /app && ./dataserver
