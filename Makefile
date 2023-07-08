#build and run dev
dorun:
	cd ui2/ && \
	yarn build
dodev:
	go run .
dobuild:
	go build
run:
	make dorun && make dodev
compile:
	make dorun && make dobuild
#build and run dev