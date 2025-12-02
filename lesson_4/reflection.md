# Рефлексия

Эталонное решение:

```java
abstract class DynArray<T>

  // конструктор
// постусловие: создан пустой массив
  public DynArray<T> DynArray();

  // команды
// предусловие: i лежит в допустимых границах массива; 
// постусловие: значение элемента i изменено на value
  public void put(i, T value);

// предусловие: i лежит в допустимых границах массива; 
// постусловие: перед элементом i добавлен 
// новый элемент с значением value; 
  public void put_left(i, T value);

// предусловие: i лежит в допустимых границах массива; 
// постусловие: после элемента i добавлен 
// новый элемент с значением value;
  public void put_right(i, T value);

// предусловие: нет; 
// постусловие: в хвост массива добавлен 
// новый элемент
  public void append(T value); 

// предусловие: i лежит в допустимых границах массива; 
// постусловие: элемент i удалён из массива;
  public void remove(int i); 

  // запросы 
// предусловие: i лежит в допустимых границах массива;
  public T get(int i); // значение i-го элемента 
  public int size(); // текущий размер массива 

  // запросы статусов (возможные значения статусов)
  public int get_put_status(); // успешно; индекс за пределами массива
  public int get_put_left_status(); // успешно; индекс за пределами массива
  public int get_put_right_status(); // успешно; индекс за пределами массива
  public int get_remove_status(); // успешно; индекс за пределами массива
  public int get_get_status(); // успешно; индекс за пределами массива
```

Проанализировав эталонное решение, я выявил недочеты:
- не были добавлены методы: put_left, put_right, append и соответствующие дополнительные запросы состояний

Решение после рефлексии:

```go
type Status int

Status (
	REMOVE_NIL = iota // команда Remove() еще не вызывалась 
	REMOVE_OK // последняя команда Remove() отработала хорошо
	REMOVE_OUT_OF_RANGE // последняя команда Remove() выполнилась с ошибкой 
						//out of range
)

Status (
	GET_NIL = iota // операция Get() еще не вызывалась
	GET_OK // последняя операции Get() отработала корректно
	GET_EMPTY_ARRAY // последняя операция Get() закончилась с ошибкой
					// доступ к пустому массиву - empty array
	GET_OUT_OF_RANGE // последняя операция Get() закончилась с ошибкой
					// out of range
)

Status (
	ADD_NIL = iota // команда Add() еще не вызывалась 
	ADD_OK // последняя команда Add() выполнилась успешно
	ADD_OUT_OF_RANGE // последняя команда Add() выполнилась с ошибкой 
					//out of range
)

Status (
	ADD_LEFT_NIL = iota // команда AddLeft() еще не вызывалась 
	ADD_LEFT_OK // последняя команда AddLeft() выполнилась успешно
	ADD_LEFT_OUT_OF_RANGE // последняя команда AddLeft() выполнилась с ошибкой 
					//out of range
)

Status (
	ADD_RIGHT_NIL = iota // команда AddRight() еще не вызывалась 
	ADD_RIGHT_OK // последняя команда AddRight() выполнилась успешно
	ADD_RIGHT_OUT_OF_RANGE // последняя команда AddRight() выполнилась с ошибкой 
					//out of range
)

const deallocateLimit = 0.25
const zeroLen = 0
const capacityLimit = 16
const reduceValue = 1.5
const raiseValue = 2

type DynArray[T any] struct {
	len int
	cap int
	array []T
	
	addStatus Status
	removeStatus Status
	getStatus Status
	
	// конструктор:
	DynArray[T](cap int) (*DynArray[T])
	
	// команды:
	
	// предусловие: вместимость списка не меньше indx
	// постусловие: элемент под индексом - indx удален
	// при необходимости редуцирует занимаемую память
	Remove(indx int)
	
	// предусловие: вместимость списка не меньше indx
	// постусловие: в динамический массив добавлен новый элемент
	// при необходимости идет реаллокация
	Add[T](value T, indx int)
	
	// предусловие: indx в допустимых диапозонах
	// постусловие: после indx вставлен элемент value
	AddRight[T](value T, indx int)
	
	// предусловие:indx в допустимых диапозонах
	// постусловие: перед indx вставлен элемент value
	AddLeft[T](value T, indx int)
	
	// постусловие: в конец массива добавлен элемент
	Append[T](value T)
	
	// запросы:
	
	// предусловие: список не пуст
	// постусловие: вернется элемент под индексом indx
	Get[T](indx int) (T)
	
	// постулосвие: вернется актуальный размер массива
	Size() (int) {
		return len
	}
	
	// постусловие: вернется актуальная вместимость массива
	Capacity() (int) {
		return cap
	}
	
	// дополнительные запросы:
	GetRemoveStatus() (Status) // успешно; длина списка меньше indx
	GetAddStatus() (Status) // успешно; длина списка меньше indx
	GetGetStatus() (Status) // успешно; список пустой
	GetGetRightStatus() (Status) // успешно; индекс за пределами массива
	GetGetLeftStatus() (Status) // успешно; индекс за пределами массива
	
	// дополнительные приватные методы:
	
	// запросы:
	
	// постусловие: возвращается новый пустой массив меньшего размера
	// в который будут переносится значения старого массива
	deallocateArray[T]() ([]T)
	
	// постусловие: возвращается новый пустой массив большего размера
	// в который будут переносится значения старого массива
	allocateArray() ([]T)
	
	// постусловие: возвращается новый, заполненный старыми значениями, массив
	copyRangeFromTo(from, to []T) ([]T)
	
}
```