# Промпт для виокремлення інтерфейсу (Розділ 2)

Вставте нижче точний текст промпту, який ви розробили для Розділу 2
домашньої роботи. Промпт має явно містити:

- **роль** ШІ (наприклад, "Ти — досвідчений Go-розробник, що
  спеціалізується на чистій архітектурі...");
- **контекст** — наведений у завданні код `OrderService`;
- **завдання** — виокремити мінімальний інтерфейс `OrderStore`,
  переписати `OrderService` на dependency injection, згенерувати мок;
- **обмеження** — маленький інтерфейс (приказка Роба Пайка),
  конкретний стиль коду, тощо;
- **бажаний формат виводу** (наприклад, "поверни лише Go-код у трьох
  блоках: інтерфейс, сервіс, мок").


## Мій промпт

lets say im golang dev, whos making microservices
create minimal interface, 
rewrite OrderService on dependecny injection through constructor
generate mock realisation for tests without real DB
also make the output to be returned in interface and service

