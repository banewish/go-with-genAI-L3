# Звіт до домашньої роботи · Урок 3

## 1. Посилання на код

https://github.com/banewish/go-with-genAI-L3

## 2. Промпт з Розділу 2

Точний текст промпту зберігається в
[`prompts/interface-extraction-prompt.md`](./prompts/interface-extraction-prompt.md).

## 3. Mixed Receivers and Consumer-Side Interfaces

Mixing value and pointer receivers on the same type is dangerous because it creates implicit, hard-to-detect mutability contracts. When some methods use value receivers (read-only, no mutations) and others use pointer receivers (mutating), developers can't tell at a glance whether calling a method will modify state or not. This leads to subtle, confusing bugs—especially in scenarios like iterating over a slice of values where `for _, w := range` creates copies. Even if `Deposit()` has a pointer receiver, calling it on a copy means mutations vanish silently. The compiler doesn't enforce consistency, so the burden falls on developers to remember: "value receiver = no mutations, pointer receiver = mutations expected."

AI helps solve this by enforcing the **Golden Rule**: if any method on a type needs to mutate, all methods should use pointer receivers for consistency and clarity. From the consumer's perspective, AI can help design minimal, focused interfaces (like `OrderStore` with just `Exec()`) and generate mock implementations that don't require a real database—making tests fast and isolated. By analyzing code patterns and suggesting dependency injection through constructors, AI teaches developers to design interfaces around what consumers actually need, not what the implementation happens to provide. This shifts the design perspective from "what does my type do?" to "what minimal contract does my consumer require?"—leading to cleaner, more testable code.

## 4. Скріншоти / відео виконання

<!-- TODO(звіт-4): видаліть цей HTML-коментар і вставте посилання на
     Google Doc / відео зі скріншотами роботи програми та виконаного
     завдання. -->
