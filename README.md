## Консольная утилита для распеределния задачи.

Программа должна обладать командой help, благодаря которой можно получить список доступных команд с кратким описанием.

Список команд для реализации:

 1. Принять задачу на старт. На вход принимается ID заказа и описания задачи. Задаче присваивается статус `at start`. Если данная задача уже  имеет статус `at start`, `in progres` или `done`, приложение должно выдать ошибку.

 2. Перевести задачу в работу, на вход подается ID задачи и ID работника.
 Задача переводится в статус `in progres`.

 3. Задача была сделана на вход подается ID. Завершить можно лишь те задачи, получили имели статус `in progres`

 4. Поменять работника в списке задач, на вход подается ID работника и список задач.

 5. Получить список всех задач/ задач конкретного пользователя, на вход передать ID пользователя(или не передавать в случае списка всех задач).

 6. Удалить задачу из хранилища, на вход подается ID задачи. Данная задача удаляется из списка.