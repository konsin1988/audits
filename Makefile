dev:
	docker compose up -d backend-dev

re:
	docker compose down
	docker image rm  konsin1988/go-server-dev:latest
	docker compose up -d --build backend-dev

prod:
	docker compose up -d --build prod

clear:
	docker compose down
	docker image rm  konsin1988/go-server-prod:1.0.0
