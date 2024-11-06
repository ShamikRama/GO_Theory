sacas
// Контекст в golang (context) нужен для того, чтобы управлять жизненном
// циклом горутин и передачи данных между ними.
// Он позволяет вам отменять операции, устанавливать таймауты и передавать значения между горутинами. 
// Контекст особенно полезен в ситуациях, когда вам нужно управлять множеством горутин над одной задачей.

// Основные возможности контекста:
// 1)отмена операции
// 2)установка таймаута для операции
// 3)передача значений - сомнительно


// Пакет context в Go полезен при взаимодействиях с API и медленными процессами, 
// особенно в production-grade системах, которые занимаются веб-запросами. 
// С его помощью можно уведомить горутины о необходимости завершить свою работу.


