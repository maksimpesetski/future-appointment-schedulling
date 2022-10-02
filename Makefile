
# remove build files
.PHONY: clean
clean:
	@rm -rf appointment-schedule

# run application locally
.PHONY: run-local
run-local: clean
run-local:
	docker-compose up
