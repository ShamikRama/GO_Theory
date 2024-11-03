
// Мапа в go - это хэш таблица ключ - значение 
// Создание мапы :
maps := make(map[string]int) // [string] - ключ int - значение
// Проврка мапы на пустоту
if _, ok := m["key"]; !ok {
	fmt.Println("map is empty")
}


maps := make(map[int]int) // память выделена 
var maps map[int]int // память не выделена 

// мы не можем добавить занчение с невыделенной памяти 

// мапа как и все в го представлена в виде структуры
// ключи в мапе неупорядочены, стоят они рандомно, чтобы не зацикливаться на порядке(безопасность)
// Мы не можем использовать мапу без блокировок параллельно, нужны блокировки
type hmap struct {
	count     int // Количество живых элементов в мапе. Это фактически размер мапы.
	flags     uint8 // Флаги, которые указывают на текущее состояние мапы. Например, флаги могут указывать на то, что мапа находится в процессе роста (ребалансировки).
	B         uint8  // Определяет размер массива бакетов.
	noverflow uint16 // Используется для отслеживания количества бакетов переполнения.
	hash0     uint32 // Семя для хеш-функции. Используется для генерации хеш-кодов ключей.

	buckets    unsafe.Pointer // Указатель на массив бакетов
	oldbuckets unsafe.Pointer // Указатель на предыдущий массив бакетов.
	nevacuate  uintptr        // Счетчик прогресса эвакуации.

	extra *mapextra //  Дополнительные поля, которые могут быть использованы для оптимизации работы мапы.
}

// Для создания уникального индекса пары ключ - значения испоьзуется хэш-функция
// Если хеш-функция создает один и тот же индекс для разных ключей, это называется коллизией. 
// В Go для решения коллизий используется метод цепочек (separate chaining). 
// Это означает, что каждый бакет (bucket) может содержать связанный список элементов(бакетов), которые хешируются в один и тот же индекс.

// buckets    unsafe.Pointer - Указатель на массив бакетов. 
// Каждый бакет содержит список элементов, которые хешируются в один и тот же индекс.

// Структура бакета
type bmap struct {
	tophash [8]uint8 // Массив хеш-кодов верхних битов ключей, хранящихся в бакете
	keys    [8]string // Массив ключей, хранящихся в бакете
	values  [8]int // Массив значений, хранящихся в бакете
	overflow *bmap // Указатель на следующий бакет переполнения(по сути такой же бакет с той же структурой) ->
	// -> Используется для хранения элементов, которые не помещаются в текущий бакет.
}
// Если бакет, соответствующий индексу, заполнен, новый элемент добавляется в бакет переполнения, на который указывает поле overflow

// Функция для вычисления хеш-кода строки
func hash(key string) uint32 {
	h := fnv.New32a() // Создает новый хеширующий объект с использованием алгоритма FNV-1a (32-битная версия).
	h.Write([]byte(key)) // Добавляет байты строки в хеширующий объект.
	return h.Sum32() // Возвращает 32-битный хеш-код строки.
}
// Хэши могут быть одинаковыми, затем можно высилить на основе хэша индекс бакета
// Индексы тоже могут быть одинаковыми(это явления коллизии)
 
// Функция для добавления элемента в хеш-таблицу
func (m *hmap) insert(key string, value int) {
	hashValue := hash(key)
	bucketIndex := hashValue % uint32(m.count) // вычисляет индекс

	bucket := m.buckets[bucketIndex] // получает бакет то есть структур bmap
	for i := 0; i < 8; i++ { // здесь мы проходимся до 8 потому что всего 8 элементов(8 ключей, 8 знач и т.д)
		if bucket.keys[i] == "" {
			bucket.keys[i] = key
			bucket.values[i] = value
			return
		}
	}

	// Если бакет заполнен, создаем новый бакет переполнения
	newBucket := &bmap{}
	bucket.overflow = newBucket
	bucket = newBucket
	bucket.keys[0] = key
	bucket.values[0] = value
}

func main() {
	// Создаем хеш-таблицу с 16 бакетами
	bucketCnt := 16
	m := &hmap{
		buckets:   make([]*bmap, bucketCnt),
		count: bucketCnt,
	}

	// Инициализируем бакеты
	for i := 0; i < bucketCnt; i++ {
		m.buckets[i] = &bmap{}
	}

	// Добавляем элементы в хеш-таблицу
	m.insert("apple", 1)
	m.insert("banana", 2)
	m.insert("cherry", 3)
	m.insert("date", 4)
	m.insert("elderberry", 5)

	// Выводим содержимое хеш-таблицы
	for i, bucket := range m.buckets {
		fmt.Printf("Bucket %d:\n", i)
		for j := 0; j < 8; j++ {
			if bucket.keys[j] != "" {
				fmt.Printf("  Key: %s, Value: %d\n", bucket.keys[j], bucket.values[j])
			}
		}
		for bucket.overflow != nil {
			bucket = bucket.overflow
			for j := 0; j < 8; j++ {
				if bucket.keys[j] != "" {
					fmt.Printf("  Overflow Key: %s, Value: %d\n", bucket.keys[j], bucket.values[j])
				}
			}
		}
	}
}