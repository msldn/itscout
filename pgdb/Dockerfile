FROM postgres:14.2-alpine
COPY  --chown=70:70 itscout-db.sql /docker-entrypoint-initdb.d/itscout-db.sql
USER postgres
EXPOSE 5432
ENTRYPOINT ["docker-entrypoint.sh"]
CMD ["postgres"]