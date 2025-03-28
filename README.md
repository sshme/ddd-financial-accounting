## Общая идея решения

Разработанная система представляет собой финансовую учетную систему, реализованную с применением принципов предметно-ориентированного проектирования (DDD) и различных паттернов проектирования. Система позволяет пользователям управлять банковскими счетами, отслеживать финансовые операции и работать с категориями расходов/доходов.

Архитектура системы разделена на слои в соответствии с принципами чистой архитектуры:
- **Доменный слой** - содержит бизнес-сущности, репозитории и фабрики
- **Прикладной слой** - содержит фасады и сервисы, реализующие бизнес-логику
- **Инфраструктурный слой** - содержит реализации репозиториев и технические детали
- **Интерфейсный слой** - содержит CLI-интерфейс для взаимодействия с пользователем

Приложение использует внедрение зависимостей через Wire для обеспечения слабой связанности компонентов.

## Реализованные принципы SOLID и GRASP

### SOLID принципы:

1. **Принцип единственной ответственности (SRP)**:
   - Классы `BankAccount`, `Category`, `Operation` в `internal/domain/entities` отвечают только за свою бизнес-логику
   - Фасады в `internal/application/services` отвечают за координацию операций
   - Репозитории в `internal/infrastructure/persistence` отвечают только за хранение данных

2. **Принцип открытости/закрытости (OCP)**:
   - Интерфейсы репозиториев в `internal/domain/repositories` позволяют расширять функциональность без изменения существующего кода
   - Состояния в `internal/interface/cli/states` могут добавляться без изменения существующих состояний

3. **Принцип подстановки Лисков (LSP)**:
   - Реализации репозиториев в `internal/infrastructure/persistence` корректно реализуют интерфейсы из `internal/domain/repositories`

4. **Принцип разделения интерфейса (ISP)**:
   - Интерфейсы репозиториев (`BankAccountRepository`, `CategoryRepository`, `OperationRepository`) в `internal/domain/repositories` специфичны для каждого типа сущности
   - Интерфейс `State` в `internal/interface/cli/states` содержит только необходимые методы

5. **Принцип инверсии зависимостей (DIP)**:
   - Высокоуровневые модули (фасады в `internal/application/services`) зависят от абстракций (интерфейсы репозиториев)
   - Зависимости внедряются через конструкторы (например, в `AccountFacade`)
   - Механизм внедрения зависимостей реализован с помощью Wire в `internal/wire.go`

### GRASP принципы:

1. **Высокая связность (High Cohesion)**:
   - Каждый класс имеет четкую цель и содержит только связанную функциональность
   - Например, фабрики в `internal/domain/factories` отвечают только за создание объектов

2. **Низкая связанность (Low Coupling)**:
   - Зависимости управляются через интерфейсы
   - Компоненты могут быть изменены независимо друг от друга

3. **Создатель (Creator)**:
   - Фабрики в `internal/domain/factories` отвечают за создание сложных объектов
   - Например, `BankAccountFactory` создает объекты `BankAccount`

4. **Информационный эксперт (Information Expert)**:
   - Методы размещены в классах, которые имеют необходимую информацию
   - Например, методы управления балансом в `BankAccount`

5. **Контроллер (Controller)**:
   - Фасады в `internal/application/services` координируют действия между слоями
   - CLI-интерфейс в `internal/interface/cli` координирует взаимодействие с пользователем

## Реализованные паттерны GoF

1. **Фасад (Facade)**:
   - Реализован в `internal/application/services` (`AccountFacade`, `CategoryFacade`, `OperationFacade`)
   - Важность: упрощает взаимодействие с подсистемами, предоставляя единый интерфейс для операций
   - Скрывает сложность взаимодействия с репозиториями и фабриками

2. **Фабрика (Factory)**:
   - Реализован в `internal/domain/factories` (`BankAccountFactory`, `CategoryFactory`, `OperationFactory`)
   - Важность: инкапсулирует логику создания объектов, обеспечивает соблюдение бизнес-правил при создании
   - Централизует логику валидации и создания объектов

3. **Репозиторий (Repository)**:
   - Реализован в `internal/infrastructure/persistence` с интерфейсами в `internal/domain/repositories`
   - Важность: абстрагирует доступ к данным, делает систему независимой от конкретного способа хранения
   - Обеспечивает единый интерфейс для работы с данными, инкапсулируя детали хранения

4. **Состояние (State)**:
   - Реализован в `internal/interface/cli/states`
   - Важность: позволяет объекту изменять свое поведение при изменении внутреннего состояния
   - Упрощает управление сложным интерфейсом с множеством экранов и переходов

5. **Внедрение зависимостей (Dependency Injection)**:
   - Реализован с использованием Google Wire в `internal/wire.go`
   - Важность: обеспечивает слабую связанность компонентов, упрощает тестирование
   - Централизует настройку зависимостей приложения

6. **Команда (Command)**:
   - Неявно реализован через состояния в `internal/interface/cli/states`
   - Важность: инкапсулирует запрос как объект, позволяя параметризовать клиентов с различными запросами
   - Упрощает работу с различными операциями пользователя

## Запуск приложения

### С помощью Taskfile
   ```bash
      task run
   ```
