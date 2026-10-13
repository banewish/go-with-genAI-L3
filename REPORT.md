# Звіт до домашньої роботи · Урок 3

## 1. Посилання на код

https://github.com/banewish/go-with-genAI-L3

## 2. Промпт з Розділу 2

lets say im golang dev, whos making microservices
create minimal interface, 
rewrite OrderService on dependecny injection through constructor
generate mock realisation for tests without real DB
also make the output to be returned in interface and service

Точний текст промпту зберігається в
[`prompts/interface-extraction-prompt.md`](./prompts/interface-extraction-prompt.md).

## 3. Міксовані приймачі та інтерфейси на стороні споживача

Дайте відповідь своїми словами (1-2 абзаци) на питання: *«Чому
міксування value та pointer ресиверів є небезпечним і як ШІ допомагає
проектувати інтерфейси на стороні споживача?»*

AI helps solve this by providing the golden rule as if any method on a type needs to mutate, all methods should use pointer receivers at least for clarity. AI can help design minimal, focused interfaces and generate mock implementations that dont require a real database—making tests fast and isolated.
 By analyzing code patterns and suggesting dependency injection through constructors, AI teaches developers to design interfaces around what consumers actually need, not what the implementation happens to provide. This leading to cleaner, more testable code.

## 4. Скріншоти / відео виконання

https://docs.google.com/document/d/1TBVtt4iMx7kTVYB7kYt0jUoajfXB_JWOXSEyQVkkNlY/edit?usp=sharing
