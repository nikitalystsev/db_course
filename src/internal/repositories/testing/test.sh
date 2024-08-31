#!/bin/bash

DATABASE=
HOST=
PORT=
USER=
PASSWORD=

function usage() { echo "Usage: $0 -h host -d database -p port -u username -w password" 1>&2; exit 1; }

while getopts d:h:p:u:w: OPTION
do
  case $OPTION in
    d)
      DATABASE=$OPTARG
      ;;
    h)
      HOST=$OPTARG
      ;;
    p)
      PORT=$OPTARG
      ;;
    u)
      USER=$OPTARG
      ;;
    w)
      PASSWORD=$OPTARG
      ;;
    *)
      usage
      ;;
  esac
done

if [[ -z $DATABASE ]] || [[ -z $HOST ]] || [[ -z $PORT ]] || [[ -z $USER ]]
then
  usage
  # shellcheck disable=SC2317
  exit 1
fi

##waiting for postgres
#while ! curl http://"$HOST":"$PORT"/ 2>&1 | grep "52"
#do
##    echo "Waiting for PostgreSQL..."
#    sleep 1
#done

#echo "Postgres is ready, installing pgTAP"
PGPASSWORD=$PASSWORD psql -h "$HOST" -p "$PORT" -d "$DATABASE" -U "$USER" -f /usr/local/share/postgresql/extension/pgtap.sql > /dev/null

rc=$?
# exit if pgtap failed to install
if [[ $rc != 0 ]] ; then
#  echo "pgTap was not installed properly. Unable to run tests!"
  exit $rc
fi

#echo "pgTAP installed, running tests"
for test_file in ./tests/*.sql; do
    PGPASSWORD=$PASSWORD pg_prove -h "$HOST" -p "$PORT" -d "$DATABASE" -U "$USER" "$test_file"

    rc=$?

    if [ $rc -ne 0 ]; then
        echo "Test failed for $test_file"
        break
    fi
done

exit $rc
