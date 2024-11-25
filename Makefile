.PHONY: artisan

artisan:
	docker compose run --rm artisan $(filter-out $@,$(MAKECMDGOALS))
