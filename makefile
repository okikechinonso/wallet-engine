mock-db:
	mockgen -source=interfaces/db/user_interface.go -destination=mock/mock_db.go -package=mock
run: 
	go run main.go