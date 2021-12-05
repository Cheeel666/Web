# Web - лабораторные работы
Данный репозиторий - переписанный курсовой проект с python на golang. 

## Цель работы
Разработать приложение-агрегатор по сбору актуальной информации о горнолыжных курортах, которое позволит пользователю получать информацию об открытых канатных дорогах и трассах из различных источников в одном месте.

## Перечень функциональных требований
- Регистрация пользователя
- Авторизация пользователя
- Просмотр личной страницы
- Просмотр списка курортов
- Просмотр списка трасс и канатных дорог
- Добавление и просмотр комментариев
- Предусмотреть страницу модератора(администратора) с возможностью удалять пользователей и менять их роли

## Use-case диаграмма системы
### Неавторизованный пользователь
![alt text](https://sun9-63.userapi.com/impg/huY5Fi399LN0tPiDx_sKe5WnCtTtyY3lw9hGIw/xlwZHTpP-u0.jpg?size=1658x474&quality=96&sign=876a2e5bc1760217bd573b07caec9df1&type=album)
### Авторизованный пользователь
![alt text](https://sun9-74.userapi.com/impg/v1tVeGs-_mgXCtItvJU--ROSKdv1W3WUxbx6Ew/w1qwtd535zM.jpg?size=1650x500&quality=96&sign=666b07993232fd83d3031da0d9961206&type=album)
### Модератор
![alt text](https://sun9-52.userapi.com/impg/uwJUfM6MwRRlA-gQ9LtIlu38ZU94yehcTgD0CQ/moYRYc392VQ.jpg?size=1654x650&quality=96&sign=de87a133f01ac272cb8c41538e99fd7c&type=album)
### Администратор
![alt text](https://sun9-81.userapi.com/impg/X4Ld-dx4s9Zk6GiG8c2khzJFChQu3GuRior2hQ/cjIrNnC6Kes.jpg?size=1654x722&quality=96&sign=4f153b54ef739399ff52e5cfe3fe9b8d&type=album)

## ER-диаграмма в нотации Чена
![alt text](https://sun9-79.userapi.com/impg/2BjWL_lTw6KzKdKdZCz8KcedPJsmxqmk9UdbwA/HrXaSL_3w8c.jpg?size=1280x700&quality=96&sign=70183586dd968593d10a28aca95e98c8&type=album)
## Диаграмма сущностей базы данных
![alt text](https://sun9-8.userapi.com/impg/uaV0oW3o3B8zPjyEqur10IwlDVefewDabHy_vQ/Z3AwxLIsPU4.jpg?size=1030x1152&quality=96&sign=a00c4f4eb4aacd11d6a54911fa03fe56&type=album)

# Ссылка на портрет в figma

https://www.figma.com/file/wZiTozg6I1r6oEsLLGKSnD/web_lab?node-id=0%3A1

# Ссылка на доску настроения

https://www.pinterest.ru/ilchel1992/%D0%B4%D0%BE%D1%81%D0%BA%D0%B0-%D0%BD%D0%B0%D1%81%D1%82%D1%80%D0%BE%D0%B5%D0%BD%D0%B8%D1%8F-%D0%B3%D0%BE%D1%80%D1%8B/%D0%B3%D0%BE%D1%80%D1%8B/

В фигме:

https://www.figma.com/file/KY3wph7XGorynR3vbTUAQg/palitra

# Шаблоны в figma

https://www.figma.com/file/wSa42WG7LVGlFupZ5RC7af/SkOpen-Templates?node-id=17%3A413

# Макет сайта в figma

https://www.figma.com/file/p25DWSneph1ZPDXVCQxsAZ/SkOpen?node-id=0%3A1

## Правки к первой лабораторной
### Первая сдача (правки):
+ добавить айдишники на ер(добавил диаграму бд, где все наглядно видно), 
+ приписать цель какой агрегатор 
+ ендпоинты поправить название
+ добавить авторизацию через хедеры
+ сделать группы в сваггере 

### Правки 2.0:
- Авторизация поменять хедер на некастомный (+)
- Доделать nginx(header server(+), gzip(+), кеширование(+))
- балансировка на 1 инстансе(-) и на трех(+) ???