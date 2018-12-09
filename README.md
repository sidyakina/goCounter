### Задание
 Программа читает строки, содержащие URL. На каждый URL нужно отправить HTTP-запрос методом GET
  и посчитать кол-во вхождений строки "Go" в теле ответа. В конце работы приложение выводит на экран
   общее кол-во найденных строк "Go" во всех переданных URL, например:
	Count for https://golang.org: 9
	Count for https://golang.org: 9
	Count for https://golang.org: 9
	Total: 27
 Каждый URL должен начать обрабатываться сразу после вычитывания и параллельно с вычитыванием следующего. 
 URL должны обрабатываться параллельно, но не более k=5 одновременно. 
 Обработчики URL не должны порождать лишних горутин, т.е. если k=1000 а обрабатываемых URL-ов нет,
  не должно создаваться 1000 горутин.
 Нужно обойтись без глобальных переменных и использовать только стандартную библиотеку.
 