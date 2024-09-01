FROM postgres:12.2-alpine

ENV PGTAP_VERSION v1.1.0

RUN apk -U add \
    alpine-sdk \
    perl \
 && git clone https://github.com/theory/pgtap \
 && cd pgtap \
 && git checkout ${PGTAP_VERSION} \
 && make \
 && make install


FROM postgres:12.2-alpine

COPY --from=0 /usr/local/share/postgresql/extension/pgtap* /usr/local/share/postgresql/extension/

WORKDIR /usr/local/src

RUN apk -U add \
    build-base \
    perl-dev \
    curl \
 && cpan TAP::Parser::SourceHandler::pgTAP \
 && apk del -r build-base

COPY ./tests ./tests
COPY ./test.sh ./test.sh
RUN chmod +x ./test.sh

CMD ["./test.sh"]
ENTRYPOINT ["./test.sh"]

# https://writeitdifferently.com/postgresql/pgtap/2020/04/08/pgtap-flyway-with-docker.html -- откуда взял идею