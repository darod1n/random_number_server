# random_number_server

# How to start server

Создайте `.env`-файл с помощью команды:
```bash
make create_env
```

Запустите сервер
```bash
make start
```

Подключитесь к `ws://localhost:8080/`

Для получения случайного числа отправьте комманду `get_random_number`

# Make commands 

`create_env` - создает файл .env на основе файла .env.example

`build` - билд 
	
`start_with_logs` - старт сервиса с выводом логов

`start` - старт сервиса

`stop` - остановить сервисы

`kill` - убить все процессы

`down` - выключить все сервисы

`restart` - рестарт 