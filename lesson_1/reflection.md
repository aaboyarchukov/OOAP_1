# Рефлексия

Эталонное решение:

```java
abstract class BoundedStack<T>

    public const int POP_NIL = 0; // push() ещё не вызывалась
    public const int POP_OK = 1; // последняя pop() отработала нормально
    public const int POP_ERR = 2; // стек пуст

    public const int PEEK_NIL = 0; // push() ещё не вызывалась
    public const int PEEK_OK = 1; // последняя peek() вернула корректное значение 
    public const int PEEK_ERR = 2; // стек пуст

    public const int PUSH_OK = 1; // последняя push() отработала нормально
    public const int PUSH_ERR = 2; // в стеке нет свободного места 

    // конструктор
// постусловие: создан новый пустой стек
    public BoundedStack<T> BoundedStack(int max_size);


    // команды:
// предусловие: в стеке менее максимального кол-ва элементов
// постусловие: в стек добавлено новое значение
    public void push(T value); 

// предусловие: стек не пустой; 
// постусловие: из стека удалён верхний элемент
    public void pop(); 

// постусловие: из стека удалятся все значения
    public void clear();


    // запросы:
// предусловие: стек не пустой
    public T peek(); 

    public int size();
    public int max_size();

    // дополнительные запросы:
    public int get_pop_status(); // возвращает значение POP_*
    public int get_peek_status(); // возвращает значение PEEK_*
    public int get_push_status(); // возвращает значение PUSH_*
```

Проанализировав эталонное решение, выделил несколько моментов:
1. Я не указал в комментариях, что именно делают дополнительные запросы для возврата статуса операций
2. Не выделял в разные группы команды и запросы, выделяя их таким образом, будто легче читается код
3. Также я добавил начальный статус для операции Push(), которая говорит о том, что стек пуст
4. Еще я обозначил некоторые возвращаемые значения и статусы именовынами типами (с помощью alias), для простоты чтения кода
5. Также я использовал встроенный функционал (iota - перечисление) для определения статусов

В целом решение было верным, так как все операции были реализованы, а самое главное отмечены ограничения (-пред и -пост условия)

Решение после рефлексии:

```go
// АТД BoundedStack
type Status int
type Size int
type Capacity int

type BoundedStack [T any]struct {
	boundedStackStorage List[T] // хранилище стека
	capacity Capacity // заданная вместимость хранилища, по-умолчанию: 32
	peekStatus Status // аттрибут состояния запроса Peek()
	popStatus Status // аттрибут состояния команды Pop()
	pushStatus Status // аттрибут состояния команды Push(), надо контроллировать переполнение
}

var DEFAULT_CAP Capacity = 32

Status (
	POP_NIL = iota // Push() еще не вызывалась (нужен для более точной диагностики)
	POP_OK // все корректно, последняя Pop() отработала нормально
	POP_ERR // стек пуст
)

Status (
	PEEK_NIL = iota // Push() еще не вызывалась (нужен для более точной диагностики)
	PEEK_OK // последняя Peek() отработала корректно
	PEEK_ERR // стек пустой
)

Status (
	PUSH_NIL = iota // стек пуст
	PUSH_OK // последняя Push() отработала корректно
	PUSH_ERR // стек заполнен
)

// конструктор (запрос)
// постусловие: создан новый пустой стек с заданной вместимостью (Capacity)
func BoundedStack[T](cap ...Capacity) (BoundedStack[T]) {
	capacity := DEFAULT_CAP // немного громоздкая обработка необязательных аргументов в Go
	if len(cap) > 0 {
		capacity = cap[0]
	}
	
	storage := List(capacity)
	return BoundedStack[T]{
		boundedStackStorage: storage,
		capacity: capacity,
		peekStatus: PEEK_NIL,
		popStatus: POP_NIL,
		pushStatus: PUSH_NIL,
	}
}

// подробная спецификация
type BoundedStackInterface[T any] interface {
	// команды
	
	// предусловие: стек не заполнен
	// постусловие: в стек помещен новый элемент
	Push[T](value T)
	
	// команда
	// предусловие: стек не пустой
	// постусловие: из стека удален верхний элемент
	Pop()
	
	// команда
	// постусловие: стек очищен -> удаляются все значения
	Clear()
	
	// запросы:
	
	
	
	// предусловие: стек не пустой
	// постусловие: вернется верхний элемент из стека
	Peek[T]() (T)
	
	// постусловие: возвращается размер стека
	Size() (Size)
	
	// дополнительные запросы:
	
	// возвращается акутальный статус последней выполненного запроса Peek()
	// -> значение PEEK_*
	GetPeekStatus() (Status)
	// возвращается акутальный статус последней выполненной команды Pop()
	// -> значение POP_*
	GetPopStatus() (Status)
	// возвращается акутальный статус последней выполненной команды Push()
	// -> значение PUSH_*
	GetPushStatus() (Status)
}

func (s *BoundedStack) Push[T](value T) {
	var size Size := s.Size()
	if size == s.capacity {
		s.pushStatus = PUSH_ERR
	} else {
		s.boundedStackStorage.Append(value) 
		s.pushStatus = PUSH_OK
	}
}

func (s *BoundedStack) GetPushStatus() (Status) {
	return s.pushStatus
}

func (s *BoundedStack) Pop() {
	var size Size := s.Size() 
	if size > 0 {
		s.boundedStackStorage.RemoveAt(-1)
		s.popStatus = POP_OK
	} else {
		s.popStatus = POP_ERR
	}
}

func (s *BoundedStack) GetPopStatus() (Status) {
	return s.popStatus
}

func (s *BoundedStack) Peek[T]() (T) {
	var size Size := s.Size()
	var result T 
	if size > 0 {
		result = s.boundedStackStorage[-1]
		s.popStatus = PEEK_OK
	} else {
		s.popStatus = POP_ERR
	}
	
	return result
}

func (s *BoundedStack) GetPeekStatus() (Status) {
	return s.peekStatus
}

func (s *BoundedStack) Clear() {
	// так как List уже реализован,
	// значит у него должен быть конструктор
	// и мы возвращаем пустой List
	s.boundedStackStorage = List(s.capacity)
	
	// устанавливаем начальные статусы
	s.popStatus = POP_NIL
	s.peekStatus = PEEK_NIL
	s.pushStatus = PUSH_NIL
}

func (s *Stack) Size() (Size) {
	return s.stackStorage.Length()
}
```