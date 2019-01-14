FROM alpine

RUN mkdir /app
RUN mkdir /app/data

COPY out /app

CMD [ "/app/main", "/app/data", "https://acruxo.sedlak.biz/reznictvi-sedlak" ]
