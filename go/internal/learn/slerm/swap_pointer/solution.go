package swap_pointer

/**
Напишите функцию, которая меняет местами две переменных через указатель.
*/

func swap(a *int, b *int) {
	*a, *b = *b, *a
}
