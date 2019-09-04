#!/bin/sh
docker build -t scrapingingo/sig-api:latest .
docker push scrapingingo/sig-api:latest
curl -X POST https://hp.just1689.co.za/api/webhooks/070da811-ae1c-411d-8436-a8321fc8addb

