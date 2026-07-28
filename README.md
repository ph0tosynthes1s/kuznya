# Кузня
![Kuznya logo](https://github.com/ph0tosynthes1s/kuznya/raw/main/docs/logo.png)
**Кузня** — CLI-инструмент для быстрого создания проектов из готовых шаблонов.

Вместо копирования старых репозиториев и ручной настройки окружения достаточно одной команды.

```bash
kuznya build my-api --template laravel
```

## Поддерживаемые шаблоны:
- Laravel;
- Gin;
- FastAPI;
- Next.js.

## Пример:
```bash
kuznya build my-api --template gin
```

Создаст готовый каркас проекта со структурой директорий, конфигурацией и базовыми файлами.

## Roadmap проекта:
- [x] инит проекта;
- [ ] пользовательские шаблоны;
- [ ] генерация отдельных компонентов (`controller`, `service`, `repository`);
- [ ] интеграция с Docker;
- [ ] git-хуки;
