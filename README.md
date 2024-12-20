## Запросы

- **Получить все продукты**
  - `GET` - `"http://127.0.0.1:8080/products"`
  - Возвращает список всех доступных продуктов.

- **Получить продукт по ID**
  - `GET` - `"http://127.0.0.1:8080/products/{id}"`
  - Возвращает информацию о продукте с указанным `ID`.

- **Создать новый продукт**
  - `POST` - `"http://127.0.0.1:8080/products"`
  - Добавляет новый продукт с переданными данными в теле запроса.

- **Обновить продукт по ID**
  - `PUT` - `"http://127.0.0.1:8080/products/{id}"`
  - Обновляет информацию о продукте с указанным `ID` на основе данных, переданных в теле запроса.

- **Удалить продукт по ID**
  - `DELETE` - `"http://127.0.0.1:8080/products/{id}"`
  - Удаляет продукт с указанным `ID` из списка.
 
- **Создать новый заказ**
  - `POST` - `"http://127.0.0.1:8080/orders/"`
  - Создает новый заказ.
 
- **Получить историю**
  - `GET` - `"http://127.0.0.1:8080/orders/"`
  - Получает историю заказов.
