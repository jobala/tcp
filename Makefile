
run:
	docker run -it -w /app --mount type=bind,src=.,target=/app  golang:1.21.5-bullseye /bin/bash
