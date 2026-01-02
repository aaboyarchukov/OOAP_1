# Рефлексия

Эталонное решение:

```java
abstract class PowerSet<T> : HashTable<T>

  // конструктор
// постусловие: создано пустое множество
// на максимальное количество элементов sz
  public PowerSet<T> PowerSet(int sz); 

  // запросы
// возвращает пересечение текущего множества
// с множеством set
  public PowerSet<T> Intersection(PowerSet<T> set);

// возвращает объединение текущего множества
// и множества set
  public PowerSet<T> Union(PowerSet<T> set);

// возвращает разницу между текущим множеством
// и множеством set
  public PowerSet<T> Difference(PowerSet<T> set);

// проверка, будет ли set подмножеством
// текущего множества
  public bool IsSubset(PowerSet<T> set);
```

Проанализировав эталонное решение выявил несколько недочетов:
- не наследовал класс HashTable, хотя это было необходимо так как большинство методов повторялись, да и по своему смыслу довольно схожие структуры
- определил постусловия для алгебраических методов на множествах, хотя они не нужны, так как данные методы полностью опредлены на данном типе
- также сделал лишний метод Equals

Решение после рефлексии:

```go
// implement HashTable class
type PowerSet[T any] struct {
	// constructor
	PowerSet[T]() PowerSet[T]
	
	// get set with all elements from two sets
	Union(powerSetFirst PowerSet[T], powerSetSecond PowerSet[T]) PowerSet[T]
	
	// get set with elements from first set which are in second set
	Intersection(powerSetFirst PowerSet[T], powerSetSecond PowerSet[T]) PowerSet[T]
	
	// get set with elements from first set which are not in second set
	Difference(powerSetFirst PowerSet[T], powerSetSecond PowerSet[T]) PowerSet[T]
	
	// check if first set contains all elements from second set 
	Issubset(powerSetFirst PowerSet[T], powerSetSecond PowerSet[T]) bool
}
```