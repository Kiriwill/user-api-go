echo "- Starting docker.."
docker compose up -d db

echo "- Waiting for mysql to be ready..."
RETRIES=10

export VERIFYMY_DATABASE_DSN="devwill:supersecret@tcp(localhost:3306)/verifymychallenge"
until docker run -it --rm --network host mysql:8.2.0 mysqld "$VERIFYMY_DATABASE_DSN" -c "select 1;" > /dev/null || [ $RETRIES -eq 0 ]; do
    echo " $((RETRIES--)) remaining attempts..."
    sleep 1;
done

echo "- Running application..."
air

