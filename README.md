# goBucket

> Есть определенный бакет с товарами. Бакет состоит из массива списка товаров. Товары расположены в списках рандомным образом, так же товары отсортированы по сроку годности. Списки в бакете разбиты символом ";". Нужно написать оптимальный алгоритм для выборки продуктов из бакета, при учете срока годности. После выборки товаров из бакета, нужно вернуть новое состояние бакета.  
>**Входные данные**  
>**Бакет:** [ [1,2,3,5,5], [2,5,4,3,1], [3,5,4,1,1], [5,1,1,1,1] ]  
>**Заказ:** 1,2,3,4,5

**Результат:** [ [2,3,5,5], [5,3,1], [5,4,1,1], [1,1,1,1] ]

## Update

1. Коментарии в коде, и стандарты написание кода. (Критично)  
  + Некоторые функции, структуры имеют комменты, а некоторые нет. Все функции, структуры, константы и т.д которые экспортируются нужно комментировать. В этом моменте тебе может помочь "линтеры" и "форматеры". go vet, gofmt, golang.org/x/lint/golint - этого набора хватает.
  
  **Ответ:** добавил комментарии ко всем экспортируемым функциям и структурам. Проверил все через go vet, golint.
  
---

2. Именование пакетов (Критично)  
  + Стоит почитать про именование пакетов в go. Сходу не совсем ясна суть пакета model, constructor, это может привести к проблемам на больших проектах. Посмотреть про именование можно тут https://talks.golang.org/2014/organizeio.slide, и в Go Effective.
  
  **Ответ:** оставил один пакет *bucket*. Подумал, если мы все делаем с бакетом, значит все в этом пакете и будет.

---

3. Тестирование (Критично)  
  + Зачем было выделена папка tests? Почему не сделал constructor_test.go и т.д?
 
  **Ответ:** поправил, теперь функции поиска заказа в bucket_test.go

---

 4. Излишние вложенности (усложняют тестирование), есть такая вещь как "unnecessary nesting logic". Старайся избегать таких моментов. (Критично)
 
  **Ответ:** не знаю правильно ли обыграл эти моменты, но по максимуму старался их переделать.

## Функции

```
func New(sizeL, sizeI, maxProdIdx int) *Bucket
```

Создает новый бакет с количеством списков *sizeL* и количеством товаров *sizeI* в каждом из списков, которые рандомно заполняются индексами от 1 до maxProdIdx.

---

```
func (b *Bucket) ShowBucket()
```

Выводит бакет на экран.

---

```
func (b *Bucket) ReceiveOrderSlow(order ...int) Bucket
```

Выполняет поиск товаров в бакете последовательно, запоминая последний оптимальный вариант, после всего поиска удаляет его из бакета, и после этого возращает новый бакет.

---

```
func (b *Bucket) ReceiveOrderFast(order ...int) Bucket
```

Выполняет поиск товаров в бакете вертикально, до первого нахождения товара, удаляет его из бакета, и после поиска всех товаров в заказе, возращает новый бакет

---

```
func (b *Bucket) ReceiveOrder(order ...int) Bucket
```

Выполняет поиск товаров в бакете также вертикально, но в отличии от ```ReceiveOrderFast(order ...int)``` использует один бесконечный цикл **for**.

## Benchmark

Benchmark Golang показал такие результаты для этих трех функций:

#### Первый прогон

```
scriptkiller@scriptkiller-X550MJ:~/go/src/github.com/BogdanYanov/goBucket/bucket$ go test -bench=. -benchmem
goos: linux
goarch: amd64
pkg: github.com/BogdanYanov/goBucket/bucket
BenchmarkBucket_ReceiveOrderSlow-2        811579              1444 ns/op             368 B/op          6 allocs/op
BenchmarkBucket_ReceiveOrderFast-2        914110              1274 ns/op             368 B/op          6 allocs/op
BenchmarkBucket_ReceiveOrder-2            869455              1320 ns/op             368 B/op          6 allocs/op
PASS
ok      github.com/BogdanYanov/goBucket/bucket  4.279s
```

#### Второй прогон

```
scriptkiller@scriptkiller-X550MJ:~/go/src/github.com/BogdanYanov/goBucket/bucket$ go test -bench=. -benchmem
goos: linux
goarch: amd64
pkg: github.com/BogdanYanov/goBucket/bucket
BenchmarkBucket_ReceiveOrderSlow-2        813232              1455 ns/op             368 B/op          6 allocs/op
BenchmarkBucket_ReceiveOrderFast-2        907250              1371 ns/op             368 B/op          6 allocs/op
BenchmarkBucket_ReceiveOrder-2            856255              1409 ns/op             368 B/op          6 allocs/op
PASS
ok      github.com/BogdanYanov/goBucket/bucket  6.376s
```

### Третий прогон

```
scriptkiller@scriptkiller-X550MJ:~/go/src/github.com/BogdanYanov/goBucket/bucket$ go test -bench=. -benchmem
goos: linux
goarch: amd64
pkg: github.com/BogdanYanov/goBucket/bucket
BenchmarkBucket_ReceiveOrderSlow-2        814268              1426 ns/op             368 B/op          6 allocs/op
BenchmarkBucket_ReceiveOrderFast-2        927988              1297 ns/op             368 B/op          6 allocs/op
BenchmarkBucket_ReceiveOrder-2            943402              1318 ns/op             368 B/op          6 allocs/op
PASS
ok      github.com/BogdanYanov/goBucket/bucket  5.388s
```
