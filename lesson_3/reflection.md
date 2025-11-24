# Рефлексия:

Эталонное решение:

```java
abstract class ParentList<T>

  // конструктор
  public ParentList<T> ParentList();

  // команды
  public void head(); 
  public void tail(); 
  public void right(); 
  public void put_right(T value); 
  public void put_left(T value); 
  public void remove();
  public void clear(); 
  public void add_tail(T value); 
  public void remove_all(T value);
  public void replace(T value); 
  public void find(T value); 

  // запросы
  public T get();
  public bool is_head();
  public bool is_tail();
  public bool is_value();
  public int size();

  // запросы статусов
  public int get_head_status();
  public int get_tail_status();
  public int get_right_status();
  public int get_put_right_status();
  public int get_put_left_status();
  public int get_remove_status();
  public int get_replace_status();
  public int get_find_status();
  public int get_get_status();
  
abstract class LinkedList<T> : ParentList<T>

  // конструктор
  public LinkedList<T> LinkedList();

abstract class TwoWayList<T> : ParentList<T>

  // конструктор
  public TwoWayList<T> TwoWayList();

// предусловие: левее курсора есть элемент; 
// постусловие: курсор сдвинут на один узел влево
  public void left();

  public int get_left_status(); // успешно; левее нету элемента
```

Проанализировав эталонное решение, я выявил один недочет:
1. Не совсем верно определил предусловие для метода left()

В целом, задание решил верно!

Решение после рефлексии:

```go
type Status int
type ParentLinkedList struct {
	// конструктор
	ParentLinkedList[T]() (ParentLinkedList[T])
	
	// команды:
	
	// предусловие: список не пуст
	// постусловие: курсор находится на первом узле списка
	Head()
	
	// предусловие: список не пуст
	// постусловие: курсор установлен на последний элемент в списке
	Tail()
	
	// предусловие: длина списка не меньше 2 
	// и курсор не стоит на последнем элементе
	// постусловие: курсор сдвинут на один элемент вправо
	Right()
	
	// предусловие: список не пуст
	// постусловие: длина списка увеличилась на 1
	// и справа от текущего узла появился узел
	PutRight[T](value T)
	
	// предусловие: список не пуст
	// постусловие: длина списка увеличилась на 1
	// и слева от текущего узла появился узел
	PutLeft[T](value T)
	
	// предусловие: список не пуст
	// постусловие: длина списка уменьшилась на единицу
	// курсор смещён к правому соседу, если он есть, 
	// в противном случае курсор смещён к левому соседу,
	// если он есть
	Remove()
	
	// постусловие: список очищен от всех элементов
	Clear()
	
	// предусловие: список пуст
	// постуловие: список имеет длину в один элемент
	AddToEmpty[T](value T)
	
	// постулосвие: курсор встанет на следующий от себя элемент, который равен значению аргумента, если такой узел найден
	Find[T](value T)
	
	// постуловие: длина списка увеличилась на 1, добавлен новый элемент в хвост
	AddTail[T](value T)
	
	// предусловие: список не пуст
	// постусловие: значения элемента текущего курсора изменилось на новое
	Replace[T](value T)
	
	// постусловие: все элементы равные значению аргумента - удалены
	RemoveAll[T](value T)
	
	// запросы:
	
	// предусловие: список имеет минимум один элемент
	// постуловие: вернется элемент под курсором в данный момент
	Get[T]() (T)
	
	// постусловие: вернется целое число - количество элементов в списке
	Size() (Size)
	
	// предусловие: список не пуст
	IsHead() ValidCase
	
	// предусловие: список не пуст
	IsTail() ValidCase
	
	// предусловие: список не пуст
	IsValue() ValidCase
	
	// дополнительные запросы
	
	GetHeadStatus() (Status) // успешно; список пуст
	GetTailStatus() (Status) // успешно; список пуст
	GetRightStatus() (Status) // успешно; правее нету элемента
	GetPutRightStatus() (Status) // успешно; список пуст
	GetPutLeftStatus() (Status) // успешно; список пуст
	GetRemoveStatus() (Status) // успешно; список пуст
	GetAddToEmptyStatus() (Status) // успешно
	GetReplaceStatus() (Status) // успешно; список пуст
	GetFindStatus() (Status) // следующий найден; 
                       // следующий не найден; список пуст
    GetGetStatus() (Status) // успешно; список пуст
}

// implements ParentLinkedList
type LinkedList struct {
	// реализация в прошлом занятии
}

// implements ParentLinkedList
type DoublyLinkedList struct {
	// конструктор
	DoublyLinkedList[T]() (DoublyLinkedList[T])
	
	// команды:
	
	// предусловие: левее курсора есть элемент
	// постусловие: курсор сдвинут на один элемент влево
	Left()
	
	// дополнительные запросы:
	GetLeftStatus() (Status) // успешно; левее нет элемента 
}
```