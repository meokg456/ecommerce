FROM kong:latest

COPY kong.yaml .

ENTRYPOINT ["/docker-entrypoint.sh"]

EXPOSE 8000

CMD ["kong", "docker-start"]