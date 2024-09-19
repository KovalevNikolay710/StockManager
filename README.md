# lamoda_tech

### _Junior test task_

Сервис предназначен для резервации товаров на складах разными пользователями.

Для запуска приложения требуется наличие _git_, _docker_ и _makefile_.
По умолчанию backend работает на `8000` порту, а postgres - на `5432` порту.

### Проверка занятости портов
```
sudo lsof -i :5432

sudo lsof -i :8000
```
Если `5432` порт занят, его необходимо освободить:
```
sudo kill -9 {PID}
```
Если `8000` порт занят, его можно освободить схожим образом или изменить переменную `SERVER_PORT` в `.env` файле _(в корне директории lamoda_tech)_ на свободный.

После настройки портов _(если такова была необходима)_ можем запускать сервис.

### Запуск сервиса
```
git clone https://github.com/Mooonsheen/lamoda_tech

cd lamoda_tech

make up
```
Если `make` up по каким либо причинам не исполняет `docker compose` команду, её необходимо запустить вручную _(весь сервис также поднимется по одной этой команде)_.

### Запуск команды 
```
sudo docker compose up
```
Ожидаем загрузку, после которой:
1) Поднимается postgres контейнер
2) Поднимается контейнер с миграциями, запускает их и отключается (работает на x86_64 архитектуре)
3) Поднимается backend контейнер

### Сервис готов к работе
