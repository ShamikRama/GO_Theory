// // ЧТО ТАКОЕ ПОТОК И ПРОЦЕСС И В ЧЕМ РАЗЛИЧИЯ ?

// // Процесс разделяется на потоки, которые работают в одном процессе
// // У разных процессов разная память, но потоки в одном и том же процессе
// // используют одну и ту же память внутри процесса, потоки могут быть парралельны

// // Многопоточность - в программе существуют множество потоков
// // Состояния потокоа:
// // 1) выполнение
// // 2) ожидание
// // 3) готовности
// // 4) переходное
// // 5) завершения

// // КТО И КАК УПРАВЛЯЕТ ГОРУТИНЫ И В RUNTIME GOLANG ?
// // СБОРЩИК МУСОРА ?

// // В чем различия между системным потоком(тредом) и горутиной?
// // Потоки системные - абстракция операционной системы, а горутины абстракция
// // рантайма гошного. Чтобы переключиться между тредами, надо залезть в ядро(это дорого)
// // Но чтобы переключиться между горутинами, надо всего лишь переключить их в треде
// // Горутины выполняются в одном треде(потоке) когда используют общую область памяти
// // Когда горутина вызвывает сискол она попадает в поток отдельный , затем в глобальную очередь

// // В ГО одноэлементное LIFO а дальше FIFO :
// // 1) поток смотрит на LIFO и берет оттуда
// // 2) Если LIFO заполнено, он кладет в FIFO

// // Network poller - отдельный тред для системных вызовов по input и output

// В Go (Golang) очереди горутин хранятся в контексте логических процессоров (P), а не в контексте процессов или потоков (M). Давайте рассмотрим, как это работает в рамках MPG модели.

// MPG модель:
// Machine (M):

// Поток (Thread): M представляет собой поток операционной системы, который выполняет код.

// Связь с P: Каждый M связан с одним логическим процессором (P) и выполняет горутины из его локальной очереди.

// Processor (P):

// Логический процессор: P представляет собой виртуальный процессор, который связан с физическим процессором (ядром) или потоком.

// Локальная очередь горутин: Каждый P имеет свою локальную очередь горутин, которые должны быть выполнены.

// Связь с M: P связан с одним M, который выполняет горутины из его локальной очереди.

// Goroutine (G):

// Горутина: G представляет собой легковесный поток, который выполняется параллельно с другими горутинами.

// Распределение: Горутины распределяются между P и выполняются на M.

// Где хранятся очереди горутин:
// Локальная очередь горутин (P):

// Каждый логический процессор (P) имеет свою локальную очередь горутин.

// Эта очередь хранит горутины, которые должны быть выполнены на данном логическом процессоре.

// Поток (M), связанный с логическим процессором, выбирает горутину из локальной очереди и выполняет ее.

// Глобальная очередь горутин:

// В дополнение к локальным очередям, существует глобальная очередь горутин.

// Глобальная очередь используется для балансировки нагрузки между логическими процессорами.

// Если локальная очередь логического процессора пуста, он может забрать горутину из глобальной очереди.

// Как работает распределение горутин:
// Создание горутин:

// Когда горутина создается с помощью ключевого слова go, она добавляется в локальную очередь связанного с ней логического процессора (P).

// Если локальная очередь P заполнена, горутина может быть добавлена в глобальную очередь горутин.

// Выполнение горутин:

// Поток (M), связанный с логическим процессором (P), выбирает горутину из локальной очереди и выполняет ее.

// Если локальная очередь пуста, M может забрать горутину из глобальной очереди или из локальной очереди другого P.

// Переключение контекста:

// Планировщик Go автоматически переключает контекст между горутинами, когда текущая горутина блокируется (например, при ожидании ввода-вывода) или когда происходит вызов time.Sleep.

// Это позволяет эффективно использовать ресурсы процессора и избежать простоев.

// Балансировка нагрузки:

// Если поток (M) завершает выполнение горутины и его локальная очередь пуста, он может забрать горутину из глобальной очереди или из локальной очереди другого P.

// Это позволяет эффективно распределять нагрузку между потоками и логическими процессорами.

// Состояние гонки — это ошибка в многопоточном программировании, которая возникает, когда результат выполнения программы зависит от порядка или времени выполнения конкурирующих операций.
// В Go состояние гонки может возникнуть,
// когда две или более горутины обращаются к одной и той же переменной или ресурсу без надлежащей синхронизации.
// Для обнаружения и исправления состояний гонки можно использовать инструмент -race и примитивы синхронизации, такие как мьютексы или каналы.