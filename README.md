# goBucket

> Есть определенный бакет с товарами. Бакет состоит из массива списка товаров. Товары расположены в списках рандомным образом, так же товары отсортированы по сроку годности. Списки в бакете разбиты символом ";". Нужно написать оптимальный алгоритм для выборки продуктов из бакета, при учете срока годности. После выборки товаров из бакета, нужно вернуть новое состояние бакета.
>**Входные данные**
>**Бакет:** 1,2,3,5,5; 2,5,4,3,1; 3,5,4,1,1; 5,1,1,1,1
>**Заказ:** 1,2,3,4,5

Для того чтобы сделать имитацию выбора, при выборе оптимального товара мы на его месте будем ставить **0**.
**Результат**
**Бакет после заказа:** 0,2,3,5,5; 0,5,0,3,1; 0,5,4,1,1; 0,1,1,1,1

## Функции

Было создано три функции, которые работают по разному, но делают одно и тоже. 

Функция **Order** использует стандартный алгоритм перебора каждого элемента в массиве. Для этой функции была создана структура оптимального товара *OptimalItem*, которая хранит в себе номер списка и номер товара в списке. При создании мы инициализируем ее поля значениями -1, для того чтобы знать, что там еще не содержиться ни одного оптимального выбора. После того как поля станут отличными от -1, начинается сравнение с элементами бакета. Если алгоритм находит нужный нам товар в бакете на первой позиции (т.е. 0 в массиве) мы его считаем оптимальным, и переходим к следующему товару в списке заказов. И так до конца перебора списка заказов.

Функция **Order2** использует алгоритм вертикального перебора, но без цикла **for**, т.е. просто прыгая по меткам. В этом алгоритме мы каждый элемент проверяем сначала на позиции 0 в каждом списке товаров. Если не находим, увеличиваем позицию и т.д.

Функция **FasterOrder** использует также алгоритм вертикального перебора, но используя цикл **for** для перебора списков товаров в бакете.

## Benchmark

Benchmark Golang показал такие результаты для этих трех функций:

#### Первый прогон

```
scriptkiller@scriptkiller-X550MJ:~/go/src/github.com/BogdanYanov/goBucket/tests$ go test -bench=. -benchmem
goos: linux
goarch: amd64
pkg: github.com/BogdanYanov/goBucket/tests
BenchmarkFasterSearch-2          2894644               393 ns/op           5.08 MB/s           0 B/op          0 allocs/op
BenchmarkSearch-2                2290870               547 ns/op           3.66 MB/s           0 B/op          0 allocs/op
BenchmarkSearch2-2               2050272               580 ns/op           3.45 MB/s           0 B/op          0 allocs/op
PASS
ok      github.com/BogdanYanov/goBucket/tests   5.175s
```

#### Второй прогон

```
scriptkiller@scriptkiller-X550MJ:~/go/src/github.com/BogdanYanov/goBucket/tests$ go test -bench=. -benchmem
goos: linux
goarch: amd64
pkg: github.com/BogdanYanov/goBucket/tests
BenchmarkFasterSearch-2          2994548               387 ns/op           5.16 MB/s           0 B/op          0 allocs/op
BenchmarkSearch-2                2402263               494 ns/op           4.05 MB/s           0 B/op          0 allocs/op
BenchmarkSearch2-2               2126692               564 ns/op           3.55 MB/s           0 B/op          0 allocs/op
PASS
ok      github.com/BogdanYanov/goBucket/tests   5.048s
```

### Третий прогон

```
scriptkiller@scriptkiller-X550MJ:~/go/src/github.com/BogdanYanov/goBucket/tests$ go test -bench=. -benchmem
goos: linux
goarch: amd64
pkg: github.com/BogdanYanov/goBucket/tests
BenchmarkFasterSearch-2          3054572               389 ns/op           5.14 MB/s           0 B/op          0 allocs/op
BenchmarkSearch-2                2415546               508 ns/op           3.94 MB/s           0 B/op          0 allocs/op
BenchmarkSearch2-2               2159890               564 ns/op           3.55 MB/s           0 B/op          0 allocs/op
PASS
ok      github.com/BogdanYanov/goBucket/tests   5.135s
```

Из этого следует что функция **FasterSearch** работает в 1,3 раза быстрее.
