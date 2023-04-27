#!/bin/bash

go build -o hotel-booking-web cmd/web/*.go && ./hotel-booking-web
./hotel-booking-web -dbname=bookings -dbuser=olhavasileva -cache=false -production=false