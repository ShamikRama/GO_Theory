// Каналы (channels) в Go — это механизм для синхронизации и обмена данными между горутинами (goroutines).
// Они обеспечивают безопасный способ передачи значений между горутинами, 
// позволяя вам координировать их работу.

// БУФЕРИЗИРОВАННЫЙ КАНАЛ

// Создание канала:
// Каналы создаются с помощью функции make. 
//Вы можете создать канал для передачи данных определенного типа:
// Каналы в HEAP(куча)

ch := make(chan int, 4) // буферризированный канал с емкостью 4(там может быть 4 элемента)

// Свойства:
// 1)потокобезопасность - с одним каналом могут оработать несколько горутин без mutex
// 2)хранение элементов, семантика FIFO
// 3)передача данных между горутинами
// 4)блокировка горутин

// С каналами работают горутины 

// Вот так устроены каналы под капотом :

type hchan struct {
	qcount   uint           // количество элементов в очереди 
	dataqsiz uint           // размерность буфера
	buf      unsafe.Pointer // ссылка на сам буфер
	elemsize uint16 // сложная тема
	closed   uint32 // закрыт ли канал
	elemtype *_type // указатель на тип данных, который может хранить канал.
	sendx    uint   // Индекс для следующей отправки данных в буферизованный канал.
	recvx    uint   //  Индекс для следующего получения данных из буферизованного канала.
	recvq    waitq  // list of recv waiters
	sendq    waitq  // очередь горутин, ссылка на структуру waitq
	lock mutex // мьютекс, используемый для операций, изменяющих состояние канала
 }
 
// У канала кольцевая очередь(до края и обратно)

// Как происходит чтение и запись в канал:
ch <- // запись
// Если отправитель(горутина) послала данные в канал:
//1)заблокируется mutex
//2) данные скопированы в канал(это копия, у сендера своя копия а у ридера своя копия, если не ссылка)
//3) разблокируется mutex
data := <-ch // чтение 
//1)заблокируется mutex
//2)ридер скопировал данные
//3)разблокируется mutex

// Buffer overflow (переполнение):
// Если буффер заполнен, горутина ставится на паузу, вызовется функция gopark(ставится на паузу)
// Если пришел ридер, он причитает данные из очереди, то есть из 0 индекса
// Потом ридер будет указывать на след ячейку
// Затем горутина-сендлер запишет данные в 0 индекс
// Потом сендлер указывает на след ячейку, на ту же что и ридер

// Все действия над спящей горутиной производит та горутина которая пришла за данными(ридер)

//РИДЕР ПРИШЕЛ ПЕРВЫЙ ЧТЕНИЕ ИЗ ПУСТОГО КАНАЛА
//если ридер пришел первый, он ставится на паузу
//затем приходит сендлер и он не передает данные в буффер канала а кладет их напрямую в стэк горутины-ридера


// ----------------------- НЕБУФЕРИЗИРОВАННЫЙ КАНАЛ ---------------------------------------------

// Создается канал без стэка без очереди
// Данные будут передаваться из стэка одной горутины к стэку другой напрямую

// .....



// ---------------------- SELECT{}---------------------------------------------------------------
// Чтение из канала не должно быть блокирующим
// Чтение из канала можно, но будет возвращать zero-value
// Запись в закрытый канал приведет к панике


// 1. Отправка в канал или получение из канала блокирует горутину

// Если отправить данные в канал, и нет другой горутины, готовой принять эти данные, то горутина заблокируется. 
// Точно так же, если попытаться получить данные из канала и в нем нет данных, то получим deadlock (взаимную блокировку).
//  Это позволяет организовать естественную синхронизацию между горутинами:



// 1. Отправка в канал или получение из канала блокирует горутину

// 2. Отправка в nil канал и получение из nil канала блокируется навсегда

// 3. Закрытие nil канала приведет к панике

// 4. Закрытие уже закрытого канала приведет к панике

// 5. Чтение из закрытого канала возвращает zero-value

// 6. Запись в закрытый канал приведет к панике