## Для миграции в терминале WbTech0/internal/storage/migrations
goose postgres "host=localhost user=admin database=db_WbTech0 password=root sslmode=disable" up

## Для запуска проекта нужно установить CONFIG_PATH окружения GO Modules
CONFIG_PATH=config/local.yaml
в GoLand:
в настройках проекта, в Go Modules в поле Environment: CONFIG_PATH=config/local.yaml

## Запуск
1. запускаем order-viewer
2. Затем order-sender
