package output

import (
	"testing"
)

func TestProcessFile(t *testing.T) {
	// Пример теста для функции processFile
	// Здесь вы можете создать временный файл и проверить, что функция работает корректно
	t.Run("Test with valid input", func(t *testing.T) {
		// Создайте временный файл и проверьте результат
	})

	t.Run("Test with invalid input", func(t *testing.T) {
		// Проверьте, как функция обрабатывает неверный ввод
	})
}

func TestPrintLine(t *testing.T) {
	// Пример теста для функции printLine
	t.Run("Test print valid line", func(t *testing.T) {
		// Проверьте, что функция корректно выводит строку
	})

	t.Run("Test print empty line", func(t *testing.T) {
		// Проверьте, как функция обрабатывает пустую строку
	})
}
