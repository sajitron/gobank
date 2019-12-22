module github.com/sajicode/gobank

require github.com/gorilla/mux v1.7.3

require github.com/sajicode/app v0.0.0

replace github.com/sajicode/app => ./app

require github.com/sajicode/controllers v0.0.0

replace github.com/sajicode/controllers => ./controllers

require github.com/sajicode/email v0.0.0

replace github.com/sajicode/email => ./email

require github.com/sajicode/models v0.0.0

replace github.com/sajicode/models => ./models

require github.com/sajicode/utils v0.0.0

replace github.com/sajicode/utils => ./utils

go 1.13
